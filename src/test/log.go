package main

import (
	"log"
	"os"
)

func main() {
	//2017/11/25 13:10:47 log-Fatal
	//exit status 1
	//log.Fatal("log-Fatal")

	//log.Fatalln("log-Fatalln")

	//log.Panic("log-Panic")

	//2017/11/25 13:16:30 log-Print
	log.Print("log-Print")

	//Warnning2017/11/25 13:19:50 log-Print
	log.SetPrefix("Warnning")
	log.Print("log-Print")

	//Warnning13:30:08 /Users/xxx/rediskin/src/test/log.go:24: log-Print
	log.SetFlags(log.Llongfile | log.Ltime)
	log.Print("log-Print")

	logFile, _ := os.Create("./logfile")
	defer logFile.Close()
	//log.SetOutput(logFile)
	//log.Print("log-print") //写入到了文件中
	//log.Println("log-print")//写入到了文件中

	logger := log.New(logFile, "[prefix]", log.Llongfile|log.LstdFlags)
	//[prefix]2017/11/25 13:37:55 /Users/didi/togoproject/rediskin/src/test/log.go:35: A debug message here
	logger.Println("A debug message here")
	//[Info]2017/11/25 13:39:07 /Users/didi/togoproject/rediskin/src/test/log.go:38: A debug message here
	logger.SetPrefix("[Info]")
	logger.Println("A debug message here")
}
