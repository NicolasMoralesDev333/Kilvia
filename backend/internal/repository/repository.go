package repository

import (
	"context"
	"errors"

	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/domain"
)

var (
	ErrNotFound = errors.New("repository: not found")
	ErrConflict = errors.New("repository: conflict")
)

type Repository interface {
	RegisterCompanyAdmin(context.Context, domain.RegisterCompanyInput, string) (domain.User, error)
	FindUserByEmail(context.Context, string) (domain.User, error)
	GetUser(context.Context, string, string) (domain.User, error)
	ListUsers(context.Context, string) ([]domain.User, error)
	CreateUser(context.Context, string, domain.CreateUserInput, string) (domain.User, error)
	UpdateUser(context.Context, string, string, domain.UpdateUserInput) (domain.User, error)

	GetCompany(context.Context, string) (domain.Company, error)
	UpdateCompany(context.Context, string, domain.Company) (domain.Company, error)

	ListDrivers(context.Context, string) ([]domain.Driver, error)
	GetDriver(context.Context, string, string) (domain.Driver, error)
	CreateDriver(context.Context, string, domain.CreateDriverInput, string) (domain.Driver, error)
	UpdateDriver(context.Context, string, string, domain.UpdateDriverInput) (domain.Driver, error)

	ListVehicleTypes(context.Context) ([]domain.VehicleType, error)
	ListVehicles(context.Context, string) ([]domain.Vehicle, error)
	GetVehicle(context.Context, string, string) (domain.Vehicle, error)
	CreateVehicle(context.Context, string, domain.CreateVehicleInput) (domain.Vehicle, error)
	UpdateVehicle(context.Context, string, string, domain.UpdateVehicleInput) (domain.Vehicle, error)

	ListTrips(context.Context, domain.Identity) ([]domain.Trip, error)
	GetTrip(context.Context, domain.Identity, string) (domain.Trip, error)
	CreateTrip(context.Context, domain.Identity, domain.CreateTripInput) (domain.Trip, error)
	UpdateTrip(context.Context, domain.Identity, string, domain.UpdateTripInput) (domain.Trip, error)

	ListLoads(context.Context, domain.Identity) ([]domain.Load, error)
	GetLoad(context.Context, domain.Identity, string) (domain.Load, error)
	CreateLoad(context.Context, domain.Identity, domain.CreateLoadInput) (domain.Load, error)
	UpdateLoad(context.Context, domain.Identity, string, domain.UpdateLoadInput) (domain.Load, error)
}
