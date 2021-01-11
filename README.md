# Meli Challange : Operacion Fuego de Quasar v1.1

## Contents
 - [Stack Tecnológico](#stack)
 - [Como utilizarlo](#como-utilizarlo)
 
 ## Stack Tecnológico (#stack)
 
 1. [Docker](https://www.docker.com/ 
 2. Testing con [Stretchr/Testify](https://github.com/stretchr/testify)
 4. Back Go 1.15
 3. Swagger con [Swag](https://github.com/swaggo/swag)
 5. Routing [Gorilla](https://github.com/gorilla)
 
 ## Cómo utilizarlo
 ### Online Herkoku Demo
  ```sh
    https://meli-challange-demo.herokuapp.com/swagger/index.html
  ```
  
 ### Local env
 1. Download fravega-challange utilizando el comando (#como-utilizarlo).
 ```sh
 $ go get -u https://github.com/miguelapabenedit/fravega-challange
 ```
 2.En el root del repositorio correr el comando docker build y run

 3.Esperar a que los contenedores se inicialicen y correr swagger 
 ```sh
http://localhost:8080/swagger/index.html
 ```
 4.Las configuraciones son manejadas dentro del archivo Dockerfile situados en
 ```sh
 ./Dockefile -- Go Api
 ./Dockerfile.web -- Heroku setting
 ```


