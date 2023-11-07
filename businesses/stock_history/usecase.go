package stockhistory

import (
	"backend-golang/app/middlewares"
	"context"
)

type stockHistoryUsecase struct {
	stockRepository Repository
	jwtAuth         *middlewares.JWTConfig
}

func NewStockHistoryUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &stockHistoryUsecase{
		stockRepository: repository,
		jwtAuth:         jwtAuth,
	}
}

func (usecase *stockHistoryUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.stockRepository.GetByID(ctx, id)
}

func (usecase *stockHistoryUsecase) Create(ctx context.Context, stockDomain *Domain) (Domain, error) {
	return usecase.stockRepository.Create(ctx, stockDomain)
}
