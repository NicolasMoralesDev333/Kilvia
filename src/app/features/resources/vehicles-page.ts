import { ChangeDetectionStrategy, Component, inject, signal } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { ApiVehicle, ApiVehicleType } from '../../core/models/api.models';
import { AuthService } from '../../core/services/auth.service';
import { KilviaApiService } from '../../core/services/kilvia-api.service';

@Component({
  selector: 'app-vehicles-page',
  imports: [ReactiveFormsModule],
  templateUrl: './vehicles-page.html',
  styleUrl: './resource-page.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class VehiclesPage {
  private readonly api = inject(KilviaApiService);
  private readonly auth = inject(AuthService);
  private readonly formBuilder = inject(FormBuilder);
  protected readonly canCreate = this.auth.currentUser()?.role === 'ADMIN' || this.auth.currentUser()?.role === 'OPERADOR';
  protected readonly vehicles = signal<readonly ApiVehicle[]>([]);
  protected readonly types = signal<readonly ApiVehicleType[]>([]);
  protected readonly loading = signal(true);
  protected readonly saving = signal(false);
  protected readonly formOpen = signal(false);
  protected readonly error = signal('');
  protected readonly form = this.formBuilder.nonNullable.group({
    typeId: ['', Validators.required],
    plate: ['', Validators.required],
    make: ['', Validators.required],
    model: ['', Validators.required],
    weightTons: [0, [Validators.required, Validators.min(0.1)]],
    volumeM3: [0, [Validators.required, Validators.min(0.1)]],
  });

  constructor() {
    this.refresh();
    this.api.listVehicleTypes().subscribe({
      next: (items) => this.types.set(items),
      error: () => this.error.set('No pudimos cargar los tipos de vehículo.'),
    });
  }

  protected refresh(): void {
    this.loading.set(true);
    this.api.listVehicles().subscribe({
      next: (items) => {
        this.vehicles.set(items);
        this.error.set('');
        this.loading.set(false);
      },
      error: () => {
        this.error.set('No pudimos cargar los vehículos. Intentá nuevamente.');
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
    this.api.createVehicle(this.form.getRawValue()).subscribe({
      next: () => {
        this.formOpen.set(false);
        this.form.reset();
        this.refresh();
        this.saving.set(false);
      },
      error: () => {
        this.error.set('No pudimos guardar el vehículo. Revisá la patente y los datos.');
        this.saving.set(false);
      },
    });
  }

  protected toggleMaintenance(vehicle: ApiVehicle): void {
    const status = vehicle.status === 'MANTENIMIENTO' ? 'DISPONIBLE' : 'MANTENIMIENTO';
    this.api.updateVehicle(vehicle.id, { status }).subscribe({
      next: () => this.refresh(),
      error: () => this.error.set('No pudimos actualizar el estado del vehículo.'),
    });
  }
}
