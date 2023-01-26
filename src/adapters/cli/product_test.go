package cli_test

import (
	"fmt"
	"testing"

	"github.com/brunohubner/fc2-hexagonal-architecture/src/adapters/cli"
	"github.com/brunohubner/fc2-hexagonal-architecture/src/application"
	mock_application "github.com/brunohubner/fc2-hexagonal-architecture/src/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productID := uuid.New().String()
	productName := "Notebook"
	productPrice := 4389.59
	productStatus := application.ENABLED

	productMock := mock_application.NewMockIProduct(ctrl)
	productMock.EXPECT().GetID().Return(productID).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockIProductService(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productID).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	expected := fmt.Sprintf(
		"Product created\nID:     %s\nName:   %s\nPrice:  %.2f\nStatus: %s\n",
		productID,
		productName,
		productPrice,
		productStatus,
	)
	result, err := cli.Run(service, cli.CREATE, "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expected, result)

	expected = fmt.Sprintf("Product %s has been enabled", productName)
	result, err = cli.Run(service, cli.ENABLE, productID, "", 0)
	require.Nil(t, err)
	require.Equal(t, expected, result)

	expected = fmt.Sprintf("Product %s has been disabled", productName)
	result, err = cli.Run(service, cli.DISABLE, productID, "", 0)
	require.Nil(t, err)
	require.Equal(t, expected, result)

	expected = fmt.Sprintf(
		"Product created\nID:     %s\nName:   %s\nPrice:  %.2f\nStatus: %s\n",
		productID,
		productName,
		productPrice,
		productStatus,
	)
	result, err = cli.Run(service, cli.GET, productID, "", 0)
	require.Nil(t, err)
	require.Equal(t, expected, result)
}
