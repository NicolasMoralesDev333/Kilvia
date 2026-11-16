import { DatePipe } from '@angular/common';
import { ChangeDetectionStrategy, Component, inject, signal } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { ApiDriver } from '../../core/models/api.models';
import { AuthService } from '../../core/services/auth.service';
import { KilviaApiService } from '../../core/services/kilvia-api.service';

@Component({
  selector: 'app-drivers-page',
  imports: [DatePipe, ReactiveFormsModule],
  templateUrl: './drivers-page.html',
  styleUrl: './resource-page.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class DriversPage {
  private readonly api = inject(KilviaApiService);
  private readonly auth = inject(AuthService);
  private readonly formBuilder = inject(FormBuilder);
  protected readonly canCreate = this.auth.currentUser()?.role === 'ADMIN';
  protected readonly drivers = signal<readonly ApiDriver[]>([]);
  protected readonly loading = signal(true);
  protected readonly saving = signal(false);
  protected readonly formOpen = signal(false);
  protected readonly error = signal('');
  protected readonly form = this.formBuilder.nonNullable.group({
    firstName: ['', Validators.required],
    lastName: ['', Validators.required],
    email: ['', [Validators.required, Validators.email]],
    password: ['', [Validators.required, Validators.minLength(10)]],
    dni: ['', Validators.required],
    licenseNumber: ['', Validators.required],
    licenseCategory: ['', Validators.required],
    licenseExpiration: ['', Validators.required],
    phone: [''],
    availability: ['DISPONIBLE', Validators.required],
  });

  constructor() {
    this.refresh();
  }

  protected refresh(): void {
    this.loading.set(true);
    this.api.listDrivers().subscribe({
      next: (items) => {
        this.drivers.set(items);
        this.error.set('');
        this.loading.set(false);
      },
      error: () => {
        this.error.set('No pudimos cargar los choferes. Intentá nuevamente.');
        this.loading.set(false);
      },
    });
  }

  protected save(): void {
    if (this.form.invalid || this.saving()) {
      this.form.markAllAsTouched();
      return;
    }
    const value = this.form.getRawValue();
    this.saving.set(true);
    this.api.createDriver({
      ...value,
      licenseExpiration: new Date(`${value.licenseExpiration}T00:00:00`).toISOString(),
    }).subscribe({
      next: () => {
        this.formOpen.set(false);
        this.form.reset({ availability: 'DISPONIBLE' });
        this.refresh();
        this.saving.set(false);
      },
      error: () => {
        this.error.set('No pudimos crear el chofer. Revisá email, DNI y licencia.');
        this.saving.set(false);
      },
    });
  }

  protected toggleStatus(driver: ApiDriver): void {
    const status = driver.driverStatus === 'INACTIVO' ? 'ACTIVO' : 'INACTIVO';
    this.api.updateDriver(driver.id, { status }).subscribe({
      next: () => this.refresh(),
      error: () => this.error.set('No pudimos actualizar el estado del chofer.'),
    });
  }
}
