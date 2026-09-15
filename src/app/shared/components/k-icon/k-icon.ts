import { ChangeDetectionStrategy, Component, computed, input } from '@angular/core';
import {
  LucideArrowLeft,
  LucideArrowUpRight,
  LucideBell,
  LucideBriefcaseBusiness,
  LucideCheck,
  LucideChevronRight,
  LucideCircleGauge,
  LucideClock,
  LucideDynamicIcon,
  LucideEllipsis,
  LucideHouse,
  LucideIcon,
  LucideMap,
  LucideMenu,
  LucidePlus,
  LucideRoute,
  LucideSettings,
  LucideTruck,
  LucideUsers,
  LucideWarehouse,
  LucideX,
} from '@lucide/angular';

export type IconName =
  | 'arrow-left'
  | 'arrow-up-right'
  | 'bell'
  | 'briefcase'
  | 'check'
  | 'chevron-right'
  | 'clock'
  | 'close'
  | 'drivers'
  | 'gauge'
  | 'home'
  | 'map'
  | 'menu'
  | 'more'
  | 'plus'
  | 'route'
  | 'settings'
  | 'truck'
  | 'warehouse';

const ICONS: Record<IconName, LucideIcon> = {
  'arrow-left': LucideArrowLeft,
  'arrow-up-right': LucideArrowUpRight,
  bell: LucideBell,
  briefcase: LucideBriefcaseBusiness,
  check: LucideCheck,
  'chevron-right': LucideChevronRight,
  clock: LucideClock,
  close: LucideX,
  drivers: LucideUsers,
  gauge: LucideCircleGauge,
  home: LucideHouse,
  map: LucideMap,
  menu: LucideMenu,
  more: LucideEllipsis,
  plus: LucidePlus,
  route: LucideRoute,
  settings: LucideSettings,
  truck: LucideTruck,
  warehouse: LucideWarehouse,
};

@Component({
  selector: 'app-k-icon',
  imports: [LucideDynamicIcon],
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `<svg [lucideIcon]="icon()" [size]="size()" [strokeWidth]="strokeWidth()"></svg>`,
})
export class KIcon {
  readonly name = input.required<IconName>();
  readonly size = input(20);
  readonly strokeWidth = input(1.8);
  protected readonly icon = computed(() => ICONS[this.name()]);
}
