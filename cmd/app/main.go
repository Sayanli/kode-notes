package main

import (
	"kode-notes/internal/app"
)

const configPath = "../../config/config.yaml"

func main() {
	app.Run(configPath)
}
