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
	DNI               string    `json:"dni"`
	LicenseNumber     string    `json:"licenseNumber"`
	LicenseCategory   string    `json:"licenseCategory"`
	LicenseExpiration time.Time `json:"licenseExpiration"`
	Phone             string    `json:"phone"`
	Availability      string    `json:"availability"`
	DriverStatus      string    `json:"driverStatus"`
}

type VehicleType struct {
	ID                  string  `json:"id"`
	Name                string  `json:"name"`
	Description         string  `json:"description"`
	ReferenceWeightTons float64 `json:"referenceWeightTons"`
	ReferenceVolumeM3   float64 `json:"referenceVolumeM3"`
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
	ID                  string    `json:"id"`
	Code                string    `json:"code"`
	OperatorID          string    `json:"operatorId"`
	DriverID            string    `json:"driverId"`
	DriverName          string    `json:"driverName"`
	VehicleID           string    `json:"vehicleId"`
	VehicleLabel        string    `json:"vehicleLabel"`
	Origin              string    `json:"origin"`
	Destination         string    `json:"destination"`
	DepartureAt         time.Time `json:"departureAt"`
	EstimatedArrivalAt  time.Time `json:"estimatedArrivalAt"`
	AvailableWeightTons float64   `json:"availableWeightTons"`
	AvailableVolumeM3   float64   `json:"availableVolumeM3"`
	Status              string    `json:"status"`
	RegisteredAt        time.Time `json:"registeredAt"`
	OriginLat           *float64  `json:"originLat,omitempty"`
	OriginLng           *float64  `json:"originLng,omitempty"`
	DestinationLat      *float64  `json:"destinationLat,omitempty"`
	DestinationLng      *float64  `json:"destinationLng,omitempty"`
	BaseDistanceKM      float64   `json:"baseDistanceKm"`
}

type Load struct {
	ID              string    `json:"id"`
	Code            string    `json:"code"`
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
	OriginLat       *float64  `json:"originLat,omitempty"`
	OriginLng       *float64  `json:"originLng,omitempty"`
	DestinationLat  *float64  `json:"destinationLat,omitempty"`
	DestinationLng  *float64  `json:"destinationLng,omitempty"`
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
	OriginLat           *float64  `json:"originLat,omitempty"`
	OriginLng           *float64  `json:"originLng,omitempty"`
	DestinationLat      *float64  `json:"destinationLat,omitempty"`
	DestinationLng      *float64  `json:"destinationLng,omitempty"`
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
	OriginLat           *float64   `json:"originLat,omitempty"`
	OriginLng           *float64   `json:"originLng,omitempty"`
	DestinationLat      *float64   `json:"destinationLat,omitempty"`
	DestinationLng      *float64   `json:"destinationLng,omitempty"`
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
	OriginLat       *float64  `json:"originLat,omitempty"`
	OriginLng       *float64  `json:"originLng,omitempty"`
	DestinationLat  *float64  `json:"destinationLat,omitempty"`
	DestinationLng  *float64  `json:"destinationLng,omitempty"`
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
	OriginLat       *float64   `json:"originLat,omitempty"`
	OriginLng       *float64   `json:"originLng,omitempty"`
	DestinationLat  *float64   `json:"destinationLat,omitempty"`
	DestinationLng  *float64   `json:"destinationLng,omitempty"`
}

type MatchCandidate struct {
	TripID, TripCode, TripOrigin, TripDestination string
	TripOriginRegion, TripDestinationRegion       string
	TripDeparture, TripArrival                    time.Time
	TripWeight, TripVolume                        float64
	TripStatus, VehicleName, VehicleLabel         string
	TripOriginLat, TripOriginLng                  *float64
	TripDestinationLat, TripDestinationLng        *float64
	LoadID, LoadCode, LoadOrigin, LoadDestination string
	LoadOriginRegion, LoadDestinationRegion       string
	LoadPickup                                    time.Time
	LoadWeight, LoadVolume                        float64
	LoadCargoType, LoadRequiredVehicle            string
	LoadRestrictions, LoadStatus                  string
	LoadOriginLat, LoadOriginLng                  *float64
	LoadDestinationLat, LoadDestinationLng        *float64
}

type MatchRecord struct {
	TripID, LoadID string
	Compatibility  float64
	DetourKM       float64
	AdditionalMins int
	UsableKM       float64
	Utilization    float64
	GeometryJSON   []byte
}

type MatchItem struct {
	ID            string           `json:"id"`
	Compatibility int              `json:"compatibility"`
	Level         string           `json:"level"`
	Trip          MatchTrip        `json:"trip"`
	Load          MatchLoad        `json:"load"`
	Capacity      MatchCapacity    `json:"capacity"`
	Metrics       MatchMetrics     `json:"metrics"`
	Criteria      []MatchCriterion `json:"criteria"`
	Geometry      MatchGeometry    `json:"geometry"`
	RouteGeometry RouteGeometry    `json:"route_geometry"`
}

type MatchTrip struct {
	ID               string        `json:"id"`
	Origin           MatchLocation `json:"origin"`
	Destination      MatchLocation `json:"destination"`
	Date             string        `json:"date"`
	DateShort        string        `json:"dateShort"`
	TimeWindow       string        `json:"timeWindow"`
	Vehicle          string        `json:"vehicle"`
	CapacityFreeTons float64       `json:"capacityFreeTons"`
	MatchCount       int           `json:"matchCount"`
	Status           string        `json:"status"`
}

type MatchLoad struct {
	ID              string        `json:"id"`
	Origin          MatchLocation `json:"origin"`
	Destination     MatchLocation `json:"destination"`
	PickupDate      string        `json:"pickupDate"`
	PickupTime      string        `json:"pickupTime"`
	CargoType       string        `json:"cargoType"`
	WeightTons      float64       `json:"weightTons"`
	VolumeM3        float64       `json:"volumeM3"`
	RequiredVehicle string        `json:"requiredVehicle"`
	Restrictions    string        `json:"restrictions"`
}

type MatchLocation struct {
	City   string `json:"city"`
	Region string `json:"region"`
}

type MatchCapacity struct {
	FreeTons           float64 `json:"freeTons"`
	LoadTons           float64 `json:"loadTons"`
	UtilizationPercent float64 `json:"utilizationPercent"`
}

type MatchMetrics struct {
	DetourKM          float64 `json:"detourKm"`
	AdditionalMinutes int     `json:"additionalMinutes"`
	UsableKM          float64 `json:"usableKm"`
	Profitability     string  `json:"profitability"`
}

type MatchCriterion struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Tone  string `json:"tone"`
}

type MatchGeometry struct {
	Planned  [][]float64 `json:"planned"`
	Utilized [][]float64 `json:"utilized"`
	Detour   [][]float64 `json:"detour"`
}

type MatchRouteWaypoint struct {
	Name       string    `json:"name"`
	Label      string    `json:"label"`
	Coordinate []float64 `json:"coordinate"`
}

type RouteGeometry struct {
	Coordinates [][]float64          `json:"coordinates"`
	Waypoints   []MatchRouteWaypoint `json:"waypoints"`
	Source      string               `json:"source"`
}
