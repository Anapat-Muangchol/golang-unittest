package services_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"gotest/repositories"
	"gotest/services"
	"testing"
)

func Test_PromotionCalculateDiscount(t *testing.T) {

	type testCase struct {
		name            string
		purchaseMin     int
		discountPercent int
		amount          int
		expected        int
	}

	cases := []testCase{
		{name: "discount case amount 100", purchaseMin: 100, discountPercent: 20, amount: 100, expected: 80},
		{name: "discount case amount 200", purchaseMin: 100, discountPercent: 20, amount: 200, expected: 160},
		{name: "discount case amount 300", purchaseMin: 100, discountPercent: 20, amount: 300, expected: 240},
		{name: "not discount case amount 50", purchaseMin: 100, discountPercent: 20, amount: 50, expected: 50},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Arrange
			promotionRepository := repositories.NewPromotionRepositoryMock()
			promotionRepository.On("GetPromotion").Return(repositories.Promotion{
				ID:              1,
				PurchaseMin:     c.purchaseMin,
				DiscountPercent: c.discountPercent,
			}, nil)

			promotionService := services.NewPromotionService(promotionRepository)

			// Act
			discount, _ := promotionService.CalculateDiscount(c.amount)
			expected := c.expected

			// Assert
			assert.Equal(t, expected, discount)
		})
	}

	t.Run("purchase amount zero", func(t *testing.T) {
		// Arrange
		promotionRepository := repositories.NewPromotionRepositoryMock()
		promotionRepository.On("GetPromotion").Return(repositories.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)

		promotionService := services.NewPromotionService(promotionRepository)

		// Act
		_, err := promotionService.CalculateDiscount(0)

		// Assert
		assert.ErrorIs(t, err, services.ErrZeroAmount)
		promotionRepository.AssertNotCalled(t, "GetPromotion")
	})

	t.Run("repository error", func(t *testing.T) {
		// Arrange
		promotionRepository := repositories.NewPromotionRepositoryMock()
		promotionRepository.On("GetPromotion").Return(repositories.Promotion{}, errors.New("repository error"))

		promotionService := services.NewPromotionService(promotionRepository)

		// Act
		_, err := promotionService.CalculateDiscount(100)

		// Assert
		assert.ErrorIs(t, err, services.ErrRepository)
	})
}
