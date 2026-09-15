import { ChangeDetectionStrategy, Component, signal } from '@angular/core';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { IconName, KIcon } from '../../components/k-icon/k-icon';

interface NavItem {
  readonly label: string;
  readonly mobileLabel?: string;
  readonly icon: IconName;
  readonly link?: string;
}

@Component({
  selector: 'app-shell',
  imports: [RouterOutlet, RouterLink, RouterLinkActive, KIcon],
  templateUrl: './app-shell.html',
  styleUrl: './app-shell.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AppShell {
  protected readonly moreOpen = signal(false);
  protected readonly desktopNav: readonly NavItem[] = [
    { label: 'Inicio', icon: 'home', link: '/home' },
    { label: 'Viajes', icon: 'truck' },
    { label: 'Oportunidades', icon: 'route', link: '/opportunities' },
    { label: 'Vehículos', icon: 'warehouse' },
    { label: 'Choferes', icon: 'drivers' },
    { label: 'Operaciones', icon: 'briefcase' },
  ];
  protected readonly mobileNav: readonly NavItem[] = [
    { label: 'Inicio', icon: 'home', link: '/home' },
    { label: 'Viajes', icon: 'truck' },
    { label: 'Oportunidades', mobileLabel: 'Oportun.', icon: 'route', link: '/opportunities' },
    { label: 'Operaciones', icon: 'briefcase' },
  ];

  protected toggleMore(): void {
    this.moreOpen.update((open) => !open);
  }
}
