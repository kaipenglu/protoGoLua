package main

import (
    "./net/tcplib"
    "./net/codec"
    "./pb/pbcpl"
    "github.com/aarzilli/golua/lua"
    "./utils"
    "path"
    "./hdl"
    "runtime"
    "os"
    "strconv"
)

type Job struct {
    msgId uint32
    msg interface{}
    c *tcplib.ComConn
}

func worker(workId int, jobs <-chan *Job) {
    L := lua.NewState()
    L.OpenLibs()
    defer L.Close()

    exePath := utils.ExecFilePath()
    luaPath := path.Join(exePath, "lua")
    utils.AddPkgPath(L, luaPath)
    if err := L.DoFile(path.Join(luaPath, "main.lua")); err != nil {
        panic(err)
    }

    for j := range jobs {
        var res interface{}

        if j.msgId == 1 {
            res = hdl.Login(L, j.msg)
        }

        c := j.c
        c.Send(res)
    }
}

func main() {
	cdc := codec.NewCodec()
	cdc.RegisterProto(1, &pbcpl.CSLogin{})

    routineNum, err := strconv.Atoi(os.Args[1])
    if err != nil {
        routineNum = runtime.NumCPU()
    }

    jobs := make(chan *Job, 100)
    for i := 0; i < routineNum; i++ {
        go worker(i, jobs)
    }

	f := func(msgId uint32, msg, client interface{}) {
        c := client.(*tcplib.ComConn)
        job := &Job{msgId, msg, c}
        jobs <- job
	}

	cdc.RegisterHandle(f)

	server := tcplib.NewTcpServer("", "3456", cdc)
	server.Start()

	client := tcplib.NewTcpClient("127.0.0.1", "3456", cdc)
	client.Start()

	p := &pbcpl.CSLogin{}
	cmd := pbcpl.PACKET_ID_PACKET_CS_LOGIN_CSLogin
	p.Cmd = &cmd
	accountName := "jack"
	p.AccountName = &accountName
	password := "123456"
	p.Password = &password
	client.Send(p)
	
	endRunning := make(chan bool, 1)
	<-endRunning
}
