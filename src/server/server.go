package server

import (
    "net"
    "strconv"
    "errors"
    "fmt"
    "container/list"
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
}

type CommandProc func(client Client) error

type Command struct {
    Name string
    Argc int
    Proc CommandProc 
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
        return errors.New("command-error")
    }
    resp := ServerInstance.Dict[argv[4]]

    //类型判定
    tmpValue, ok := (resp.Value).(string)
    fmt.Println(tmpValue, ok)
    if !ok {
        conn.Write([]byte("$12\r\ninvalid-type\r\n"))
    } else {
        conn.Write([]byte("$" + strconv.Itoa(len(tmpValue)) + "\r\n" + tmpValue + "\r\n"))
    }
    return nil
}

func SetCommand(client Client) error{
    conn := client.Conn
    argv := client.CommandArgv
    err := checkCommandProtocol(&client)
    if err != nil {
        return errors.New("command-error")
    }
    conn.Write([]byte("$2\r\nok\r\n"))
    ServerInstance.Dict[argv[4]] = RedisObj{argv[6]}
    return nil
}

func LpushCommand(client Client) error {
    conn := client.Conn
    argv := client.CommandArgv
    err := checkCommandProtocol(&client)
    if err != nil {
        return errors.New("command-error")
    }
    if _, ok := ServerInstance.Dict[argv[4]]; !ok {
        ls := list.New()
        ServerInstance.Dict[argv[4]] = RedisObj{ls}
    }
    tmpls := ServerInstance.Dict[argv[4]].Value.(*list.List)
    tmpls.PushBack(argv[6])
    conn.Write([]byte("$2\r\nok\r\n"))
    return nil
}

func LpopCommand(client Client) error {
    conn := client.Conn
    argv := client.CommandArgv
    err := checkCommandProtocol(&client)
    if err != nil {
        return errors.New("command-error")
    }
    if _, ok := ServerInstance.Dict[argv[4]]; !ok {
        conn.Write([]byte("$10\r\nempty-list\r\n"))
        return nil
    }
    ls := ServerInstance.Dict[argv[4]].Value.(*list.List)
    tmpValue := ls.Back().Value.(string)
    conn.Write([]byte("$" + strconv.Itoa(len(tmpValue)) + "\r\n" + tmpValue + "\r\n"))
    return nil
}

var ServerInstance Server
var CommandList = [4]Command{{"get", 1, GetCommand},{"set", 2, SetCommand},{"lpush", 2, LpushCommand},{"lpop", 1, LpopCommand}}
