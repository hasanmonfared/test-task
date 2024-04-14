package mysqluser

import "app/adapter/mysql"

type DB struct {
	adapter mysql.Adapter
}

func New(adapter mysql.Adapter) DB {
	return DB{adapter: adapter}
}
