package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_users_schema   = "mysql_users_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	schema   = os.Getenv(mysql_users_schema)
)

// As soon as the package is imported, the function fires
func init() {
	// username:password@tcp(host)/schema?charset=utf8
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema)
	var err error

	// import https://github.com/go-sql-driver/mysql/
	Client, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err)
	}
	//mysql.SetLogger()
	log.Println("Database loaded successfully")
}
