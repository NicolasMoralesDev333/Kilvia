import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';

export interface CreateOfferRequest {
  readonly matchId: string;
  readonly amount: number;
  readonly message: string;
}

@Injectable({ providedIn: 'root' })
export class OfferService {
  createOffer(request: CreateOfferRequest): Observable<{ readonly id: string }> {
    // Adaptador demo. La futura implementación reemplaza este retorno por POST /api/v1/matches/:matchId/offers.
    void request;
    return of({ id: 'OF-DEMO' });
  }
}
