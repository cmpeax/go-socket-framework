package routerHandler

import (
	"fmt"

	"cmpeax.tech/lower-machine/lib/routerDI"
)

func WSServiceExport() routerDI.MapOfWSCallbackJSONFunc {
	return routerDI.MapOfWSCallbackJSONFunc{
		"0x01": func(jsonData routerDI.Message) {
			fmt.Println("0x01调用")
		},
		"0x02": func(jsonData routerDI.Message) {
			fmt.Println("0x02调用")
		},
		"0x03": func(jsonData routerDI.Message) {
			fmt.Println("0x03调用")
		},
		"0x04": func(jsonData routerDI.Message) {
			fmt.Println("0x04调用")
		},
		"0x05": func(jsonData routerDI.Message) {
			fmt.Println("0x05调用")
		},
		"0x06": func(jsonData routerDI.Message) {
			fmt.Println("0x06调用")
		},
		"0x07": func(jsonData routerDI.Message) {
			fmt.Println("0x07调用")
		},
	}
}
