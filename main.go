package main

import (
	"database/sql"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

const(
	host ="localhost"
	port =0000
	user = "postgres"
	password = " "
	dbname = "driverdb"
)

func main(){
	pqsqlconn := fmt.Sprintf("host = %s port = %d password = %s dbname= %s sslmode-disable",host,port,password,dbname)

	db, err := sql.Open()

}
