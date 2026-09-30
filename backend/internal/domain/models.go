package domain

import "time"

type Role string

const (
	RoleAdmin    Role = "ADMIN"
	RoleOperator Role = "OPERADOR"
	RoleDriver   Role = "CHOFER"
)

type Identity struct {
	UserID    string
	CompanyID string
	Role      Role
}

type Company struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	CUIT         string    `json:"cuit"`
	Address      string    `json:"address"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	Status       string    `json:"status"`
	RegisteredAt time.Time `json:"registeredAt"`
}

type User struct {
	ID           string    `json:"id"`
	CompanyID    string    `json:"companyId"`
	CompanyName  string    `json:"companyName,omitempty"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Email        string    `json:"email"`
	Role         Role      `json:"role"`
	Status       string    `json:"status"`
	RegisteredAt time.Time `json:"registeredAt"`
	PasswordHash string    `json:"-"`
}

type Driver struct {
	User
	DNI               string     `json:"dni"`
	LicenseNumber     string     `json:"licenseNumber"`
	LicenseCategory   string     `json:"licenseCategory"`
	LicenseExpiration time.Time  `json:"licenseExpiration"`
	Phone             string     `json:"phone"`
	Availability      string     `json:"availability"`
	DriverStatus      string     `json:"driverStatus"`
}

type VehicleType struct {
	ID                    string  `json:"id"`
	Name                  string  `json:"name"`
	Description           string  `json:"description"`
	ReferenceWeightTons   float64 `json:"referenceWeightTons"`
	ReferenceVolumeM3     float64 `json:"referenceVolumeM3"`
}

type Vehicle struct {
	ID           string    `json:"id"`
	CompanyID    string    `json:"companyId"`
	TypeID       string    `json:"typeId"`
	TypeName     string    `json:"typeName"`
	Plate        string    `json:"plate"`
	Make         string    `json:"make"`
	Model        string    `json:"model"`
	WeightTons   float64   `json:"weightTons"`
	VolumeM3     float64   `json:"volumeM3"`
	Status       string    `json:"status"`
	RegisteredAt time.Time `json:"registeredAt"`
}

type Trip struct {
	ID                 string    `json:"id"`
	OperatorID         string    `json:"operatorId"`
	DriverID           string    `json:"driverId"`
	DriverName         string    `json:"driverName"`
	VehicleID          string    `json:"vehicleId"`
	VehicleLabel       string    `json:"vehicleLabel"`
	Origin             string    `json:"origin"`
	Destination        string    `json:"destination"`
	DepartureAt        time.Time `json:"departureAt"`
	EstimatedArrivalAt time.Time `json:"estimatedArrivalAt"`
	AvailableWeightTons float64  `json:"availableWeightTons"`
	AvailableVolumeM3  float64   `json:"availableVolumeM3"`
	Status             string    `json:"status"`
	RegisteredAt       time.Time `json:"registeredAt"`
}

type Load struct {
	ID              string    `json:"id"`
	UserID          string    `json:"userId"`
	Origin          string    `json:"origin"`
	Destination     string    `json:"destination"`
	PickupAt        time.Time `json:"pickupAt"`
	CargoType       string    `json:"cargoType"`
	WeightTons      float64   `json:"weightTons"`
	VolumeM3        float64   `json:"volumeM3"`
	RequiredVehicle string    `json:"requiredVehicle"`
	Restrictions    string    `json:"restrictions"`
	Status          string    `json:"status"`
	RegisteredAt    time.Time `json:"registeredAt"`
}

type RegisterCompanyInput struct {
	CompanyName string `json:"companyName"`
	CUIT        string `json:"cuit"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      Role   `json:"role"`
}

type UpdateUserInput struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Status    *string `json:"status"`
}

type CreateDriverInput struct {
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	DNI               string    `json:"dni"`
	LicenseNumber     string    `json:"licenseNumber"`
	LicenseCategory   string    `json:"licenseCategory"`
	LicenseExpiration time.Time `json:"licenseExpiration"`
	Phone             string    `json:"phone"`
	Availability      string    `json:"availability"`
}

type UpdateDriverInput struct {
	Phone        *string `json:"phone"`
	Availability *string `json:"availability"`
	Status       *string `json:"status"`
}

type CreateVehicleInput struct {
	TypeID     string  `json:"typeId"`
	Plate      string  `json:"plate"`
	Make       string  `json:"make"`
	Model      string  `json:"model"`
	WeightTons float64 `json:"weightTons"`
	VolumeM3   float64 `json:"volumeM3"`
}

type UpdateVehicleInput struct {
	TypeID     *string  `json:"typeId"`
	Make       *string  `json:"make"`
	Model      *string  `json:"model"`
	WeightTons *float64 `json:"weightTons"`
	VolumeM3   *float64 `json:"volumeM3"`
	Status     *string  `json:"status"`
}

type CreateTripInput struct {
	DriverID            string    `json:"driverId"`
	VehicleID           string    `json:"vehicleId"`
	Origin              string    `json:"origin"`
	Destination         string    `json:"destination"`
	DepartureAt         time.Time `json:"departureAt"`
	EstimatedArrivalAt  time.Time `json:"estimatedArrivalAt"`
	AvailableWeightTons float64   `json:"availableWeightTons"`
	AvailableVolumeM3   float64   `json:"availableVolumeM3"`
}

type UpdateTripInput struct {
	DriverID            *string    `json:"driverId"`
	VehicleID           *string    `json:"vehicleId"`
	Origin              *string    `json:"origin"`
	Destination         *string    `json:"destination"`
	DepartureAt         *time.Time `json:"departureAt"`
	EstimatedArrivalAt  *time.Time `json:"estimatedArrivalAt"`
	AvailableWeightTons *float64   `json:"availableWeightTons"`
	AvailableVolumeM3   *float64   `json:"availableVolumeM3"`
	Status              *string    `json:"status"`
}

type CreateLoadInput struct {
	Origin          string    `json:"origin"`
	Destination     string    `json:"destination"`
	PickupAt        time.Time `json:"pickupAt"`
	CargoType       string    `json:"cargoType"`
	WeightTons      float64   `json:"weightTons"`
	VolumeM3        float64   `json:"volumeM3"`
	RequiredVehicle string    `json:"requiredVehicle"`
	Restrictions    string    `json:"restrictions"`
}

type UpdateLoadInput struct {
	Origin          *string    `json:"origin"`
	Destination     *string    `json:"destination"`
	PickupAt        *time.Time `json:"pickupAt"`
	CargoType       *string    `json:"cargoType"`
	WeightTons      *float64   `json:"weightTons"`
	VolumeM3        *float64   `json:"volumeM3"`
	RequiredVehicle *string    `json:"requiredVehicle"`
	Restrictions    *string    `json:"restrictions"`
	Status          *string    `json:"status"`
}
