package server

import (
    "net"
    "strconv"
    "errors"
    "fmt"
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
}

type CommandProc func(client Client) error

type Command struct {
    Name string
    Argc int
    Proc CommandProc 
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
    conn.Write([]byte("$10\r\nsetcommand\r\n"))
    err := checkCommandProtocol(&client)
    if err != nil {
        return errors.New("command-error")
    }
    ServerInstance.Dict[argv[4]] = RedisObj{argv[6]}
    return nil
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

var ServerInstance Server
var CommandList = [2]Command{{"get", 1, GetCommand},{"set", 2, SetCommand}}
