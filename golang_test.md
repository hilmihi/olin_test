# Golang Test

## Prerequisites

Before running the application, ensure you have the following installed:

- [Docker](https://www.docker.com/) (for containerization)
- [Go](https://golang.org/) (for running the backend service locally)

## Running Backend Service
### Option 1: Local Development

1. Navigate to the './olin_test' directory.
2. Run the Backend service:

    ```bash
    go run ./app/main.go
    ```
### Option 2: Docker
1. Navigate to the './olin_test' directory.
2. Build the Backend image:

    ```bash
    docker build -t diyo_be .
    ```
3. Run the Backend image:

    ```bash
    docker run -p 8912:8912 --env-file .env diyo_be
    ```


## Accessing the Application
- Backend service: http://localhost:9000
- Hit endpoint POST http://localhost:9000/soal-1 for the question no 1, with body like this:
    ```bash
    {
        "nums": [2,7,11,15],
        "target": 9
    }
    ```

- Hit endpoint POST http://localhost:9000/soal-2 for the question no 2 (complexity is O(n^2)), with body like this:
    ```bash
    {
        "nums": [-1,0,1,2,-1,-4]
    }
    ```
 
- Hit endpoint POST http://localhost:9000/soal-3 for the question no 3, with body like this:
    ```bash
    {
        "str": "wordgoodgoodgoodbestword",
        "words": ["word", "good", "best", "word"]
    }
    ```   

- Or you can import from postman collection in file OLIN Test.postman_collection.json