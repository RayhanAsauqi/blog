package handler

import (
	"net/http"

	"github.com/RayhanAsauqi/blog-app/internal/dto"
	"github.com/RayhanAsauqi/blog-app/internal/service"
	"github.com/RayhanAsauqi/blog-app/pkg/response"
	"github.com/labstack/echo/v4"
	"log"
)

type AuthHandler struct {
	Service service.AuthService
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

func (h *AuthHandler) Register(c echo.Context) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}
	}()

	// Validasi Service tidak nil
	if h.Service == nil {
		log.Println("Service is nil")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	var request dto.RegisterRequest
	if err := c.Bind(&request); err != nil {
		return response.ErrorResponsesAuth(c, http.StatusBadRequest, err.Error())
	}

	if err := h.Service.Register(c.Request().Context(), request); err != nil {
		return response.ErrorResponsesAuth(c, http.StatusBadRequest, err.Error())
	}

	return response.RegisterSuccess(c, http.StatusCreated, "User registered successfully")
}

func (h *AuthHandler) Login(c echo.Context) error {
	var request dto.LoginRequest
	if err := c.Bind(&request); err != nil {
		return response.ErrorResponsesAuth(c, http.StatusBadRequest, err.Error())
	}

	token, err := h.Service.Login(c.Request().Context(), request)
	
	if err != nil {
		return response.ErrorResponsesAuth(c, http.StatusUnauthorized, err.Error())
	}

	return response.Success(c, http.StatusOK, token)
}