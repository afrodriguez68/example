# example

Microservicio HTTP en Go para laboratorio DevOps.

## Resumen

- Lenguaje: Go `1.23.x`
- Módulo: `devops-lab-micro`
- Puerto por defecto: `8080` (configurable con variable `PORT`)
- Entrada principal: `cmd/server/main.go`

## Endpoints

- `GET /` → mensaje base del servicio
- `GET /health` → estado de salud
- `GET /version` → versión del servicio y versión de Go en runtime

### Ejemplos con curl

```bash
curl http://localhost:8080/
curl http://localhost:8080/health
curl http://localhost:8080/version
```

## Ejecutar localmente

### Requisitos

- Go `1.23+`

### Build y ejecución

```bash
go build ./...
go run ./cmd/server
```

Con puerto personalizado:

```bash
PORT=9090 go run ./cmd/server
```

## Tests

```bash
go test ./...
```

## Docker

### Build de imagen

```bash
docker build -t afrodriguez68/example:local .
```

### Ejecutar contenedor

```bash
docker run --rm -p 8080:8080 --name example afrodriguez68/example:local
```

## CI/CD con GitHub Actions

Workflow: `.github/workflows/ci.yml`

El pipeline ejecuta:

1. Checkout del repositorio
2. Build de Go (`go build ./...`)
3. Análisis SonarQube
4. Login a Docker Hub
5. Build y push de imagen Docker con tags:
   - `docker.io/afrodriguez68/example:${{ github.run_number }}`
   - `docker.io/afrodriguez68/example:latest`

### Secrets requeridos en GitHub

Configura estos secrets en el repositorio (Settings → Secrets and variables → Actions):

- `SONAR_TOKEN`
- `SONAR_HOST_URL`
- `DOCKERHUB_USER`
- `DOCKERHUB_TOKEN`

### Cómo ejecutar el workflow

- Automático al hacer push a `main`
- Manual en GitHub: Actions → `CI/CD` → `Run workflow`

### Comportamiento ante errores de conexión

El workflow está configurado para continuar si falla:

- SonarQube scan
- Login a Docker Hub
- Push de la imagen

En esos casos se registra un warning y el job continúa.

## Estructura del proyecto

```text
cmd/
  server/
    main.go
internal/
  httpapi/
    server.go
    server_test.go
Dockerfile
go.mod
```
