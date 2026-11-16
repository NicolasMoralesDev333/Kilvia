ALTER TABLE viaje
    ADD COLUMN lat_origen NUMERIC(9,6),
    ADD COLUMN lng_origen NUMERIC(9,6),
    ADD COLUMN lat_destino NUMERIC(9,6),
    ADD COLUMN lng_destino NUMERIC(9,6),
    ADD COLUMN distance_base_km NUMERIC(10,2),
    ADD COLUMN codigo_viaje VARCHAR(24) NOT NULL
        DEFAULT ('VJ-' || upper(substr(replace(gen_random_uuid()::text, '-', ''), 1, 8)));

ALTER TABLE carga
    ADD COLUMN lat_origen NUMERIC(9,6),
    ADD COLUMN lng_origen NUMERIC(9,6),
    ADD COLUMN lat_destino NUMERIC(9,6),
    ADD COLUMN lng_destino NUMERIC(9,6),
    ADD COLUMN codigo_carga VARCHAR(24) NOT NULL
        DEFAULT ('CG-' || upper(substr(replace(gen_random_uuid()::text, '-', ''), 1, 8)));

ALTER TABLE match_logistico
    ADD COLUMN route_geometry JSONB;

ALTER TABLE viaje DROP CONSTRAINT viaje_estado_check;
ALTER TABLE viaje ADD CONSTRAINT viaje_estado_check
    CHECK (estado IN ('PROGRAMADO', 'ACTIVO', 'EN_CURSO', 'COMPLETADO', 'CANCELADO'));

CREATE UNIQUE INDEX uq_viaje_codigo ON viaje(codigo_viaje);
CREATE UNIQUE INDEX uq_carga_codigo ON carga(codigo_carga);
CREATE INDEX idx_match_compatibilidad_fecha
    ON match_logistico(compatibilidad DESC, fecha_creacion DESC);
