package main

import (
	"github.com/usmaarn/blogg_api/cmd/config"
	"github.com/usmaarn/blogg_api/internal/router"
)

func main() {
	config.Load()
	router.Initialize(":8080")
}
