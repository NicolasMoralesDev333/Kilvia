import { HttpClient } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { Observable } from 'rxjs';
import {
  ApiDriver,
  ApiCompany,
  ApiLoad,
  ApiTrip,
  ApiUser,
  ApiVehicle,
  ApiVehicleType,
} from '../models/api.models';

const API_BASE_URL = '/api/v1';

@Injectable({ providedIn: 'root' })
export class KilviaApiService {
  private readonly http = inject(HttpClient);

  listTrips(): Observable<ApiTrip[]> {
    return this.http.get<ApiTrip[]>(`${API_BASE_URL}/trips`);
  }

  updateTrip(id: string, input: { readonly status: string }): Observable<ApiTrip> {
    return this.http.patch<ApiTrip>(`${API_BASE_URL}/trips/${id}`, input);
  }

  createTrip(input: Omit<ApiTrip, 'id' | 'operatorId' | 'driverName' | 'vehicleLabel' | 'status'>): Observable<ApiTrip> {
    return this.http.post<ApiTrip>(`${API_BASE_URL}/trips`, input);
  }

  listVehicles(): Observable<ApiVehicle[]> {
    return this.http.get<ApiVehicle[]>(`${API_BASE_URL}/vehicles`);
  }

  listVehicleTypes(): Observable<ApiVehicleType[]> {
    return this.http.get<ApiVehicleType[]>(`${API_BASE_URL}/vehicle-types`);
  }

  createVehicle(input: Pick<ApiVehicle, 'typeId' | 'plate' | 'make' | 'model' | 'weightTons' | 'volumeM3'>): Observable<ApiVehicle> {
    return this.http.post<ApiVehicle>(`${API_BASE_URL}/vehicles`, input);
  }

  updateVehicle(id: string, input: { readonly status: string }): Observable<ApiVehicle> {
    return this.http.patch<ApiVehicle>(`${API_BASE_URL}/vehicles/${id}`, input);
  }

  listDrivers(): Observable<ApiDriver[]> {
    return this.http.get<ApiDriver[]>(`${API_BASE_URL}/drivers`);
  }

  createDriver(input: {
    readonly firstName: string;
    readonly lastName: string;
    readonly email: string;
    readonly password: string;
    readonly dni: string;
    readonly licenseNumber: string;
    readonly licenseCategory: string;
    readonly licenseExpiration: string;
    readonly phone: string;
    readonly availability: string;
  }): Observable<ApiDriver> {
    return this.http.post<ApiDriver>(`${API_BASE_URL}/drivers`, input);
  }

  updateDriver(id: string, input: { readonly status: string }): Observable<ApiDriver> {
    return this.http.patch<ApiDriver>(`${API_BASE_URL}/drivers/${id}`, input);
  }

  listLoads(): Observable<ApiLoad[]> {
    return this.http.get<ApiLoad[]>(`${API_BASE_URL}/loads`);
  }

  createLoad(input: Omit<ApiLoad, 'id' | 'userId' | 'status'>): Observable<ApiLoad> {
    return this.http.post<ApiLoad>(`${API_BASE_URL}/loads`, input);
  }

  updateLoad(id: string, input: { readonly status: string }): Observable<ApiLoad> {
    return this.http.patch<ApiLoad>(`${API_BASE_URL}/loads/${id}`, input);
  }

  listUsers(): Observable<ApiUser[]> {
    return this.http.get<ApiUser[]>(`${API_BASE_URL}/users/`);
  }

  createOperator(input: { readonly firstName: string; readonly lastName: string; readonly email: string; readonly password: string }): Observable<ApiUser> {
    return this.http.post<ApiUser>(`${API_BASE_URL}/users/`, { ...input, role: 'OPERADOR' });
  }

  updateUser(id: string, input: { readonly status: string }): Observable<ApiUser> {
    return this.http.patch<ApiUser>(`${API_BASE_URL}/users/${id}`, input);
  }

  getCompany(): Observable<ApiCompany> {
    return this.http.get<ApiCompany>(`${API_BASE_URL}/company`);
  }

  updateCompany(input: Pick<ApiCompany, 'name' | 'cuit' | 'address' | 'phone' | 'email'>): Observable<ApiCompany> {
    return this.http.patch<ApiCompany>(`${API_BASE_URL}/company`, input);
  }
}
