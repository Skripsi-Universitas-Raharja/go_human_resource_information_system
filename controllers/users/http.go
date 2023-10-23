package users

import (
	"backend-golang/businesses/users"
	"backend-golang/controllers"
	"backend-golang/controllers/users/request"
	"backend-golang/controllers/users/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authUseCase users.Usecase
}

func NewAuthController(authUC users.Usecase) *AuthController {
	return &AuthController{
		authUseCase: authUC,
	}
}

func (ctrl *AuthController) Register(c echo.Context) error {
	userInput := request.UserRegistration{}
	ctx := c.Request().Context()

	if err := c.Bind(&userInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, http.StatusBadRequest, true, "invalid request", "")
	}

	err := userInput.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, http.StatusBadRequest, true, "validation failed", "")
	}

	user, err := ctrl.authUseCase.Register(ctx, userInput.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "error when inserting data", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, http.StatusCreated, false, "user registered", response.FromDomain(user))
}

func (ctrl *AuthController) Login(c echo.Context) error {
	userInput := request.UserLogin{}
	ctx := c.Request().Context()

	if err := c.Bind(&userInput); err != nil {
		return controllers.NewResponseLogin(c, http.StatusBadRequest, http.StatusBadRequest, true, "invalid request", "")
	}

	err := userInput.Validate()

	if err != nil {
		return controllers.NewResponseLogin(c, http.StatusBadRequest, http.StatusBadRequest, true, "validation failed", "")
	}

	token, err := ctrl.authUseCase.Login(ctx, userInput.ToDomainLogin())

	var isFailed bool = err != nil || token == ""

	if isFailed {
		return controllers.NewResponseLogin(c, http.StatusUnauthorized, http.StatusUnauthorized, true, "invalid email or password", "")
	}

	return controllers.NewResponseLogin(c, http.StatusOK, http.StatusOK, false, "token created", token)
}
