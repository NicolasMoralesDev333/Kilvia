export type UserRole = 'ADMIN' | 'OPERADOR' | 'CHOFER';

export interface AuthUser {
  readonly id: string;
  readonly companyId: string;
  readonly companyName?: string;
  readonly firstName: string;
  readonly lastName: string;
  readonly email: string;
  readonly role: UserRole;
  readonly status: string;
  readonly registeredAt: string;
}

export type ApiUser = AuthUser;

export interface ApiCompany {
  readonly id: string;
  readonly name: string;
  readonly cuit: string;
  readonly address: string;
  readonly phone: string;
  readonly email: string;
  readonly status: string;
  readonly registeredAt: string;
}

export interface AuthResult {
  readonly token: string;
  readonly user: AuthUser;
}

export interface LoginRequest {
  readonly email: string;
  readonly password: string;
}

export interface RegisterRequest {
  readonly companyName: string;
  readonly cuit: string;
  readonly firstName: string;
  readonly lastName: string;
  readonly email: string;
  readonly password: string;
}

export interface ApiVehicle {
  readonly id: string;
  readonly typeId: string;
  readonly typeName: string;
  readonly plate: string;
  readonly make: string;
  readonly model: string;
  readonly weightTons: number;
  readonly volumeM3: number;
  readonly status: string;
}

export interface ApiDriver {
  readonly id: string;
  readonly firstName: string;
  readonly lastName: string;
  readonly email: string;
  readonly dni: string;
  readonly licenseNumber: string;
  readonly licenseCategory: string;
  readonly licenseExpiration: string;
  readonly phone: string;
  readonly availability: string;
  readonly driverStatus: string;
}

export interface ApiTrip {
  readonly id: string;
  readonly operatorId: string;
  readonly driverId: string;
  readonly driverName: string;
  readonly vehicleId: string;
  readonly vehicleLabel: string;
  readonly origin: string;
  readonly destination: string;
  readonly departureAt: string;
  readonly estimatedArrivalAt: string;
  readonly availableWeightTons: number;
  readonly availableVolumeM3: number;
  readonly status: string;
}

export interface ApiLoad {
  readonly id: string;
  readonly userId: string;
  readonly origin: string;
  readonly destination: string;
  readonly pickupAt: string;
  readonly cargoType: string;
  readonly weightTons: number;
  readonly volumeM3: number;
  readonly requiredVehicle: string;
  readonly restrictions: string;
  readonly status: string;
}

export interface ApiVehicleType {
  readonly id: string;
  readonly name: string;
  readonly description: string;
  readonly referenceWeightTons: number;
  readonly referenceVolumeM3: number;
}
