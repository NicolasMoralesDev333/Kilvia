import { HttpClient } from '@angular/common/http';
import { inject, Injectable, InjectionToken } from '@angular/core';
import { Observable, map, of, throwError } from 'rxjs';
import { MOCK_OPPORTUNITIES, MOCK_TRIPS } from '../data/mock-data';
import { MatchOpportunity, Trip } from '../models/kilvia.models';

export interface OpportunityRepository {
  getMatches(): Observable<readonly MatchOpportunity[]>;
  getMatch(id: string): Observable<MatchOpportunity>;
  getTrips(): Observable<readonly Trip[]>;
}

export const OPPORTUNITY_REPOSITORY = new InjectionToken<OpportunityRepository>('OPPORTUNITY_REPOSITORY');

@Injectable()
export class MockOpportunityRepository implements OpportunityRepository {
  getMatches(): Observable<readonly MatchOpportunity[]> {
    return of(MOCK_OPPORTUNITIES);
  }

  getMatch(id: string): Observable<MatchOpportunity> {
    const match = MOCK_OPPORTUNITIES.find((item) => item.id === id);
    return match ? of(match) : throwError(() => new Error('Match not found'));
  }

  getTrips(): Observable<readonly Trip[]> {
    return of(MOCK_TRIPS);
  }

}

@Injectable()
export class ApiOpportunityRepository implements OpportunityRepository {
  private readonly http = inject(HttpClient);
  private readonly baseUrl = '/api/v1';

  getMatches(): Observable<readonly MatchOpportunity[]> {
    return this.http.get<readonly MatchOpportunity[]>(`${this.baseUrl}/matches`);
  }

  getMatch(id: string): Observable<MatchOpportunity> {
    return this.http.get<MatchOpportunity>(`${this.baseUrl}/matches/${id}`);
  }

  getTrips(): Observable<readonly Trip[]> {
    return this.http
      .get<readonly ApiTrip[]>(`${this.baseUrl}/trips`)
      .pipe(map((trips) => trips.map((trip) => mapTrip(trip))));
  }
}

interface ApiTrip {
  readonly id: string;
  readonly code?: string;
  readonly origin: string;
  readonly destination: string;
  readonly departureAt: string;
  readonly estimatedArrivalAt: string;
  readonly vehicleLabel: string;
  readonly availableWeightTons: number;
  readonly status: string;
}

function mapTrip(trip: ApiTrip): Trip {
  const date = new Date(trip.departureAt);
  const dateText = new Intl.DateTimeFormat('es-AR', { day: '2-digit', month: 'short' }).format(date);
  return {
    id: trip.code ?? trip.id,
    tripId: trip.id,
    origin: { city: trip.origin, region: '' },
    destination: { city: trip.destination, region: '' },
    date: new Intl.DateTimeFormat('es-AR', { dateStyle: 'medium' }).format(date),
    dateShort: dateText.replace('.', ''),
    timeWindow: `${new Intl.DateTimeFormat('es-AR', { hour: '2-digit', minute: '2-digit' }).format(date)}–${new Intl.DateTimeFormat('es-AR', { hour: '2-digit', minute: '2-digit' }).format(new Date(trip.estimatedArrivalAt))}`,
    vehicle: trip.vehicleLabel,
    capacityFreeTons: trip.availableWeightTons,
    matchCount: 0,
    status: trip.status === 'PROGRAMADO' || trip.status === 'ACTIVO' ? 'unmatched' : 'pending',
  };
}
