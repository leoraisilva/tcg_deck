package usecase

import "tcg_deck/go/repository"

type Usecase struct {
	r repository.Repository
}

func NewUsecase(r repository.Repository) Usecase {
	return Usecase{r: r}
}
