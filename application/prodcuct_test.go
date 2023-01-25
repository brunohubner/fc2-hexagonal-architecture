package application_test

import (
	"testing"

	"github.com/brunohubner/fc2-hexagonal-architecture/application"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enabled(t *testing.T) {
	product := application.Product{
		"1",
		"Phone",
		10,
		application.ENABLED,
	}

	err := product.Enabled()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enabled()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{
		uuid.New().String(),
		"Phone",
		0,
		application.ENABLED,
	}

	err := product.Disabled()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disabled()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	// product := application.Product{
	// 	uuid.New().String(),
	// 	"Phone",
	// 	0,
	// 	application.ENABLED,
	// }

	// _, err := product.IsValid()
	// require.Nil(t, err)

	// product.Status = application.ENABLED
	// product.Price = -10

	// _, err = product.IsValid()
	// require.Equal(t, "the status must be greater or equal zero", err.Error())
}
