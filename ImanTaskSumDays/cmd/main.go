package main

import (
	"sum_days/config"
	"sum_days/src/handler"
	"fmt"
)

func main() {
	cfg := config.Load()
	r := handler.SetupRouter()
	r.Run(fmt.Sprintf(":%s", cfg.HTTPPort))
}