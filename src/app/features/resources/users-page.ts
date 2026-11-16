import { ChangeDetectionStrategy, Component, inject, signal } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { ApiUser } from '../../core/models/api.models';
import { AuthService } from '../../core/services/auth.service';
import { KilviaApiService } from '../../core/services/kilvia-api.service';

@Component({
  selector: 'app-users-page',
  imports: [ReactiveFormsModule],
  templateUrl: './users-page.html',
  styleUrl: './resource-page.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class UsersPage {
  private readonly api = inject(KilviaApiService);
  private readonly auth = inject(AuthService);
  private readonly formBuilder = inject(FormBuilder);
  protected readonly currentUser = this.auth.currentUser;
  protected readonly users = signal<readonly ApiUser[]>([]);
  protected readonly loading = signal(true);
  protected readonly saving = signal(false);
  protected readonly formOpen = signal(false);
  protected readonly error = signal('');
  protected readonly form = this.formBuilder.nonNullable.group({
    firstName: ['', Validators.required],
    lastName: ['', Validators.required],
    email: ['', [Validators.required, Validators.email]],
    password: ['', [Validators.required, Validators.minLength(10)]],
  });

  constructor() {
    this.refresh();
  }

  protected refresh(): void {
    this.loading.set(true);
    this.api.listUsers().subscribe({
      next: (items) => {
        this.users.set(items);
        this.error.set('');
        this.loading.set(false);
      },
      error: () => {
        this.error.set('No pudimos cargar los usuarios.');
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
    this.api.createOperator(this.form.getRawValue()).subscribe({
      next: () => {
        this.formOpen.set(false);
        this.form.reset();
        this.refresh();
        this.saving.set(false);
      },
      error: () => {
        this.error.set('No pudimos crear el operador. Verificá el email y la contraseña.');
        this.saving.set(false);
      },
    });
  }

  protected toggleStatus(user: ApiUser): void {
    const status = user.status === 'INACTIVO' ? 'ACTIVO' : 'INACTIVO';
    this.api.updateUser(user.id, { status }).subscribe({
      next: () => this.refresh(),
      error: () => this.error.set('No pudimos actualizar el estado del usuario.'),
    });
  }
}
