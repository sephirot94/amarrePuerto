# Sistema controlador de Amarres de un Puerto
## Trabajo Practico Final para Materia Sistemas Concurrentes de la Universidad de Belgrano

### Para instalar el compilador de go en caso de querer modificar el codigo y compilar, siga los pasos de este tutorial
* #### Instalar compilador Golang (https://www.tutorialesprogramacionya.com/goya/detalleconcepto.php?punto=1&codigo=1&inicio=0)

### Luego, una vez instalado el compilador, ejecute las siguientes lineas de codigo:

* #### Para compilar (esta operacion gener aun archivo binario ejecutable):

   `go build /src/api/main.go`

* #### Para ejecutar:

   `go run /src/api/main.go`

### Estructura de directorios:

#### /src/app contiene el codigo del programa. El mismo esta separado por:

* ##### main.go : la ejecucion del programa empieza aqui.

* ##### application: donde se construye y ejecuta la aplicacion.

* ##### controllers: contienen las pseudo clases (comparando con Java) de nuestras entidades del sistema, Amarras y Navios.

* ##### helpers: contienen servicios reutilizables para el manejo de concurrencia entre procesos.

* ##### models: define los modelos de datos que el sistema necesita.

#### /pkg contiene los binarios compilados separados por versiones. La ultima version en la carpeta es la actual.

### Autores:

* ##### Agustin Candenas - matricula #50211725
* ##### Nereo Candenas - matricula #50111219
* ##### Gabriel Ferreira - matricula #50111222
* ##### Ivan Jinkus - matricula #50111220
* ##### Mariano Martin - matricula #50111221



