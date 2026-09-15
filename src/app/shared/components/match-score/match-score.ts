import { ChangeDetectionStrategy, Component, input } from '@angular/core';

@Component({
  selector: 'app-match-score',
  changeDetection: ChangeDetectionStrategy.OnPush,
  host: { class: 'block min-w-0' },
  template: `
    <div class="match-score" [class.match-score--compact]="compact()" [attr.aria-label]="score() + '% de match. ' + label()">
      <div class="match-score__gauge" aria-hidden="true">
        <svg viewBox="0 0 148 92"><path class="gauge-track" pathLength="100" d="M14 78A60 60 0 0 1 134 78" /><path class="gauge-value" pathLength="100" d="M14 78A60 60 0 0 1 134 78" [attr.stroke-dasharray]="score() + ' 100'" /></svg>
        <div class="match-score__value">{{ score() }}<span>%</span></div>
      </div>
      <div class="match-score__meta"><span>MATCH</span><strong>{{ label() }}</strong></div>
    </div>
  `,
  styles: `
    .match-score{display:grid;grid-template-columns:7rem minmax(0,1fr);align-items:end;gap:.85rem;color:#fff}.match-score__gauge{height:4.4rem;position:relative}.match-score svg{position:absolute;inset:0;width:100%;height:100%;overflow:visible}.gauge-track,.gauge-value{fill:none;stroke-width:8;stroke-linecap:square}.gauge-track{stroke:#29405f}.gauge-value{stroke:#22c06e}.match-score__value{position:absolute;right:0;bottom:0;left:0;text-align:center;font-size:2.35rem;line-height:1;font-weight:800;letter-spacing:-.075em}.match-score__value span{margin-left:.08rem;color:#7fe3aa;font-size:.95rem}.match-score__meta{display:flex;flex-direction:column;padding-bottom:.25rem}.match-score__meta span{color:#7fe3aa;font-size:.61rem;font-weight:800;letter-spacing:.14em}.match-score__meta strong{max-width:9rem;margin-top:.25rem;font-size:.78rem;line-height:1.2}.match-score--compact{grid-template-columns:3.1rem minmax(0,1fr);gap:.6rem;color:#0b1d3a}.match-score--compact .match-score__gauge{height:2rem}.match-score--compact .gauge-track{stroke:#cbd6e1}.match-score--compact .gauge-track,.match-score--compact .gauge-value{stroke-width:10}.match-score--compact .match-score__value{bottom:-.08rem;font-size:1rem;letter-spacing:-.03em}.match-score--compact .match-score__value span{color:#087b4b;font-size:.6rem}.match-score--compact .match-score__meta{padding:0}.match-score--compact .match-score__meta span{color:#087b4b}.match-score--compact .match-score__meta strong{overflow:hidden;max-width:7rem;font-size:.65rem;text-overflow:ellipsis;white-space:nowrap}@media(min-width:640px){.match-score{grid-template-columns:8rem minmax(0,1fr);gap:1rem}.match-score__gauge{height:5rem}.match-score__value{font-size:2.75rem}}
  `,
})
export class MatchScore {
  readonly score = input.required<number>();
  readonly label = input('Alta compatibilidad');
  readonly compact = input(false);
}
