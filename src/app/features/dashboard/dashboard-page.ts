import { ChangeDetectionStrategy, Component, computed, inject } from '@angular/core';
import { RouterLink } from '@angular/router';
import { MatchService } from '../../core/services/match.service';
import { KIcon } from '../../shared/components/k-icon/k-icon';
import { MatchScore } from '../../shared/components/match-score/match-score';

@Component({
  selector: 'app-dashboard-page',
  imports: [RouterLink, KIcon, MatchScore],
  templateUrl: './dashboard-page.html',
  styleUrl: './dashboard-page.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class DashboardPage {
  private readonly matchService = inject(MatchService);
  protected readonly metrics = this.matchService.operationalMetrics;
  protected readonly featured = this.matchService.featuredOpportunity;
  protected readonly trips = this.matchService.trips;
  protected readonly totalOpportunities = computed(() => this.matchService.opportunities().length);
  protected readonly loading = this.matchService.loading;
  protected readonly error = this.matchService.error;

  protected tons(value: number): string {
    return value.toFixed(1).replace('.', ',');
  }
}
