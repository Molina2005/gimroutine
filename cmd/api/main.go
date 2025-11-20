package main

import (
	"fmt"
	"log"
	"modulo/internal/database"
)

func main() {
	connect, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer connect.Close()
	fmt.Println("conexion exitosa a postgres")
}
