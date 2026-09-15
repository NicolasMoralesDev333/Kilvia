import { Routes } from '@angular/router';
import { AppShell } from './shared/layout/app-shell/app-shell';

export const routes: Routes = [
  {
    path: '',
    component: AppShell,
    children: [
      {
        path: 'home',
        loadComponent: () => import('./features/dashboard/dashboard-page').then((component) => component.DashboardPage),
        title: 'Inicio operativo · Kilvia',
      },
      { path: 'dashboard', pathMatch: 'full', redirectTo: 'home' },
      {
        path: 'opportunities',
        loadComponent: () => import('./features/opportunities/opportunities-page').then((component) => component.OpportunitiesPage),
        title: 'Oportunidades · Kilvia',
      },
      {
        path: 'opportunities/:id',
        loadComponent: () => import('./features/opportunities/opportunity-detail-page').then((component) => component.OpportunityDetailPage),
        title: 'Detalle del match · Kilvia',
      },
      { path: '', pathMatch: 'full', redirectTo: 'home' },
    ],
  },
  { path: '**', redirectTo: 'home' },
];
