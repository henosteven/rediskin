package server

import (
    "net"
    "strconv"
    "errors"
)

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
    Dict map[string]string
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
    argvFlag := argv[0][0:1]
    argvCount, _:= strconv.Atoi(argv[0][1:2])
    if (argvFlag != "*") {
        conn.Write([]byte("$21\r\ninvalid-argument-flag\r\n"))
        return errors.New("invliad-aggument-flag")
    }
    if(argvCount != 2) {
        conn.Write([]byte("$22\r\ninvalid-argument-count\r\n"))
        return errors.New("invliad-aggument-count")
    } else {
        resp := ServerInstance.Dict[argv[4]]
        conn.Write([]byte("$" + strconv.Itoa(len(resp)) + "\r\n" + resp + "\r\n"))
        return nil
    }
}

func SetCommand(client Client) error{
    conn := client.Conn
    conn.Write([]byte("$10\r\nsetcommand\r\n"))
    argv := client.CommandArgv
    ServerInstance.Dict[argv[4]] = argv[6]
    return nil
}

var ServerInstance Server
var CommandList = [2]Command{{"get", 1, GetCommand},{"set", 2, SetCommand}}
