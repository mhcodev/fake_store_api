# FAKE STORE API

## User Endpoints
### GET /api/v1/users
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
        "users": [
            {
                "id": 3,
                "userTypeID": 2,
                "name": "Shannon",
                "email": "Shannon_Hegmann@gmail.com",
                "password": "12243a1d90981d89c50d76f93d344a2e55a8f65bd27bdc53d8743927fe6537c8",
                "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/724.jpg",
                "phone": "+1 (732) 597-8261",
                "status": 1,
                "createdAt": "2024-09-23T08:19:49Z",
                "updatedAt": "2024-09-23T08:19:49Z"
            }
        ]
      }
     ```
### GET /api/v1/users/{id}
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
              "name": "Sadie",
              "email": "Sadie.West@yahoo.com",
              "password": "7669ac4d663d618438006c1d51750d1386722c65bbe45335de8dbdbedcb97a61",
              "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/742.jpg",
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
        "userTypeID": 3,
        "name": "Milena",
        "email": "mila.26@example.com",
        "password": "PassWordSecret",
        "avatar": "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png",
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
              "id": 122,
              "userTypeID": 3,
              "name": "Milena",
              "email": "mila.26@example.com",
              "password": "$2a$10$tWFU3Z363uycuk3ekfMbD.IqqRiLNBH3BoBvCoR2O.XwtQDwzUlHy",
              "avatar": "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png",
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
        "email": "sadie.west@yahoo.com"
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
        "userTypeID": 3,
        "name": "Milena Z.",
        "email": "mila.26@example.com",
        "password": "PassWordSecret",
        "avatar": "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png",
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
              "id": 122,
              "userTypeID": 3,
              "name": "Milena Z.",
              "email": "mila.26@example.com",
              "password": "$2a$10$sNaXildzZZX2aPG.uysyf.XNZ15ygo9PPQB0.dLnRmgK6OLngUW.e",
              "avatar": "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png",
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
### GET /api/v1/category
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
                  "name": "Electronics",
                  "imageURL": "https://caring-sunlight.info",
                  "status": 0
              },
              {
                  "id": 2,
                  "name": "Forniture",
                  "imageURL": "https://self-reliant-undertaker.org/",
                  "status": 0
              }
          ],
          "code": 200,
          "success": true
      }
     ```
### GET /api/v1/category/:id
- **Description**: Fetch category by ID
- **URL Parameters**:
  - `id` (integer, required): id of the category
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
      {
          "category": {
              "id": 2,
              "name": "Forniture",
              "imageURL": "https://self-reliant-undertaker.org/",
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
        "category": {
            "name": "Tecnology",
            "imageURL": "https://totebagfactory.com/cdn/shop/products/Cotton-Canvas-Tote-Babgs-Natural.jpg?v=1552815941&width=1214"
        }
    }
    ```
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
      {
          "category": {
              "id": 39,
              "name": "Tecnology",
              "imageURL": "https://totebagfactory.com/cdn/shop/products/Cotton-Canvas-Tote-Babgs-Natural.jpg?v=1552815941&width=1214",
              "status": 1
          },
          "code": 200,
          "success": true
      }
     ```
### PUT /api/v1/category/:id
- **Description**: Update a category by ID
- **URL Parameters**:
  - `id` (string, required): id of the category
- **JSON Body Parameters**:
  - `name` (string, required): Name of the category
  - `imageURL` (string, required): Image of the category
 - Example
    ```json
    {
        "category": {
            "name": "Tecnology v2",
            "imageURL": "https://totebagfactory.com/cdn/shop/products/Cotton-Canvas-Tote-Babgs-Natural.jpg?v=1552815941&width=1214"
        }
    }
    ```
- **Response**:
   - **Status Code**: 200 OK
   - **Body**:
     ```json
      {
          "category": {
              "id": 39,
              "name": "Tecnology v2",
              "imageURL": "https://totebagfactory.com/cdn/shop/products/Cotton-Canvas-Tote-Babgs-Natural.jpg?v=1552815941&width=1214",
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