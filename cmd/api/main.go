package main

import (
	"backend/internal/bootstrap"
	"time"
)

const idleTimeout = 5 * time.Second

func main() {

	if err := bootstrap.Start(); err != nil {
		panic(err)
	}

}
