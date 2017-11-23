package server

import (
    "net"
    "strconv"
    "errors"
    "fmt"
    "container/list"
    "strings"
)

type RedisObj struct {
    Value interface{}
}

type Client struct{
    IP string
    Port int
    Conn net.Conn
    Inputbuffer []byte
    CurrentCommand Command
    CommandArgv  []string
}

type Server struct {
    Port int
    Addr string
    ClientList []Client
    Dict map[string]RedisObj
    ExpireDict map[string]int
    Webport int
    MasterAddr string
    MasterPort int
    Role uint8
    SlaveList []Client
}

type CommandProc func(client Client) error

type Command struct {
    Name string
    Argc int
    Proc CommandProc 
    NeedProgate uint8
}

func checkCommandProtocol(client *Client) error {
    argv := client.CommandArgv
    argvFlag := argv[0][0:1]
    argvCount, _:= strconv.Atoi(argv[0][1:2])
    fmt.Println(argvCount, client.CurrentCommand, argv)
    if (argvFlag != "*") {
        client.Conn.Write([]byte("$21\r\ninvalid-argument-flag\r\n"))
        return errors.New("invliad-aggument-flag")
    }
    if(argvCount != (client.CurrentCommand.Argc + 1 )) {
        client.Conn.Write([]byte("$22\r\ninvalid-argument-count\r\n"))
        return errors.New("invliad-aggument-count")
    }
    return nil
}



//[*2 $3 get $4 name ]
func GetCommand(client Client) error{
    conn := client.Conn
    argv := client.CommandArgv
    err := checkCommandProtocol(&client)
    if err != nil {
        return err
    }
    resp := ServerInstance.Dict[argv[4]]

    //类型判定
    tmpValue, ok := (resp.Value).(string)
    if !ok {
        responseNil(conn)
    } else {
        responseValue(tmpValue, conn)
    }
    return nil
}

func SetCommand(client Client) error{
    conn := client.Conn
    argv := client.CommandArgv
    err := checkCommandProtocol(&client)
    if err != nil {
        return err
    }
    ServerInstance.Dict[argv[4]] = createRedisObj(argv[6])
    responseOK(conn)
    return nil
}

func LpushCommand(client Client) error {
    conn := client.Conn
    argv := client.CommandArgv
    err := checkCommandProtocol(&client)
    if err != nil {
        return err
    }
    if _, ok := ServerInstance.Dict[argv[4]]; !ok {
        ls := list.New()
        ServerInstance.Dict[argv[4]] = createRedisObj(ls)
    }
    tmpls := ServerInstance.Dict[argv[4]].Value.(*list.List)
    tmpls.PushBack(argv[6])
    responseOK(conn)
    return nil
}

func LpopCommand(client Client) error {
    conn := client.Conn
    argv := client.CommandArgv
    err := checkCommandProtocol(&client)
    if err != nil {
        return err
    }
    if _, ok := ServerInstance.Dict[argv[4]]; !ok {
        responseNil(conn)
    } else {
        ls := ServerInstance.Dict[argv[4]].Value.(*list.List)
        tmpValue := ls.Back().Value.(string)
        responseValue(tmpValue, conn)
    }
    return nil
}

func SlaveofCommand(client Client) error {
    conn := client.Conn
    argv := client.CommandArgv
    err := checkCommandProtocol(&client)
    if err != nil {
        return err
    }
    masterAddr := argv[4]
    masterPort := argv[6]
    go func(addr string, port string, conn net.Conn) error{
        masterConn, err := net.Dial("tcp", addr + ":" + port)
        if err != nil {
           return errors.New("faile-to-connect-to-master")
        }
        fmt.Fprintf(masterConn, "*1\r\n$5\r\nslave\r\n")
        for {
            tmpBuffer := make([]byte, 1024)
            _, err := masterConn.Read(tmpBuffer)
            if err != nil {
                fmt.Println("reader-from-master-error")
                return err
            }
            tmp := string(tmpBuffer)
            tmpArgv := strings.Split(tmp, "\r\n")
            fmt.Println("======", tmpArgv)
            robj := createRedisObj(string(tmpArgv[1]))
            ServerInstance.Dict["name"] = robj
        }
        return nil
    }(masterAddr, masterPort, conn)
    conn.Write([]byte("$15\r\nslave-of-master\r\n"))
    return nil
}

func SlaveCommand(client Client) error {
    conn := client.Conn
    ServerInstance.SlaveList = append(ServerInstance.SlaveList, client)
    conn.Write([]byte("$11\r\nsendtoslave\r\n"))
    return nil
}

func HSetCommand(client Client) error {
    conn := client.Conn
    argv  := client.CommandArgv
    err := checkCommandProtocol(&client)
    if err != nil {
        return err
    }
    key := argv[4]
    field := argv[6]
    value := argv[8]
    
    var curMap map[string]string
    tmpMap, ok := ServerInstance.Dict[key]
    if (!ok) {
        curMap = make(map[string]string)
    } else {
        curMap, _ = tmpMap.Value.(map[string]string)
    }
    curMap[field] = value
    ServerInstance.Dict[key] = createRedisObj(curMap)
    responseOK(conn)
    return nil
}

func HGetCommand(client Client) error {
    conn := client.Conn
    argv  := client.CommandArgv
    err := checkCommandProtocol(&client)
    if err != nil {
        return err
    }
    key := argv[4]
    field := argv[6]

    tmpMap, ok := ServerInstance.Dict[key]
    if (!ok) {
        responseNil(conn)
    } else {
        tmpObj, tok := tmpMap.Value.(map[string]string)
        tmpValue := tmpObj[field]
        if tok {
            responseValue(tmpValue, conn)
        } else {
            responseNil(conn)
        }
    }
    return nil
}

func Progate() error {
    for _, client := range ServerInstance.SlaveList {
        conn := client.Conn
        conn.Write([]byte("$15\r\nget-from-master\r\n"))
    }
    return nil
}

func createRedisObj(v interface{}) RedisObj {
    var obj RedisObj
    obj.Value = v
    return obj
}

func responseOK(conn net.Conn) error {
    conn.Write([]byte("$2\r\nok\r\n"))
    return nil
}

func responseNil(conn net.Conn) error {
    conn.Write([]byte("$3\r\nnil\r\n"))
    return nil
}

func responseValue(str string, conn net.Conn) error {
    conn.Write([]byte("$" + strconv.Itoa(len(str)) + "\r\n" + str+ "\r\n"))
    return nil
}

var ServerInstance Server
var CommandList = [...]Command{
    {Name:"get",  Argc:1, Proc:GetCommand},
    {Name:"set",  Argc:2, Proc:SetCommand, NeedProgate: 1},
    {Name:"hget",  Argc:2, Proc:HGetCommand},
    {Name:"hset",  Argc:3, Proc:HSetCommand, NeedProgate: 1},
    {Name:"lpush",Argc:2, Proc:LpushCommand, NeedProgate: 1},
    {Name:"lpop", Argc:1, Proc:LpopCommand},
    {Name:"slaveof", Argc:2, Proc:SlaveofCommand},
    {Name:"slave", Argc:0, Proc:SlaveCommand},
}
