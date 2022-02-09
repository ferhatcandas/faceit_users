

## Intoduction
This api application is used for user creation, update, deletion and listing.

## Folder Structure

```
├── api/httputils
├── cmd
│   └── usersapi/*
│   └── main.go
├── configs
│   └── api.yaml
├── internal
│   └── usersapi/*
├── pkg/*
└── main.go
```

## API Doc

http://localhost:8000/swagger/index.html

<hr>

## API  
* **Description:** Get users with filter
* **URL**
  http://localhost:8000/users?country=UK&pageIndex=1&pageSize=20
* **Method:** `GET` 
 
   **Required:**
   `country`

   **Non Required:**
   - `pageIndex` default is 1
   - `pageSize`  default is 20

* **Success Response:**
   * **Code:** 200 <br />  
   * **Body:** 
```json 
[
  {
    "id": "c9c88101-b2f1-41bd-b671-05996a12ebdf",
    "firstname": "string",
    "lastname": "string",
    "nickname": "string",
    "email": "string",
    "country": "string"
  },
  {
    "id": "7d97ef62-0225-45ce-874e-b4ca947e552d",
    "firstname": "string",
    "lastname": "string",
    "nickname": "string2",
    "email": "string",
    "country": "string"
  }
]
```

* **Error Response:**
  * **Code:** 500 Internal Server Error<br />
  * **Body:** string 
  * **Code:** 400 Bad Request <br />
  * **Body:** string 
    

* **Sample Call:**

```curl
curl --location --request GET 'http://localhost:8000/users?country=string&pageIndex=1&pageSize=20' 
```

<hr>

* **Description:** Creates a user if nickname is unique
* **URL**
  http://localhost:8000/users
* **Method:** `POST` 
 
   **Request Body:**
 
 ```json
{
  "country": "string",
  "email": "string",
  "firstname": "string",
  "lastname": "string",
  "nickname": "string",
  "password": "string"
}
 ```

* **Success Response:**
   * **Code:** 201 <br />  
* **Error Response:**
  * **Code:** 500 Internal Server Error<br />
  * **Body:** string 
  * **Code:** 400 Bad Request<br />
  * **Body:** string 
  * **Code:** 409 Conflict<br />
  * **Body:** string 

* **Sample Call:**
```curl
 curl -X 'POST' \
  'http://localhost:8000/users' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "country": "string",
  "email": "string",
  "firstname": "string",
  "lastname": "string",
  "nickname": "string",
  "password": "string"
}'
```
 
<hr>


* **Description:** Creates a user if nickname is unique
* **URL**
  http://localhost:8000/users/{id}
* **Method:** `PATCH` 
 
   **Request Body:**
 
 ```json
{
  "country": "string",
  "firstname": "string",
  "lastname": "string"
}
 ```

* **Success Response:**
   * **Code:** 200 <br />  
* **Error Response:**
  * **Code:** 500 Internal Server Error<br />
  * **Body:** string 
  * **Code:** 400 Bad Request<br />
  * **Body:** string 
  * **Code:** 404 Not Found<br />
  * **Body:** string 

* **Sample Call:**
```curl
 curl -X 'PATCH' \
  'http://localhost:8000/users/6c47b457-75db-49b9-b733-2edf0b3bde74' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "country": "string",
  "firstname": "string",
  "lastname": "string"
}'
```
 
<hr>

* **Description:** Deletes a user if id is exist
* **URL**
  http://localhost:8000/users/{id}
* **Method:** `DELETE` 

  **Required:**
   `id`
* **Success Response:**
   * **Code:** 204 <br />  
* **Error Response:**
  * **Code:** 500 Internal Server Error<br />
  * **Body:** string 
  * **Code:** 404 Not Found<br />
  * **Body:** string 
* **Sample Call:**

```curl
curl -X 'DELETE' \
  'http://localhost:8000/users/6c47b457-75db-49b9-b733-2edf0b3bde74' 
```

<hr>

* **Description:** Health check
* **URL**
  http://localhost:8000/hc
* **Method:** `GET` 
 

* **Success Response:**
   * **Code:** 500 <br />  
   * **Body:** 
```json 
{
  "mongo": "FAIL",
  "serverPort": "8000",
  "database": "faceit"
}
```

*   * **Code:** 200 <br />  
*  * **Body:** 
```json
{
  "mongo": "OK",
  "serverPort": "8000",
  "database": "faceit"
}
```
* **Sample Call:**

```curl
curl -X 'GET' \
  'http://localhost:8000/hc' 
```

<hr>
