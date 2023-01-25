package db_test

import (
	"testing"

	"github.com/brunohubner/fc2-hexagonal-architecture/src/adapters/db"
	"github.com/brunohubner/fc2-hexagonal-architecture/src/application"
	"github.com/stretchr/testify/require"
)

func TestProductDb_Get(t *testing.T) {
	db.SetUpTests()
	defer db.Close()
	productDb := db.NewProductDb(db.Db)

	product, err := productDb.Get("9235d7aa-7854-4f50-8d4d-0fd4d3263c0a")

	require.Nil(t, err)
	require.Equal(t, "Product test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	db.SetUpTests()
	defer db.Close()
	productDb := db.NewProductDb(db.Db)

	product := application.NewProduct()
	product.Name = "Phone"
	product.Price = 1899.89

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Name = "Phone - updated"
	product.Price = 2199.86
	product.Status = application.ENABLED

	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())
}
