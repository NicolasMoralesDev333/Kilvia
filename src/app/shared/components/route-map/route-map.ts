import { afterNextRender, ChangeDetectionStrategy, Component, effect, ElementRef, input, OnDestroy, ViewEncapsulation, viewChild } from '@angular/core';
import type { Feature, LineString } from 'geojson';
import { GeoJSONSource, LngLatBounds, Map as MapLibreMap, Marker, NavigationControl, StyleSpecification } from 'maplibre-gl';
import { MatchOpportunity, RouteCoordinate } from '../../../core/models/kilvia.models';

@Component({
  selector: 'app-route-map',
  template: `
    <section class="kilvia-map" aria-label="Mapa de relación entre el viaje y la carga">
      <div #map class="kilvia-map__canvas"></div>
      <div class="kilvia-map__legend" aria-label="Referencias del mapa">
        <span><i class="planned"></i>Ruta planificada</span>
        <span><i class="utilized"></i>Tramo aprovechado</span>
        <span><i class="detour"></i>Desvío</span>
      </div>
      <div class="kilvia-map__demo">Geometría demo<span class="kilvia-map__demo-detail"> · Preparado para OSRM</span></div>
    </section>
  `,
  styles: `
    .kilvia-map{height:100%;min-height:25rem;position:relative;overflow:hidden;background:#dfe8f2}.kilvia-map__canvas{position:absolute;inset:0}.kilvia-map__legend{position:absolute;z-index:2;right:.75rem;bottom:.75rem;left:.75rem;display:flex;flex-wrap:wrap;gap:.55rem 1rem;padding:.75rem;color:#344054;background:rgba(255,255,255,.94);font-size:.72rem;font-weight:700}.kilvia-map__legend span{display:flex;align-items:center;gap:.4rem}.kilvia-map__legend i{display:block;width:1.25rem;height:3px}.kilvia-map__legend .planned{background:#1e6bff}.kilvia-map__legend .utilized{background:#22c06e}.kilvia-map__legend .detour{background:#f79009}.kilvia-map__demo{position:absolute;z-index:2;top:.75rem;left:.75rem;padding:.45rem .6rem;color:#526173;background:rgba(255,255,255,.9);font-size:.65rem;font-weight:700}.kilvia-map__demo-detail{display:none}.kilvia-map .maplibregl-ctrl-top-right{top:.75rem;right:.75rem}.kilvia-map .maplibregl-ctrl-group{border-radius:0;box-shadow:0 2px 10px rgba(11,29,58,.16)}.kilvia-map .maplibregl-ctrl-group button{width:2.75rem;height:2.75rem}.kilvia-map .maplibregl-ctrl-attrib{font-size:9px}.kilvia-map__marker-wrap{width:1rem;height:1rem;position:relative}.kilvia-map__marker{position:absolute;min-width:0;max-width:5.5rem;border-left:3px solid #1e6bff;padding:.38rem .45rem;color:#0b1d3a;background:rgba(255,255,255,.96);box-shadow:0 3px 10px rgba(11,29,58,.16);font-family:Manrope,sans-serif;white-space:nowrap}.kilvia-map__marker--load{border-color:#22c06e}.kilvia-map__marker--origin{right:1.15rem;top:.95rem}.kilvia-map__marker--cargo{left:1.15rem;top:.95rem}.kilvia-map__marker--destination{left:1.15rem;bottom:.95rem}.kilvia-map__marker small{display:none;color:#667085;font-size:.55rem;font-weight:800;letter-spacing:.06em}.kilvia-map__marker strong{display:block;font-size:.64rem}.kilvia-map__dot{display:block;width:1rem;height:1rem;border:4px solid #fff;border-radius:50%;background:#1e6bff;box-shadow:0 0 0 2px #1e6bff}.kilvia-map__dot--load{background:#22c06e;box-shadow:0 0 0 2px #22c06e}@media(min-width:768px){.kilvia-map{min-height:34rem}.kilvia-map__legend{right:auto;left:1.25rem;bottom:1.25rem}.kilvia-map__demo{top:1.25rem;left:1.25rem}.kilvia-map__demo-detail{display:inline}.kilvia-map__marker{min-width:7rem;max-width:none;border-left-width:4px;padding:.45rem .55rem;white-space:normal}.kilvia-map__marker--origin{top:1.35rem}.kilvia-map__marker--cargo{right:1.15rem;left:auto;top:1.35rem}.kilvia-map__marker--destination{bottom:1.35rem}.kilvia-map__marker small{display:block}.kilvia-map__marker strong{margin-top:.12rem;font-size:.68rem}}
  `,
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class RouteMap implements OnDestroy {
  readonly opportunity = input.required<MatchOpportunity>();
  private readonly mapHost = viewChild.required<ElementRef<HTMLDivElement>>('map');
  private map?: MapLibreMap;
  private markers: Marker[] = [];
  private resizeObserver?: ResizeObserver;

  constructor() {
    effect(() => {
      const opportunity = this.opportunity();
      if (this.map?.loaded()) queueMicrotask(() => this.drawRoute(opportunity));
    });
    afterNextRender(() => this.initializeMap());
  }

  ngOnDestroy(): void {
    this.resizeObserver?.disconnect();
    this.markers.forEach((marker) => marker.remove());
    this.map?.remove();
  }

  private initializeMap(): void {
    const style: StyleSpecification = {
      version: 8,
      sources: {
        osm: {
          type: 'raster',
          tiles: ['https://tile.openstreetmap.org/{z}/{x}/{y}.png'],
          tileSize: 256,
          attribution: '© OpenStreetMap contributors',
        },
      },
      layers: [
        { id: 'background', type: 'background', paint: { 'background-color': '#dfe8f2' } },
        { id: 'osm', type: 'raster', source: 'osm', paint: { 'raster-saturation': -0.85, 'raster-opacity': 0.68 } },
      ],
    };

    this.map = new MapLibreMap({
      container: this.mapHost().nativeElement,
      style,
      center: [-64.5, -34],
      zoom: 4,
      attributionControl: { compact: true },
    });
    this.map.addControl(new NavigationControl({ showCompass: false }), 'top-right');
    this.map.on('load', () => {
      const opportunity = this.opportunity();
      this.addRouteLayers(opportunity);
      this.drawRoute(opportunity);
    });
    this.resizeObserver = new ResizeObserver(() => this.map?.resize());
    this.resizeObserver.observe(this.mapHost().nativeElement);
  }

  private addRouteLayers(opportunity: MatchOpportunity): void {
    if (!this.map) return;
    this.addLine('planned', opportunity.geometry.planned, '#1E6BFF', 10);
    this.addLine('utilized', opportunity.geometry.utilized, '#22C06E', 11);
    this.addLine('detour', opportunity.geometry.detour, '#F79009', 6, [1.5, 1.5]);
  }

  private addLine(id: string, coordinates: readonly RouteCoordinate[], color: string, width: number, dasharray?: number[]): void {
    if (!this.map) return;
    this.map.addSource(id, { type: 'geojson', data: this.lineFeature(coordinates) });
    this.map.addLayer({
      id,
      type: 'line',
      source: id,
      layout: { 'line-cap': 'round', 'line-join': 'round' },
      paint: { 'line-color': color, 'line-width': width, ...(dasharray ? { 'line-dasharray': dasharray } : {}) },
    });
  }

  private drawRoute(opportunity: MatchOpportunity): void {
    if (!this.map?.getSource('planned')) return;
    const { planned, utilized, detour } = opportunity.geometry;
    (this.map.getSource('planned') as GeoJSONSource).setData(this.lineFeature(planned));
    (this.map.getSource('utilized') as GeoJSONSource).setData(this.lineFeature(utilized));
    (this.map.getSource('detour') as GeoJSONSource).setData(this.lineFeature(detour));

    this.markers.forEach((marker) => marker.remove());
    this.markers = [
      this.createMarker(planned[0], 'Origen compartido', opportunity.trip.origin.city, 'origin'),
      this.createMarker(utilized[utilized.length - 1], 'Destino carga', opportunity.load.destination.city, 'cargo'),
      this.createMarker(planned[planned.length - 1], 'Fin del viaje', opportunity.trip.destination.city, 'destination'),
    ];

    const bounds = new LngLatBounds();
    [...planned, ...detour].forEach((coordinate) => bounds.extend([coordinate[0], coordinate[1]]));
    const padding = this.mapHost().nativeElement.clientWidth < 600 ? 56 : 120;
    this.map.fitBounds(bounds, { padding, duration: 0, maxZoom: 7 });
  }

  private createMarker(coordinate: RouteCoordinate, label: string, city: string, role: 'origin' | 'cargo' | 'destination'): Marker {
    const wrapper = document.createElement('div');
    const dot = document.createElement('span');
    const text = document.createElement('span');
    const eyebrow = document.createElement('small');
    const cityLabel = document.createElement('strong');
    const load = role !== 'destination';
    wrapper.className = 'kilvia-map__marker-wrap';
    dot.className = `kilvia-map__dot${load ? ' kilvia-map__dot--load' : ''}`;
    text.className = `kilvia-map__marker kilvia-map__marker--${role}${load ? ' kilvia-map__marker--load' : ''}`;
    eyebrow.textContent = label;
    cityLabel.textContent = city;
    text.append(eyebrow, cityLabel);
    wrapper.append(dot, text);
    return new Marker({ element: wrapper, anchor: 'center' }).setLngLat([coordinate[0], coordinate[1]]).addTo(this.map!);
  }

  private lineFeature(coordinates: readonly RouteCoordinate[]): Feature<LineString> {
    return {
      type: 'Feature',
      properties: {},
      geometry: { type: 'LineString', coordinates: coordinates.map(([longitude, latitude]) => [longitude, latitude]) },
    };
  }
}
