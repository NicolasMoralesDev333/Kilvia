package service

import (
	"context"
	"strings"

	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) GetCompany(ctx context.Context, identity domain.Identity) (domain.Company, error) {
	item, err := s.repository.GetCompany(ctx, identity.CompanyID)
	return item, translateRepositoryError(err)
}

func (s *Service) UpdateCompany(ctx context.Context, identity domain.Identity, company domain.Company) (domain.Company, error) {
	if identity.Role != domain.RoleAdmin {
		return domain.Company{}, ErrForbidden
	}
	if strings.TrimSpace(company.Name) == "" {
		return domain.Company{}, ErrInvalidInput
	}
	item, err := s.repository.UpdateCompany(ctx, identity.CompanyID, company)
	return item, translateRepositoryError(err)
}

func (s *Service) ListUsers(ctx context.Context, identity domain.Identity) ([]domain.User, error) {
	items, err := s.repository.ListUsers(ctx, identity.CompanyID)
	return items, translateRepositoryError(err)
}

func (s *Service) GetUser(ctx context.Context, identity domain.Identity, userID string) (domain.User, error) {
	item, err := s.repository.GetUser(ctx, identity.CompanyID, userID)
	return item, translateRepositoryError(err)
}

func (s *Service) CreateUser(ctx context.Context, identity domain.Identity, input domain.CreateUserInput) (domain.User, error) {
	if identity.Role != domain.RoleAdmin {
		return domain.User{}, ErrForbidden
	}
	if input.Role != domain.RoleOperator || strings.TrimSpace(input.FirstName) == "" ||
		strings.TrimSpace(input.LastName) == "" || !validEmail(input.Email) || len(input.Password) < 10 {
		return domain.User{}, ErrInvalidInput
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}
	item, err := s.repository.CreateUser(ctx, identity.CompanyID, input, string(hash))
	return item, translateRepositoryError(err)
}

func (s *Service) UpdateUser(ctx context.Context, identity domain.Identity, userID string, input domain.UpdateUserInput) (domain.User, error) {
	if identity.Role != domain.RoleAdmin {
		return domain.User{}, ErrForbidden
	}
	item, err := s.repository.UpdateUser(ctx, identity.CompanyID, userID, input)
	return item, translateRepositoryError(err)
}

func (s *Service) ListDrivers(ctx context.Context, identity domain.Identity) ([]domain.Driver, error) {
	items, err := s.repository.ListDrivers(ctx, identity.CompanyID)
	return items, translateRepositoryError(err)
}

func (s *Service) GetDriver(ctx context.Context, identity domain.Identity, driverID string) (domain.Driver, error) {
	item, err := s.repository.GetDriver(ctx, identity.CompanyID, driverID)
	return item, translateRepositoryError(err)
}

func (s *Service) CreateDriver(ctx context.Context, identity domain.Identity, input domain.CreateDriverInput) (domain.Driver, error) {
	if identity.Role != domain.RoleAdmin {
		return domain.Driver{}, ErrForbidden
	}
	if strings.TrimSpace(input.FirstName) == "" || strings.TrimSpace(input.LastName) == "" ||
		!validEmail(input.Email) || len(input.Password) < 10 || strings.TrimSpace(input.DNI) == "" ||
		strings.TrimSpace(input.LicenseNumber) == "" || input.LicenseExpiration.IsZero() {
		return domain.Driver{}, ErrInvalidInput
	}
	if input.Availability == "" {
		input.Availability = "DISPONIBLE"
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.Driver{}, err
	}
	item, err := s.repository.CreateDriver(ctx, identity.CompanyID, input, string(hash))
	return item, translateRepositoryError(err)
}

func (s *Service) UpdateDriver(ctx context.Context, identity domain.Identity, driverID string, input domain.UpdateDriverInput) (domain.Driver, error) {
	if identity.Role != domain.RoleAdmin {
		return domain.Driver{}, ErrForbidden
	}
	item, err := s.repository.UpdateDriver(ctx, identity.CompanyID, driverID, input)
	return item, translateRepositoryError(err)
}

func (s *Service) ListVehicleTypes(ctx context.Context) ([]domain.VehicleType, error) {
	return s.repository.ListVehicleTypes(ctx)
}

func (s *Service) ListVehicles(ctx context.Context, identity domain.Identity) ([]domain.Vehicle, error) {
	items, err := s.repository.ListVehicles(ctx, identity.CompanyID)
	return items, translateRepositoryError(err)
}

func (s *Service) GetVehicle(ctx context.Context, identity domain.Identity, vehicleID string) (domain.Vehicle, error) {
	item, err := s.repository.GetVehicle(ctx, identity.CompanyID, vehicleID)
	return item, translateRepositoryError(err)
}

func (s *Service) CreateVehicle(ctx context.Context, identity domain.Identity, input domain.CreateVehicleInput) (domain.Vehicle, error) {
	if strings.TrimSpace(input.TypeID) == "" || strings.TrimSpace(input.Plate) == "" ||
		strings.TrimSpace(input.Make) == "" || strings.TrimSpace(input.Model) == "" ||
		input.WeightTons <= 0 || input.VolumeM3 <= 0 {
		return domain.Vehicle{}, ErrInvalidInput
	}
	item, err := s.repository.CreateVehicle(ctx, identity.CompanyID, input)
	return item, translateRepositoryError(err)
}

func (s *Service) UpdateVehicle(ctx context.Context, identity domain.Identity, vehicleID string, input domain.UpdateVehicleInput) (domain.Vehicle, error) {
	item, err := s.repository.UpdateVehicle(ctx, identity.CompanyID, vehicleID, input)
	return item, translateRepositoryError(err)
}

func (s *Service) ListTrips(ctx context.Context, identity domain.Identity) ([]domain.Trip, error) {
	items, err := s.repository.ListTrips(ctx, identity)
	return items, translateRepositoryError(err)
}

func (s *Service) GetTrip(ctx context.Context, identity domain.Identity, tripID string) (domain.Trip, error) {
	item, err := s.repository.GetTrip(ctx, identity, tripID)
	return item, translateRepositoryError(err)
}

func (s *Service) CreateTrip(ctx context.Context, identity domain.Identity, input domain.CreateTripInput) (domain.Trip, error) {
	if identity.Role == domain.RoleDriver {
		return domain.Trip{}, ErrForbidden
	}
	if strings.TrimSpace(input.Origin) == "" || strings.TrimSpace(input.Destination) == "" ||
		!input.EstimatedArrivalAt.After(input.DepartureAt) ||
		input.AvailableWeightTons < 0 || input.AvailableVolumeM3 < 0 {
		return domain.Trip{}, ErrInvalidInput
	}
	if _, err := s.repository.GetDriver(ctx, identity.CompanyID, input.DriverID); err != nil {
		return domain.Trip{}, translateRepositoryError(err)
	}
	if _, err := s.repository.GetVehicle(ctx, identity.CompanyID, input.VehicleID); err != nil {
		return domain.Trip{}, translateRepositoryError(err)
	}
	item, err := s.repository.CreateTrip(ctx, identity, input)
	return item, translateRepositoryError(err)
}

func (s *Service) UpdateTrip(ctx context.Context, identity domain.Identity, tripID string, input domain.UpdateTripInput) (domain.Trip, error) {
	if identity.Role == domain.RoleDriver {
		return domain.Trip{}, ErrForbidden
	}
	current, err := s.repository.GetTrip(ctx, identity, tripID)
	if err != nil {
		return domain.Trip{}, translateRepositoryError(err)
	}
	if input.DriverID != nil {
		if _, err := s.repository.GetDriver(ctx, identity.CompanyID, *input.DriverID); err != nil {
			return domain.Trip{}, translateRepositoryError(err)
		}
	}
	if input.VehicleID != nil {
		if _, err := s.repository.GetVehicle(ctx, identity.CompanyID, *input.VehicleID); err != nil {
			return domain.Trip{}, translateRepositoryError(err)
		}
	}
	departure, arrival := current.DepartureAt, current.EstimatedArrivalAt
	if input.DepartureAt != nil {
		departure = *input.DepartureAt
	}
	if input.EstimatedArrivalAt != nil {
		arrival = *input.EstimatedArrivalAt
	}
	if !arrival.After(departure) {
		return domain.Trip{}, ErrInvalidInput
	}
	item, err := s.repository.UpdateTrip(ctx, identity, tripID, input)
	return item, translateRepositoryError(err)
}

func (s *Service) ListLoads(ctx context.Context, identity domain.Identity) ([]domain.Load, error) {
	if identity.Role == domain.RoleDriver {
		return nil, ErrForbidden
	}
	items, err := s.repository.ListLoads(ctx, identity)
	return items, translateRepositoryError(err)
}

func (s *Service) GetLoad(ctx context.Context, identity domain.Identity, loadID string) (domain.Load, error) {
	if identity.Role == domain.RoleDriver {
		return domain.Load{}, ErrForbidden
	}
	item, err := s.repository.GetLoad(ctx, identity, loadID)
	return item, translateRepositoryError(err)
}

func (s *Service) CreateLoad(ctx context.Context, identity domain.Identity, input domain.CreateLoadInput) (domain.Load, error) {
	if identity.Role == domain.RoleDriver {
		return domain.Load{}, ErrForbidden
	}
	if strings.TrimSpace(input.Origin) == "" || strings.TrimSpace(input.Destination) == "" ||
		input.PickupAt.IsZero() || input.WeightTons <= 0 || input.VolumeM3 <= 0 {
		return domain.Load{}, ErrInvalidInput
	}
	item, err := s.repository.CreateLoad(ctx, identity, input)
	return item, translateRepositoryError(err)
}

func (s *Service) UpdateLoad(ctx context.Context, identity domain.Identity, loadID string, input domain.UpdateLoadInput) (domain.Load, error) {
	if identity.Role == domain.RoleDriver {
		return domain.Load{}, ErrForbidden
	}
	item, err := s.repository.UpdateLoad(ctx, identity, loadID, input)
	return item, translateRepositoryError(err)
}
