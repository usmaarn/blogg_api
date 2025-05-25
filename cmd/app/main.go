package main

import (
	"github.com/usmaarn/blogg_api/cmd/config"
	"github.com/usmaarn/blogg_api/internal/router"
	"github.com/usmaarn/blogg_api/package/validators"
)

func main() {
	config.Load()
	validators.InitializeValidator()
	router.Initialize(":8080")
}
