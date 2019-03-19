package mysqldb

import (
	"database/sql"
	"fmt"
	"sync"
)

type Mysqldb struct {
	db *sql.DB
}

var mysqlObject *Mysqldb
var once sync.Once

func GetInstance() *Mysqldb {
	once.Do(func() {
		mysqlObject = &Mysqldb{}
	})
	return mysqlObject
}

func (mdb *Mysqldb) OpenMysqlDateBase() *sql.DB {
	var err error
	mdb.db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/upperMonitorF?charset=utf8")

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return mdb.db
}

func (mdb *Mysqldb) CloseDB() {
	mdb.db.Close()
}
