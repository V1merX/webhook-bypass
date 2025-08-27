package main

import (
	"os"

	"github.com/V1merX/webhook-bypass/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		os.Exit(1)
	}
}
