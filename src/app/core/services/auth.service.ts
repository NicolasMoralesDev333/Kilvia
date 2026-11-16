import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Injectable, computed, inject, signal } from '@angular/core';
import { Observable, catchError, map, of, tap, throwError } from 'rxjs';
import { AuthResult, AuthUser, LoginRequest, RegisterRequest } from '../models/api.models';

const API_BASE_URL = '/api/v1';
const TOKEN_KEY = 'kilvia_access_token';

@Injectable({ providedIn: 'root' })
export class AuthService {
  private readonly http = inject(HttpClient);
  private readonly userState = signal<AuthUser | null>(null);

  readonly currentUser = this.userState.asReadonly();
  readonly isAuthenticated = computed(() => this.userState() !== null);
  readonly role = computed(() => this.userState()?.role ?? null);

  login(input: LoginRequest): Observable<AuthUser> {
    return this.http.post<AuthResult>(`${API_BASE_URL}/auth/login`, input).pipe(
      tap((result) => this.saveSession(result)),
      map((result) => result.user),
    );
  }

  register(input: RegisterRequest): Observable<AuthUser> {
    return this.http.post<AuthResult>(`${API_BASE_URL}/auth/register`, input).pipe(
      tap((result) => this.saveSession(result)),
      map((result) => result.user),
    );
  }

  ensureSession(): Observable<boolean> {
    if (this.userState()) {
      return of(true);
    }
    if (!this.token()) {
      return of(false);
    }
    return this.http.get<AuthUser>(`${API_BASE_URL}/auth/me`).pipe(
      tap((user) => this.userState.set(user)),
      map(() => true),
      catchError((error: unknown) => {
        if (error instanceof HttpErrorResponse && error.status === 401) {
          this.clearSession();
          return of(false);
        }
        return throwError(() => error);
      }),
    );
  }

  token(): string | null {
    return localStorage.getItem(TOKEN_KEY);
  }

  clearSession(): void {
    localStorage.removeItem(TOKEN_KEY);
    this.userState.set(null);
  }

  private saveSession(result: AuthResult): void {
    localStorage.setItem(TOKEN_KEY, result.token);
    this.userState.set(result.user);
  }
}
