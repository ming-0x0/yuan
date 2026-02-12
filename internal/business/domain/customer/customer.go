package customer

import (
	"net/mail"
	"strings"

	"github.com/ming-0x0/yuan/internal/common/domain"
	"github.com/ming-0x0/yuan/internal/common/domain/id"
)

type Customer = *customer

type customer struct {
	ID           id.ID
	CustomerName string
	Email        string
	PhoneNumber  string
	Message      string
	Note         string
	ServiceType  int
	Status       int
}

func New(name, email, phone string, serviceType int) (Customer, error) {
	if strings.TrimSpace(name) == "" || strings.TrimSpace(phone) == "" {
		return nil, domain.ErrRequiredField
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, domain.ErrInvalidEmail
	}

	return &customer{
		ID:           id.MustNew(),
		CustomerName: name,
		Email:        email,
		PhoneNumber:  phone,
		ServiceType:  serviceType,
		Status:       2,
	}, nil
}
