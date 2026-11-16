import { ChangeDetectionStrategy, Component, inject, signal } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { catchError, map, of, switchMap, tap } from 'rxjs';
import { MatchService } from '../../core/services/match.service';
import { OfferService } from '../../core/services/offer.service';
import { MatchOpportunity } from '../../core/models/kilvia.models';
import { CapacityFit } from '../../shared/components/capacity-fit/capacity-fit';
import { KIcon } from '../../shared/components/k-icon/k-icon';
import { MatchCriteria } from '../../shared/components/match-criteria/match-criteria';
import { MatchScore } from '../../shared/components/match-score/match-score';
import { RouteMap } from '../../shared/components/route-map/route-map';

@Component({
  selector: 'app-opportunity-detail-page',
  imports: [RouterLink, ReactiveFormsModule, CapacityFit, KIcon, MatchCriteria, MatchScore, RouteMap],
  templateUrl: './opportunity-detail-page.html',
  styleUrl: './opportunity-detail-page.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class OpportunityDetailPage {
  private readonly route = inject(ActivatedRoute);
  private readonly matchService = inject(MatchService);
  private readonly offerService = inject(OfferService);
  private readonly formBuilder = inject(FormBuilder);
  protected readonly loading = signal(true);
  protected readonly loadError = signal('');
  protected readonly opportunity = toSignal(
    this.route.paramMap.pipe(
      map((params) => params.get('id') ?? ''),
      tap(() => {
        this.loading.set(true);
        this.loadError.set('');
      }),
      switchMap((id) =>
        this.matchService.getMatch(id).pipe(
          tap(() => this.loading.set(false)),
          catchError(() => {
            this.loading.set(false);
            this.loadError.set('No pudimos cargar el detalle de esta oportunidad.');
            return of(null as MatchOpportunity | null);
          }),
        ),
      ),
    ),
    { initialValue: null },
  );
  protected readonly offerOpen = signal(false);
  protected readonly offerSaved = signal(false);
  protected readonly decision = signal<'discarded' | 'saved' | null>(null);
  protected readonly offerForm = this.formBuilder.nonNullable.group({
    amount: [0, [Validators.required, Validators.min(1)]],
    message: ['', [Validators.required, Validators.minLength(10), Validators.maxLength(300)]],
  });

  protected openOffer(): void {
    this.offerSaved.set(false);
    this.offerOpen.set(true);
  }

  protected scrollToDecision(): void {
    document.getElementById('decision')?.scrollIntoView({ behavior: 'smooth', block: 'start' });
  }

  protected closeOffer(): void {
    this.offerOpen.set(false);
  }

  protected submitOffer(matchId: string): void {
    if (this.offerForm.invalid) {
      this.offerForm.markAllAsTouched();
      return;
    }
    this.offerService.createOffer({ matchId, ...this.offerForm.getRawValue() }).subscribe(() => {
      this.offerSaved.set(true);
      this.offerForm.reset();
    });
  }

  protected chooseDecision(decision: 'discarded' | 'saved'): void {
    this.decision.set(decision);
  }

  protected tons(value: number): string {
    return value.toFixed(1).replace('.', ',');
  }
}
