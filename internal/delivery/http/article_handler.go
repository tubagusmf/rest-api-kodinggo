package handler

import (
	"golang-rest-api-articles/internal/model"
	"net/http"
	"strconv"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	articleUsecase model.IArticleUsecase
}

func NewArticleHandler(e *echo.Group, us model.IArticleUsecase) {
	handlers := &ArticleHandler{
		articleUsecase: us,
	}

	articles := e.Group("/articles")

	articles.Use(echojwt.WithConfig(jwtConfig()))

	articles.GET("", handlers.GetArticles)
	articles.GET("/:id", handlers.GetArticle)
	articles.POST("", handlers.CreateArticle)
	articles.PUT("/:id", handlers.UpdateArticle)
	articles.DELETE("/:id", handlers.DeleteArticle)
}

func (s *ArticleHandler) GetArticles(c echo.Context) error {
	claims := claimSession(c)
	if claims == nil {
		return c.JSON(http.StatusUnauthorized, response{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	if claims.Role != "admin" {
		return c.JSON(http.StatusUnauthorized, response{
			Status:  http.StatusForbidden,
			Message: "Forbidden",
		})
	}

	reqLimit := c.QueryParam("limit")
	reqOffset := c.QueryParam("offset")

	var limit, offset int32
	if reqLimit == "" {
		limit = 10
	}
	if reqOffset == "" {
		offset = 0
	}

	articles, err := s.articleUsecase.FindAll(c.Request().Context(), model.ArticleFilter{
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    articles,
	})

}

func (s *ArticleHandler) GetArticle(c echo.Context) error {
	claims := claimSession(c)
	if claims == nil {
		return c.JSON(http.StatusUnauthorized, response{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	id := c.Param("id")
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	article, err := s.articleUsecase.FindById(c.Request().Context(), int64(parsedId))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    article,
	})
}

func (s *ArticleHandler) CreateArticle(c echo.Context) error {
	claims := claimSession(c)
	if claims == nil {
		return c.JSON(http.StatusUnauthorized, response{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	var in model.CreateArticleInput
	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := s.articleUsecase.Create(c.Request().Context(), in); err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response{
		Message: "Success",
	})
}

func (s *ArticleHandler) UpdateArticle(c echo.Context) error {
	claims := claimSession(c)
	if claims == nil {
		return c.JSON(http.StatusUnauthorized, response{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	articleId := c.Param("id")
	parsedId, err := strconv.Atoi(articleId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var in model.UpdateArticleInput
	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if parsedId > 0 {
		in.Id = int64(parsedId)
	}

	if err := s.articleUsecase.Update(c.Request().Context(), in); err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Status:  http.StatusOK,
		Message: "Success",
	})

}

func (s *ArticleHandler) DeleteArticle(c echo.Context) error {
	claims := claimSession(c)
	if claims == nil {
		return c.JSON(http.StatusUnauthorized, response{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	id := c.Param("id")
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.articleUsecase.Delete(c.Request().Context(), int64(parsedId)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, response{
		Status:  http.StatusOK,
		Message: "Success",
	})
}
