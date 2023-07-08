package main

import (
	"Projects/go-url-shortener/cmd/internal/config"
	"fmt"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}