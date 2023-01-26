package cli

import (
	"fmt"

	"github.com/brunohubner/fc2-hexagonal-architecture/src/application"
)

const (
	CREATE  = "create"
	ENABLE  = "enable"
	DISABLE = "disable"
	GET     = "get"
)

func Run(
	service application.IProductService,
	action string,
	productID string,
	productName string,
	price float64,
) (string, error) {
	result := ""

	switch action {
	case CREATE:
		product, err := service.Create(productName, price)
		if err != nil {
			return result, nil
		}

		result = fmt.Sprintf(
			"Product created\nID:     %s\nName:   %s\nPrice:  %.2f\nStatus: %s\n",
			product.GetID(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus(),
		)

	case ENABLE:
		product, err := service.Get(productID)
		if err != nil {
			return result, nil
		}

		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s has been enabled", res.GetName())

	case DISABLE:
		product, err := service.Get(productID)
		if err != nil {
			return result, nil
		}

		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s has been disabled", res.GetName())

	default:
		res, err := service.Get(productID)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf(
			"Product created\nID:     %s\nName:   %s\nPrice:  %.2f\nStatus: %s\n",
			res.GetID(),
			res.GetName(),
			res.GetPrice(),
			res.GetStatus(),
		)
	}
	return result, nil
}
