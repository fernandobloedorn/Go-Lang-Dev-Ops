package main

import (
	"github.com/fernandobloedorn/Go-Lang-Dev-Ops/config"
	"github.com/fernandobloedorn/Go-Lang-Dev-Ops/router"
)

var (
	logger config.Logger
)

func main() {

	logger := config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		// panic(err)
		return
	}

	router.Initialize()
}
