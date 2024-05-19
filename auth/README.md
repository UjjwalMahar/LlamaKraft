# llamaKraft Auth Service

This Go package provides authentication functionalities for the llamaKraft application using the Gin web framework and JWT tokens.

## Features

- **Sign-Up:** Allows new users to create an account with a username, email, and password.
- **Login:** Authenticates users and returns a JWT token for session management.
- **Token Validation:** Validates the JWT token to ensure the user is authenticated.
- **Health Check:** A simple endpoint to check if the service is up and running.

## Endpoints

### Sign-Up

`POST /signup`

Registers a new user with a username, email, and password.

**Request Body:**
```json
{
  "username": "your_username",
  "email": "your_email",
  "password": "your_password"
}
```

**Responses:**
- `200 OK` if the user is successfully created.
- `400 Bad Request` if there is an error in the request body or during user creation.

### Login

`POST /login`

Authenticates the user using their username and password, and returns a JWT token.

**Request Body:**
```json
{
  "username": "your_username",
  "password": "your_password"
}
```

**Responses:**
- `200 OK` if the login is successful and a token is generated.
- `400 Bad Request` if there is an error in the request body or authentication fails.

### Validate Token

`GET /validate`

Validates the provided JWT token to ensure the user is authenticated.

**Responses:**
- `200 OK` with the user information if the token is valid.
- `401 Unauthorized` if the token is invalid or missing.

### Health Check

`GET /health`

Checks if the service is up and running.

**Responses:**
- `200 OK` with a message "UP" if the service is running.

## Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/UjjwalMahar/llamakraft
   cd llamakraft/auth
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Environment Variables:**
   Create a `.env` file and set the following environment variables:
   ```
   DB_HOST=your_database_host
   DB_USER=your_database_user
   DB_PASSWORD=your_database_password
   DB_NAME=your_database_name
   SECRET=your_jwt_secret
   ```

4. **Run the application:**
   ```bash
   go run main.go
   ```

## Running Using Kraft Cloud 

**Linux**

1. Install the [kraft tool](https://docs.kraft.cloud/quickstart/)

    ```curl -sSfL https://get.kraftkit.sh | sh```

2. Once you have installed kraft, set the KraftCloud access token you received during your on-boarding:

    ```
    export KRAFTCLOUD_TOKEN=token
    export KRAFTCLOUD_METRO=fra0 # set globally, or set via the cmd line as below
    ```
3. Lauch Instance 
    ```
    kraft cloud deploy -p 443:8080 -e 'PWD=/' .
    ```


## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.


