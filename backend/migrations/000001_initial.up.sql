CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE empresa (
    id_empresa UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nombre VARCHAR(160) NOT NULL,
    cuit VARCHAR(16) NOT NULL UNIQUE,
    direccion VARCHAR(240) NOT NULL DEFAULT '',
    telefono VARCHAR(40) NOT NULL DEFAULT '',
    email VARCHAR(254) NOT NULL DEFAULT '',
    estado VARCHAR(24) NOT NULL DEFAULT 'ACTIVA' CHECK (estado IN ('ACTIVA', 'INACTIVA')),
    fecha_registro TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE usuario (
    id_usuario UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_empresa UUID NOT NULL REFERENCES empresa(id_empresa),
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    email VARCHAR(254) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    tipo_usuario VARCHAR(16) NOT NULL CHECK (tipo_usuario IN ('ADMIN', 'OPERADOR', 'CHOFER')),
    estado VARCHAR(24) NOT NULL DEFAULT 'ACTIVO' CHECK (estado IN ('ACTIVO', 'INACTIVO')),
    fecha_registro TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE chofer (
    id_usuario UUID PRIMARY KEY REFERENCES usuario(id_usuario) ON DELETE CASCADE,
    dni VARCHAR(20) NOT NULL,
    numero_licencia VARCHAR(50) NOT NULL,
    categoria_licencia VARCHAR(20) NOT NULL,
    vencimiento_licencia DATE NOT NULL,
    telefono VARCHAR(40) NOT NULL DEFAULT '',
    disponibilidad VARCHAR(24) NOT NULL DEFAULT 'DISPONIBLE'
        CHECK (disponibilidad IN ('DISPONIBLE', 'EN_VIAJE', 'NO_DISPONIBLE')),
    estado VARCHAR(24) NOT NULL DEFAULT 'ACTIVO' CHECK (estado IN ('ACTIVO', 'INACTIVO')),
    UNIQUE (dni),
    UNIQUE (numero_licencia)
);

CREATE TABLE tipo_vehiculo (
    id_tipo_vehiculo UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nombre VARCHAR(80) NOT NULL UNIQUE,
    descripcion VARCHAR(240) NOT NULL DEFAULT '',
    capacidad_peso_referencia NUMERIC(10,2) NOT NULL CHECK (capacidad_peso_referencia > 0),
    capacidad_volumen_referencia NUMERIC(10,2) NOT NULL CHECK (capacidad_volumen_referencia > 0)
);

-- VEHICULO necesita propiedad directa para aislar una flota todavía no asignada a viajes.
-- Esta FK no es redundante: el MER no ofrece otro camino estable hacia EMPRESA.
CREATE TABLE vehiculo (
    id_vehiculo UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_empresa UUID NOT NULL REFERENCES empresa(id_empresa),
    id_tipo_vehiculo UUID NOT NULL REFERENCES tipo_vehiculo(id_tipo_vehiculo),
    patente VARCHAR(16) NOT NULL UNIQUE,
    marca VARCHAR(80) NOT NULL,
    modelo VARCHAR(80) NOT NULL,
    capacidad_peso NUMERIC(10,2) NOT NULL CHECK (capacidad_peso > 0),
    capacidad_volumen NUMERIC(10,2) NOT NULL CHECK (capacidad_volumen > 0),
    estado VARCHAR(24) NOT NULL DEFAULT 'DISPONIBLE'
        CHECK (estado IN ('DISPONIBLE', 'EN_VIAJE', 'MANTENIMIENTO', 'INACTIVO')),
    fecha_registro TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE viaje (
    id_viaje UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_usuario_operador UUID NOT NULL REFERENCES usuario(id_usuario),
    id_chofer UUID NOT NULL REFERENCES chofer(id_usuario),
    id_vehiculo UUID NOT NULL REFERENCES vehiculo(id_vehiculo),
    origen VARCHAR(180) NOT NULL,
    destino VARCHAR(180) NOT NULL,
    fecha_salida TIMESTAMPTZ NOT NULL,
    fecha_llegada_estimada TIMESTAMPTZ NOT NULL,
    capacidad_disponible NUMERIC(10,2) NOT NULL CHECK (capacidad_disponible >= 0),
    volumen_disponible NUMERIC(10,2) NOT NULL CHECK (volumen_disponible >= 0),
    estado VARCHAR(24) NOT NULL DEFAULT 'PROGRAMADO'
        CHECK (estado IN ('PROGRAMADO', 'EN_CURSO', 'COMPLETADO', 'CANCELADO')),
    fecha_registro TIMESTAMPTZ NOT NULL DEFAULT now(),
    CHECK (fecha_llegada_estimada > fecha_salida)
);

CREATE TABLE carga (
    id_carga UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_usuario UUID NOT NULL REFERENCES usuario(id_usuario),
    origen VARCHAR(180) NOT NULL,
    destino VARCHAR(180) NOT NULL,
    fecha TIMESTAMPTZ NOT NULL,
    tipo_carga VARCHAR(100) NOT NULL,
    peso NUMERIC(10,2) NOT NULL CHECK (peso > 0),
    volumen NUMERIC(10,2) NOT NULL CHECK (volumen > 0),
    vehiculo_requerido VARCHAR(100) NOT NULL,
    restricciones TEXT NOT NULL DEFAULT '',
    estado VARCHAR(24) NOT NULL DEFAULT 'PENDIENTE'
        CHECK (estado IN ('PENDIENTE', 'ASIGNADA', 'EN_TRANSITO', 'ENTREGADA', 'CANCELADA')),
    fecha_registro TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE match_logistico (
    id_match UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_viaje UUID NOT NULL REFERENCES viaje(id_viaje),
    id_carga UUID NOT NULL REFERENCES carga(id_carga),
    compatibilidad NUMERIC(5,2) NOT NULL CHECK (compatibilidad BETWEEN 0 AND 100),
    desvio_km NUMERIC(10,2) NOT NULL CHECK (desvio_km >= 0),
    tiempo_adicional INTEGER NOT NULL CHECK (tiempo_adicional >= 0),
    km_aprovechables NUMERIC(10,2) NOT NULL CHECK (km_aprovechables >= 0),
    capacidad_utilizada NUMERIC(5,2) NOT NULL CHECK (capacidad_utilizada BETWEEN 0 AND 100),
    estado VARCHAR(24) NOT NULL DEFAULT 'PROPUESTO',
    fecha_creacion TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (id_viaje, id_carga)
);

CREATE TABLE oferta (
    id_oferta UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_match UUID NOT NULL REFERENCES match_logistico(id_match),
    id_usuario UUID NOT NULL REFERENCES usuario(id_usuario),
    monto NUMERIC(14,2) NOT NULL CHECK (monto > 0),
    mensaje TEXT NOT NULL DEFAULT '',
    estado VARCHAR(24) NOT NULL DEFAULT 'ENVIADA',
    fecha_envio TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE operacion (
    id_operacion UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_oferta UUID NOT NULL UNIQUE REFERENCES oferta(id_oferta),
    fecha_confirmacion TIMESTAMPTZ NOT NULL DEFAULT now(),
    estado VARCHAR(24) NOT NULL DEFAULT 'CONFIRMADA',
    observaciones TEXT NOT NULL DEFAULT ''
);

CREATE INDEX idx_usuario_empresa ON usuario(id_empresa);
CREATE INDEX idx_chofer_disponibilidad ON chofer(disponibilidad);
CREATE INDEX idx_vehiculo_empresa_estado ON vehiculo(id_empresa, estado);
CREATE INDEX idx_viaje_operador_salida ON viaje(id_usuario_operador, fecha_salida);
CREATE INDEX idx_viaje_chofer_salida ON viaje(id_chofer, fecha_salida);
CREATE INDEX idx_carga_usuario_fecha ON carga(id_usuario, fecha);
CREATE INDEX idx_match_viaje ON match_logistico(id_viaje);
CREATE INDEX idx_match_carga ON match_logistico(id_carga);
CREATE INDEX idx_oferta_match ON oferta(id_match);

INSERT INTO tipo_vehiculo (
    nombre, descripcion, capacidad_peso_referencia, capacidad_volumen_referencia
) VALUES
    ('Semirremolque', 'Configuración de carga general para larga distancia', 28.00, 90.00),
    ('Sider', 'Semirremolque con apertura lateral', 27.00, 90.00),
    ('Furgón', 'Unidad cerrada para carga general', 8.00, 40.00)
ON CONFLICT (nombre) DO NOTHING;
