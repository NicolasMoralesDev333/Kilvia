import { HttpClient } from '@angular/common/http';
import { inject, Injectable, InjectionToken } from '@angular/core';
import { Observable, of } from 'rxjs';
import { MOCK_OPERATIONAL_METRICS, MOCK_OPPORTUNITIES, MOCK_TRIPS } from '../data/mock-data';
import { MatchOpportunity, OperationalMetric, Trip } from '../models/kilvia.models';

export interface OpportunityRepository {
  getMatches(): Observable<readonly MatchOpportunity[]>;
  getTrips(): Observable<readonly Trip[]>;
  getOperationalMetrics(): Observable<readonly OperationalMetric[]>;
}

export const OPPORTUNITY_REPOSITORY = new InjectionToken<OpportunityRepository>('OPPORTUNITY_REPOSITORY');

@Injectable()
export class MockOpportunityRepository implements OpportunityRepository {
  getMatches(): Observable<readonly MatchOpportunity[]> {
    return of(MOCK_OPPORTUNITIES);
  }

  getTrips(): Observable<readonly Trip[]> {
    return of(MOCK_TRIPS);
  }

  getOperationalMetrics(): Observable<readonly OperationalMetric[]> {
    return of(MOCK_OPERATIONAL_METRICS);
  }
}

@Injectable()
export class ApiOpportunityRepository implements OpportunityRepository {
  private readonly http = inject(HttpClient);
  private readonly baseUrl = '/api/v1';

  getMatches(): Observable<readonly MatchOpportunity[]> {
    return this.http.get<readonly MatchOpportunity[]>(`${this.baseUrl}/matches`);
  }

  getTrips(): Observable<readonly Trip[]> {
    return this.http.get<readonly Trip[]>(`${this.baseUrl}/trips`);
  }

  getOperationalMetrics(): Observable<readonly OperationalMetric[]> {
    return this.http.get<readonly OperationalMetric[]>(`${this.baseUrl}/operations/summary`);
  }
}
