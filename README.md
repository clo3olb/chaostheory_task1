# Chaos Theory Internship - Take Home Task.

### Hyeonwoo KIM(clo3olb)

-   City Universiry of Hong Kong | Computer Science
-   clo3olb@gmail.com
-   josephkim3-c@my.cityu.edu.hk
-   +852 6465-9192 (Hong Kong)
-   +82 10-9192-9527 (Korea)

## Getting Started

To start using this package, install Docker and use `git clone`:

```sh
$ git clone github.com/clo3olb/chaostheory_task1
```

Use commands as follows.

```sh
$ docker build -t test-server .
```

```sh
$ docker run --rm -p 80:80 test-server
```

# REST API

The REST API to the JSON Server is described below.

## Documentation(Home)

Displays paths and documentation of the REST API

### Request

```sh
$ curl -X GET http://localhost:80/
```

### Response

```
[
   {
      "path":"http://localhost:80/",
      "method":"GET",
      "description":"Displays paths and documentation of the REST API"
   },
   {
      "path":"http://localhost:80/list",
      "method":"GET",
      "description":"Lists all the data in database as an array"
   },
   {
      "path":"http://localhost:80/add",
      "method":"POST",
      "description":"Adds data to the database. Example Format : { \"key\": \"string\", \"value\": \"string\" }"
   }
]
```

## List Data

Lists all the data in database as an array

### Request

```sh
$ curl -X GET http://localhost:80/list
```

### Response

```
[
   {
      "timestamp":"2022-02-16T17:01:27Z",
      "key":"another sample key",
      "value":"another sample value"
   },
   {
      "timestamp":"2022-02-16T17:00:30Z",
      "key":"sample key",
      "value":"sample value"
   }
]
```

## Add Data

Adds data to the database.

### Request

| Parameter | type   |
| --------- | ------ |
| key       | string |
| value     | string |

```sh
$ curl -X POST -H "Content-Type: application/json" \
    -d '{"key": "sample key", "value": "sample value"}' \
    http://localhost:80/add
```

### Response

```
{
    "data":"Data created."
}
```
