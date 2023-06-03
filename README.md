# Erply-test-task

This project is a Golang-based API endpoint/middleware that interacts with the Erply API to read and write customer data. It provides simple authentication of requests using a token and utilizes a local database for storage for getting customers.

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Documentation](#api-documentation)
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
- Clcck on the **Swagger Docs** button in the UI.

## Testing
Unit tests can be run in the following directories:
- database
- utils

To run unit tests write command:
    ```go test``` or ```go test -v``` for the verbose version.
