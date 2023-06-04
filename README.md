# Erply-test-task

This project is a Golang-based API endpoint/middleware that interacts with the Erply API to read and write customer data. It provides simple authentication of requests using a token and utilizes a local database (sqlite3) for storage for getting customers.

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Documentation](#api-documentation)
- [Authentication](#authentication)
- [Local storage](#local-storage)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:
   ```bash git clone https://github.com/jserva90/Erply-test-task.git```
2. Change into the project directory:   
   ```cd Erply-test-task```
3. Install dependencies:
    ```go mod download```

## Configuration
Setup a 32-byte Secret Key for encryption in the .env file (For example):
    ```0123456789abcdef0123456789abcdef```

## Usage
1. Start the application:
    ```make```
2. The application will start at: [http://localhost:8080](http://localhost:8080)

## API Documentation
The Swagger API documentation for the endpoints can be accessed using the following methods:
- Open http://localhost:8080/swagger/index.html in your browser.
- Click on the **Swagger Docs** button in the UI.

## Authentication
Authentication in this project requires providing the Erply client code, username (mail address), and password. Upon successful authentication, the login credentials and session key are encrypted and stored in the SQLite3 database. For subsequent requests to get or add customers, the encrypted session key and client code are retrieved from the database, decrypted, and used to authenticate the request to the Erply API. This ensures secure storage and retrieval of authentication information, facilitating seamless and secure API interactions.

## Local Storage
In this project, I utilize SQLite3 as a local storage mechanism. When a customer is requested from the Erply API, the data is stored in the SQLite3 database. Subsequent requests for the same customer within 10 minutes retrieve the data from the database instead of making a new API request. After 10 minutes, a new API request is made to update the local storage with the latest customer data. This approach improves response time and reduces unnecessary API calls.

## Testing
Unit tests can be run in the following directories:
- database
- utils

To run unit tests, go to the corresponging directory and run command:
    ```go test``` or ```go test -v``` for the verbose version.

#### Written in [Go](https://go.dev/) version 1.20

##### Author [Jserva90](https://github.com/jserva90)