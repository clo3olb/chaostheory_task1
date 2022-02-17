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
git clone https://github.com/clo3olb/chaostheory_task1.git
```

```sh
cd chaostheory_task1/
```

Use commands as follows to build docker image and create container.

```sh
docker build -t test-server .
```

```sh
docker run --rm -p 80:80 test-server
```

# REST API

The REST API to the JSON Server is described below.

## List Data

Lists all the data in the database as an array

### Request

|  Path   | Method |
| :-----: | :----: |
| `/list` |  GET   |

```sh
curl -X GET http://localhost:80/list
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

|  Path  | Method |
| :----: | :----: |
| `/add` |  POST  |

```sh
curl -X POST \
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

# RESTful Approaches for Data

## list

Adds data to the database.

### Request

|  Path   | Method |
| :-----: | :----: |
| `/data` |  GET   |

```sh
curl -X POST http://localhost:80/data
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

## Create(POST)

Adds data to the database.

### Request

|     Path      | Method |
| :-----------: | :----: |
| `/data/{key}` |  POST  |

```sh
curl -X POST \
   -d '{"value": "sample value"}'
   http://localhost:80/data/samplekey
```

### Payload

| Parameter | type   |
| --------- | ------ |
| value     | string |

### Response

```
{
  "timestamp": "2022-02-17T14:16:19Z",
  "key": "samplekey",
  "value": "some other value"
}
```

## Read(GET)

Adds data to the database.

### Request

|     Path      | Method |
| :-----------: | :----: |
| `/data/{key}` |  GET   |

```sh
curl -X GET http://localhost:80/data/samplekey
```

### Response

```
{
  "timestamp": "2022-02-17T14:16:19Z",
  "key": "samplekey",
  "value": "some other value"
}
```

## Update(PUT)

Adds data to the database.

### Request

|     Path      | Method |
| :-----------: | :----: |
| `/data/{key}` |  PUT   |

```sh
curl -X PUT \
   -d '{"value": "sample value"}'
   http://localhost:80/data/samplekey
```

### Payload

| Parameter | type   |
| --------- | ------ |
| value     | string |

### Response

```
{
  "data": "Data updated."
}
```

## Read(DELETE)

Adds data to the database.

### Request

|     Path      | Method |
| :-----------: | :----: |
| `/data/{key}` | DELETE |

```sh
curl -X DELETE http://localhost:80/data/samplekey
```

### Response

```
{
  "data": "Data deleted."
}
```
