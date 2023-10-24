# CRUD Tutoriales
## Instrucciones de Instalación y Stack Tecnológico
La aplicación requiere de [GO](https://go.dev/) y [PostgreSQL](https://www.postgresql.org/)

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

- db #Conexión a Base de datos
