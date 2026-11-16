import { DatePipe } from '@angular/common';
import { ChangeDetectionStrategy, Component, inject, signal } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { ApiLoad } from '../../core/models/api.models';
import { KilviaApiService } from '../../core/services/kilvia-api.service';

@Component({
  selector: 'app-loads-page',
  imports: [DatePipe, ReactiveFormsModule],
  templateUrl: './loads-page.html',
  styleUrl: './resource-page.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class LoadsPage {
  private readonly api = inject(KilviaApiService);
  private readonly formBuilder = inject(FormBuilder);
  protected readonly loads = signal<readonly ApiLoad[]>([]);
  protected readonly loading = signal(true);
  protected readonly saving = signal(false);
  protected readonly formOpen = signal(false);
  protected readonly error = signal('');
  protected readonly form = this.formBuilder.nonNullable.group({
    origin: ['', Validators.required],
    destination: ['', Validators.required],
    pickupAt: ['', Validators.required],
    cargoType: ['', Validators.required],
    weightTons: [0, [Validators.required, Validators.min(0.1)]],
    volumeM3: [0, [Validators.required, Validators.min(0.1)]],
    requiredVehicle: ['', Validators.required],
    restrictions: [''],
  });

  constructor() {
    this.refresh();
  }

  protected refresh(): void {
    this.loading.set(true);
    this.api.listLoads().subscribe({
      next: (items) => {
        this.loads.set(items);
        this.error.set('');
        this.loading.set(false);
      },
      error: () => {
        this.error.set('No pudimos cargar las cargas. Intentá nuevamente.');
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
    this.api.createLoad({
      ...value,
      pickupAt: new Date(value.pickupAt).toISOString(),
    }).subscribe({
      next: () => {
        this.formOpen.set(false);
        this.form.reset();
        this.refresh();
        this.saving.set(false);
      },
      error: () => {
        this.error.set('No pudimos registrar la carga. Verificá los datos ingresados.');
        this.saving.set(false);
      },
    });
  }

  protected cancel(load: ApiLoad): void {
    this.api.updateLoad(load.id, { status: 'CANCELADA' }).subscribe({
      next: () => this.refresh(),
      error: () => this.error.set('No pudimos actualizar el estado de la carga.'),
    });
  }
}
