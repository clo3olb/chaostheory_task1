# Chaos Theory Internship - Take Home Task.

### Hyeonwoo KIM(clo3olb)

-   City Universiry of Hong Kong | Computer Science
-   clo3olb@gmail.com
-   josephkim3-c@my.cityu.edu.hk
-   +852 6465-9192 (Hong Kong)
-   +82 10-9192-9527 (Korea)

## Getting Started

To start using this package, install Docker and use `git clone` and move to cloned repository:

```sh
$ git clone https://github.com/clo3olb/chaostheory_task1.git
```

```sh
$ cd chaostheory_task1/
```

Use commands as follows to build docker image and create container.

```sh
$ docker build -t test-server .
```

```sh
$ docker run --rm -p 80:80 test-server
```

# REST API

The REST API to the JSON Server is described below.

## Documentation(Home)

Displays paths and descriptions of each endpoints

### Request

| Path | Method |
| :--: | :----: |
| `/`  |  GET   |

```sh
$ curl -X GET http://localhost:80/
```

### Response

```
[
   {
      "path":"http://localhost:80/",
      "method":"GET",
      "description":Displays paths and descriptions of each endpoints"
   },
   {
      "path":"http://localhost:80/list",
      "method":"GET",
      "description":"Lists all the data in the database as an array"
   },
   {
      "path":"http://localhost:80/add",
      "method":"POST",
      "description":"Adds data to the database. Example Format : { \"key\": \"string\", \"value\": \"string\" }"
   }
]
```

## List Data

Lists all the data in the database as an array

### Request

|  Path   | Method |
| :-----: | :----: |
| `/list` |  GET   |

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

|  Path  | Method | Payload |
| :----: | :----: | :-----: |
| `/add` |  POST  |  true   |

```sh
$ curl -X POST \
    -H "Content-Type: application/json" \
    -d '{"key": "sample key", "value": "sample value"}' \
    http://localhost:80/add
```

### Payload

| Parameter | type   |
| --------- | ------ |
| key       | string |
| value     | string |

### Response

```
{
    "data":"Data created."
}
```
