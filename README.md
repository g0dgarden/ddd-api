# ddd-api

Sample Application Structure for Domain-Driven Design.


## Installation

```bash

$ go get -u github.com/labstack/echo
$ go get -u github.com/valyala/fasthttp
$ go get -u github.com/BurntSushi/toml
$ go get -u gopkg.in/gorp.v1
$ go get -u github.com/go-sql-driver/mysql

```

## Exsample Database Scheme

```bash
CREATE DATABASE `sample` /*!40100 DEFAULT CHARACTER SET utf8 */;

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `status` varchar(30) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
```

## HTTP Request Sample

### POST


### GET a record

```bash

curl -H "Content-Type: application/json" -X GET http://localhost:8888/v1/users/1

# =>
{"Id":1,"Name":"suzuki","Email":"ddd@exsample.com","Status":"active","Created_at":"2016-10-24T01:15:05+09:00","Updated_at":"2016-10-24T01:15:05+09:00"}

```

### GET records

```bash

curl -H "Content-Type: application/json" -X GET http://localhost:8888/v1/users

# =>
[
    {"Id":1,"Name":"suzuki","Email":"ddd@exsample.com","Status":"active","Created_at":"2016-10-24T01:15:05+09:00","Updated_at":"2016-10-24T01:15:05+09:00"},
    {"Id":2,"Name":"saitou","Email":"ddd@exsample.com","Status":"inactive","Created_at":"2016-10-24T01:15:21+09:00","Updated_at":"2016-10-24T01:15:21+09:00"}
]

```
