import { ChangeDetectionStrategy, Component, computed, inject, signal } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { map } from 'rxjs';
import { MatchService } from '../../core/services/match.service';
import { OfferService } from '../../core/services/offer.service';
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
  private readonly routeId = toSignal(this.route.paramMap.pipe(map((params) => params.get('id') ?? '')), {
    initialValue: this.route.snapshot.paramMap.get('id') ?? '',
  });

  protected readonly opportunity = computed(() => this.matchService.getById(this.routeId()) ?? this.matchService.featuredOpportunity());
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
