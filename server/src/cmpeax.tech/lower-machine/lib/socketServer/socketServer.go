package socketServer

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"cmpeax.tech/lower-machine/lib/DataParser"
	"cmpeax.tech/lower-machine/lib/routerDI"
)

type Socket struct {
	addr       string
	routerList map[string]routerDI.CallbackFunc
	db         *sql.DB
}

func NewSocket(addr string, routerList map[string]routerDI.CallbackFunc, dbobj *sql.DB) *Socket {
	return &Socket{
		addr:       addr,
		routerList: routerList,
		db:         dbobj,
	}
}

func (s *Socket) StartSocket() {
	listener, err := net.Listen("tcp", s.addr)

	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close() //关闭监听的端口
	for {
		conn, err := listener.Accept() //用conn接收链接
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("connect:" + conn.RemoteAddr().String())
		go handleConnection(conn, s)
	}
}

func handleConnection(conn net.Conn, s *Socket) {
	frouterDI := routerDI.InitRouter(conn, s.routerList, s.db)
	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), "ERR!", err)
			return
		}

		frouterDI.StartMatch(DataParser.Parser(string(buffer[:n])))
	}
}
