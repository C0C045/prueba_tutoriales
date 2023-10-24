# CRUD Tutoriales
## Instrucciones de Instalación y Stack Tecnológico
La aplicación requiere de [GO](https://go.dev/), [PostgreSQL](https://www.postgresql.org/) y [Postman](https://www.postman.com/)

En dónde, se usaron las siguientes librerías:

[Gorilla Mux:](https://github.com/gorilla/mux)
```sh
go get -u github.com/gorilla/mux
```

[GORM:](https://gorm.io/index.html)
```sh
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

Se implemento el siguiente ERD en [PostgreSQL:](https://www.postgresql.org/)

![image](https://github.com/C0C045/prueba_tutoriales/assets/55455142/095b635a-8001-4aca-9ea0-2df506b84722)

## Estructura del Proyecto
Se implemento la siguiente estructura:

- db -> # Conexión a Base de datos
- models -> # Modelos de cada una de las tablas
- routes -> # Cada una de las rutas del CRUD
- ...

![image](https://github.com/C0C045/prueba_tutoriales/assets/55455142/a08acc44-30b5-42fa-a8e3-b75108b3ee3e)

## Funcionamiento
En general se hace cada una de los metodos CRUD con ayuda del aplicativo de [Postman](https://www.postman.com/), en dónde se cumple lo siguiente:

`1.` Se implementa la creación usuarios:
https://github.com/C0C045/prueba_tutoriales/assets/55455142/7a3273dc-dfd0-42c9-ad5c-15ebe5107eaa

