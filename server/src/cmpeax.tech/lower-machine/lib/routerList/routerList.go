package routerList

import (
	"cmpeax.tech/lower-machine/lib/routerDI"
)

type RouterList struct {
	list routerDI.MapOfCallbackJSONFunc
}

type WSRouterList struct {
	list routerDI.MapOfWSCallbackJSONFunc
}

func BuildRouterList() *RouterList {
	return &RouterList{
		list: routerDI.MapOfCallbackJSONFunc{},
	}
}

func BuildWSRouterList() *WSRouterList {
	return &WSRouterList{
		list: routerDI.MapOfWSCallbackJSONFunc{},
	}
}

//导入list表
func (l *RouterList) Push(matchStr string, callback routerDI.CallbackJSONFunc) {
	l.Get()[matchStr] = callback
}

func (l *RouterList) PushBox(pushList routerDI.MapOfCallbackJSONFunc) {
	for key, callback := range pushList {
		l.list[key] = callback
	}
}
func (l *RouterList) Get() routerDI.MapOfCallbackJSONFunc {
	return l.list
}

//WS:导入list表
func (l *WSRouterList) Push(matchStr string, callback routerDI.WSCallbackJSONFunc) {
	l.Get()[matchStr] = callback
}

func (l *WSRouterList) PushBox(pushList routerDI.MapOfWSCallbackJSONFunc) {
	for key, callback := range pushList {
		l.list[key] = callback
	}
}
func (l *WSRouterList) Get() routerDI.MapOfWSCallbackJSONFunc {
	return l.list
}
