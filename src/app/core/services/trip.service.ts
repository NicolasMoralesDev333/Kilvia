import { Injectable, signal } from '@angular/core';
import { Observable, of } from 'rxjs';
import { CreateTripRequest } from '../models/kilvia.models';

export interface VehicleOption {
  readonly id: string;
  readonly label: string;
}

@Injectable({ providedIn: 'root' })
export class TripService {
  private readonly vehiclesState = signal<readonly VehicleOption[]>([
    { id: 'VH-023', label: 'AC 908 YT · Semirremolque' },
    { id: 'VH-017', label: 'AF 324 JR · Sider' },
  ]);

  readonly vehicles = this.vehiclesState.asReadonly();

  createTrip(request: CreateTripRequest): Observable<{ readonly id: string }> {
    // Adaptador demo. La futura implementación reemplaza este retorno por POST /api/v1/trips.
    void request;
    return of({ id: 'VJ-DEMO' });
  }
}
