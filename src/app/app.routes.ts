import { Routes } from '@angular/router';
import { AppShell } from './shared/layout/app-shell/app-shell';
import { authGuard, guestGuard, roleGuard } from './core/guards/auth.guards';

export const routes: Routes = [
  {
    path: 'login',
    canActivate: [guestGuard],
    loadComponent: () => import('./features/auth/auth-page').then((component) => component.AuthPage),
    data: { mode: 'login' },
    title: 'Iniciar sesión · Kilvia',
  },
  {
    path: 'register',
    canActivate: [guestGuard],
    loadComponent: () => import('./features/auth/auth-page').then((component) => component.AuthPage),
    data: { mode: 'register' },
    title: 'Crear cuenta · Kilvia',
  },
  {
    path: '',
    component: AppShell,
    canActivate: [authGuard],
    children: [
      {
        path: 'home',
        canActivate: [roleGuard],
        data: { roles: ['ADMIN', 'OPERADOR'] },
        loadComponent: () => import('./features/dashboard/dashboard-page').then((component) => component.DashboardPage),
        title: 'Inicio operativo · Kilvia',
      },
      { path: 'dashboard', pathMatch: 'full', redirectTo: 'home' },
      {
        path: 'opportunities',
        canActivate: [roleGuard],
        data: { roles: ['ADMIN', 'OPERADOR'] },
        loadComponent: () => import('./features/opportunities/opportunities-page').then((component) => component.OpportunitiesPage),
        title: 'Oportunidades · Kilvia',
      },
      {
        path: 'opportunities/:id',
        canActivate: [roleGuard],
        data: { roles: ['ADMIN', 'OPERADOR'] },
        loadComponent: () => import('./features/opportunities/opportunity-detail-page').then((component) => component.OpportunityDetailPage),
        title: 'Detalle del match · Kilvia',
      },
      {
        path: 'trips',
        canActivate: [roleGuard],
        data: { roles: ['ADMIN', 'OPERADOR', 'CHOFER'] },
        loadComponent: () => import('./features/trips/trips-page').then((component) => component.TripsPage),
        title: 'Viajes · Kilvia',
      },
      {
        path: 'vehicles',
        canActivate: [roleGuard],
        data: { roles: ['ADMIN', 'OPERADOR'] },
        loadComponent: () => import('./features/resources/vehicles-page').then((component) => component.VehiclesPage),
        title: 'Vehículos · Kilvia',
      },
      {
        path: 'drivers',
        canActivate: [roleGuard],
        data: { roles: ['ADMIN', 'OPERADOR'] },
        loadComponent: () => import('./features/resources/drivers-page').then((component) => component.DriversPage),
        title: 'Choferes · Kilvia',
      },
      {
        path: 'loads',
        canActivate: [roleGuard],
        data: { roles: ['ADMIN', 'OPERADOR'] },
        loadComponent: () => import('./features/resources/loads-page').then((component) => component.LoadsPage),
        title: 'Cargas · Kilvia',
      },
      {
        path: 'users',
        canActivate: [roleGuard],
        data: { roles: ['ADMIN'] },
        loadComponent: () => import('./features/resources/users-page').then((component) => component.UsersPage),
        title: 'Usuarios · Kilvia',
      },
      {
        path: 'company',
        canActivate: [roleGuard],
        data: { roles: ['ADMIN'] },
        loadComponent: () => import('./features/resources/company-page').then((component) => component.CompanyPage),
        title: 'Empresa · Kilvia',
      },
      { path: '', pathMatch: 'full', redirectTo: 'home' },
    ],
  },
  { path: '**', redirectTo: 'home' },
];
