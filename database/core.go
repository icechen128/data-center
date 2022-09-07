package database

import (
	"data-center/database/common"
	"data-center/database/mysql"
	"sync"
)

type Manager struct {
	dbs map[string]common.Database
}

var manager Manager

var once sync.Once
var lock sync.RWMutex

func GetManager() Manager {
	once.Do(func() {
		manager.dbs = make(map[string]common.Database)
	})
	return manager
}

func (m *Manager) AddDB(driverName string, host, port, user, password, dbName string) {
	var db common.Database
	switch driverName {
	case "mysql":
		db = mysql.New(host, port, user, password, dbName)
	default:
		panic("unknown driver name")
	}

	lock.Lock()
	defer lock.Unlock()
	manager.dbs[dbName] = db
}
