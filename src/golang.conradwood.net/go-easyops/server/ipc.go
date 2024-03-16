package server

import (
	"fmt"
	"golang.conradwood.net/go-easyops/cmdline"
	"golang.yacloud.eu/unixipc"
	"strconv"
	"sync"
)

var (
	ipc_fd_env  = cmdline.ENV("GE_AUTODEPLOYER_IPC_FD", "if set it is assumed to be a filedescriptor over which an IPC can be initiated with the autodeployer")
	ipc_lock    sync.Mutex
	ipc_started = false
	srv         *unixipc.IPCServer
)

func start_ipc() {
	ipc_lock.Lock()
	defer ipc_lock.Unlock()
	if ipc_started {
		return
	}
	ipc_started = true
	if ipc_fd_env.Value() == "" {
		//fmt.Printf("[go-easyops] no ipc fd\n")
		return
	}
	fd, err := strconv.Atoi(ipc_fd_env.Value())
	if err != nil {
		panic(fmt.Sprintf("GE_AUTODEPLOYER_IPC_FD invalid value: %s", err))
	}
	srv, err = unixipc.NewConnectedServer(fd)
	if err != nil {
		panic(fmt.Sprintf("failed to start autodeployer IPC: %s", err))
	}
	srv.Send("startup", nil)
	//fmt.Printf("[go-easyops] Started IPC Server\n")
}
