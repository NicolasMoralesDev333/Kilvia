import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { RouterLink } from '@angular/router';
import { MatchService } from '../../core/services/match.service';
import { KIcon } from '../../shared/components/k-icon/k-icon';
import { MatchScore } from '../../shared/components/match-score/match-score';

@Component({
  selector: 'app-opportunities-page',
  imports: [RouterLink, KIcon, MatchScore],
  templateUrl: './opportunities-page.html',
  styleUrl: './opportunities-page.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class OpportunitiesPage {
  private readonly matchService = inject(MatchService);
  protected readonly opportunities = this.matchService.opportunities;

  protected tons(value: number): string {
    return value.toFixed(1).replace('.', ',');
  }
}
