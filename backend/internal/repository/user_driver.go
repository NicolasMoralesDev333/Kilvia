package repository

import (
	"context"
	"strings"

	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/domain"
)

func (r *Postgres) GetUser(ctx context.Context, companyID, userID string) (domain.User, error) {
	var user domain.User
	err := r.pool.QueryRow(ctx, `
		SELECT u.id_usuario, u.id_empresa, e.nombre, u.nombre, u.apellido, u.email,
		       u.tipo_usuario, u.estado, u.fecha_registro
		FROM usuario u
		JOIN empresa e ON e.id_empresa = u.id_empresa
		WHERE u.id_empresa = $1 AND u.id_usuario = $2
	`, companyID, userID).Scan(
		&user.ID, &user.CompanyID, &user.CompanyName, &user.FirstName, &user.LastName,
		&user.Email, &user.Role, &user.Status, &user.RegisteredAt,
	)
	return user, normalizeError(err)
}

func (r *Postgres) ListUsers(ctx context.Context, companyID string) ([]domain.User, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT u.id_usuario, u.id_empresa, e.nombre, u.nombre, u.apellido, u.email,
		       u.tipo_usuario, u.estado, u.fecha_registro
		FROM usuario u
		JOIN empresa e ON e.id_empresa = u.id_empresa
		WHERE u.id_empresa = $1
		ORDER BY u.apellido, u.nombre
	`, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]domain.User, 0)
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(
			&user.ID, &user.CompanyID, &user.CompanyName, &user.FirstName, &user.LastName,
			&user.Email, &user.Role, &user.Status, &user.RegisteredAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, rows.Err()
}

func (r *Postgres) CreateUser(ctx context.Context, companyID string, input domain.CreateUserInput, passwordHash string) (domain.User, error) {
	var user domain.User
	err := r.pool.QueryRow(ctx, `
		INSERT INTO usuario (id_empresa, nombre, apellido, email, password_hash, tipo_usuario)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id_usuario, id_empresa, nombre, apellido, email, tipo_usuario, estado, fecha_registro
	`, companyID, strings.TrimSpace(input.FirstName), strings.TrimSpace(input.LastName),
		strings.ToLower(strings.TrimSpace(input.Email)), passwordHash, input.Role,
	).Scan(&user.ID, &user.CompanyID, &user.FirstName, &user.LastName, &user.Email, &user.Role, &user.Status, &user.RegisteredAt)
	if err != nil {
		return domain.User{}, normalizeError(err)
	}
	company, err := r.GetCompany(ctx, companyID)
	if err != nil {
		return domain.User{}, err
	}
	user.CompanyName = company.Name
	return user, nil
}

func (r *Postgres) UpdateUser(ctx context.Context, companyID, userID string, input domain.UpdateUserInput) (domain.User, error) {
	var user domain.User
	err := r.pool.QueryRow(ctx, `
		UPDATE usuario
		SET nombre = COALESCE($3, nombre),
		    apellido = COALESCE($4, apellido),
		    estado = COALESCE($5, estado)
		WHERE id_empresa = $1 AND id_usuario = $2
		RETURNING id_usuario, id_empresa, nombre, apellido, email, tipo_usuario, estado, fecha_registro
	`, companyID, userID, input.FirstName, input.LastName, input.Status).Scan(
		&user.ID, &user.CompanyID, &user.FirstName, &user.LastName, &user.Email,
		&user.Role, &user.Status, &user.RegisteredAt,
	)
	if err != nil {
		return domain.User{}, normalizeError(err)
	}
	company, err := r.GetCompany(ctx, companyID)
	if err != nil {
		return domain.User{}, err
	}
	user.CompanyName = company.Name
	return user, nil
}

func (r *Postgres) ListDrivers(ctx context.Context, companyID string) ([]domain.Driver, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT u.id_usuario, u.id_empresa, e.nombre, u.nombre, u.apellido, u.email,
		       u.tipo_usuario, u.estado, u.fecha_registro,
		       c.dni, c.numero_licencia, c.categoria_licencia, c.vencimiento_licencia,
		       c.telefono, c.disponibilidad, c.estado
		FROM chofer c
		JOIN usuario u ON u.id_usuario = c.id_usuario
		JOIN empresa e ON e.id_empresa = u.id_empresa
		WHERE u.id_empresa = $1
		ORDER BY u.apellido, u.nombre
	`, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	drivers := make([]domain.Driver, 0)
	for rows.Next() {
		var driver domain.Driver
		if err := rows.Scan(
			&driver.ID, &driver.CompanyID, &driver.CompanyName, &driver.FirstName,
			&driver.LastName, &driver.Email, &driver.Role, &driver.Status, &driver.RegisteredAt,
			&driver.DNI, &driver.LicenseNumber, &driver.LicenseCategory, &driver.LicenseExpiration,
			&driver.Phone, &driver.Availability, &driver.DriverStatus,
		); err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	return drivers, rows.Err()
}

func (r *Postgres) GetDriver(ctx context.Context, companyID, driverID string) (domain.Driver, error) {
	var driver domain.Driver
	err := r.pool.QueryRow(ctx, `
		SELECT u.id_usuario, u.id_empresa, e.nombre, u.nombre, u.apellido, u.email,
		       u.tipo_usuario, u.estado, u.fecha_registro,
		       c.dni, c.numero_licencia, c.categoria_licencia, c.vencimiento_licencia,
		       c.telefono, c.disponibilidad, c.estado
		FROM chofer c
		JOIN usuario u ON u.id_usuario = c.id_usuario
		JOIN empresa e ON e.id_empresa = u.id_empresa
		WHERE u.id_empresa = $1 AND u.id_usuario = $2
	`, companyID, driverID).Scan(
		&driver.ID, &driver.CompanyID, &driver.CompanyName, &driver.FirstName,
		&driver.LastName, &driver.Email, &driver.Role, &driver.Status, &driver.RegisteredAt,
		&driver.DNI, &driver.LicenseNumber, &driver.LicenseCategory, &driver.LicenseExpiration,
		&driver.Phone, &driver.Availability, &driver.DriverStatus,
	)
	return driver, normalizeError(err)
}

func (r *Postgres) CreateDriver(ctx context.Context, companyID string, input domain.CreateDriverInput, passwordHash string) (domain.Driver, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return domain.Driver{}, err
	}
	defer tx.Rollback(ctx)

	var userID string
	err = tx.QueryRow(ctx, `
		INSERT INTO usuario (id_empresa, nombre, apellido, email, password_hash, tipo_usuario)
		VALUES ($1, $2, $3, $4, $5, 'CHOFER')
		RETURNING id_usuario
	`, companyID, strings.TrimSpace(input.FirstName), strings.TrimSpace(input.LastName),
		strings.ToLower(strings.TrimSpace(input.Email)), passwordHash,
	).Scan(&userID)
	if err != nil {
		return domain.Driver{}, normalizeError(err)
	}
	_, err = tx.Exec(ctx, `
		INSERT INTO chofer (
			id_usuario, dni, numero_licencia, categoria_licencia,
			vencimiento_licencia, telefono, disponibilidad
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, userID, strings.TrimSpace(input.DNI), strings.TrimSpace(input.LicenseNumber),
		strings.TrimSpace(input.LicenseCategory), input.LicenseExpiration,
		strings.TrimSpace(input.Phone), input.Availability,
	)
	if err != nil {
		return domain.Driver{}, normalizeError(err)
	}
	if err := tx.Commit(ctx); err != nil {
		return domain.Driver{}, err
	}
	return r.GetDriver(ctx, companyID, userID)
}

func (r *Postgres) UpdateDriver(ctx context.Context, companyID, driverID string, input domain.UpdateDriverInput) (domain.Driver, error) {
	command, err := r.pool.Exec(ctx, `
		UPDATE chofer c
		SET telefono = COALESCE($3, c.telefono),
		    disponibilidad = COALESCE($4, c.disponibilidad),
		    estado = COALESCE($5, c.estado)
		FROM usuario u
		WHERE c.id_usuario = u.id_usuario
		  AND u.id_empresa = $1
		  AND c.id_usuario = $2
	`, companyID, driverID, input.Phone, input.Availability, input.Status)
	if err != nil {
		return domain.Driver{}, normalizeError(err)
	}
	if command.RowsAffected() == 0 {
		return domain.Driver{}, ErrNotFound
	}
	return r.GetDriver(ctx, companyID, driverID)
}
