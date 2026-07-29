# GO-ApiRest-Web

Aplicación web en Go que consume la API pública de Chuck Norris para mostrar chistes por distintas modalidades (random, categoría y búsqueda por palabra), renderizando HTML en el navegador.

## 🚀 Descripción

Este proyecto levanta un servidor HTTP local en Go y expone rutas que:
- muestran una página principal (`/`)
- consultan un chiste random (`/api/random`)
- consultan categorías (`/api/categoria`)
- consultan chistes por categoría (`/api/categoria/{nombre}`)
- consultan chistes por texto (`/api/search`)

La lógica está separada en `handlers` y `services` dentro de `internal/`.

## 🧱 Estructura del proyecto

```text
GO-ApiRest-Web/
├── main.go
├── go.mod
├── go.sum
├── internal/
│   ├── handlers/
│   └── services/
├── web/
└── GO-ApiRest-Web.exe
```

## ⚙️ Requisitos

- Go **1.25.0** (según `go.mod`)
- Conexión a internet (para consumir la API externa de Chuck Norris)

> Nota: si no tenés Go 1.25 instalado, podés probar con la versión estable más cercana y ajustar `go.mod` si hace falta.

## ▶️ Cómo ejecutar

### Opción A: ejecutar desde código fuente
```bash
go run main.go
```

### Opción B: compilar y ejecutar
```bash
go build -o app .
./app
```

En Windows:
```powershell
go build -o app.exe .
.\app.exe
```

El servidor queda escuchando en:

- http://localhost:8080

## 🌐 Endpoints registrados

- `GET /` → Home
- `GET /api/random` → chiste random
- `GET /api/categoria` → listado de categorías
- `GET /api/categoria/` → ruta base para categoría
- `GET /api/categoria/{categoria}` → chistes por categoría
- `GET /api/search` → búsqueda por palabra (normalmente vía query string)

## 🛠️ Tecnologías

- Go (net/http estándar)
- Arquitectura por capas simple (`handlers` + `services`)
- Consumo de API REST externa (Chuck Norris)

## 📌 Estado del proyecto

Proyecto funcional a nivel de estructura y arranque del servidor, pendiente de validación final de comportamiento de endpoints y vistas según implementación de `internal/` y `web/`.

## 🧪 Pruebas recomendadas

Con el server corriendo:

- Abrir `http://localhost:8080`
- Probar:
  - `http://localhost:8080/api/random`
  - `http://localhost:8080/api/categoria`
  - `http://localhost:8080/api/categoria/dev` (o alguna categoría válida)
  - `http://localhost:8080/api/search?query=car`

## ❗Problemas comunes

- **Puerto 8080 en uso**: cambiar el puerto en `main.go`.
- **Sin conexión a internet**: fallarán llamadas a la API externa.
- **Versión de Go** no compatible con `go.mod`.

## 📄 Licencia

No definida actualmente en el repositorio.
