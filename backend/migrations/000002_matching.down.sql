DROP INDEX IF EXISTS idx_match_compatibilidad_fecha;
DROP INDEX IF EXISTS uq_carga_codigo;
DROP INDEX IF EXISTS uq_viaje_codigo;

ALTER TABLE match_logistico DROP COLUMN IF EXISTS route_geometry;
ALTER TABLE carga DROP COLUMN IF EXISTS codigo_carga,
    DROP COLUMN IF EXISTS lat_origen,
    DROP COLUMN IF EXISTS lng_origen,
    DROP COLUMN IF EXISTS lat_destino,
    DROP COLUMN IF EXISTS lng_destino;
ALTER TABLE viaje DROP COLUMN IF EXISTS codigo_viaje,
    DROP COLUMN IF EXISTS distance_base_km,
    DROP COLUMN IF EXISTS lat_origen,
    DROP COLUMN IF EXISTS lng_origen,
    DROP COLUMN IF EXISTS lat_destino,
    DROP COLUMN IF EXISTS lng_destino;

ALTER TABLE viaje DROP CONSTRAINT viaje_estado_check;
ALTER TABLE viaje ADD CONSTRAINT viaje_estado_check
    CHECK (estado IN ('PROGRAMADO', 'EN_CURSO', 'COMPLETADO', 'CANCELADO'));
