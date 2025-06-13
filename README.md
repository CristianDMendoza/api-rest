# API REST en Go 

Este proyecto implementa una **API REST básica** utilizando el lenguaje de programación **Go** junto al paquete `gorilla/mux`. Su objetivo es permitir operaciones CRUD (Crear, Leer, Actualizar, Eliminar) sobre una lista de registros almacenados en memoria.

## Funcionalidades

- **`GET /`**: Muestra un mensaje de bienvenida.
- **`GET /Datos`**: Devuelve todos los registros disponibles.
- **`GET /Datos/{id}`**: Devuelve un registro específico por ID.
- **`POST /Datos`**: Crea un nuevo registro.
- **`PUT /Datos/{id}`**: Actualiza un registro existente por ID.
- **`DELETE /Datos/{id}`**: Elimina un registro existente por ID.

## Estructura de los datos (JSON)

json
{
  "ID": "1",
  "Nombre": "Juan",
  "Genero": "Masculino"
}
## Tecnologías utilizadas
Go (versión 1.24.3)

Gorilla Mux (enrutador HTTP para Go)

JSON para intercambio de datos

##  Instalación y uso
Clona o copia este repositorio localmente:

makefile
Copiar código
C:\xampp\htdocs\Proyectos\api-rest
Abre tu terminal y ubícate en el directorio del proyecto:

bash
Copiar código
cd /c/xampp/htdocs/Proyectos/api-rest
Instala el paquete gorilla/mux (si aún no lo tienes):

bash
Copiar código
go get -u github.com/gorilla/mux
Ejecuta el servidor:

bash
Copiar código
go run main.go
Accede a http://localhost:8000 desde el navegador o prueba las rutas con Postman o cURL.

## Autor
Cristian David Mendoza
GitHub: CristianDMendoza


