# REST API Documentation

## User Routes

### Register User

- **Endpoint:** `/user/register`
- **Method:** `POST`
- **Description:** Registers a new user.
- **Request Body:**
  ```json
  {
    "username": "string",
    "email": "string",
    "password": "string"
  }

- **Response**
  
  ```json
  {
    "id": "uuid",
    "username": "string",
    "email": "string",
    "created_at": "timestamp",
    "updated_at": "timestamp"
  }
  ```


### Login User

- **Endpoint**: `/user/login`
- **Method**: POST
- **Description**: Authenticates a user.
- **Request Body**:
  ```json
  {
    "email": "string",
    "password": "string"
  }

- **Response**
  
  ```json
  {
    "token": "string",
  }
  ```
  

### Get User Info

- **Endpoint**: `/user/`
- **Method**: GET
- **Description**: Retrieves information about the authenticated user.
- **Headers**:
    - **Authorization: Bearer `<token>`**
  
- **Response**
  
  ```json
  {
    "id": "uuid",
    "username": "string",
    "email": "string",
    "photo": "string",
    "created_at": "timestamp",
    "updated_at": "timestamp"
  }
  ```

### Update User

- **Endpoint**: `/user/update`
- **Method**: `PUT`
- **Description**:  Updates the username and email of the authenticated user.
- **Headers**:
    - **Authorization: Bearer `<token>`**
  
- **Response**
  
  ```json
    {
      "id": "efed36b2-ebbf-4681-9f59-739ee687aff6",
      "username": "martis",
      "email": "mortix@gmail.com",
      "created_at": "2024-06-30T18:23:06.917144+07:00",
      "updated_at": "2024-07-01T07:51:02.162766+07:00"
    }
  ```

### Update Password

- **Endpoint:** `/user/password`
- **Method:** `PUT`
- **Description:** Updates the password of the authenticated user.
- **Headers:**
  - `Authorization: Bearer <token>`
- **Request Body:**
  ```json
  {
    "oldPassword": "string",
    "newPassword": "string"
  }
- **Response**
  
  ```json
  {
    "message": "Password updated successfully"
  }
  ```


## Photo Routes

### View Photo Profile

- **Endpoint:** `/photo/:fileName`
- **Method:** `GET`
- **Description:** Retrieves the photo profile by file name.
- **Parameters:**
  - `fileName` (path parameter): The name of the photo file to retrieve.
- **Response:**
  - **200 OK:** Returns the photo file.
  - **404 Not Found:** Photo not found.
  

### Add Photo Profile

- **Endpoint:** `/photo/upload`
- **Method:** `POST`
- **Description:** Uploads a new photo profile.
- **Headers:**
  - `Authorization: Bearer <token>`
- **Request Body:**
  - Single file upload using multipart/form-data with a file field named `file`
- **Response:**
  - **200 OK:**
    ```json
    {
      "message": "Profile photo uploaded successfully"
    }
    ```
  - **400 Bad Request:** Failed to upload photo
  - **401 Unauthorized:** Missing or invalid token
  - **404 Not Found:** User not found
  - **500 Internal Server Error:** Failed to save uploaded file or update user photo


### Update Photo Profile

- **Endpoint:** `/photo/update`
- **Method:** `PUT`
- **Description:** Updates the photo profile of the authenticated user.
- **Headers:**
  - `Authorization: Bearer <token>`
- **Request Body:**
  - Single file upload using multipart/form-data with a file field named `file`
- **Response:**
  - **200 OK:**
    ```json
    {
      "message": "Profile photo updated successfully"
    }
    ```
  - **400 Bad Request:** Failed to upload photo
  - **401 Unauthorized:** Missing or invalid token
  - **404 Not Found:** Photo not found
  - **500 Internal Server Error:** Failed to save uploaded file, delete old photo, or update photo record
  

### Delete Photo Profile

- **Endpoint:** `/photo/delete`
- **Method:** `DELETE`
- **Description:** Deletes the photo profile of the authenticated user.
- **Headers:**
  - `Authorization: Bearer <token>`
- **Response:**
  - **200 OK:**
    ```json
    {
      "message": "Profile photo deleted successfully"
    }
    ```
  - **404 Not Found:** Photo not found
  - **500 Internal Server Error:** Failed to delete photo file or photo record