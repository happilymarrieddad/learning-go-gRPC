package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	api "github.com/happilymarrieddad/learning-go-gRPC/api"
)

func main() {
	db, err := xorm.NewEngine("mysql", "root:pass@tcp(localhost:3306)/grpc")
	if err != nil {
		panic(err)
	}

	_, err = db.DBMetas()
	if err != nil {
		panic(err)
	}

	api.Run(50051, db)
}
