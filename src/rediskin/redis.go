package main

import(
    "fmt"
    "net"
    "runtime"
    "server"
    "strconv"
    "strings"
    "errors"
    "flag"
    "github.com/BurntSushi/toml"
    "io/ioutil"
    "os"
    "os/signal"
    "syscall"
    "webserver"
    "service"
)

func initServer() {
    fp, err := os.Open("../src/conf/conf.toml") 
    if err != nil {
        panic("failed to open conf file")
    }

    fcontent, err := ioutil.ReadAll(fp)
    if err != nil {
        panic("failed to read conf file")
    }

    err = toml.Unmarshal(fcontent, &server.ServerInstance)
    if err != nil {
        fmt.Println("toml.Unmarshal error ", err)
        panic("failed to parase conf file")
    }
    server.ServerInstance.Dict = make(map[string]server.RedisObj)
    server.ServerInstance.ExpireDict = make(map[string]int)
}

func createClient(conn net.Conn) server.Client {
    var client server.Client
    client.IP = conn.RemoteAddr().String()
    client.Conn = conn
    return client
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    signal.Notify(c, syscall.SIGTERM)
    go func() {
        <-c
        shutDown()
    } ()

    initServer()
    var port *int = flag.Int("port", 1024, "input port")
    flag.Parse()
    fmt.Println(*port)
    server.ServerInstance.Port = *port

    service.Wg.Add(1)
    go webserver.StartWebServer()

    service.Wg.Add(1)
    go startRedisServer()

    service.Wg.Wait()
}

func shutDown() {
    fmt.Println("ctrl-c or SIGTERM found.~bye~bye~")
    os.Exit(1)
}

func startRedisServer() {
    defer service.Wg.Done()
    netListen, err := net.Listen("tcp", server.ServerInstance.Addr + ":" + strconv.Itoa(server.ServerInstance.Port)) 
    if (err != nil) {
        netListen.Close()
    }
    defer netListen.Close()
    for {
        conn, err := netListen.Accept()
        if err != nil {
            continue
        }

        client := createClient(conn)
        server.ServerInstance.ClientList = append(server.ServerInstance.ClientList, client)

        Log(conn.RemoteAddr().String(), " tcp connect success")
        go handleClient(client)
    }
}

func handleClient(client server.Client) {
    conn := client.Conn
    buffer := make([]byte, 2048)
    for {
        err := readFromClient(conn, buffer)
        client.Inputbuffer = buffer
        if (err != nil) {
            break
        } else {
            processInputBuffer(client)
        }
    }
}

func processInputBuffer(client server.Client) {
    conn := client.Conn
    buffer := client.Inputbuffer
    inputString := string(buffer)
    tmpList := strings.Split(inputString, "\r\n")
    if (len(tmpList) < 3) {
        conn.Write([]byte("$12\r\ninvalid-data\r\n"))
    } else {
        commandName := tmpList[2]
        command, err := findCommand(commandName, tmpList)
        if err!= nil {
            conn.Write([]byte("$4\r\noooo\r\n"))
        } else {
            client.CurrentCommand = command 
            client.CommandArgv = tmpList
            processCommand(command, client)
        }
    }
}

func processCommand(command server.Command, client server.Client) {
    command.Proc(client)
    if command.NeedProgate == 1 {
        server.Progate()
    }
}

func findCommand(command string, tmpList []string) (server.Command, error){
    var scommand server.Command
    var findCommand bool
    for _, scommand = range server.CommandList {
        fmt.Println(scommand, command)
        if (scommand.Name == strings.TrimRight(command, "\r\n")) {
            findCommand = true
            break
        }
    }
    if (findCommand) {
        return scommand, nil
    } else {
        return scommand, errors.New("invalid-command")
    }
}

func readFromClient(conn net.Conn, buffer []byte) error{
    _, err := conn.Read(buffer)
    if err != nil {
        Log(conn.RemoteAddr().String(), " connection error: ", err)
        conn.Close();
        return err
    }
    return nil
}

func Log(v ...interface{}) {
    fmt.Println(v...)
}
