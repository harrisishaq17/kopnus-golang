package controller

import (
	"auth-service/model"
	"auth-service/service"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *userController {
	return &userController{service}
}

func (controller *userController) UserRoutes(e *echo.Echo) {
	e.Use(middleware.CORS())

	// User EP
	var userRoute = e.Group("/auth")
	userRoute.Use(middleware.BodyDump(Logger))
	userRoute.POST("/info-user", controller.InfoUser)
	userRoute.POST("/login", controller.Login)
}

func (ctrl *userController) InfoUser(c echo.Context) error {
	request := new(model.GetDataUserRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewJsonResponse(false).SetError("400", "Bad Request"))
	} else if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewJsonResponse(false).SetError("400", err.Error()))
	}

	resp, err := ctrl.service.GetDataUser(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.(*model.JsonResponse))
	}

	return c.JSON(http.StatusOK, model.NewJsonResponse(true).SetData(resp))
}

func (ctrl *userController) Login(c echo.Context) error {
	request := new(model.LoginUserRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewJsonResponse(false).SetError("400", "Bad Request"))
	} else if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewJsonResponse(false).SetError("400", err.Error()))
	}

	token, err := ctrl.service.Login(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.(*model.JsonResponse))
	}

	return c.JSON(http.StatusOK, model.NewJsonResponse(true).SetData(token))
}
