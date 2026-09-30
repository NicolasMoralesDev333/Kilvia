package repository

import (
	"context"
	"strings"

	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/domain"
)

func (r *Postgres) RegisterCompanyAdmin(ctx context.Context, input domain.RegisterCompanyInput, passwordHash string) (domain.User, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return domain.User{}, err
	}
	defer tx.Rollback(ctx)

	var companyID string
	err = tx.QueryRow(ctx, `
		INSERT INTO empresa (nombre, cuit, email)
		VALUES ($1, $2, $3)
		RETURNING id_empresa
	`, strings.TrimSpace(input.CompanyName), strings.TrimSpace(input.CUIT), strings.ToLower(strings.TrimSpace(input.Email))).Scan(&companyID)
	if err != nil {
		return domain.User{}, normalizeError(err)
	}

	var user domain.User
	err = tx.QueryRow(ctx, `
		INSERT INTO usuario (id_empresa, nombre, apellido, email, password_hash, tipo_usuario)
		VALUES ($1, $2, $3, $4, $5, 'ADMIN')
		RETURNING id_usuario, id_empresa, nombre, apellido, email, tipo_usuario, estado, fecha_registro
	`, companyID, strings.TrimSpace(input.FirstName), strings.TrimSpace(input.LastName), strings.ToLower(strings.TrimSpace(input.Email)), passwordHash).
		Scan(&user.ID, &user.CompanyID, &user.FirstName, &user.LastName, &user.Email, &user.Role, &user.Status, &user.RegisteredAt)
	if err != nil {
		return domain.User{}, normalizeError(err)
	}
	user.CompanyName = strings.TrimSpace(input.CompanyName)
	if err := tx.Commit(ctx); err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *Postgres) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	err := r.pool.QueryRow(ctx, `
		SELECT u.id_usuario, u.id_empresa, e.nombre, u.nombre, u.apellido, u.email,
		       u.tipo_usuario, u.estado, u.fecha_registro, u.password_hash
		FROM usuario u
		JOIN empresa e ON e.id_empresa = u.id_empresa
		WHERE lower(u.email) = lower($1)
	`, strings.TrimSpace(email)).Scan(
		&user.ID, &user.CompanyID, &user.CompanyName, &user.FirstName, &user.LastName,
		&user.Email, &user.Role, &user.Status, &user.RegisteredAt, &user.PasswordHash,
	)
	return user, normalizeError(err)
}

func (r *Postgres) GetCompany(ctx context.Context, companyID string) (domain.Company, error) {
	var company domain.Company
	err := r.pool.QueryRow(ctx, `
		SELECT id_empresa, nombre, cuit, direccion, telefono, email, estado, fecha_registro
		FROM empresa WHERE id_empresa = $1
	`, companyID).Scan(
		&company.ID, &company.Name, &company.CUIT, &company.Address, &company.Phone,
		&company.Email, &company.Status, &company.RegisteredAt,
	)
	return company, normalizeError(err)
}

func (r *Postgres) UpdateCompany(ctx context.Context, companyID string, company domain.Company) (domain.Company, error) {
	var updated domain.Company
	err := r.pool.QueryRow(ctx, `
		UPDATE empresa
		SET nombre = $2, direccion = $3, telefono = $4, email = $5
		WHERE id_empresa = $1
		RETURNING id_empresa, nombre, cuit, direccion, telefono, email, estado, fecha_registro
	`, companyID, strings.TrimSpace(company.Name), strings.TrimSpace(company.Address), strings.TrimSpace(company.Phone), strings.ToLower(strings.TrimSpace(company.Email))).Scan(
		&updated.ID, &updated.Name, &updated.CUIT, &updated.Address, &updated.Phone,
		&updated.Email, &updated.Status, &updated.RegisteredAt,
	)
	return updated, normalizeError(err)
}
