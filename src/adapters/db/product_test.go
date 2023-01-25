package db_test

import (
	"testing"

	"github.com/brunohubner/fc2-hexagonal-architecture/src/adapters/db"
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
