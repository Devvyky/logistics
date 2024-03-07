package db

import (
	"context"
	"testing"

	"github.com/devvyky/logistics/util"
	"github.com/stretchr/testify/require"
)

func TestCreatePackSize(t *testing.T) {
	createRandomPackSize(t)
}

func createRandomPackSize(t *testing.T) ProductPackSize {
	arg := CreatePackSizeParams{
		ProductLine: util.RandomString(8),
		PackSize:    util.RandomInt(250, 5000),
	}

	productPackSize, err := testQueries.CreatePackSize(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, productPackSize)
	require.Equal(t, arg.ProductLine, productPackSize.ProductLine)
	require.Equal(t, arg.PackSize, productPackSize.PackSize)
	require.NotZero(t, productPackSize.ID)
	require.NotZero(t, productPackSize.CreatedAt)

	return productPackSize
}

func TestGetPackSize(t *testing.T) {
	pksize1 := createRandomPackSize(t)
	pksize2, err := testQueries.GetPackSize(context.Background(), pksize1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, pksize2)

	require.Equal(t, pksize1, pksize2)
}
