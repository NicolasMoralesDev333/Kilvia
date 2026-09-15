import { computed, DestroyRef, inject, Injectable, signal } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { MatchOpportunity, OperationalMetric, Trip } from '../models/kilvia.models';
import { OPPORTUNITY_REPOSITORY } from './opportunity.repository';

@Injectable({ providedIn: 'root' })
export class MatchService {
  private readonly repository = inject(OPPORTUNITY_REPOSITORY);
  private readonly destroyRef = inject(DestroyRef);
  private readonly opportunitiesState = signal<readonly MatchOpportunity[]>([]);
  private readonly tripsState = signal<readonly Trip[]>([]);
  private readonly metricsState = signal<readonly OperationalMetric[]>([]);

  readonly opportunities = this.opportunitiesState.asReadonly();
  readonly trips = this.tripsState.asReadonly();
  readonly operationalMetrics = this.metricsState.asReadonly();
  readonly featuredOpportunity = computed(() => this.opportunitiesState()[0] ?? null);
  readonly secondaryOpportunities = computed(() => this.opportunitiesState().slice(1));

  constructor() {
    this.repository.getMatches().pipe(takeUntilDestroyed(this.destroyRef)).subscribe((items) => this.opportunitiesState.set(items));
    this.repository.getTrips().pipe(takeUntilDestroyed(this.destroyRef)).subscribe((items) => this.tripsState.set(items));
    this.repository
      .getOperationalMetrics()
      .pipe(takeUntilDestroyed(this.destroyRef))
      .subscribe((items) => this.metricsState.set(items));
  }

  getById(id: string): MatchOpportunity | null {
    return this.opportunitiesState().find((opportunity) => opportunity.id === id) ?? null;
  }
}
