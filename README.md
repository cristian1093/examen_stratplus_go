Instrucciones para ejecutar el programa, Configurar GOPATH Y GOROOT
https://sourabhbajaj.com/mac-setup/Go/

Descargar framework de gin
https://github.com/gin-gonic/gin

El archivo .env contiene los datos de conexión de la bd mysql

Iniciar el proyecto en terminal
go run main.go

La base de datos que se utilizo para el proyecto fue mysql, ya se realizan relaciones entre tablas.

1.-Tabla users:

Contiene información sobre los usuarios (como Id, User, Email, Phone).

2.- Tabla bonds:

Almacena detalles sobre los bonos (por ejemplo, Id, Name, Description, CurrentPrice).

3.-Tabla orders:

Representa las órdenes realizadas por los usuarios para comprar o vender bonos.
Contiene claves foráneas (UserId y BondId) que probablemente se relacionan con las tablas users y bonds respectivamente.
Se utiliza enlaces (User y Bond) para almacenar información detallada de usuario y bono relacionada con cada orden.
El código muestra el uso de procedimientos almacenados para realizar consultas (UpdateAndShowOrders) que pueden involucrar operaciones relacionadas en estas tablas.

En este link se muestra el esquema de las tablas utilizadas y sus relaciones.

https://1drv.ms/i/s!Alkd4o1OlKu9gZ57TJkluru3eaxgWQ?e=fAJgXe

En el la carpeta scripts del proyecto contiene los scripts/mysql de la base de datos, como tambien un procedimiento almacenado el cual realiza la cancelacion de las ordenes cuando el tiempo de expiracion se haya completado. 

Adjunto el documento json de Postman para que puedan importar las apis

https://1drv.ms/u/s!Alkd4o1OlKu9gZ58i97XVugLDCXEsw?e=NWK6XT

Decisiones técnicas tomadas en la construcción del API

La arquitectura del código sigue el patrón de arquitectura de tres capas, donde cada capa tiene una responsabilidad específica y las capas están separadas para facilitar la escalabilidad y el mantenimiento. A continuación, se describe la arquitectura y cómo interactúan los diferentes componentes:


La arquitectura del código que has proporcionado sigue el patrón de arquitectura de tres capas, donde cada capa tiene una responsabilidad específica y las capas están separadas para facilitar la escalabilidad y el mantenimiento. A continuación, se describe la arquitectura y cómo interactúan los diferentes componentes:

1.-Capa de Controladores (controllers):

Esta capa maneja las solicitudes HTTP, interactúa con el framework Gin y contiene la lógica de control de la aplicación.
Los controladores reciben solicitudes HTTP, realizan la validación necesaria y llaman a funciones en la capa de modelos para interactuar con los datos.
Utiliza la estructura Order para representar los datos de las órdenes y manejar su creación y actualización.

2.-Capa de Modelos (models):

Esta capa se encarga de interactuar con la capa de almacenamiento de datos (base de datos, sistema de archivos, etc.).
Contiene funciones que realizan operaciones de lectura y escritura en la fuente de datos.
La capa de modelos se comunica con la capa de controladores para devolver resultados o errores.

Sincronización (sync.RWMutex):


sync.RWMutex se utiliza para controlar el acceso concurrente, permitiendo múltiples lecturas simultáneas (RLock) pero asegurando la exclusión mutua para escrituras (Lock y Unlock).
Esta sincronización es esencial para evitar condiciones de carrera cuando varias solicitudes intentan leer o escribir en el OrderBook simultáneamente.

3.-Framework Gin:

Gin se utiliza como el marco web para manejar las solicitudes HTTP y las respuestas.
Los controladores se registran con las rutas definidas en Gin, y Gin dirige las solicitudes a los controladores correspondientes.
Facilita la manipulación de parámetros de URL, decodificación de solicitudes JSON, y la generación de respuestas HTTP.

4.-Base de Datos:

No se proporciona el código específico de la conexión y manipulación de la base de datos. Sin embargo, se utilizan consultas SQL para interactuar con la base de datos.
Parece que las funciones acceden a la base de datos a través de una conexión denominada database.db.


Se utilizaron bibliotecas externas las cuales son las siguientes:

1.- "github.com/gin-gonic/gin"  (Gin)
Gin es un framework web para el lenguaje de programación Go (Golang) que se ha vuelto popular debido a su rendimiento rápido y su diseño minimalista. Algunas de las ventajas de usar Gin en el desarrollo de aplicaciones web con Go incluyen:

-Rendimiento rápido
-Bajo acoplamiento
-Sintaxis sencilla
-Middleware eficiente
-Enrutamiento rápido
-Soporte para grupos de rutas
-Documentación clara
-Amplia comunidad

2.- "github.com/tools/viper" es un paquete en Go (Golang) que proporciona una solución completa y robusta para la gestión de la configuración en aplicaciones. Esta biblioteca es especialmente útil para leer y escribir configuraciones en diferentes formatos, como JSON, YAML, TOML, entre otros. Aquí hay algunas características clave de Viper:

-Formatos de configuración compatibles
-Soporte para múltiples fuentes de configuración
-Configuración predeterminada
-Recarga automática de configuración
-Manipulación de configuración anidada
-Soporte para lectura desde directorios
-Funciones de expansión de variables
-Documentación completa y ejemplos


3.- "github.com/hako/branca" es una biblioteca en Go (Golang) que implementa el formato de token Branca. Branca es un formato de token seguro y compacto basado en el cifrado de chacha20 y la autenticación de Poly1305. Este formato es adecuado para la creación y verificación de tokens en aplicaciones web y sistemas distribuidos. Aquí hay algunas características clave de la biblioteca Branca en Go:

-Seguridad
-Compacto
-Sin dependencias externas
-Tokens autenticados y encriptados
-Compatibilidad con otros lenguajes

4.- "github.com/go-sql-driver/mysql" es un controlador (driver) MySQL para Go (Golang). Este controlador proporciona la funcionalidad necesaria para que las aplicaciones escritas en Go se conecten, interactúen y gestionen bases de datos MySQL. Aquí hay algunas características clave de este controlador:
-Compatibilidad con MySQL
-Soporte para funcionalidades avanzadas
-Conexiones seguras
-Configuración flexible
-Preparación de consultas
-Escaneo directo en estructuras
-Seguimiento de errores detallado

5.- "github.com/dgrijalva/jwt-go" es una biblioteca en Go (Golang) para la creación y verificación de JSON Web Tokens (JWT). Los JWT son un estándar abierto (RFC 7519) que define una forma compacta y autónoma de representar información entre dos partes. Los JWT son comúnmente utilizados para autenticación y autorización en aplicaciones web y servicios.

Algunas características clave de la biblioteca jwt-go incluyen:
-Compatibilidad con el estándar JWT
-Soporte para firmado y verificación
-Fácil creación de tokens
-Verificación de tokens
-Soporte para claims estándar y personalizados
-Gestión de expiración y emisión