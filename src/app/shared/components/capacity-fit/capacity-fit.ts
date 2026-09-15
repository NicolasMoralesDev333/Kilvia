import { ChangeDetectionStrategy, Component, computed, input } from '@angular/core';
import { VehicleCapacity } from '../../../core/models/kilvia.models';

@Component({
  selector: 'app-capacity-fit',
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <section class="capacity" aria-labelledby="capacity-title">
      <div class="capacity__intro"><p class="eyebrow">Ajuste de capacidad</p><h3 id="capacity-title">La carga ocupa {{ capacity().utilizationPercent }}% de la capacidad libre.</h3><p>La visualización compara el peso de la carga con el espacio disponible antes del match.</p></div>
      <div class="capacity__visual" role="img" [attr.aria-label]="capacity().loadTons + ' toneladas de carga sobre ' + capacity().freeTons + ' toneladas disponibles'">
        <div class="capacity__numbers"><div><span>Carga</span><strong>{{ formattedLoad() }} t</strong></div><div><span>Capacidad libre</span><strong>{{ formattedFree() }} t</strong></div></div>
        <div class="capacity__bay"><div class="capacity__fill" [style.width.%]="capacity().utilizationPercent"><span>{{ capacity().utilizationPercent }}% EN USO</span></div><div class="capacity__remaining"><span>{{ remainingTons() }} t LIBRES</span></div><i [style.left.%]="capacity().utilizationPercent"></i></div>
        <div class="capacity__scale"><span>0 t</span><span>Espacio utilizado vs. disponible</span><span>{{ formattedFree() }} t</span></div>
      </div>
    </section>
  `,
  styles: `
    .capacity{display:grid;gap:1.5rem;padding:.25rem 0}.capacity__intro h3{max-width:29rem;margin:.4rem 0 0;color:#0b1d3a;font-size:1.35rem;line-height:1.3;font-weight:800;letter-spacing:-.03em}.capacity__intro>p:last-child{max-width:30rem;margin:.65rem 0 0;color:#667085;font-size:.78rem;line-height:1.55}.eyebrow{margin:0;color:#087b4b;font-size:.67rem;font-weight:800;letter-spacing:.1em;text-transform:uppercase}.capacity__numbers{display:flex;justify-content:space-between;margin-bottom:.7rem}.capacity__numbers div{display:flex;flex-direction:column}.capacity__numbers div:last-child{text-align:right}.capacity__numbers span{color:#667085;font-size:.65rem}.capacity__numbers strong{color:#0b1d3a;font-size:.95rem}.capacity__bay{display:flex;height:4.75rem;position:relative;overflow:hidden;border:2px solid #0b1d3a;background:repeating-linear-gradient(90deg,#e7edf4 0,#e7edf4 calc(25% - 1px),#cbd6e1 25%,#e7edf4 calc(25% + 1px))}.capacity__fill,.capacity__remaining{display:flex;align-items:center;justify-content:center}.capacity__fill{min-width:3rem;color:#fff;background:repeating-linear-gradient(-45deg,#159b59 0,#159b59 10px,#22c06e 10px,#22c06e 20px)}.capacity__remaining{flex:1;color:#475467}.capacity__fill span,.capacity__remaining span{font-size:.58rem;font-weight:800;letter-spacing:.06em}.capacity__bay i{position:absolute;inset:-.4rem auto -.4rem;width:3px;background:#0b1d3a}.capacity__scale{display:flex;justify-content:space-between;gap:.5rem;margin-top:.5rem;color:#667085;font-size:.6rem}.capacity__scale span:nth-child(2){text-align:center}@media(min-width:780px){.capacity{grid-template-columns:minmax(14rem,.75fr) minmax(20rem,1.25fr);align-items:center;gap:2.5rem;padding:1rem 0}.capacity__bay{height:5.25rem}}
  `,
})
export class CapacityFit {
  readonly capacity = input.required<VehicleCapacity>();
  protected readonly formattedLoad = computed(() => this.format(this.capacity().loadTons));
  protected readonly formattedFree = computed(() => this.format(this.capacity().freeTons));

  protected remainingTons(): string {
    return this.format(this.capacity().freeTons - this.capacity().loadTons);
  }

  private format(value: number): string {
    return value.toFixed(1).replace('.', ',');
  }
}
