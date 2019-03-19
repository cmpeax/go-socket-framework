package routerList

import "cmpeax.tech/lower-machine/lib/routerDI"

type RouterList struct {
	list map[string]routerDI.CallbackFunc
}

func BuildRouterList() *RouterList {
	return &RouterList{
		list: map[string]routerDI.CallbackFunc{},
	}
}

//导入list表
func (l *RouterList) Push(matchStr string, callback routerDI.CallbackFunc) {
	l.Get()[matchStr] = callback
}

func (l *RouterList) PushBox(pushList map[string]routerDI.CallbackFunc) {
	for key, callback := range pushList {
		l.list[key] = callback
	}
}
func (l *RouterList) Get() map[string]routerDI.CallbackFunc {
	return l.list
}
