package main

import "log"

func main() {
	var connErr = Connect()
	if connErr != nil {
		log.Fatal(connErr)
	}
	srv := GetServer()
	srv.Run("localhost:8081")
}
