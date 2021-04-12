# CRUD_Golang
Proyecto CRUD con Golang y MySQL
 
## Comenzando 🚀
Para ejecutar el proyecto es necesario crear una base de datos en SQL (el nombre de la bd debe ser tickets) con los siguientes atributos

el usuario y la contraseña que se deben configurar en la creacion de la bd son

usuario : root
contraseña: admin

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

 Home
![Home](https://user-images.githubusercontent.com/71616837/114416775-a1ec5b80-9b76-11eb-8b6f-715e90223d53.PNG)

 Editar
 
 ![EditarTicket](https://user-images.githubusercontent.com/71616837/114416774-a1ec5b80-9b76-11eb-805e-4b1d9e1bbd93.PNG)
 
 Buscar Por numero de Ticket
 
 ![BuscarTicket](https://user-images.githubusercontent.com/71616837/114416773-a153c500-9b76-11eb-82d1-787f09db9af8.PNG)
 
 Agregar Nuevo Ticket
 
 ![Agregar_Ticket](https://user-images.githubusercontent.com/71616837/114416771-a153c500-9b76-11eb-9f4e-de1452bacb6b.PNG)
 
 Visualizar ticket
 
 ![Visualizar_Ticket](https://user-images.githubusercontent.com/71616837/114416777-a284f200-9b76-11eb-815f-56c9d5bc1bcc.PNG)

 ## Construido con 🛠️

*[Golang](https://golang.org/)
*[MySQL](https://www.mysql.com/)
 
 
 ## Autor ✒️
 
 Erick Yoan Ahumada
 
 ## Referencia 📖
 
 He Aprendido un poco de Golang haciendo uso del siguiente Tutorial
 [Tutorial](https://www.youtube.com/watch?v=G58gN0lIbyI&t=1586s)
 
