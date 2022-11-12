package main

import (
	"FP2/router"
	"os"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := router.StartAPP()
	PORT := os.Getenv("SERVICE_PORT")
	if PORT == "" {
		PORT = "9000"
	}
	err := r.Run()
	if err != nil {
		log.Println("Error Graceful", err)
		os.Exit(1)
	}
}
