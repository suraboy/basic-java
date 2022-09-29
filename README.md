# Getting Started

### Reference Documentation

For further reference, please consider the following sections:

* [Official Apache Maven documentation](https://maven.apache.org/guides/index.html)
* [Spring Boot Maven Plugin Reference Guide](https://docs.spring.io/spring-boot/docs/2.7.4/maven-plugin/reference/html/)
* [Create an OCI image](https://docs.spring.io/spring-boot/docs/2.7.4/maven-plugin/reference/html/#build-image)
* [Spring Data JDBC](https://docs.spring.io/spring-boot/docs/2.7.4/reference/htmlsingle/#data.sql.jdbc)
* [JDBC API](https://docs.spring.io/spring-boot/docs/2.7.4/reference/htmlsingle/#data.sql)
* [Spring Data JPA](https://docs.spring.io/spring-boot/docs/2.7.4/reference/htmlsingle/#data.sql.jpa-and-spring-data)

### Guides

The following guides illustrate how to use some features concretely:

* [Using Spring Data JDBC](https://github.com/spring-projects/spring-data-examples/tree/master/jdbc/basics)
* [Accessing Relational Data using JDBC with Spring](https://spring.io/guides/gs/relational-data-access/)
* [Managing Transactions](https://spring.io/guides/gs/managing-transactions/)
* [Accessing Data with JPA](https://spring.io/guides/gs/accessing-data-jpa/)
* [Accessing data with MySQL](https://spring.io/guides/gs/accessing-data-mysql/)


Database Setup
------------------------------------
You need to create database `mydb` in `http://localhost:8880` and Run the migration artisan command:
```bash
docker-compose up -d
```


Installing Dependencies via Composer
------------------------------------
Run the composer installer (m1):

```bash
arch -x86_64 /usr/local/homebrew/bin/mvn spring-boot:run
```


Running Application
------------------------------------
Open the browser
```bash
http://localhost:8081
```