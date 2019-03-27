package routerList

import "cmpeax.tech/lower-machine/lib/routerDI"

type RouterList struct {
	list map[string]routerDI.CallbackJSONFunc
}

func BuildRouterList() *RouterList {
	return &RouterList{
		list: map[string]routerDI.CallbackJSONFunc{},
	}
}

//导入list表
func (l *RouterList) Push(matchStr string, callback routerDI.CallbackJSONFunc) {
	l.Get()[matchStr] = callback
}

func (l *RouterList) PushBox(pushList map[string]routerDI.CallbackJSONFunc) {
	for key, callback := range pushList {
		l.list[key] = callback
	}
}
func (l *RouterList) Get() map[string]routerDI.CallbackJSONFunc {
	return l.list
}
