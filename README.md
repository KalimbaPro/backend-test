# Japhy Backend Test in Golang

## Technical Stack
- Go
- Docker
- MySQL

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Docker](https://www.docker.com/products/docker-desktop/)
- [Git](https://git-scm.com/downloads)

## Tasks
you are a backend developer in a pet food startup. A new functionality will be implemented, we want to be able to easily manage the breeds of dogs and cats that are registered in the database in our back office,
you must implement a CRUD api to manage the breeds of pets. The breeds are stored in a CSV file located at `./breeds.csv`.
the aim of this test is to demonstrate backend development skills using the Go programming language. The application implements a simple REST API for managing resources
you are free to take initiatives and make improvements to the codebase.
Have fun and good luck!

### you need to implement the following tasks:
- create a new table in the core database to store the breeds of pets, to do this you must create a new migration file in the `database_actions` directory.
- store the breeds of pets in the database (list of breeds are on `breeds.csv`).
- implement CRUD functionality for the breeds resource (GET, POST, PUT, DELETE).
- search functionality to filter breeds based on pet characteristics (weight and species).


## Installation

1. Fork the project repository
2. Copy the `.env.example` file to `.env`
3. Build the application `docker compose build`
4. Run docker compose to start the application `docker compose up -d`
5. Once the application is up and running, you can access the REST API at http://localhost:50010. Use tools like Postman or curl to interact with the API.
6. `curl -v http://localhost:50010/health` to ensure your application is running.
7. send us the link to your repository with the api.

# Japhy Backend Test in Golang

## Technical Stack
- Go
- Docker
- MySQL

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Docker](https://www.docker.com/products/docker-desktop/)
- [Git](https://git-scm.com/downloads)

## Tasks
you are a backend developer in a pet food startup. A new functionality will be implemented, we want to be able to easily manage the breeds of dogs and cats that are registered in the database in our back office,
you must implement a CRUD api to manage the breeds of pets. The breeds are stored in a CSV file located at `./breeds.csv`.
the aim of this test is to demonstrate backend development skills using the Go programming language. The application implements a simple REST API for managing resources
you are free to take initiatives and make improvements to the codebase.
Have fun and good luck!

### you need to implement the following tasks:
- create a new table in the core database to store the breeds of pets, to do this you must create a new migration file in the `database_actions` directory.
- store the breeds of pets in the database (list of breeds are on `breeds.csv`).
- implement CRUD functionality for the breeds resource (GET, POST, PUT, DELETE).
- search functionality to filter breeds based on pet characteristics (weight and species).


## Installation

1. Fork the project repository
2. Copy the `.env.example` file to `.env`
3. Build the application `docker compose build`
4. Run docker compose to start the application `docker compose up -d`
5. Once the application is up and running, you can access the REST API at http://localhost:50010. Use tools like Postman or curl to interact with the API.
6. `curl -v http://localhost:50010/health` to ensure your application is running.
7. send us the link to your repository with the api.

8. # API Endpoints

## 1. **GET /v1/**
Retrieve all pet breeds.

- **Method**: `GET`
- **Response Example**:
    ```json
    [
      {
        "id": 1,
        "species": "dog",
        "pet_size": "small",
        "name": "affenpinscher",
        "average_male_adult_weight": 6000,
        "average_female_adult_weight": 5000
      },
      {
        "id": 2,
        "species": "dog",
        "pet_size": "medium",
        "name": "miniature_american_shepherd",
        "average_male_adult_weight": 12000,
        "average_female_adult_weight": 12000
      }
    ]
    ```

## 2. **GET /v1/{id}**
Retrieve a specific breed by its ID.

- **Method**: `GET`
- **URL Parameter**:
    - `id`: The breed's unique ID.
- **Response Example**:
    ```json
    {
      "id": 1,
      "species": "dog",
      "pet_size": "small",
      "name": "affenpinscher",
      "average_male_adult_weight": 6000,
      "average_female_adult_weight": 5000
    }
    ```

## 3. **POST /v1/**
Create a new pet breed.

- **Method**: `POST`
- **Request Body Example**:
    ```json
    {
      "species": "dog",
      "pet_size": "small",
      "name": "poodle",
      "average_male_adult_weight": 10000,
      "average_female_adult_weight": 8000
    }
    ```
- **Response Example**:
    ```json
    {
      "id": 4,
      "species": "dog",
      "pet_size": "small",
      "name": "poodle",
      "average_male_adult_weight": 10000,
      "average_female_adult_weight": 8000
    }
    ```

## 4. **PUT /v1/{id}**
Update a specific breed by its ID.

- **Method**: `PUT`
- **URL Parameter**:
    - `id`: The breed's unique ID.
- **Request Body Example**:
    ```json
    {
      "species": "dog",
      "pet_size": "small",
      "name": "poodle",
      "average_male_adult_weight": 10500,
      "average_female_adult_weight": 8500
    }
    ```
- **Response Example**:
    ```json
    {
      "id": 4,
      "species": "dog",
      "pet_size": "small",
      "name": "poodle",
      "average_male_adult_weight": 10500,
      "average_female_adult_weight": 8500
    }
    ```

## 5. **DELETE /v1/{id}**
Delete a specific breed by its ID.

- **Method**: `DELETE`
- **URL Parameter**:
    - `id`: The breed's unique ID.
- **Response Example**:
    - `Breed with id: 2 deleted successfully`

## 6. **GET /v1/search**
Search for breeds based on pet characteristics (weight and species).

- **Method**: `GET`
- **Query Parameters**:
    - `species`: (optional) Filter by species (e.g., `dog`, `cat`).
    - `weight`: (optional) Filter by pet weight.
- **Example Request**:
    ```bash
    curl -X GET "http://localhost:50010/api/breeds/search?species=dog&weight=5000
    ```
- **Response Example**:
    ```json
    [
      {
        "id": 1,
        "species": "dog",
        "pet_size": "small",
        "name": "affenpinscher",
        "average_male_adult_weight": 6000,
        "average_female_adult_weight": 5000
      }
    ]
    ```

