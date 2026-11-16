import { computed, DestroyRef, inject, Injectable, signal } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { catchError, forkJoin, of } from 'rxjs';
import { MatchOpportunity, OperationalMetric, Trip } from '../models/kilvia.models';
import { OPPORTUNITY_REPOSITORY } from './opportunity.repository';
import { KilviaApiService } from './kilvia-api.service';

@Injectable({ providedIn: 'root' })
export class MatchService {
  private readonly repository = inject(OPPORTUNITY_REPOSITORY);
  private readonly api = inject(KilviaApiService);
  private readonly destroyRef = inject(DestroyRef);
  private readonly opportunitiesState = signal<readonly MatchOpportunity[]>([]);
  private readonly tripsState = signal<readonly Trip[]>([]);
  private readonly metricsState = signal<readonly OperationalMetric[]>([]);
  private readonly loadingState = signal(true);
  private readonly errorState = signal('');

  readonly opportunities = this.opportunitiesState.asReadonly();
  readonly trips = this.tripsState.asReadonly();
  readonly operationalMetrics = this.metricsState.asReadonly();
  readonly featuredOpportunity = computed(() => this.opportunitiesState()[0] ?? null);
  readonly secondaryOpportunities = computed(() => this.opportunitiesState().slice(1));
  readonly loading = this.loadingState.asReadonly();
  readonly error = this.errorState.asReadonly();

  constructor() {
    forkJoin({
      matches: this.repository.getMatches(),
      trips: this.repository.getTrips(),
    })
      .pipe(
        catchError(() => {
          this.errorState.set('No pudimos cargar los datos de planificación.');
          return of({ matches: [], trips: [] });
        }),
        takeUntilDestroyed(this.destroyRef),
      )
      .subscribe(({ matches, trips }) => {
        this.opportunitiesState.set(matches);
        this.setPlanningData(matches, trips);
        this.loadingState.set(false);
      });
  }

  getMatch(id: string) {
    return this.repository.getMatch(id);
  }

  recalculate(): void {
    this.loadingState.set(true);
    this.errorState.set('');
    this.api.recalculateMatches().subscribe({
      next: () => this.reload(),
      error: () => {
        this.errorState.set('No pudimos recalcular las oportunidades.');
        this.loadingState.set(false);
      },
    });
  }

  private reload(): void {
    forkJoin({
      matches: this.repository.getMatches(),
      trips: this.repository.getTrips(),
    })
      .pipe(takeUntilDestroyed(this.destroyRef))
      .subscribe({
        next: ({ matches, trips }) => {
          this.opportunitiesState.set(matches);
          this.setPlanningData(matches, trips);
          this.loadingState.set(false);
        },
        error: () => {
          this.errorState.set('No pudimos cargar los datos de planificación.');
          this.loadingState.set(false);
        },
      });
  }

  getById(id: string): MatchOpportunity | null {
    return this.opportunitiesState().find((opportunity) => opportunity.id === id) ?? null;
  }

  private setPlanningData(matches: readonly MatchOpportunity[], trips: readonly Trip[]): void {
    const mappedTrips = trips.map((trip) => {
      const matchCount = matches.filter((item) =>
        item.trip.tripId ? item.trip.tripId === trip.tripId || item.trip.tripId === trip.id : item.trip.id === trip.id,
      ).length;
      return {
        ...trip,
        matchCount,
        status: matchCount > 0 ? 'matched' as const : trip.status === 'pending' ? 'pending' as const : 'unmatched' as const,
      };
    });
    this.tripsState.set(mappedTrips);
    this.metricsState.set([
      { label: 'Viajes programados', value: trips.length },
      { label: 'Oportunidades compatibles', value: matches.length },
      { label: 'Viajes por optimizar', value: mappedTrips.filter((trip) => trip.matchCount > 0).length },
      { label: 'Operaciones activas', value: 0 },
    ]);
  }
}
