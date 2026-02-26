package customer

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/godruoyi/go-snowflake"
)

type Status uint8

const (
	StatusPending   Status = 1
	StatusProcessed Status = 2
)

type Customer = *customer

type customer struct {
	ID           uint64
	CustomerName string
	Email        string
	PhoneNumber  string
	Message      string
	Note         string
	ServiceType  int
	Status       Status
}

func New(name, email, phone string, serviceType int) (Customer, error) {
	c := &customer{
		ID:           snowflake.ID(),
		CustomerName: name,
		Email:        email,
		PhoneNumber:  phone,
		ServiceType:  serviceType,
		Status:       StatusPending,
	}

	if err := c.validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *customer) validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.CustomerName, validation.Required),
		validation.Field(&c.Email, validation.Required, is.Email),
		validation.Field(&c.PhoneNumber, validation.Required),
		validation.Field(&c.ServiceType, validation.Required),
	)
}
