package users

import (
	"backend-golang/app/middlewares"
	"backend-golang/businesses/users"

	"backend-golang/controllers"
	"backend-golang/controllers/users/request"
	"backend-golang/controllers/users/response"

	"net/http"

	"github.com/golang-jwt/jwt/v4"
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

func (ctrl *AuthController) Logout(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)

	isListed := middlewares.CheckToken(token.Raw)

	if !isListed {
		return controllers.NewResponse(c, http.StatusUnauthorized, http.StatusUnauthorized, true, "invalid token", isListed)
	}
	// Invalidate the token by removing it from the whitelist
	isLoggedOut := middlewares.Logout(token.Raw)

	return controllers.NewResponse(c, http.StatusOK, http.StatusOK, false, "logout success", isLoggedOut)
}

// func (ctrl *AuthController) UpdateProfileUser(c echo.Context) error {
// 	var userID string = c.Param("id")
// 	input := request.UserProfile{}

// 	ctx := c.Request().Context()

// 	if err := c.Bind(&input); err != nil {
// 		return controllers.NewResponse(c, http.StatusBadRequest, http.StatusBadRequest, true, "invalid request", "")
// 	}

// 	err := input.Validate()
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusBadRequest, http.StatusBadRequest, true, "invalid request", "")
// 	}

// 	user, err := ctrl.authUseCase.UpdateProfileUser(ctx, input.ToDomain(), userID)

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusNotFound, http.StatusNotFound, true, err.Error(), "")
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, http.StatusOK, false, "customer updated", response.FromDomain(user))
// }

// func (ctrl *AuthController) UploadProfileImage(c echo.Context) error {
// 	userID := c.Param("id")
// 	file, err := c.FormFile("image")

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusBadRequest, http.StatusBadRequest, true, "error handling file upload", "")
// 	}

// 	src, err := file.Open()
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "Unable to open file", "")
// 	}
// 	defer src.Close()

// 	avatarPath := file.Filename
// 	// user, _, err := ctrl.authUseCase.UploadProfileImage(c.Request().Context(), avatarPath, userID)
// 	user, _, err := ctrl.authUseCase.UploadProfileImage(ctx, input.ToDomain(), avatarPath, userID)

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusNotFound, http.StatusNotFound, true, err.Error(), "")
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, http.StatusOK, false, "customer profile image updated", user)
// }
