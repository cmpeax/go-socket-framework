package main

import (
	"database/sql"
	"fmt"

	"cmpeax.tech/lower-machine/lib/WSService"

	"cmpeax.tech/lower-machine/lib/SService"

	"cmpeax.tech/lower-machine/lib/routerList"
	"cmpeax.tech/lower-machine/routerHandler"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	container := routerList.BuildRouterList()
	container.PushBox(routerHandler.ServiceExport())

	wscontainer := routerList.BuildWSRouterList()
	wscontainer.PushBox(routerHandler.WSServiceExport())

	fmt.Printf("start Server")
	db, err := sql.Open("mysql", "root:test123456@tcp(127.0.0.1:3306)/upperMonitorF?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	addr := ":3972"
	wsaddr := ":5000"
	sS := SService.NewSService(addr, container.Get(), db) //表示监听本地所有ip的8080端口，也可以这样写：addr := ":8080"
	wsS := WSService.NewWSService(wsaddr, wscontainer.Get(), db, sS.PDevices())
	go wsS.StartWService()
	sS.StartService()
}
