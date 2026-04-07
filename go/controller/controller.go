package controller

import "tcg_deck/go/usecase"

type Controller struct {
	u usecase.Usecase
}

func NewController(u usecase.Usecase) Controller {
	return Controller{u: u}
}
