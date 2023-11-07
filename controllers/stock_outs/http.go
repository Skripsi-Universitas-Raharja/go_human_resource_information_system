package stockouts

import (
	stockouts "backend-golang/businesses/stock_outs"

	"backend-golang/controllers"
	"backend-golang/controllers/stock_outs/request"
	"backend-golang/controllers/stock_outs/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StockOutController struct {
	stockUseCase stockouts.Usecase
}

func NewStockOutController(authUC stockouts.Usecase) *StockOutController {
	return &StockOutController{
		stockUseCase: authUC,
	}
}

func (cc *StockOutController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	stockID := c.Param("id")

	stock, err := cc.stockUseCase.GetByID(ctx, stockID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, http.StatusNotFound, true, "category not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, http.StatusOK, false, "category found", response.FromDomain(stock))
}

func (sc *StockOutController) StockOut(c echo.Context) error {
	input := request.StockOut{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, http.StatusBadRequest, true, "invalid request", "")
	}

	stock, err := sc.stockUseCase.StockOut(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "failed to add a stock", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, http.StatusCreated, false, "stock registered", response.FromDomain(stock))
}
