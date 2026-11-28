# Users API
API REST en Go para gestionar usuarios (nombre, apellidos y email). Usa Gin para la API y GORM para hablar con MySQL, y se arranca entera con Docker Compose.

# Cómo arrancarla
Hace falta tener Docker instalado.
1. Copia `.env.example` a `.env` y cambia los valores (contraseñas y API key).
2. Arranca todo:`docker compose up -d --build` en la terminal.

3. La API queda en `http://localhost:8080`.

Para pararla: `docker compose down`.

# Cómo usarla

Todas las peticiones necesitan la cabecera `X-API-Key` con la clave que pusiste en el `.env`.

- GET `/users`: lista los usuarios
- POST `/users`: crea un usuario
- GET `/users/:id`: devuelve un usuario
- PUT `/users/:id`: actualiza un usuario
- DELETE `/users/:id`: borra un usuario

Ejemplo para crear un usuario:

- Método: POST
- URL: `http://localhost:8080/users`
- Header: `X-API-Key` con la clave del `.env`
- Body: raw, formato JSON 
{
  "name": "Hugo",
  "surname": "Merino",
  "email": "hugom@gmail.com"
}

`name` y `surname` son obligatorios (máximo 100 caracteres) y `email` tiene que ser un email válido.

# Estructura

Sigue el patrón MVC:

- `models/`: cómo es un usuario.
- `controllers/`: lo que hace cada operación (crear, listar, buscar, actualizar, borrar).
- `routes/`: une cada URL con su controller. La respuesta en JSON hace de vista.
- `middleware/`: código que se ejecuta antes de los controllers (API key, rate limit, tamaño del body).
- `database/`: conexión con MySQL.
- `main.go`: arranca el servidor.

# Seguridad

Cosas que he pensado y lo que he hecho para cada una:

- **Saturar la API (DoS):** rate limit por IP. Máximo 5 peticiones por segundo, si se pasa responde 429.
- **Peticiones gigantes:** el body no puede pasar de 64 KB.
- **Datos malos:** validación en el modelo campos obligatorios, longitud máxima, formato de email.
- **Acceso sin permiso:** hace falta la API key en todas las rutas. Sin ella responde 401.
- **Inyección SQL:** GORM usa consultas parametrizadas.
- **Que el cliente elija el id:** el id lo pone siempre la base de datos.
- **Falsear la IP para saltarse el rate limit:** Gin no se fía de las cabeceras de proxy.
- **Secretos en GitHub:** las contraseñas y la API key están en `.env`, que no se sube. Solo se sube `.env.example`.
- **Puertos abiertos:** la API y la base de datos solo escuchan en `127.0.0.1`.
- **Contenedor con permisos de más:** la API corre con un usuario sin privilegios, no como root.
