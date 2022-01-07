# FIFA 21 Ultimate Team
![Logo](https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB)

![Logo](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

### Estructura del proyecto
    .
    ├── backend                    # Archivos del backend desarrollado en Go
    ├── frontend                   # Archivos del fronted desarrollado en React
    └── README.md
 
## Configuraciones para GO
## Variables de entorno para GO

Para ejecutar este proyecto se necesita las siguiente variables de entorno en el archivo .env

`X_API_KEY`  
`USERNAME`  
`PASSWORD`  
`HOSTNAME`  
`DBNAME`  
`PAGINATION`   
`REWRITEDATA` 

La Variable `REWRITEDATA` por defecto esta en 1, lo que significa que se **cargaran los datos de la api oficial de FIFA y formatear los datos para almacenarla en la base de datos local.**  
`Puerto por defecto del servidor: ` **8000**

### Estructura del backend
    backend
        ├── api_rest                    # Servidor api rest
        ├── cmd                         # Carpeta de main.go 
        ├── db                          # Modelo de la bases de datos y coneccion
        ├── rest_fut21                  # Consumo de la api oficial
        └── .env                        # Variables de entorno
        └── go.mod
        └── go.sum
 

## Variables de entorno para React

Para ejecutar este proyecto se necesita las siguiente variables de entorno en el archivo .env

`REACT_APP_URL_PLAYER`  
`REACT_APP_URL_TEAM`

### Estructura del frontend
    frontend
        ├── web-fut21                    # Recursos raiz
            ├── node_modules             # Dependecias
            ├── public                   # Punto de entrada
            ├── src                      # Recursos locales
            └── .env                     # Variables locales
            └── .gitignore
            └── package-lock.json
            └── package.json
            

