

## Intoduction
If I briefly describe this application infrastructure, it saves, updates, deletes and lists a user's information, any addition, update, deletion is stored in each event database, the relevant exchange publishes the event related to a separate application on rabbitmq and consumes these events. There are separate applications.

## Run

```
docker-compose up -d
```

## Documentation 

1 - [API Readme](cmd/usersapi/README.md)

2 - [Producer Readme](cmd/producer/README.md)

3 - [Consumer Readme](cmd/consumer/README.md)


## Collection schemes

users

```json
  {
    "_id": "c9c88101-b2f1-41bd-b671-05996a12ebdf",
    "country": "string",
    "created_at": {"$date": "2022-02-09T13:54:19.756Z"},
    "email": "string",
    "first_name": "string",
    "last_name": "string",
    "nickname": "string",
    "password": "$2a$10$lLLU0sTOaIIWdU5sz9FPbezZa8oArU49E4vdaoKAtD3djGDkvtThu"
  }
```

events
```json
  {
    "_id": {"$oid": "6203c78bfb23228a37342a3a"},
    "correlationId": "string",
    "createdAt": {"$date": "2022-02-09T13:54:19.777Z"},
    "eventType": "Created",
    "exchange": "User:Created",
    "payload": "{\"id\":\"c9c88101-b2f1-41bd-b671-05996a12ebdf\",\"firstName\":\"string\",\"lastName\":\"string\",\"nickName\":\"string\",\"password\":\"$2a$10$lLLU0sTOaIIWdU5sz9FPbezZa8oArU49E4vdaoKAtD3djGDkvtThu\",\"email\":\"string\",\"country\":\"string\",\"createdAt\":\"2022-02-09T13:54:19.7561881Z\"}"
  }

```

eventhistories
```json
  {
    "_id": {"$oid": "6203c78bfb23228a37342a3a"},
    "correlationId": "string",
    "createdAt": {"$date": "2022-02-09T13:54:19.777Z"},
    "eventType": "Created",
    "exchange": "User:Created",
    "payload": "{\"id\":\"c9c88101-b2f1-41bd-b671-05996a12ebdf\",\"firstName\":\"string\",\"lastName\":\"string\",\"nickName\":\"string\",\"password\":\"$2a$10$lLLU0sTOaIIWdU5sz9FPbezZa8oArU49E4vdaoKAtD3djGDkvtThu\",\"email\":\"string\",\"country\":\"string\",\"createdAt\":\"2022-02-09T13:54:19.7561881Z\"}"
  }

```

## Transactional Outbox Pattern

[Source]("https://microservices.io/patterns/data/transactional-outbox.html")



## Infastructure


<center><img src="./docs/diagram.png" /></center>

<hr>

```license
MIT License

Copyright (c) 2022 Ferhat Candaş

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```
