import { ChangeDetectionStrategy, Component, input } from '@angular/core';
import { MatchCriterion } from '../../../core/models/kilvia.models';
import { KIcon } from '../k-icon/k-icon';

@Component({
  selector: 'app-match-criteria',
  imports: [KIcon],
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <section aria-labelledby="criteria-title">
      <div class="section-heading"><p class="eyebrow">Validación operativa</p><h2 id="criteria-title">¿Por qué es compatible?</h2><p>Señales principales para evaluar este match.</p></div>
      <ul class="criteria-list">
        @for (criterion of criteria(); track criterion.label) {
          <li [class.warning]="criterion.tone === 'warning'"><span class="criterion-icon"><app-k-icon [name]="criterion.tone === 'positive' ? 'check' : 'route'" [size]="14" /></span><span>{{ criterion.label }}</span><strong>{{ criterion.value }}</strong></li>
        }
      </ul>
    </section>
  `,
  styles: `
    .section-heading{padding-bottom:1.25rem;border-bottom:2px solid #0b1d3a}.section-heading h2{margin:.3rem 0 0;color:#0b1d3a;font-size:1.4rem;line-height:1.25;font-weight:800;letter-spacing:-.03em}.section-heading>p:last-child{max-width:18rem;margin:.55rem 0 0;color:#667085;font-size:.72rem;line-height:1.5}.eyebrow{margin:0;color:#1e6bff;font-size:.66rem;font-weight:800;letter-spacing:.1em;text-transform:uppercase}.criteria-list{margin:0;padding:0;list-style:none}.criteria-list li{display:grid;grid-template-columns:1.4rem 1fr auto;align-items:center;gap:.65rem;min-height:3.45rem;border-bottom:1px solid #bdcbd9;color:#344054;font-size:.76rem}.criteria-list strong{color:#087b4b;font-size:.7rem}.criterion-icon{display:grid;width:1.3rem;height:1.3rem;place-items:center;color:#087b4b;border:1px solid #75d9a0}.criteria-list li.warning strong,.criteria-list li.warning .criterion-icon{color:#a15c07}.criteria-list li.warning .criterion-icon{border-color:#f4b64c}@media(min-width:768px){.criteria-list li{min-height:3.75rem}}
  `,
})
export class MatchCriteria {
  readonly criteria = input.required<readonly MatchCriterion[]>();
}
