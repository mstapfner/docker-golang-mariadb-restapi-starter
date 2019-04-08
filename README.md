![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)
# Docker-Compose - Golang-Mariadb-Web-Starterkit
Starterkit for a Golang Web-Server with Connection to MariaDB (MySQL Engine) and a REST-API 

Docker-Compose based Setup including installation and startup script
It uses the following technologies:
* Docker-Compose for the Deployment of all components
* MariaDB (with MySQL Database engine) running in its own container
* Golang application that connects to the database
* GORM (Object-Relational-Mapper) for storing and retrieving data from database
* Gorilla Mux as a HTTP-Framework for serving the REST API
* Govendor as a vendoring tool

To simply start the application, use: `sudo sh startup.sh`. 
This script includes the Docker and Docker-compose instalaltion for Ubuntu. 
For a simple run use `docker-compose up`

The valid HTTP Requests can be found under: `api_requests.postman_collection.json`.
It can be imported into POSTMAN.

