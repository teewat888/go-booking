package main

import (
	"fmt"

	"github.com/teewat888/go-booking/boilerplate/internal/config"
	"github.com/teewat888/go-booking/boilerplate/internal/dependencies"
)

func main() {
	fmt.Println("Hello, World!")

	cfg := config.FromEnv()

	deps := dependencies.InitDependencies(&cfg)

	fmt.Printf("%+v", deps)
}
