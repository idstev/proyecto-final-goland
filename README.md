# proyecto-final-goland
proyecto de sistema de gestion de marketplace

Este proyecto es una plataforma backend desarrollada en Go (Golang) con conexión a PostgreSQL, que permite la gestión de un marketplace de productos con funcionalidad de usuarios, autenticación, productos y pedidos. Incluye tanto vistas HTML como servicios web tipo API RESTful para integración con otras aplicaciones.

Equipo de desarrollo

- Nombre del proyecto: Sistema de gestion de marketplace
- Integrantes:
  - Mariana Defáz, Johaan Cuenca
- Materia / Docente: Programacion Orientada A Objetos - Milton Palacios
  
Objetivos del sistema

El objetivo principal de este proyecto es desarrollar un sistema de gestión de un *Marketplace* funcional que:

- Permita el registro y autenticación de usuarios.
- Diferencie entre *vendedores* y *compradores* mediante roles.
- Habilite a los vendedores a publicar productos.
- Permita a los compradores realizar pedidos.
- Ofrezca vistas HTML limpias para navegación y formularios.
- Exponga servicios web con *JSON* (API REST) para consumo externo.

 Estructura del proyecto

marketplace/
├── cmd/                # Punto de entrada (main.go)
├── config/             # Conexión a base de datos
├── controllers/        # Lógica de rutas y controladores
├── models/             # Modelos de datos y lógica SQL
├── static/             # Archivos CSS
├── templates/          # Vistas HTML
├── go.mod / go.sum     # Dependencias del proyecto
```

 Funcionalidades principales

 Usuarios
- Registro de usuarios vía formulario HTML y JSON.
- Inicio de sesión con cookies de sesión.
- Gestión de roles: `buyer` (comprador) y `seller` (vendedor).
- Visualización del usuario actual mediante `/api/me`.

Productos
- Crear, editar, eliminar y listar productos por usuario.
- Productos filtrados por vendedor.
- Vista HTML y servicios REST (`/api/products`).

Pedidos (Diseñado)
- Relación de compradores con productos comprados.
- Estructura lista para implementar CRUD de `orders`.
- Total calculado automáticamente.

 API REST (JSON)
- `/api/register` – Crear usuario
- `/api/login` – Login con sesión
- `/api/me` – Obtener usuario autenticado
- `/api/products` – Listar o crear productos
- `/api/products/{id}` – Obtener, actualizar o eliminar producto

Tecnologías utilizadas

- Lenguaje: Go (Golang)
- Base de datos:PostgreSQL
- Frontend:HTML + CSS
- Servicios REST: net/http + JSON
- IDE: GoLand / Visual Studio Code

 Modelo de base de datos

Entidades:
- `users`: Usuarios con autenticación y roles.
- `products`: Productos ofrecidos por vendedores.
- `orders`: Pedidos realizados por compradores (estructura lista).

 Cómo ejecutar

1. Clona el repositorio:
   ```bash
   git clone https://github.com/idstev/marketplace.git
   ```

2. Configura PostgreSQL con el esquema base.

3. Ejecuta el servidor:
   ```bash
   go run cmd/main.go
   ```

4. Abre en el navegador: [http://localhost:8080](http://localhost:8080)
   
 Estado del desarrollo

- [x] Registro y login de usuarios
- [x] CRUD de productos
- [x] Servicios REST JSON
- [x] Sesión por cookies
- [x] Modelo ER visual
- [ ] CRUD de pedidos (`orders`) en desarrollo

