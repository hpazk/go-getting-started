package auth

import (
	"net/http"

	"github.com/hpazk/go-echo-rest-api/app/helpers"
	"github.com/hpazk/go-echo-rest-api/app/services"

	"github.com/labstack/echo/v4"
)

type (
	AuthController struct {
	}

	RegisterRequest struct {
		Email    string `json:"email" form:"email" query:"email" validate:"email,required"`
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	LoginRequest struct {
		Email    string `json:"email" form:"email" query:"email" validate:"email,required"`
		Password string `json:"password" validate:"required"`
	}
)

func (controller AuthController) Routes() []helpers.Route {
	return []helpers.Route{
		{
			Method:  echo.POST,
			Path:    "/auth/login",
			Handler: controller.Login,
		},
		{
			Method:  echo.POST,
			Path:    "/auth/register",
			Handler: controller.Register,
		},
	}
}

func (controller AuthController) Register(c echo.Context) error {
	params := new(RegisterRequest)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := c.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if user := services.GetUsersService().FindUserByEmail(params.Email); user != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "email is already used")
	}

	// need Response Formatter
	user := services.GetUsersService().AddUser(params.Name, params.Email, params.Password)
	response := helpers.ResponseFormatter{
		Code:    201,
		Status:  "success",
		Message: "user successfully registered",
		Data:    user,
	}
	return c.JSON(http.StatusCreated, response)
}

func (controller AuthController) Login(c echo.Context) error {
	params := new(LoginRequest)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := c.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user := services.GetUsersService().FindUserByEmail(params.Email)
	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}
	if matched := helpers.GetPasswordUtil().CheckPasswordHash(params.Password, user.Password); !matched {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}
	token, _ := GetAuthService().GetAccessToken(user)

	return c.JSON(http.StatusOK, map[string]string{
		"auth_token": token,
	})
}
