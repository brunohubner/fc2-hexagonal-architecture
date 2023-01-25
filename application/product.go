package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type IProduct interface {
	IsValid() (bool, error)
	Enabled() error
	Disabled() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("the status must be enabled or disabled")
	}

	if p.Price < 0 {
		return false, errors.New("the status must be greater or equal zero")
	}

	if _, err := govalidator.ValidateStruct(p); err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enabled() error {
	if p.Price <= 0 {
		return errors.New("the price must be greater than zero to enable the product")
	}

	p.Status = ENABLED

	return nil
}

func (p *Product) Disabled() error {
	if p.Price > 0 {
		return errors.New("the price must be zero in order to have the product disabled")
	}

	p.Status = DISABLED

	return nil
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
