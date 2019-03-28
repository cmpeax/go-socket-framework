package routerHandler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net"

	"cmpeax.tech/lower-machine/lib/DataParser"
	"cmpeax.tech/lower-machine/lib/routerDI"
	"cmpeax.tech/lower-machine/struct/ACS"
)

func ServiceExport() routerDI.MapOfCallbackJSONFunc {
	return routerDI.MapOfCallbackJSONFunc{
		"0x11": func(jsondata map[string]interface{}, conn net.Conn, db *sql.DB) {
			fmt.Println("...")
			rows, err := db.Query("select employeeID,employeeName from validateList")
			if err != nil {
				fmt.Print(err)
				return
			}
			employeeID := ""
			employeeName := ""
			for rows.Next() {
				err := rows.Scan(&employeeID, &employeeName)
				if err != nil {
					fmt.Println(err)

				}

				fmt.Println(employeeID, ",", employeeName)

			}
		},
		"0x01": func(jsondata map[string]interface{}, conn net.Conn, db *sql.DB) {
			for key, value := range jsondata {
				fmt.Println("key:", key, "value:", value)
			}

			m := ACS.NewACS0x01(jsondata["token"].(string))
			m.Code = "0x01"
			m.IsPass = "PASS"
			m.PassData = ACS.AcsPassData{}

			jsonBytes, err := json.Marshal(m)
			if err != nil {
				fmt.Println(err)
			}

			conn.Write([]byte(DataParser.ParserToGbk(string(jsonBytes))))
		},
		"0x02": func(jsondata map[string]interface{}, conn net.Conn, db *sql.DB) {
			for key, value := range jsondata {
				fmt.Println("key:", key, "value:", value)
			}

			m := ACS.NewACS0x02(jsondata["token"].(string))
			m.Code = "0x02"
			m.Result = "OK"

			jsonBytes, err := json.Marshal(m)
			if err != nil {
				fmt.Println(err)
			}

			conn.Write([]byte(DataParser.ParserToGbk(string(jsonBytes))))
		},
		"0x03": func(jsondata map[string]interface{}, conn net.Conn, db *sql.DB) {
			for key, value := range jsondata {
				fmt.Println("key:", key, "value:", value)
			}

			m := ACS.NewACS0x02(jsondata["token"].(string))
			m.Code = "0x03"
			m.Result = "OK"

			jsonBytes, err := json.Marshal(m)
			if err != nil {
				fmt.Println(err)
			}

			conn.Write([]byte(DataParser.ParserToGbk(string(jsonBytes))))
		},
		"0x05": func(jsondata map[string]interface{}, conn net.Conn, db *sql.DB) {
			for key, value := range jsondata {
				fmt.Println("key:", key, "value:", value)
			}

			m := ACS.NewACS0x02(jsondata["token"].(string))
			m.Code = "0x05"
			m.Result = "OK"

			jsonBytes, err := json.Marshal(m)
			if err != nil {
				fmt.Println(err)
			}

			conn.Write([]byte(DataParser.ParserToGbk(string(jsonBytes))))
		},
	}
}
