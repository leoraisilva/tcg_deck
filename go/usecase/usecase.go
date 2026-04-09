package usecase

import (
	"tcg_deck/go/model"
	"tcg_deck/go/repository"
)

type Usecase struct {
	r repository.Repository
}

func NewUsecase(r repository.Repository) Usecase {
	return Usecase{r: r}
}

func (u *Usecase) CreateDeck(response model.Response) (int, error) {
	return u.r.CreateDeck(response)
}

func (u *Usecase) GetDeck(id int) (model.Response, error) {
	return u.r.GetDeck(id)
}
