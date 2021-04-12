# CRUD_Golang
Proyecto CRUD con Golang y MySQL
 
## Comenzando 🚀
Para ejecutar el proyecto es necesario crear una base de datos en SQL (el nombre de la bd debe ser tickets) con los siguientes atributos

```
CREATE TABLE `tickets` (
  `id` int NOT NULL AUTO_INCREMENT,
  `usuario` varchar(15) NOT NULL,
  `descripcion` varchar(200) NOT NULL,
  `fechaCreacion` date NOT NULL,
  `fechaActualizacion` date NOT NULL,
  `numTicket` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
```

Una vez creada la bd, el paso a seguir es ejecutar la API. Para ello desde consola ejecutamos

```
 go run .\main.go

```

Y en cualquier navegador digitar " http://localhost:8080/ "

## Despliegue 📦

![Home](https://user-images.githubusercontent.com/71616837/114416775-a1ec5b80-9b76-11eb-8b6f-715e90223d53.PNG)
