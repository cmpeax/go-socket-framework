package routerDI

import (
	"database/sql"
	"fmt"
	"net"
)

type CallbackFunc func(conn net.Conn, db *sql.DB)

type Router struct {
	conn      net.Conn
	db        *sql.DB
	matchList map[string]CallbackFunc
}

func InitRouter(dconn net.Conn, matchList map[string]CallbackFunc, dbobj *sql.DB) *Router {
	return &Router{
		conn:      dconn,
		matchList: matchList,
		db:        dbobj,
	}
}

func (r *Router) StartMatch(matchCodeStr string) {
	for key, callback := range r.matchList {
		if matchCodeStr == key {
			fmt.Printf("match: %s", key)
			callback(r.conn, r.db)
		}
	}
}
