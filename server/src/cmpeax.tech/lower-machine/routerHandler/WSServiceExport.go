package routerHandler

import (
	"encoding/json"
	"fmt"
	"net"

	"cmpeax.tech/lower-machine/lib/DataParser"
	"cmpeax.tech/lower-machine/lib/routerDI"
	"cmpeax.tech/lower-machine/struct/ACS"
)

func WSServiceExport() routerDI.MapOfWSCallbackJSONFunc {
	return routerDI.MapOfWSCallbackJSONFunc{

		"0x03": func(jsonData routerDI.Message, conn net.Conn) {

			m := ACS.NewACS0x03("123456")
			m.Code = "0x03"

			jsonBytes, err := json.Marshal(m)
			if err != nil {
				fmt.Println(err)
			}
			conn.Write([]byte(DataParser.ParserToGbk(string(jsonBytes))))

			fmt.Println("0x03调e用")
		},
		"0x04": func(jsonData routerDI.Message, conn net.Conn) {

			m := ACS.NewACS0x04("123456")
			m.Code = "0x04"

			jsonBytes, err := json.Marshal(m)
			if err != nil {
				fmt.Println(err)
			}
			conn.Write([]byte(DataParser.ParserToGbk(string(jsonBytes))))

			fmt.Println("0x04e调用")
		},
	}
}
