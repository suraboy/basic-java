# Go-fiber-api
This project use in learning Golang , If any mistake, apologize.

Get Started
-----------

#### Requirements

To run this application on your machine, you need at least:
* golang


Application flow pattern:
---------------------
https://github.com/suraboy/go-fiber-api.git

Run the docker for development:
---------------------
You can start the application and run the containers in the background, use following command inside project root:

```bash
docker-compose up -d
```
```bash
docker-compose down
```
Setup Database
------------------------------------
Check ip database contrainer
```bash
 docker inspect pg_container | grep IPAddress
```

Running Application
------------------------------------
Open the browser
```bash
http://localhost:7304
```
Open the database
```bash
http://localhost:5050
```
