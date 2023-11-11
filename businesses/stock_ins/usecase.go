package stockins

import (
	"backend-golang/app/middlewares"
	// "backend-golang/businesses/stocks"

	"context"
)

type stockInUsecase struct {
	stockInRepository Repository
	jwtAuth           *middlewares.JWTConfig
}

func NewStockInUsecase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &stockInUsecase{
		stockInRepository: repository,
		jwtAuth:           jwtAuth,
	}
}

// func (usecase *stockInUsecase) GetAll(ctx context.Context) ([]Domain, error) {
// 	return usecase.stockInRepository.GetAll(ctx)
// }

func (usecase *stockInUsecase) Create(ctx context.Context, stockInDomain *Domain) (Domain, error) {
	return usecase.stockInRepository.Create(ctx, stockInDomain)
}

func (usecase *stockInUsecase) StockIn(ctx context.Context, stockInDomain *Domain) (Domain, error) {
	return usecase.stockInRepository.StockIn(ctx, stockInDomain)
}
