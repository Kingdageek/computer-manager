package controllers

import "computer-manager/internal/services"

type Controllers struct {
	Computer *ComputerController
}

func NewControllers(svcs *services.Services) *Controllers {
	return &Controllers{
		Computer: NewComputerController(svcs),
	}
}
