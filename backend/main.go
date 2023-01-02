package main

import (
	"math/rand"
	"time"

	"backend/loaders/fiber"
	"backend/loaders/hub"
	"backend/loaders/mysql"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	hub.Init()
	mysql.Init()
	fiber.Init()
}
