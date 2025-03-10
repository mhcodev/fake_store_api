# FAKE STORE API
## URL BASE: https://fake-store-api.mhdeploy.com

![image](https://github.com/user-attachments/assets/a79e871b-1120-4e03-a649-bc2d667efccf)
![image](https://github.com/user-attachments/assets/791c13df-008c-4824-bb41-20d955c0630d)

## This project contains these tools:
- Golang
- PostgreSQL
- Prometheus
- Grafana
- Redis
- Docker Compose

## How to run the project

Install make on Linux (Debian/Ubuntu)
```bash
sudo apt install make
```

Install htpasswd on Linux (Debian/Ubuntu) for Prometheus password
```bash
sudo apt install apache2-utils
```

Install PostgreSQL 16
>[!IMPORTANT]
>* sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list
>* curl -fsSL https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo gpg --dearmor -o /etc/apt/trusted.gpg.d/postgresql.gpg
>* sudo apt update
>* sudo apt install postgresql-16 postgresql-contrib-16
>* Configure your user & password db
>* sudo systemctl start postgresql
>* sudo systemctl enable postgresql
>* Change ".env.example" to ".env"
>* Configure ".env" file with your credentials
>* Exec the schemas file `db/schemas.sql` in PostgreSQL DB


Run this command to build the image
```bash
make docker-build
```

Finally run the container
```bash
make docker-run
```

## Run in development
Install Golang dependencies
```bash
go mod tidy
```
Ensure you have air package installed
```bash
go install github.com/air-verse/air@latest
```

Run go api with air
```bash
air
```


## Auth Endpoints
### GET /api/v1/auth/data
- **Description**: Get data from token
- **Header Parameters**:
  - `Authorization` (string, required): Bearer token generated when user logged in
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
      {
          "code": 200,
          "data": {
              "exp": 1732444079,
              "userID": 115,
              "userTypeID": 3
          },
          "success": true
      }
     ```
### POST /api/v1/auth/login
- **Description**: Login using email & password
- **JSON Body Parameters**:
  - `email` (string, required): Email of the user
  - `password` (string, required): Password of the user
  - Example
    ```json
    {
        "email": "Jill_Rogahn94@gmail.com",
        "password": "123456"
    }
    ```
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
      {
          "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDAwOTYzMDcsInVzZXJJRCI6MSwidXNlclR5cGVJRCI6Mn0.kRdGvmjymhRjc2NGcR02QciyHTpOlIWUPdhev4-hx0I",
          "code": 200,
          "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDA3MDAyMDcsInVzZXJJRCI6MSwidXNlclR5cGVJRCI6Mn0.iVHzXaixP4dGeHu9T7-yOu_aD0IauGEOyjHz98syEDc",
          "success": true,
          "user": {
              "id": 1,
              "userTypeID": 2,
              "name": "Jill",
              "email": "Jill_Rogahn94@gmail.com",
              "password": "$2a$10$Hhv4LRJQZd62XYxu/KbAzO3XAWJDJkYRe0GvlCWbeFWxvnVF9A08u",
              "avatar": "http://172.22.56.31:3000/uploads/PROFILE2.png",
              "phone": "+1 (993) 554-2275",
              "status": 1,
              "createdAt": "2024-11-24T19:16:51Z",
              "updatedAt": "2024-11-24T19:16:51Z"
          }
      }
     ```
### POST /api/v1/auth/refresh
- **Description**: Generate a new access token for user
- **JSON Body Parameters**:
  - `refreshToken` (string, required): Refresh token generated when user logged in
  - Example
    ```json
    {
      "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzA2MTg3NzksInJlZnJlc2giOnRydWUsInVzZXJfaWQiOjExNX0.E1QsBT9yU-oJcOvJMFVIdYwOrUOTnRicKGsY643rqFg"
    }
    ```
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
      {
          "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzI0NDQ1MTYsInVzZXJJRCI6MTE1LCJ1c2VyVHlwZUlEIjozfQ.g4lXy_M0JJ5QMOeCJBgrkDS_muhWmy6cnzGq6dR6IKM",
          "code": 200,
          "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzMwNDg0MTYsInVzZXJJRCI6MTE1LCJ1c2VyVHlwZUlEIjozfQ.K7n9ZoxajuIUsIl3Ifwjh3L9q6rk2u_9oZrG7lkcgA4",
          "success": true
      }
     ```
## User Endpoints
### GET /api/v1/user - [view example](https://fake-store-api.mhdeploy.com/api/v1/user)
- **Description**: Fetch user by Params
- **URL Parameters**:
  - `limit` (integer, optional): Quantity of rows (default 15)
  - `offset` (integer, optional): Offset of rows (default 0)
  - `name` (string, optional): User's name
  - `type` (integer, optional): User's name
  - `email` (string, optional): ID of the user
  - `status` (integer, optional): ID of the user
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
      {
        "code": 200,
        "success": true,
        "count": 10,
        "users": [
            {
                "id": 1,
                "userTypeID": 2,
                "name": "Jill",
                "email": "Jill_Rogahn94@gmail.com",
                "password": "12243a1d90981d89c50d76f93d344a2e55a8f65bd27bdc53d8743927fe6537c8",
                "avatar": "http://[BASE_URL]/uploads/PROFILE2.png",
                "phone": "+1 (993) 554-2275",
                "status": 1,
                "createdAt": "2024-09-23T08:19:49Z",
                "updatedAt": "2024-09-23T08:19:49Z"
            }
        ]
      }
     ```
### GET /api/v1/user/{id} - [view example](https://fake-store-api.mhdeploy.com/api/v1/user/1)
- **Description**: Fetch user by ID
- **URL Parameters**:
  - `id` (integer, required): ID of the user
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
     {
          "code": 200,
          "success": true,
          "user": {
              "id": 1,
              "userTypeID": 1,
              "name": "Jill",
              "email": "Jill_Rogahn94@gmail.com",
              "password": "df85418c5fef51124e9f255b723d9ace2761ba44d00542a815eafb106bfd6092",
              "avatar": "http://[BASE_URL]/uploads/PROFILE2.png",
              "phone": "+1 (641) 196-0431",
              "status": 1,
              "createdAt": "2024-09-23T08:19:49Z",
              "updatedAt": "2024-09-23T08:19:49Z"
          }
      }
     ```
### POST /api/v1/users
- **Description**: Create an user
- **JSON Body Parameters**:
  - `userTypeID` (integer, required): Type or Rol of the user
  - `name` (string, required): Name of the user
  - `email` (string, required): Email of the user
  - `password` (string, required): ID of the user
  - `avatar` (string, required): Profile photo of the user
  - `phone` (string, optional): Phone of the user
  - `status` (integer, optional): Status of the user
  - Example
    ```json
    {
        "userTypeID": 1,
        "name": "Sadie",
        "email": "Sadie.West@yahoo.com",
        "password": "PassWordSecret",
        "avatar": "http://[BASE_URL]/uploads/PROFILE1.png",
        "phone": "+1 (641) 196-0431"
    }
    ```
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
     {
          "code": 200,
          "success": true,
          "user": {
              "id": 2,
              "userTypeID": 1,
              "name": "Sadie",
              "email": "Sadie.West@yahoo.com",
              "password": "7669ac4d663d618438006c1d51750d1386722c65bbe45335de8dbdbedcb97a61",
              "avatar": "http://[BASE_URL]/uploads/PROFILE1.png",
              "phone": "987123987",
              "status": 1,
              "createdAt": "0001-01-01T00:00:00Z",
              "updatedAt": "0001-01-01T00:00:00Z"
          }
      }
     ```

### POST /api/v1/user/email/is-available
- **Description**: Valid if user email for registration
- **JSON Body Parameters**:
  - `email` (string, required): Email user for registration
  - Example
    ```json
    {
        "email": "Shannon_Hegmann@gmail.com"
    }
    ```
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
     {
          "code": 200,
          "formatIsOK": true,
          "isAvailable": true, // Indicates if email is available for registration
          "success": true
      }
     ```

### PUT /api/v1/user/:id
- **Description**: Update an user by ID
- **URL Parameters**:
  - `id` (integer, required): ID of the user
- **JSON Body Parameters**:
  - `userTypeID` (integer, optional): Type or Rol of the user
  - `name` (string, optional): Name of the user
  - `email` (string, optional): Email of the user
  - `password` (string, optional): Password of the user
  - `avatar` (string, optional): Photo of the user
  - `phone` (string, optional): Phone of the user
  - `status` (integer, optional): Status of the user
  - Example
    ```json
    {
        "userTypeID": 3,
        "name": "Milena Z.",
        "email": "mila.26@example.com",
        "password": "PassWordSecret",
        "avatar": "http://[BASE_URL]/uploads/PROFILE2.png",
        "phone": "987123987",
        "status": 1
    }
    ```
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
     {
          "code": 200,
          "success": true,
          "user": {
              "id": 2,
              "userTypeID": 3,
              "name": "Milena Z.",
              "email": "mila.26@example.com",
              "password": "$2a$10$sNaXildzZZX2aPG.uysyf.XNZ15ygo9PPQB0.dLnRmgK6OLngUW.e",
              "avatar": "http://[BASE_URL]/uploads/PROFILE2.png",
              "phone": "987123987",
              "status": 1,
              "createdAt": "2024-11-09T19:20:30Z",
              "updatedAt": "2024-11-10T02:19:16Z"
          }
      }
     ```
### DELETE /api/v1/user/:id
- **Description**: Delete an user by ID
- **URL Parameters**:
  - `id` (integer, required): ID of the user
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
     {
          "code": 200,
          "msg": "user 55 deleted",
          "success": true
      }
     ```

## Category Endpoints
### GET /api/v1/category - [view example](https://fake-store-api.mhdeploy.com/api/v1/category)
- **Description**: Fetch all categories
- **URL Parameters**:
  - `limit` (integer, optional): Quantity of rows (default 15)
  - `offset` (integer, optional): Offset of rows (default 0)
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
      {
          "categories": [
              {
                  "id": 1,
                  "name": "Clothes",
                  "imageURL": "http://[BASE_URL]/uploads/category_1.jpeg",
                  "status": 1
              },
              {
                  "id": 2,
                  "name": "Games",
                  "imageURL": "http://[BASE_URL]/uploads/category_2.webp",
                  "status": 1
              },
              {
                  "id": 3,
                  "name": "PC Components",
                  "imageURL": "http://[BASE_URL]/uploads/category_3.png",
                  "status": 1
              },
              {
                  "id": 4,
                  "name": "Smartphones",
                  "imageURL": "http://[BASE_URL]/uploads/category_4.webp",
                  "status": 1
              }
          ],
          "code": 200,
          "count": 4,
          "success": true
      }
     ```
### GET /api/v1/category/:id - [view example](https://fake-store-api.mhdeploy.com/api/v1/category/1)
- **Description**: Fetch category by ID
- **URL Parameters**:
  - `id` (integer, required): id of the category
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
      {
          "category": {
              "id": 1,
              "name": "Clothes",
              "imageURL": "http://[BASE_URL]/uploads/category_1.jpeg",
              "status": 1
          },
          "code": 200,
          "success": true
      }
     ```
### POST /api/v1/category
- **Description**: Create a category
- **JSON Body Parameters**:
  - `name` (string, required): Name of the category
  - `imageURL` (string, required): Image of the category
  - Example
    ```json
    {
        "name": "Tecnology",
        "imageURL": "http://[BASE_URL]/uploads/category_5_.jpeg",
    }
    ```
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
      {
          "category": {
              "id": 5,
              "name": "Tecnology",
              "imageURL": "http://[BASE_URL]/uploads/category_5_.jpeg",
              "status": 1
          },
          "code": 200,
          "success": true
      }
     ```
### PUT /api/v1/category/:id
- **Description**: Update a category by ID
- **URL Parameters**:
  - `id` (integer, required): id of the category
- **JSON Body Parameters**:
  - `name` (string, required): Name of the category
  - `imageURL` (string, required): Image of the category
 - Example
    ```json
    {
        "name": "Tecnology v2",
        "imageURL": "http://[BASE_URL]/uploads/category_5_.jpeg",
    }
    ```
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
      {
          "category": {
              "id": 5,
              "name": "Tecnology v2",
              "imageURL": "http://[BASE_URL]/uploads/category_5_.jpeg",
              "status": 1
          },
          "code": 200,
          "success": true
      }
     ```
### DELETE /api/v1/category/:id
- **Description**: Delete a category by ID
- **URL Parameters**:
  - `id` (string, required): id of the category
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
      {
          "code": 200,
          "msg": "category deleted",
          "success": true
      }
     ```
## Product Endpoints
### GET /api/v1/product/:id - [view example](https://fake-store-api.mhdeploy.com/api/v1/product/1)
- **Description**: Returns a product by ID
- **URL Parameters**:
  - `id` (integer, required): id of the product
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
      {
          "code": 200,
          "product": {
              "id": 1,
              "categoryID": 1,
              "sku": "CLOTH12345",
              "name": "Casual T-Shirt",
              "slug": "casual-t-shirt",
              "stock": 25,
              "description": "A comfortable cotton t-shirt perfect for everyday wear.",
              "price": 19.99,
              "discount": 0.1,
              "images": [
                  "http://[BASE_URL]/uploads/clothes_1.jpeg"
              ],
              "category": {
                  "id": 0,
                  "name": "",
                  "imageURL": "",
                  "status": 0
              },
              "status": 1,
              "createdAt": "2024-11-24T23:40:44Z",
              "updatedAt": "2024-11-24T23:40:44Z"
          },
          "success": true
      }
     ```
### GET /api/v1/product - [view example](https://fake-store-api.mhdeploy.com/api/v1/product)
- **Description**: Returns a list of products
- **URL Parameters**:
  - `id` (integer, required): id of the product
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
      {
          "code": 200,
          "count": 21,
          "products": [
              {
                  "id": 1,
                  "categoryID": 1,
                  "sku": "CLOTH12345",
                  "name": "Casual T-Shirt",
                  "slug": "casual-t-shirt",
                  "stock": 25,
                  "description": "A comfortable cotton t-shirt perfect for everyday wear.",
                  "price": 19.99,
                  "discount": 0.1,
                  "images": [
                      "http://[BASE_URL]/uploads/clothes_1.jpeg"
                  ],
                  "category": {
                      "id": 0,
                      "name": "",
                      "imageURL": "",
                      "status": 0
                  },
                  "status": 1,
                  "createdAt": "2024-11-24T23:40:44Z",
                  "updatedAt": "2024-11-24T23:40:44Z"
              }
          ],
          "success": true
      }
     ```
### POST /api/v1/product
- **Description**: Create a product
- **JSON Body Parameters**:
  - `categoryID` (string, required): ID of category
  - `name` (string, required): Name of the product
  - `description` (string, required): Description of the product
  - `price` (double, required): Price of the product
  - `sku` (string, optional): SKU of the product (generated if not provided)
  - `stock` (integer, optional): Stock of the product (default 0)
  - `discount` (double, optional): Discount of the product (default 0)
  - `images` (array[string], optional): Images of the product (default empty array)
  - Example
    ```json
    {
        "categoryID": 1,
        "sku": "03973ASDQWE",
        "name": "Generic T-Shirt",
        "stock": 111,
        "description": "Gray/Blue/Red T-Shirt, great if you want to do a sport",
        "price": 100.00,
        "discount": 0.00,
        "images": [
        ]
    }
    ```
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
      {
          "code": 200,
          "product": {
              "id": 21,
              "categoryID": 1,
              "sku": "03973ASDQWE",
              "name": "Generic T-Shirt",
              "slug": "generic-granite-chair",
              "stock": 111,
              "description": "Gray/Blue/Red T-Shirt, great if you want to do a sport",
              "price": 100.00,
              "discount": 0.00,
              "images": [
                  "http://[BASE_URL]/uploads/03973ASDQWE.png",
              ],
              "category": {
                  "id": 1,
                  "name": "Clothes",
                  "imageURL": "http://[BASE_URL]/uploads/category_1.jpeg",
                  "status": 1
              },
              "status": 1,
              "createdAt": "2024-11-24T04:11:24Z",
              "updatedAt": "2024-11-24T04:11:24Z"
          },
          "success": true
      }
     ```
### PUT /api/v1/product/:id
- **Description**: Update a product by ID
- **URL Parameters**:
  - `id` (string, required): id of the product
- **JSON Body Parameters**:
  - `categoryID` (string, optional): ID of category
  - `name` (string, optional): Name of the product
  - `description` (string, optional): Description of the product
  - `price` (double, optional): Price of the product
  - `sku` (string, optional): SKU of the product (generated if not provided)
  - `stock` (integer, optional): Stock of the product (default 0)
  - `discount` (double, optional): Discount of the product (default 0)
  - `images` (array[string], optional): Images of the product (default empty array)
 - Example
    ```json
    {
        "categoryID": 1,
        "sku": "CLOTH12345",
        "name": "Casual Black T-Shirt",
        "stock": 20,
        "description": "A comfortable cotton t-shirt perfect for everyday wear.",
        "price": 19.99,
        "discount": 0.1,
        "images": [
            "http://[BASE_URL]/uploads/clothes_1.jpeg"
        ]
    }
    ```
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
      {
          "code": 200,
          "product": {
              "id": 1,
              "categoryID": 1,
              "sku": "CLOTH12345",
              "name": "Casual Black T-Shirt",
              "slug": "casual-black-t-shirt",
              "stock": 20,
              "description": "A comfortable cotton t-shirt perfect for everyday wear.",
              "price": 19.99,
              "discount": 0.1,
              "images": [
                  "http://[BASE_URL]/uploads/clothes_1.jpeg"
              ],
              "category": {
                  "id": 0,
                  "name": "",
                  "imageURL": "",
                  "status": 0
              },
              "status": 1,
              "createdAt": "2024-11-24T23:40:44Z",
              "updatedAt": "2024-11-26T00:53:05Z"
          },
          "success": true
      }
     ```
### DELETE /api/v1/product/:id
- **Description**: Delete a product by ID
- **URL Parameters**:
  - `id` (string, required): id of the product
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
      {
          "code": 200,
          "msg": "product 1 deleted",
          "success": true
      }
     ```

## File Endpoints
### POST /api/v1/file/upload
- **Description**: Create a product
- **Form Data Parameters**:
  - `images` (array[string], required): Images of the product
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
      {
          "code": 200,
          "errors": [],
          "files": [
              {
                  "id": 38,
                  "originalName": "user_4.png",
                  "filename": "HIOCJ0.png",
                  "type": ".png",
                  "baseURL": "http://[BASE_URL]",
                  "url": "http://[BASE_URL]/uploads/HIOCJ0.png",
                  "status": 0,
                  "createdAt": "0001-01-01T00:00:00Z",
                  "updatedAt": "0001-01-01T00:00:00Z"
              }
          ],
          "msg": "Files uploaded successfully",
          "success": true
      }
     ```
