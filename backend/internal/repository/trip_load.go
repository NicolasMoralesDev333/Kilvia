package repository

import (
	"context"
	"strings"

	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/domain"
)

const tripSelect = `
	SELECT v.id_viaje, v.id_usuario_operador, v.id_chofer,
	       concat(cu.nombre, ' ', cu.apellido), v.id_vehiculo,
	       concat(ve.patente, ' · ', tv.nombre), v.origen, v.destino,
	       v.fecha_salida, v.fecha_llegada_estimada, v.capacidad_disponible,
	       v.volumen_disponible, v.estado, v.fecha_registro
	FROM viaje v
	JOIN usuario op ON op.id_usuario = v.id_usuario_operador
	JOIN usuario cu ON cu.id_usuario = v.id_chofer
	JOIN vehiculo ve ON ve.id_vehiculo = v.id_vehiculo
	JOIN tipo_vehiculo tv ON tv.id_tipo_vehiculo = ve.id_tipo_vehiculo
`

func (r *Postgres) ListTrips(ctx context.Context, identity domain.Identity) ([]domain.Trip, error) {
	rows, err := r.pool.Query(ctx, tripSelect+`
		WHERE op.id_empresa = $1
		  AND ($2 <> 'CHOFER' OR v.id_chofer = $3)
		ORDER BY v.fecha_salida
	`, identity.CompanyID, identity.Role, identity.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := make([]domain.Trip, 0)
	for rows.Next() {
		var item domain.Trip
		if err := scanTrip(rows, &item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *Postgres) GetTrip(ctx context.Context, identity domain.Identity, tripID string) (domain.Trip, error) {
	var item domain.Trip
	err := scanTrip(r.pool.QueryRow(ctx, tripSelect+`
		WHERE op.id_empresa = $1
		  AND v.id_viaje = $2
		  AND ($3 <> 'CHOFER' OR v.id_chofer = $4)
	`, identity.CompanyID, tripID, identity.Role, identity.UserID), &item)
	return item, normalizeError(err)
}

func (r *Postgres) CreateTrip(ctx context.Context, identity domain.Identity, input domain.CreateTripInput) (domain.Trip, error) {
	var tripID string
	err := r.pool.QueryRow(ctx, `
		INSERT INTO viaje (
			id_usuario_operador, id_chofer, id_vehiculo, origen, destino,
			fecha_salida, fecha_llegada_estimada, capacidad_disponible, volumen_disponible
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id_viaje
	`, identity.UserID, input.DriverID, input.VehicleID, strings.TrimSpace(input.Origin),
		strings.TrimSpace(input.Destination), input.DepartureAt, input.EstimatedArrivalAt,
		input.AvailableWeightTons, input.AvailableVolumeM3,
	).Scan(&tripID)
	if err != nil {
		return domain.Trip{}, normalizeError(err)
	}
	return r.GetTrip(ctx, identity, tripID)
}

func (r *Postgres) UpdateTrip(ctx context.Context, identity domain.Identity, tripID string, input domain.UpdateTripInput) (domain.Trip, error) {
	command, err := r.pool.Exec(ctx, `
		UPDATE viaje v
		SET id_chofer = COALESCE($3, v.id_chofer),
		    id_vehiculo = COALESCE($4, v.id_vehiculo),
		    origen = COALESCE($5, v.origen),
		    destino = COALESCE($6, v.destino),
		    fecha_salida = COALESCE($7, v.fecha_salida),
		    fecha_llegada_estimada = COALESCE($8, v.fecha_llegada_estimada),
		    capacidad_disponible = COALESCE($9, v.capacidad_disponible),
		    volumen_disponible = COALESCE($10, v.volumen_disponible),
		    estado = COALESCE($11, v.estado)
		WHERE v.id_viaje = $2
		  AND EXISTS (
		      SELECT 1 FROM usuario op
		      WHERE op.id_usuario = v.id_usuario_operador AND op.id_empresa = $1
		  )
	`, identity.CompanyID, tripID, input.DriverID, input.VehicleID, input.Origin,
		input.Destination, input.DepartureAt, input.EstimatedArrivalAt,
		input.AvailableWeightTons, input.AvailableVolumeM3, input.Status,
	)
	if err != nil {
		return domain.Trip{}, normalizeError(err)
	}
	if command.RowsAffected() == 0 {
		return domain.Trip{}, ErrNotFound
	}
	return r.GetTrip(ctx, identity, tripID)
}

type rowScanner interface {
	Scan(dest ...any) error
}

func scanTrip(row rowScanner, item *domain.Trip) error {
	return row.Scan(
		&item.ID, &item.OperatorID, &item.DriverID, &item.DriverName,
		&item.VehicleID, &item.VehicleLabel, &item.Origin, &item.Destination,
		&item.DepartureAt, &item.EstimatedArrivalAt, &item.AvailableWeightTons,
		&item.AvailableVolumeM3, &item.Status, &item.RegisteredAt,
	)
}

const loadSelect = `
	SELECT c.id_carga, c.id_usuario, c.origen, c.destino, c.fecha, c.tipo_carga,
	       c.peso, c.volumen, c.vehiculo_requerido, c.restricciones, c.estado, c.fecha_registro
	FROM carga c
	JOIN usuario u ON u.id_usuario = c.id_usuario
`

func (r *Postgres) ListLoads(ctx context.Context, identity domain.Identity) ([]domain.Load, error) {
	rows, err := r.pool.Query(ctx, loadSelect+`
		WHERE u.id_empresa = $1
		ORDER BY c.fecha
	`, identity.CompanyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := make([]domain.Load, 0)
	for rows.Next() {
		var item domain.Load
		if err := scanLoad(rows, &item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *Postgres) GetLoad(ctx context.Context, identity domain.Identity, loadID string) (domain.Load, error) {
	var item domain.Load
	err := scanLoad(r.pool.QueryRow(ctx, loadSelect+`
		WHERE u.id_empresa = $1 AND c.id_carga = $2
	`, identity.CompanyID, loadID), &item)
	return item, normalizeError(err)
}

func (r *Postgres) CreateLoad(ctx context.Context, identity domain.Identity, input domain.CreateLoadInput) (domain.Load, error) {
	var loadID string
	err := r.pool.QueryRow(ctx, `
		INSERT INTO carga (
			id_usuario, origen, destino, fecha, tipo_carga, peso, volumen,
			vehiculo_requerido, restricciones
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id_carga
	`, identity.UserID, strings.TrimSpace(input.Origin), strings.TrimSpace(input.Destination),
		input.PickupAt, strings.TrimSpace(input.CargoType), input.WeightTons,
		input.VolumeM3, strings.TrimSpace(input.RequiredVehicle), strings.TrimSpace(input.Restrictions),
	).Scan(&loadID)
	if err != nil {
		return domain.Load{}, normalizeError(err)
	}
	return r.GetLoad(ctx, identity, loadID)
}

func (r *Postgres) UpdateLoad(ctx context.Context, identity domain.Identity, loadID string, input domain.UpdateLoadInput) (domain.Load, error) {
	command, err := r.pool.Exec(ctx, `
		UPDATE carga c
		SET origen = COALESCE($3, c.origen),
		    destino = COALESCE($4, c.destino),
		    fecha = COALESCE($5, c.fecha),
		    tipo_carga = COALESCE($6, c.tipo_carga),
		    peso = COALESCE($7, c.peso),
		    volumen = COALESCE($8, c.volumen),
		    vehiculo_requerido = COALESCE($9, c.vehiculo_requerido),
		    restricciones = COALESCE($10, c.restricciones),
		    estado = COALESCE($11, c.estado)
		WHERE c.id_carga = $2
		  AND EXISTS (
		      SELECT 1 FROM usuario u
		      WHERE u.id_usuario = c.id_usuario AND u.id_empresa = $1
		  )
	`, identity.CompanyID, loadID, input.Origin, input.Destination, input.PickupAt,
		input.CargoType, input.WeightTons, input.VolumeM3, input.RequiredVehicle,
		input.Restrictions, input.Status,
	)
	if err != nil {
		return domain.Load{}, normalizeError(err)
	}
	if command.RowsAffected() == 0 {
		return domain.Load{}, ErrNotFound
	}
	return r.GetLoad(ctx, identity, loadID)
}

func scanLoad(row rowScanner, item *domain.Load) error {
	return row.Scan(
		&item.ID, &item.UserID, &item.Origin, &item.Destination, &item.PickupAt,
		&item.CargoType, &item.WeightTons, &item.VolumeM3, &item.RequiredVehicle,
		&item.Restrictions, &item.Status, &item.RegisteredAt,
	)
}
