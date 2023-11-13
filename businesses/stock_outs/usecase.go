package stockouts

import (
	"backend-golang/app/middlewares"
	"context"
)

type stockoutUsecase struct {
	stockRepository Repository
	jwtAuth         *middlewares.JWTConfig
}

func NewStockOutUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &stockoutUsecase{
		stockRepository: repository,
		jwtAuth:         jwtAuth,
	}
}

func (usecase *stockoutUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.stockRepository.GetAll(ctx)
}
func (usecase *stockoutUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.stockRepository.GetByID(ctx, id)
}

func (usecase *stockoutUsecase) StockOut(ctx context.Context, stockDomain *Domain) (Domain, error) {
	return usecase.stockRepository.StockOut(ctx, stockDomain)
}

func (usecase *stockoutUsecase) ExportToExcel(ctx context.Context) ([]Domain, error) {
	return usecase.stockRepository.ExportToExcel(ctx)
}
