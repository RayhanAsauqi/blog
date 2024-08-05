package handler

import (
	"net/http"
	"strconv"

	"github.com/RayhanAsauqi/blog-app/internal/dto"
	"github.com/RayhanAsauqi/blog-app/internal/service"
	"github.com/RayhanAsauqi/blog-app/pkg/response"
	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	Service service.ArticleService
}

// Add this constructor function
func NewArticleHandler(service service.ArticleService) *ArticleHandler {
	return &ArticleHandler{Service: service}
}

func (h *ArticleHandler) GetArticles(c echo.Context) error {
	articles, err := h.Service.FindAll(c.Request().Context())
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, articles)
}

func (h *ArticleHandler) GetArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid article ID")
	}

	article, err := h.Service.FindByID(c.Request().Context(), int64(id))
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, article)
}

func (h *ArticleHandler) CreateArticle(c echo.Context) error {
	var request dto.CreateArticleRequest
	if err := c.Bind(&request); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	if err := h.Service.Create(c.Request().Context(), request); err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusCreated, "Article created successfully")
}