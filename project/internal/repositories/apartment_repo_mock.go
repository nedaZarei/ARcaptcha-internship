package repositories

import (
	"context"

	"github.com/nedaZarei/arcaptcha-internship-2025/neda-arcaptcha-internship-2025/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockApartmentRepo struct {
	mock.Mock
}

func (m *MockApartmentRepo) CreateApartment(ctx context.Context, apartment models.Apartment) (int, error) {
	args := m.Called(ctx, apartment)
	return args.Int(0), args.Error(1)
}

func (m *MockApartmentRepo) GetApartmentByID(id int) (*models.Apartment, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Apartment), args.Error(1)
}

func (m *MockApartmentRepo) UpdateApartment(ctx context.Context, apartment models.Apartment) error {
	args := m.Called(ctx, apartment)
	return args.Error(0)
}

func (m *MockApartmentRepo) DeleteApartment(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
