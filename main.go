package main

import (
	"assigment-2/config"
	"assigment-2/routes"
)

func main() {
	config.StartDB()
	routes.Routes()
}
