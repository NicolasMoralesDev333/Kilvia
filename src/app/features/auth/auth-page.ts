import { ChangeDetectionStrategy, Component, computed, inject, signal } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { HttpErrorResponse } from '@angular/common/http';
import { AuthService } from '../../core/services/auth.service';

@Component({
  selector: 'app-auth-page',
  imports: [ReactiveFormsModule, RouterLink],
  templateUrl: './auth-page.html',
  styleUrl: './auth-page.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AuthPage {
  private readonly auth = inject(AuthService);
  private readonly router = inject(Router);
  private readonly route = inject(ActivatedRoute);
  private readonly formBuilder = inject(FormBuilder);

  protected readonly isRegister = this.route.snapshot.data['mode'] === 'register';
  protected readonly heading = computed(() => (this.isRegister ? 'Tu empresa, en movimiento.' : 'Volvé a la ruta.'));
  protected readonly loading = signal(false);
  protected readonly errorMessage = signal('');

  protected readonly form = this.formBuilder.nonNullable.group({
    companyName: ['', this.isRegister ? [Validators.required, Validators.minLength(2)] : []],
    cuit: ['', this.isRegister ? [Validators.required, Validators.minLength(8)] : []],
    firstName: ['', this.isRegister ? Validators.required : []],
    lastName: ['', this.isRegister ? Validators.required : []],
    email: ['', [Validators.required, Validators.email]],
    password: ['', [Validators.required, Validators.minLength(10)]],
  });

  constructor() {
    if (this.route.snapshot.queryParamMap.get('session') === 'unavailable') {
      this.errorMessage.set('No pudimos validar tu sesión. Iniciá sesión nuevamente cuando la API esté disponible.');
    }
  }

  protected submit(): void {
    if (this.form.invalid || this.loading()) {
      this.form.markAllAsTouched();
      return;
    }

    this.loading.set(true);
    this.errorMessage.set('');
    const values = this.form.getRawValue();
    const request = this.isRegister
      ? this.auth.register(values)
      : this.auth.login({ email: values.email, password: values.password });

    request.subscribe({
      next: (user) => {
        void this.router.navigateByUrl(user.role === 'CHOFER' ? '/trips' : '/home');
      },
      error: (error: unknown) => {
        this.loading.set(false);
        this.errorMessage.set(this.errorMessageFrom(error));
      },
      complete: () => this.loading.set(false),
    });
  }

  private errorMessageFrom(error: unknown): string {
    if (error instanceof HttpErrorResponse && error.status === 409) {
      return 'Ya existe una cuenta con esos datos. Revisá el email o el CUIT.';
    }
    if (error instanceof HttpErrorResponse && error.status === 0) {
      return 'No se pudo conectar con Kilvia. Verificá que la API esté disponible.';
    }
    return this.isRegister
      ? 'No pudimos crear la cuenta. Revisá los datos e intentá nuevamente.'
      : 'Email o contraseña incorrectos.';
  }
}
