package handlers_test

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gotest/handlers"
	"gotest/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"
)

func Test_PromotionCalculateDiscount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		amount := 100
		expected := 80

		promotionService := services.NewPromotionServiceMock()
		promotionService.On("CalculateDiscount", amount).Return(expected, nil)

		promotionHandler := handlers.NewPromotionHandler(promotionService)

		// http://localhost:8000/calculate?amount=100
		app := fiber.New()
		app.Get("/calculate", promotionHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		// Act
		res, _ := app.Test(req)

		// Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}
	})

	t.Run("bad request", func(t *testing.T) {
		// Arrange
		amount := "errorAmount"

		promotionService := services.NewPromotionServiceMock()
		promotionHandler := handlers.NewPromotionHandler(promotionService)

		// http://localhost:8000/calculate?amount=errorAmount
		app := fiber.New()
		app.Get("/calculate", promotionHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		// Act
		res, _ := app.Test(req)

		// Assert
		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})

	t.Run("internal server error", func(t *testing.T) {
		// Arrange
		amount := 100

		promotionService := services.NewPromotionServiceMock()
		promotionService.On("CalculateDiscount", amount).Return(0, services.ErrRepository)

		promotionHandler := handlers.NewPromotionHandler(promotionService)

		// http://localhost:8000/calculate?amount=100
		app := fiber.New()
		app.Get("/calculate", promotionHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		// Act
		res, _ := app.Test(req)

		// Assert
		assert.Equal(t, fiber.StatusInternalServerError, res.StatusCode)
	})
}
