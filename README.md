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

`1.` Se implementa el metodo POST de usuarios:

https://github.com/C0C045/prueba_tutoriales/assets/55455142/7a3273dc-dfd0-42c9-ad5c-15ebe5107eaa

`2.` Se implementa el metodo GET de todos los usuarios:

https://github.com/C0C045/prueba_tutoriales/assets/55455142/2d5778f8-640b-40d5-ab46-00eb43072885

`3.` Se implementa el metodo POST de detalle con un usuario existente:

https://github.com/C0C045/prueba_tutoriales/assets/55455142/160d6343-801b-4480-9752-aa0cfef5042d

`4.` Se desarrolla el metodo GET de todos los detalles:

https://github.com/C0C045/prueba_tutoriales/assets/55455142/d7ecdc62-57fa-4c4f-b3fb-a5b420be44d6

`5.` Se desarrolla el metodo POST de tutorial con un detalle existente:

https://github.com/C0C045/prueba_tutoriales/assets/55455142/3c181589-4c3f-4fd1-b6fe-ce4f835480fe

`6.` Se desarrolla el metodo GET de todos los tutoriales:

https://github.com/C0C045/prueba_tutoriales/assets/55455142/1c50f00b-162c-45d8-80fb-03d6a27730ff

`7.` Se desarrolla el metodo PUT de un tutorial en especifico:

https://github.com/C0C045/prueba_tutoriales/assets/55455142/0ff001a3-442f-46a2-bfe9-ee907187d3d0

`8.` Se desarrolla el metodo DELETE de un tutorial en especifico:

https://github.com/C0C045/prueba_tutoriales/assets/55455142/174786f6-65e0-4a81-8a91-6720fe38a53c

