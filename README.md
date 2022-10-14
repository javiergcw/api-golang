<!-- # GOLANG REST API SIMPLE

# Package
* gorilla/mux: `go get github.com/gorilla/mux` -->

<h1 align="center"><!--  -->
  <b> API GOLANG </b>
  <br>
</h1>

<h4 align="center">API de ticket, desarrollada en <a href="https://go.dev" target="_blank">GO</a>.</h4>

<p align="center">
  <a href="https://github.com/gorilla/mux">
    <img src="https://camo.githubusercontent.com/a62a5e2040257dd8787001ffa5d95964d7bc77024aa2ba3d94e64ec1e151228e/68747470733a2f2f636c6f75642d63646e2e7175657374696f6e61626c652e73657276696365732f676f72696c6c612d69636f6e2d36342e706e67"
         alt="Gorilla">

</a>

</p>
<p align="center">
  <a href="#requisitos">Requisitos</a> •
  <a href="#descargar">Descargar</a> •
  <a href="#postman">Postman</a> 

</p>

## Funciones

* Permite realizar las siguientes acciones con el ticket:
    - [x] Crear
    - [x] Eliminar
    - [x] Editar
    - [x] Recuperar
        - [x] Con paginación

* Caracteristicas

- [x] http RESTFUL

* Estructura de los tickets

```go
type Ticket struct {
ID           int    `json:"id"`
Name         string `json:"Name"`
Status       string `json:"Status""`
CreationDate string `json:"CreationDate"`
UpdateDate   string `json:"UpdateDate"`
}
```

* Puntos extras
    - [ ] Pruebas unitarias
    - [ ] Docker
    - [ ] GRPC y GraphQL.
    - [X] SOLID, patrones de diseño

## Descargar

Para descargar y compilar es necesario [Git](https://git-scm.com), [Node.js](https://nodejs.org/en/download/)
, [Go](https://go.dev)
Luego de clonarlo, es de suma importancia realizar paso 2 y 3.

```bash
# Paso 1: Clonar el repositorio
$ git clone https://github.com/javiergcw/api-golang.git

# Paso 2: Dependecias gorilla 
$ go get github.com/gorilla/mux

# Paso 3: Iniciar el servidor (Asegurarse que está en la ruta correcta del proyecto)
$ go run main.go
```

## Postman

Para realizar las solicitudes es recomendable el uso de [Postman](https://www.postman.com), posteriormente de compilar
el proyecto añadir en tu workspace el socket http://localhost:3000/tickets , y realizar las respectivas solicitudes por
medio de
GET, POST, PUT, DELETE.

Si quiere, puede realizar pruebas con el siguiente modelo en el body de postman

```json lines
{
  "Name": "Jesus",
  "Status": "Cerrado",
  "CreationDate": "2022-10-12 12:08:38",
  "UpdateDate": "2022-10-12 12:08:38"
}
```

- **CONSULTAR:** Metodo GET | Socket http://localhost:3000/tickets
- **PAGINAR:** Metodo GET | Socket http://localhost:3000/tickets/ID_QUE_REQUIERA_CONSULTAR
- **CREAR:** Metodo POST | Socket http://localhost:3000/tickets | añadir en el body de postman un json con
  Name,Status,CreationDate y UpdateDate
- **ELIMINAR:** Metodo DELETE | Socket http://localhost:3000/tickets/ID_QUE_REQUIERA_ELIMINAR
- **ACTUALIZAR:** Metodo PUT | Socket http://localhost:3000/tickets/ID_QUE_REQUIERA_EDITAR | añadir en el body del
  postman, el json editado, en caso de que no añada todo el json completo, este tomara los atributos no obtenidos como vacios.
