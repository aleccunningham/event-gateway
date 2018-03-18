package controller

import (
	"log"

	"github.com/spotahome/kooper/operator/controller"
)

type (
	Controller struct {
		controller.Controller

		config Config
		logger log.Logger
	}
)
