import { ChangeDetectionStrategy, Component, inject, signal } from '@angular/core';
import { DatePipe } from '@angular/common';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { ApiDriver, ApiTrip, ApiVehicle } from '../../core/models/api.models';
import { AuthService } from '../../core/services/auth.service';
import { KilviaApiService } from '../../core/services/kilvia-api.service';

@Component({
  selector: 'app-trips-page',
  imports: [DatePipe, ReactiveFormsModule],
  templateUrl: './trips-page.html',
  styleUrl: '../resources/resource-page.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class TripsPage {
  private readonly api = inject(KilviaApiService);
  private readonly auth = inject(AuthService);
  private readonly formBuilder = inject(FormBuilder);
  protected readonly user = this.auth.currentUser;
  protected readonly canManage = this.user()?.role !== 'CHOFER';
  protected readonly trips = signal<readonly ApiTrip[]>([]);
  protected readonly drivers = signal<readonly ApiDriver[]>([]);
  protected readonly vehicles = signal<readonly ApiVehicle[]>([]);
  protected readonly loading = signal(true);
  protected readonly saving = signal(false);
  protected readonly formOpen = signal(false);
  protected readonly error = signal('');
  protected readonly saved = signal(false);
  protected readonly form = this.formBuilder.nonNullable.group({
    driverId: ['', Validators.required],
    vehicleId: ['', Validators.required],
    origin: ['', [Validators.required, Validators.minLength(2)]],
    destination: ['', [Validators.required, Validators.minLength(2)]],
    departureAt: ['', Validators.required],
    estimatedArrivalAt: ['', Validators.required],
    availableWeightTons: [0, [Validators.required, Validators.min(0)]],
    availableVolumeM3: [0, [Validators.required, Validators.min(0)]],
  });

  constructor() {
    this.refresh();
    if (this.canManage) {
      this.api.listDrivers().subscribe({ next: (items) => this.drivers.set(items), error: () => this.error.set('No pudimos cargar los choferes.') });
      this.api.listVehicles().subscribe({ next: (items) => this.vehicles.set(items), error: () => this.error.set('No pudimos cargar los vehículos.') });
    }
  }

  protected refresh(): void {
    this.loading.set(true);
    this.api.listTrips().subscribe({
      next: (items) => {
        this.trips.set(items);
        this.error.set('');
        this.loading.set(false);
      },
      error: () => {
        this.error.set('No pudimos cargar los viajes. Intentá nuevamente.');
        this.loading.set(false);
      },
    });
  }

  protected save(): void {
    if (this.form.invalid || this.saving()) {
      this.form.markAllAsTouched();
      return;
    }
    this.saving.set(true);
    this.error.set('');
    const values = this.form.getRawValue();
    this.api.createTrip({
      ...values,
      departureAt: new Date(values.departureAt).toISOString(),
      estimatedArrivalAt: new Date(values.estimatedArrivalAt).toISOString(),
    }).subscribe({
      next: () => {
        this.saved.set(true);
        this.formOpen.set(false);
        this.form.reset();
        this.refresh();
        this.saving.set(false);
      },
      error: () => {
        this.error.set('No pudimos guardar el viaje. Verificá los datos y las asignaciones.');
        this.saving.set(false);
      },
    });
  }

  protected cancel(trip: ApiTrip): void {
    this.api.updateTrip(trip.id, { status: 'CANCELADO' }).subscribe({
      next: () => this.refresh(),
      error: () => this.error.set('No pudimos actualizar el estado del viaje.'),
    });
  }
}
