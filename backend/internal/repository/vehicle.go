package repository

import (
	"context"
	"strings"

	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/domain"
)

func (r *Postgres) ListVehicleTypes(ctx context.Context) ([]domain.VehicleType, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id_tipo_vehiculo, nombre, descripcion,
		       capacidad_peso_referencia, capacidad_volumen_referencia
		FROM tipo_vehiculo ORDER BY nombre
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := make([]domain.VehicleType, 0)
	for rows.Next() {
		var item domain.VehicleType
		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.ReferenceWeightTons, &item.ReferenceVolumeM3); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *Postgres) ListVehicles(ctx context.Context, companyID string) ([]domain.Vehicle, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT v.id_vehiculo, v.id_empresa, v.id_tipo_vehiculo, t.nombre,
		       v.patente, v.marca, v.modelo, v.capacidad_peso, v.capacidad_volumen,
		       v.estado, v.fecha_registro
		FROM vehiculo v
		JOIN tipo_vehiculo t ON t.id_tipo_vehiculo = v.id_tipo_vehiculo
		WHERE v.id_empresa = $1
		ORDER BY v.fecha_registro DESC
	`, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := make([]domain.Vehicle, 0)
	for rows.Next() {
		var item domain.Vehicle
		if err := rows.Scan(
			&item.ID, &item.CompanyID, &item.TypeID, &item.TypeName, &item.Plate,
			&item.Make, &item.Model, &item.WeightTons, &item.VolumeM3, &item.Status, &item.RegisteredAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *Postgres) GetVehicle(ctx context.Context, companyID, vehicleID string) (domain.Vehicle, error) {
	var item domain.Vehicle
	err := r.pool.QueryRow(ctx, `
		SELECT v.id_vehiculo, v.id_empresa, v.id_tipo_vehiculo, t.nombre,
		       v.patente, v.marca, v.modelo, v.capacidad_peso, v.capacidad_volumen,
		       v.estado, v.fecha_registro
		FROM vehiculo v
		JOIN tipo_vehiculo t ON t.id_tipo_vehiculo = v.id_tipo_vehiculo
		WHERE v.id_empresa = $1 AND v.id_vehiculo = $2
	`, companyID, vehicleID).Scan(
		&item.ID, &item.CompanyID, &item.TypeID, &item.TypeName, &item.Plate,
		&item.Make, &item.Model, &item.WeightTons, &item.VolumeM3, &item.Status, &item.RegisteredAt,
	)
	return item, normalizeError(err)
}

func (r *Postgres) CreateVehicle(ctx context.Context, companyID string, input domain.CreateVehicleInput) (domain.Vehicle, error) {
	var vehicleID string
	err := r.pool.QueryRow(ctx, `
		INSERT INTO vehiculo (
			id_empresa, id_tipo_vehiculo, patente, marca, modelo,
			capacidad_peso, capacidad_volumen
		) VALUES ($1, $2, upper($3), $4, $5, $6, $7)
		RETURNING id_vehiculo
	`, companyID, input.TypeID, strings.TrimSpace(input.Plate), strings.TrimSpace(input.Make),
		strings.TrimSpace(input.Model), input.WeightTons, input.VolumeM3,
	).Scan(&vehicleID)
	if err != nil {
		return domain.Vehicle{}, normalizeError(err)
	}
	return r.GetVehicle(ctx, companyID, vehicleID)
}

func (r *Postgres) UpdateVehicle(ctx context.Context, companyID, vehicleID string, input domain.UpdateVehicleInput) (domain.Vehicle, error) {
	command, err := r.pool.Exec(ctx, `
		UPDATE vehiculo
		SET id_tipo_vehiculo = COALESCE($3, id_tipo_vehiculo),
		    marca = COALESCE($4, marca),
		    modelo = COALESCE($5, modelo),
		    capacidad_peso = COALESCE($6, capacidad_peso),
		    capacidad_volumen = COALESCE($7, capacidad_volumen),
		    estado = COALESCE($8, estado)
		WHERE id_empresa = $1 AND id_vehiculo = $2
	`, companyID, vehicleID, input.TypeID, input.Make, input.Model, input.WeightTons, input.VolumeM3, input.Status)
	if err != nil {
		return domain.Vehicle{}, normalizeError(err)
	}
	if command.RowsAffected() == 0 {
		return domain.Vehicle{}, ErrNotFound
	}
	return r.GetVehicle(ctx, companyID, vehicleID)
}
