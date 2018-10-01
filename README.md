# gouser
A REST API in GO to manipulate user data -- uses MUX and mysql

##Setup the database

###Postgres
```
docker run --name=postgres-gouser -e POSTGRES_PASSWORD=**** -p 5432:5432 -d postgres:latest
```
Login through the psql client:
```
psql -h localhost -U gouser
```

And run the following 3 commands:
```
Create database gouser_db;
```
```
create user gouser with encrypted password 'gouser';
```
```
grant all priveleges on gouser_db to gouser;
```

Create table
```
CREATE TABLE users (
  id SERIAL,
  name TEXT NOT NULL,
  age NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  CONSTRAINT users_pkey PRIMARY KEY (id)
);
```

###Mysql
docker pull mysql/mysql-server:latest

```
docker run --name=mysql-gouser -d mysql/mysql-server:latest
```

Get password: docker logs mysql-gouser
```
docker exec -it mysql-gouser mysql -uroot -p
```

Change the password of the root user
```
ALTER USER 'root'@'localhost' IDENTIFIED BY '****';
```

Create database and table
```
CREATE DATABASE rest_api_example;
USE rest_api_example;
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    age INT NOT NULL
);
```
```
docker stop mysql-gouser
docker start mysql-gouser
```
