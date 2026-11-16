export interface RoutePoint {
  readonly city: string;
  readonly region: string;
}

export interface VehicleCapacity {
  readonly freeTons: number;
  readonly loadTons: number;
  readonly utilizationPercent: number;
}

export interface Trip {
  readonly id: string;
  readonly tripId?: string;
  readonly origin: RoutePoint;
  readonly destination: RoutePoint;
  readonly date: string;
  readonly dateShort: string;
  readonly timeWindow: string;
  readonly vehicle: string;
  readonly capacityFreeTons: number;
  readonly matchCount: number;
  readonly status: 'matched' | 'pending' | 'unmatched';
}

export interface Load {
  readonly id: string;
  readonly origin: RoutePoint;
  readonly destination: RoutePoint;
  readonly pickupDate: string;
  readonly pickupTime: string;
  readonly cargoType: string;
  readonly weightTons: number;
  readonly volumeM3: number;
  readonly requiredVehicle: string;
  readonly restrictions: string;
}

export interface MatchCriterion {
  readonly label: 'Ruta' | 'Fecha' | 'Vehículo' | 'Capacidad' | 'Desvío';
  readonly value: string;
  readonly tone: 'positive' | 'warning';
}

export interface RouteMetrics {
  readonly detourKm: number;
  readonly additionalMinutes: number;
  readonly usableKm: number;
  readonly profitability: string;
}

export type RouteCoordinate = readonly [longitude: number, latitude: number];

export interface RouteGeometry {
  readonly planned: readonly RouteCoordinate[];
  readonly utilized: readonly RouteCoordinate[];
  readonly detour: readonly RouteCoordinate[];
}

export interface MatchOpportunity {
  readonly id: string;
  readonly compatibility: number;
  readonly level: string;
  readonly trip: Trip;
  readonly load: Load;
  readonly capacity: VehicleCapacity;
  readonly metrics: RouteMetrics;
  readonly criteria: readonly MatchCriterion[];
  readonly geometry: RouteGeometry;
  readonly route_geometry?: MatchRouteGeometry;
}

export interface MatchRouteWaypoint {
  readonly name: 'origen_viaje' | 'origen_carga' | 'destino_carga' | 'destino_viaje';
  readonly label: string;
  readonly coordinate: RouteCoordinate;
}

export interface MatchRouteGeometry {
  readonly coordinates: readonly RouteCoordinate[];
  readonly waypoints: readonly MatchRouteWaypoint[];
  readonly source: 'OSRM' | 'FALLBACK';
}

export interface OperationalMetric {
  readonly label: string;
  readonly value: number;
  readonly context?: string;
  readonly tone?: 'default' | 'positive' | 'warning';
}
