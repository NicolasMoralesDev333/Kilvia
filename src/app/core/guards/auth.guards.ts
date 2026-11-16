import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { catchError, map, of } from 'rxjs';
import { AuthService } from '../services/auth.service';
import { UserRole } from '../models/api.models';

function homeForRole(role: UserRole): string {
  return role === 'CHOFER' ? '/trips' : '/home';
}

export const authGuard: CanActivateFn = () => {
  const auth = inject(AuthService);
  const router = inject(Router);
  return auth.ensureSession().pipe(
    map((valid) => valid || router.createUrlTree(['/login'])),
    catchError(() => of(router.createUrlTree(['/login'], { queryParams: { session: 'unavailable' } }))),
  );
};

export const guestGuard: CanActivateFn = (_route, state) => {
  const auth = inject(AuthService);
  const router = inject(Router);
  if (state.url.includes('session=unavailable')) return true;
  return auth.ensureSession().pipe(
    map((valid) => (valid ? router.createUrlTree([homeForRole(auth.currentUser()!.role)]) : true)),
    catchError(() => of(router.createUrlTree(['/login'], { queryParams: { session: 'unavailable' } }))),
  );
};

export const roleGuard: CanActivateFn = (route) => {
  const auth = inject(AuthService);
  const router = inject(Router);
  const user = auth.currentUser();
  const roles = route.data['roles'] as UserRole[] | undefined;
  if (!user || !roles?.includes(user.role)) {
    return router.createUrlTree([user ? homeForRole(user.role) : '/login']);
  }
  return true;
};
