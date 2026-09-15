import { ChangeDetectionStrategy, Component, computed, inject, signal } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { RouterLink } from '@angular/router';
import { MatchService } from '../../core/services/match.service';
import { TripService } from '../../core/services/trip.service';
import { KIcon } from '../../shared/components/k-icon/k-icon';
import { MatchScore } from '../../shared/components/match-score/match-score';

@Component({
  selector: 'app-dashboard-page',
  imports: [RouterLink, ReactiveFormsModule, KIcon, MatchScore],
  templateUrl: './dashboard-page.html',
  styleUrl: './dashboard-page.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class DashboardPage {
  private readonly matchService = inject(MatchService);
  private readonly tripService = inject(TripService);
  private readonly formBuilder = inject(FormBuilder);

  protected readonly formOpen = signal(false);
  protected readonly formSaved = signal(false);
  protected readonly metrics = this.matchService.operationalMetrics;
  protected readonly featured = this.matchService.featuredOpportunity;
  protected readonly trips = this.matchService.trips;
  protected readonly vehicles = this.tripService.vehicles;
  protected readonly totalOpportunities = computed(() => this.matchService.opportunities().length);
  protected readonly tripForm = this.formBuilder.nonNullable.group({
    origin: ['', [Validators.required, Validators.minLength(2)]],
    destination: ['', [Validators.required, Validators.minLength(2)]],
    departureDate: ['', Validators.required],
    vehicleId: ['', Validators.required],
    freeCapacityTons: [0, [Validators.required, Validators.min(0.1)]],
  });

  protected openTripForm(): void {
    this.formSaved.set(false);
    this.formOpen.set(true);
  }

  protected closeTripForm(): void {
    this.formOpen.set(false);
  }

  protected submitTrip(): void {
    if (this.tripForm.invalid) {
      this.tripForm.markAllAsTouched();
      return;
    }

    this.tripService.createTrip(this.tripForm.getRawValue()).subscribe(() => {
      this.formSaved.set(true);
      this.tripForm.reset();
    });
  }

  protected tons(value: number): string {
    return value.toFixed(1).replace('.', ',');
  }
}
