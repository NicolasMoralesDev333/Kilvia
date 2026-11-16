import { ApplicationConfig, provideBrowserGlobalErrorListeners } from '@angular/core';
import { provideHttpClient, withInterceptors } from '@angular/common/http';
import { provideRouter, withInMemoryScrolling } from '@angular/router';

import { routes } from './app.routes';
import { MockOpportunityRepository, OPPORTUNITY_REPOSITORY } from './core/services/opportunity.repository';
import { authInterceptor } from './core/services/auth.interceptor';

export const appConfig: ApplicationConfig = {
  providers: [
    provideBrowserGlobalErrorListeners(),
    provideHttpClient(withInterceptors([authInterceptor])),
    provideRouter(routes, withInMemoryScrolling({ scrollPositionRestoration: 'top' })),
    MockOpportunityRepository,
    { provide: OPPORTUNITY_REPOSITORY, useExisting: MockOpportunityRepository },
  ],
};
