package stockins

import (
	stockins "backend-golang/businesses/stock_ins"
	"backend-golang/controllers"
	"backend-golang/controllers/stock_ins/request"
	"backend-golang/controllers/stock_ins/response"

	// _reqStocks "backend-golang/controllers/stocks/request"
	// _resStocks "backend-golang/controllers/stocks/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StockInController struct {
	stockUsecase stockins.Usecase
}

func NewStockInController(courseUC stockins.Usecase) *StockInController {
	return &StockInController{
		stockUsecase: courseUC,
	}
}

func (sc *StockInController) Create(c echo.Context) error {
	input := request.StockIn{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, http.StatusBadRequest, true, "invalid request", "")
	}

	stockIn, err := sc.stockUsecase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "failed to add a stock", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, http.StatusCreated, false, "stock registered", response.FromDomain(stockIn))
}

func (cc *StockInController) StockIn(c echo.Context) error {
	input := request.StockIn{}
	ctx := c.Request().Context()
	// token := c.Get("user").(*jwt.Token)

	// isListed := middlewares.CheckToken(token.Raw)

	// if !isListed {
	// 	return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "invalid token", isListed)
	// }

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, http.StatusBadRequest, true, "invalid request", "")
	}

	stockIn, err := cc.stockUsecase.StockIn(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "failed to add a stock", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, http.StatusCreated, false, "stock transaction created", response.FromDomain(stockIn))
}
