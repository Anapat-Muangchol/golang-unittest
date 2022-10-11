//go:build integration

package handlers_test

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gotest/handlers"
	"gotest/repositories"
	"gotest/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"
)

func Test_PromotionCalculateDiscountIntegrationService(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		amount := 100
		expected := 80

		promotionRepo := repositories.NewPromotionRepositoryMock()
		promotionRepo.On("GetPromotion").Return(repositories.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)

		promotionService := services.NewPromotionService(promotionRepo)
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

}
