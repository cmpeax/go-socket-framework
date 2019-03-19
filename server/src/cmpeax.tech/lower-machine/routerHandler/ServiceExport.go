package routerHandler

import (
	"database/sql"
	"fmt"
	"net"

	"cmpeax.tech/lower-machine/lib/routerDI"
)

func ServiceExport() map[string]routerDI.CallbackFunc {
	return map[string]routerDI.CallbackFunc{
		"abc": func(conn net.Conn, db *sql.DB) {
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
		"def": func(conn net.Conn, db *sql.DB) {
			fmt.Println("def")
		},
	}
}
