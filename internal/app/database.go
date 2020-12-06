package app

import (
	"bifrost/gen/entmodels"
	"database/sql"
	"log"

	// load db driver
	_ "go.elastic.co/apm/module/apmsql/mysql"
	_ "go.elastic.co/apm/module/apmsql/sqlite3"

	entSQL "github.com/facebook/ent/dialect/sql"
)


// EntClient 全局数据库client
var EntClient *entmodels.Client

//InitDatabase 初始化数据库连接
func InitDatabase()  {
	var db *sql.DB
	var err error
	db, err = sql.Open("mysql", "root:mysql123@tcp(127.0.0.1:3306)/bifrost")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	driver := entSQL.OpenDB("mysql", db)
	EntClient = entmodels.NewClient(entmodels.Driver(driver))
}
