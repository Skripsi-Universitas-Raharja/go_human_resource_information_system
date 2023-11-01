package profiles

import (
	// "backend-golang/app/middlewares"
	"backend-golang/businesses/profiles"
	"fmt"

	"backend-golang/controllers"
	"backend-golang/controllers/profiles/request"
	"backend-golang/controllers/profiles/response"
	"net/http"

	// "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ProfilesController struct {
	authUseCase profiles.Usecase
}

func NewProfilesController(authUC profiles.Usecase) *ProfilesController {
	return &ProfilesController{
		authUseCase: authUC,
	}
}

func (ctrl *ProfilesController) GetByID(c echo.Context) error {
	var userID string = c.Param("id")
	// token := c.Get("user").(*jwt.Token)

	// isListed := middlewares.CheckToken(token.Raw)

	// if !isListed {
	// 	return controllers.NewResponse(c, http.StatusUnauthorized, http.StatusUnauthorized, true, "invalid token", isListed)
	// }

	ctx := c.Request().Context()

	user, err := ctrl.authUseCase.GetByID(ctx, userID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, http.StatusNotFound, true, err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusOK, http.StatusOK, true, "get user by id", response.FromDomain(user))
}

func (ctrl *ProfilesController) UpdateProfileUser(c echo.Context) error {
	profileInput := request.UserProfile{}
	ctx := c.Request().Context()

	var userID string = c.Param("id")

	if err := c.Bind(&profileInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, http.StatusBadRequest, true, "invalid request", "")
	}

	err := profileInput.Validate()
	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, http.StatusBadRequest, true, "invalid request", "")
	}
	//percobaan
	if userID == "" {
		return controllers.NewResponse(c, http.StatusBadRequest, http.StatusBadRequest, true, "invalid user ID", "")
	}

	// nil
	profile, err := ctrl.authUseCase.UpdateProfileUser(ctx, profileInput.ToDomain(), userID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, http.StatusNotFound, true, "No data found", "")
	}

	// nil
	return controllers.NewResponse(c, http.StatusOK, http.StatusOK, false, "customer updated", response.FromDomain(profile))
}

func (ctrl *ProfilesController) UploadProfileImage(c echo.Context) error {
	profileInput := request.UserProfile{}
	ctx := c.Request().Context()

	userID := c.Param("id")
	file, err := c.FormFile("image")

	if err != nil {

		fmt.Println("Unable to open file:", err.Error())

		return controllers.NewResponse(c, http.StatusBadRequest, http.StatusBadRequest, true, "error handling file upload", "")
	}

	src, err := file.Open()
	if err != nil {
		fmt.Println("Unable to open file:", err.Error())

		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "Unable to open file", "")
	}
	defer src.Close()

	avatarPath := file.Filename
	user, _, err := ctrl.authUseCase.UploadProfileImage(ctx, profileInput.ToDomain(), avatarPath, userID)

	if err != nil {
		fmt.Println("Unable to open file:", err.Error())

		return controllers.NewResponse(c, http.StatusNotFound, http.StatusNotFound, true, err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusOK, http.StatusOK, false, "customer profile image updated", user)
}
