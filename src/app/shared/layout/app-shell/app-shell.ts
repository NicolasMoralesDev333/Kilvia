import { ChangeDetectionStrategy, Component, computed, inject, signal } from '@angular/core';
import { Router, RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { AuthService } from '../../../core/services/auth.service';
import { IconName, KIcon } from '../../components/k-icon/k-icon';

interface NavItem {
  readonly label: string;
  readonly mobileLabel?: string;
  readonly icon: IconName;
  readonly link?: string;
  readonly roles?: readonly ('ADMIN' | 'OPERADOR' | 'CHOFER')[];
}

@Component({
  selector: 'app-shell',
  imports: [RouterOutlet, RouterLink, RouterLinkActive, KIcon],
  templateUrl: './app-shell.html',
  styleUrl: './app-shell.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AppShell {
  private readonly auth = inject(AuthService);
  private readonly router = inject(Router);
  protected readonly moreOpen = signal(false);
  protected readonly user = this.auth.currentUser;
  protected readonly initials = computed(() => {
    const user = this.user();
    return user ? `${user.firstName.charAt(0)}${user.lastName.charAt(0)}`.toUpperCase() : 'K';
  });
  protected readonly desktopNav = computed<readonly NavItem[]>(() => {
    const role = this.user()?.role;
    if (role === 'CHOFER') return [{ label: 'Mis viajes', icon: 'truck', link: '/trips' }];
    return [
      { label: 'Inicio', icon: 'home', link: '/home' },
      { label: 'Viajes', icon: 'truck', link: '/trips' },
      { label: 'Oportunidades', icon: 'route', link: '/opportunities' },
      { label: 'Vehículos', icon: 'warehouse', link: '/vehicles' },
      { label: 'Choferes', icon: 'drivers', link: '/drivers' },
      { label: 'Cargas', icon: 'briefcase', link: '/loads' },
    ];
  });
  protected readonly mobileNav = computed<readonly NavItem[]>(() => {
    const role = this.user()?.role;
    if (role === 'CHOFER') return [{ label: 'Mis viajes', icon: 'truck', link: '/trips' }];
    return [
      { label: 'Inicio', icon: 'home', link: '/home' },
      { label: 'Viajes', icon: 'truck', link: '/trips' },
      { label: 'Oportunidades', mobileLabel: 'Oportun.', icon: 'route', link: '/opportunities' },
      { label: 'Cargas', icon: 'briefcase', link: '/loads' },
    ];
  });

  protected toggleMore(): void {
    this.moreOpen.update((open) => !open);
  }

  protected logout(): void {
    this.auth.clearSession();
    void this.router.navigate(['/login']);
  }
}
