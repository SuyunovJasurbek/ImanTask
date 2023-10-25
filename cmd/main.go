package main

import (
	"days_limit/config"
	"days_limit/src/handler"
	"fmt"
)

func main() {
	cfg := config.Load()
	r := handler.SetupRouter()
	r.Run(fmt.Sprintf(":%s", cfg.HTTPPort))
}