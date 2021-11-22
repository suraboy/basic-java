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

```
.
├── cmd                     # Application entry point
├── config                  # Application configuration
├── docker                  # Build files for instance Dockerfile.
├── infrastructure          # Database schema example.
├── internal                # Golang's standard internal package
│   ├── core     
│   │   ├── domain          # The models' structure and rules for the application.
│   │   ├── port            # Act like interactors to connect to the outside of the core.
│   │   └── service         # The business logic of the application.
│   │
│   ├── handler             # The handlers, to handle the incoming requests.
│   └── repostiory          # The repositories such as Database, a microservice API exposed via gRPC or REST, or just a simple CSV file.
│   └── routes              # The route groups, 
│
├── pkg                     # utility packages.
└── protocol                # The protocols to serve for client.
```

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
