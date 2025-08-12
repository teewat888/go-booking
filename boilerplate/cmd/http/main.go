package main

import (
	"fmt"

	"github.com/teewat888/go-booking/boilerplate/internal/config"
)

func main() {
	fmt.Println("Hello, World!")

	cfg := config.FromEnv()

	fmt.Printf("%+v", cfg)
}
