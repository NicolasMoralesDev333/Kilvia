import { ChangeDetectionStrategy, Component, inject, signal } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { ApiCompany } from '../../core/models/api.models';
import { KilviaApiService } from '../../core/services/kilvia-api.service';

@Component({
  selector: 'app-company-page',
  imports: [ReactiveFormsModule],
  templateUrl: './company-page.html',
  styleUrl: './resource-page.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class CompanyPage {
  private readonly api = inject(KilviaApiService);
  private readonly formBuilder = inject(FormBuilder);
  protected readonly company = signal<ApiCompany | null>(null);
  protected readonly loading = signal(true);
  protected readonly saving = signal(false);
  protected readonly error = signal('');
  protected readonly saved = signal(false);
  protected readonly form = this.formBuilder.nonNullable.group({
    name: ['', Validators.required],
    cuit: ['', Validators.required],
    address: [''],
    phone: [''],
    email: ['', Validators.email],
  });

  constructor() {
    this.api.getCompany().subscribe({
      next: (company) => {
        this.company.set(company);
        this.form.patchValue(company);
        this.loading.set(false);
      },
      error: () => {
        this.error.set('No pudimos cargar los datos de la empresa.');
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
    this.saved.set(false);
    this.api.updateCompany(this.form.getRawValue()).subscribe({
      next: (company) => {
        this.company.set(company);
        this.saved.set(true);
        this.saving.set(false);
      },
      error: () => {
        this.error.set('No pudimos guardar los datos de la empresa.');
        this.saving.set(false);
      },
    });
  }
}
