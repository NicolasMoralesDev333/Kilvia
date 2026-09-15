import { ApplicationConfig, provideBrowserGlobalErrorListeners } from '@angular/core';
import { provideHttpClient } from '@angular/common/http';
import { provideRouter, withInMemoryScrolling } from '@angular/router';

import { routes } from './app.routes';
import { MockOpportunityRepository, OPPORTUNITY_REPOSITORY } from './core/services/opportunity.repository';

export const appConfig: ApplicationConfig = {
  providers: [
    provideBrowserGlobalErrorListeners(),
    provideHttpClient(),
    provideRouter(routes, withInMemoryScrolling({ scrollPositionRestoration: 'top' })),
    MockOpportunityRepository,
    { provide: OPPORTUNITY_REPOSITORY, useExisting: MockOpportunityRepository },
  ],
};
