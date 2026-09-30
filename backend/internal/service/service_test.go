package service

import (
	"context"
	"testing"
	"time"

	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/domain"
	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type stubRepository struct {
	repository.Repository
	registeredHash string
	companySeen    string
}

func (s *stubRepository) RegisterCompanyAdmin(_ context.Context, input domain.RegisterCompanyInput, hash string) (domain.User, error) {
	s.registeredHash = hash
	return domain.User{ID: "user-1", CompanyID: "company-1", FirstName: input.FirstName, Email: input.Email, Role: domain.RoleAdmin, Status: "ACTIVO"}, nil
}

func (s *stubRepository) GetDriver(_ context.Context, companyID, driverID string) (domain.Driver, error) {
	s.companySeen = companyID
	return domain.Driver{User: domain.User{ID: driverID, CompanyID: companyID, Role: domain.RoleDriver}}, nil
}

func (s *stubRepository) GetVehicle(_ context.Context, companyID, vehicleID string) (domain.Vehicle, error) {
	s.companySeen = companyID
	return domain.Vehicle{ID: vehicleID, CompanyID: companyID}, nil
}

func (s *stubRepository) CreateTrip(_ context.Context, identity domain.Identity, input domain.CreateTripInput) (domain.Trip, error) {
	s.companySeen = identity.CompanyID
	return domain.Trip{ID: "trip-1", OperatorID: identity.UserID, DriverID: input.DriverID, VehicleID: input.VehicleID}, nil
}

func TestRegisterHashesPasswordAndForcesAdmin(t *testing.T) {
	repo := &stubRepository{}
	svc := New(repo, "a-secret-that-is-longer-than-thirty-two-characters", time.Hour)
	result, err := svc.Register(context.Background(), domain.RegisterCompanyInput{
		CompanyName: "Empresa", CUIT: "30-12345678-9", FirstName: "Ana",
		LastName: "Pérez", Email: "ana@example.com", Password: "Segura12345",
	})
	if err != nil {
		t.Fatalf("register returned error: %v", err)
	}
	if result.User.Role != domain.RoleAdmin {
		t.Fatalf("expected ADMIN, got %s", result.User.Role)
	}
	if repo.registeredHash == "Segura12345" || bcrypt.CompareHashAndPassword([]byte(repo.registeredHash), []byte("Segura12345")) != nil {
		t.Fatal("password was not stored as a valid bcrypt hash")
	}
	if _, err := svc.ParseToken(result.Token); err != nil {
		t.Fatalf("issued token is invalid: %v", err)
	}
}

func TestCreateTripUsesAuthenticatedCompany(t *testing.T) {
	repo := &stubRepository{}
	svc := New(repo, "a-secret-that-is-longer-than-thirty-two-characters", time.Hour)
	identity := domain.Identity{UserID: "operator-1", CompanyID: "company-a", Role: domain.RoleOperator}
	input := domain.CreateTripInput{
		DriverID: "driver-1", VehicleID: "vehicle-1", Origin: "Buenos Aires", Destination: "Mendoza",
		DepartureAt: time.Now().Add(time.Hour), EstimatedArrivalAt: time.Now().Add(8 * time.Hour),
		AvailableWeightTons: 10, AvailableVolumeM3: 30,
	}
	if _, err := svc.CreateTrip(context.Background(), identity, input); err != nil {
		t.Fatalf("create trip returned error: %v", err)
	}
	if repo.companySeen != identity.CompanyID {
		t.Fatalf("expected company %s, got %s", identity.CompanyID, repo.companySeen)
	}
}

func TestDriverCannotCreateTrip(t *testing.T) {
	svc := New(&stubRepository{}, "a-secret-that-is-longer-than-thirty-two-characters", time.Hour)
	_, err := svc.CreateTrip(context.Background(), domain.Identity{Role: domain.RoleDriver}, domain.CreateTripInput{})
	if err != ErrForbidden {
		t.Fatalf("expected forbidden, got %v", err)
	}
}
