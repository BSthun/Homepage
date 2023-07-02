package main

import (
	"backend/loaders/fiber"
	"backend/loaders/hub"
	"backend/loaders/mysql"
)

func main() {
	mysql.Init()
	hub.Init()
	fiber.Init()
}
