# Guía para ejecutar Kilvia en otra PC

Esta guía sirve para mostrar Kilvia localmente en una computadora Windows usando Docker. No hace falta instalar Go, Node.js ni Visual Studio Code.

## 1. Preparar y pasar el proyecto

1. Comprimí la carpeta del proyecto en un archivo `.zip`.
2. Pasá el ZIP a la otra PC, por ejemplo con un pendrive o Drive.
3. Extraelo en una carpeta fácil de encontrar, por ejemplo `Escritorio\Kilvia`.

## 2. Instalar y abrir Docker Desktop

1. Instalá Docker Desktop para Windows si todavía no está instalado.
2. Si el instalador solicita habilitar WSL 2 o la virtualización, completá esos pasos y reiniciá la PC cuando lo pida.
3. Abrí Docker Desktop y esperá a que indique que el motor está iniciado.

## 3. Iniciar Kilvia

1. Abrí la carpeta extraída y verificá que contenga `docker-compose.yml`.
2. Abrí PowerShell en esa carpeta. Una forma sencilla es hacer clic derecho en un espacio vacío de la carpeta y elegir **Abrir en Terminal**.
3. Ejecutá:

   ```powershell
   docker compose up --build -d
   ```

La primera vez puede tardar varios minutos mientras Docker descarga imágenes y construye la aplicación. Dejá Docker Desktop abierto.

No necesitás crear ni copiar un archivo `.env` para esta demo local: Docker Compose usa valores predeterminados adecuados para ejecutar la aplicación en esa PC.

## 4. Abrir la aplicación

Cuando termine el comando, abrí en el navegador:

**http://localhost:4200**

Elegí **Crear cuenta** para registrar la empresa y el usuario administrador. La base de datos de esa PC empieza vacía; cargá los vehículos, choferes, viajes y cargas que quieras mostrar.

La vista de mapa puede necesitar conexión a internet para cargar los mapas y las rutas.

## 5. Comprobar o detener los servicios

Para comprobar que los contenedores están activos, ejecutá en la misma carpeta:

```powershell
docker compose ps
```

Para detener Kilvia al terminar:

```powershell
docker compose down
```

Esto conserva los datos en esa PC. No uses `docker compose down -v`, porque esa opción elimina también el volumen de la base de datos.

## Si algo falla

- **Docker no está iniciado:** abrí Docker Desktop y esperá a que el motor esté listo; después repetí el comando de inicio.
- **El puerto ya está en uso:** cerrá la otra aplicación que esté usando los puertos `4200`, `5432` o `8080`, y volvé a intentar.
- **La web todavía no abre:** esperá un poco y comprobá los servicios con `docker compose ps`.
- **Querés ver el error de inicio:** ejecutá `docker compose logs`.

Para reiniciar los servicios sin borrar los datos, podés ejecutar:

```powershell
docker compose down
docker compose up --build -d
```
