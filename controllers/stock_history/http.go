package stockhistory

import (
	stockhistory "backend-golang/businesses/stock_history"

	"backend-golang/controllers"
	"backend-golang/controllers/stock_history/request"
	"backend-golang/controllers/stock_history/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StockHistoryController struct {
	stockHistoryUseCase stockhistory.Usecase
}

func NewStockHistoryController(authUC stockhistory.Usecase) *StockHistoryController {
	return &StockHistoryController{
		stockHistoryUseCase: authUC,
	}
}

func (shc *StockHistoryController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	stockID := c.Param("id")

	stock, err := shc.stockHistoryUseCase.GetByID(ctx, stockID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, http.StatusNotFound, true, "category not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, http.StatusOK, false, "category found", response.FromDomain(stock))
}

func (shc *StockHistoryController) Create(c echo.Context) error {
	input := request.StockHistory{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, http.StatusBadRequest, true, "invalid request", "")
	}

	stock, err := shc.stockHistoryUseCase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "failed to add a stock", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, http.StatusCreated, false, "stock registered", response.FromDomain(stock))
}
