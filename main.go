package main

import (
	"FP2/router"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := router.StartAPP()
	PORT := os.Getenv("MYSQLPORT")
	r.Run(":" + PORT)
}
