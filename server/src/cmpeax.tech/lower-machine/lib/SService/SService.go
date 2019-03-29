package SService

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"strings"

	"cmpeax.tech/lower-machine/lib/DataParser"
	"cmpeax.tech/lower-machine/lib/routerDI"
)

type SService struct {
	addr        string
	routerList  routerDI.MapOfCallbackJSONFunc
	wrouterList routerDI.MapOfWSCallbackJSONFunc
	db          *sql.DB
	devices     map[IPAddress]*DeviceConn
	RequestOfWS chan routerDI.Message
}

type DeviceConn struct {
	deviceID     string
	ipAddress    string
	conn         *net.Conn
	isConnecting string
}

type IPAddress string

type SDevicesPoint *map[IPAddress]DeviceConn

func NewSService(addr string, routerList routerDI.MapOfCallbackJSONFunc, wrouterList routerDI.MapOfWSCallbackJSONFunc, dbobj *sql.DB) *SService {
	return &SService{
		addr:        addr,
		routerList:  routerList,
		wrouterList: wrouterList,
		db:          dbobj,
		devices:     map[IPAddress]*DeviceConn{},
		RequestOfWS: make(chan routerDI.Message),
	}

}

func (s *SService) Devices() map[IPAddress]*DeviceConn {
	return s.devices
}

func (s *SService) PDevices() *map[IPAddress]*DeviceConn {
	return &s.devices
}

func (s *SService) StartService() {
	listener, err := net.Listen("tcp", s.addr)

	if err != nil {
		log.Fatal(err)
	}

	go s.HandleWebSocketTransmit()

	defer listener.Close() //关闭监听的端口
	for {
		tconn, err := listener.Accept() //用conn接收链接
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("connect:" + tconn.RemoteAddr().String())
		var tempIP IPAddress = IPAddress(strings.Split(tconn.RemoteAddr().String(), ":")[0])
		if s.Devices()[tempIP] != nil {

		}
		s.Devices()[tempIP] = &DeviceConn{
			ipAddress:    string(tempIP),
			conn:         &tconn,
			isConnecting: "true",
		}

		//测试输出
		fmt.Println("新设备")
		for key, value := range s.devices {
			fmt.Println("IP:", key, "DATA:", value)
		}

		go handleConnection(tconn, s)

	}
}

func (s *SService) HandleWebSocketTransmit() {

	fmt.Println("开始监听")
	for {
		getMsg := <-s.RequestOfWS

		s.StartWSMatchJson(getMsg)

	}
}

func (s *SService) StartWSMatchJson(msg routerDI.Message) {
	wsfrouterDI := routerDI.InitWSRouter(s.wrouterList, s.db)

	devicePoint := s.Devices()
	tempMsgIp := IPAddress(msg.Ip)
	fmt.Println("得到的地址" + msg.Ip)

	for item := range devicePoint {
		fmt.Println("\n item:" + item)
	}
	if devicePoint[tempMsgIp] != nil {
		for key, callback := range wsfrouterDI.WSGetRouterList() {
			if msg.Code == key {
				fmt.Printf("match: %s\n", key)
				callback(msg, *devicePoint[tempMsgIp].conn)
			}

		}
	}

}

func handleConnection(conn net.Conn, s *SService) {

	frouterDI := routerDI.InitRouter(&conn, s.routerList, s.db)
	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), "ERR!", err)
			return
		}

		fmt.Printf(DataParser.Parser(string(buffer[:n])), "<-getit")
		frouterDI.StartMatchJson(DataParser.Parser(string(buffer[:n])))
	}

}
