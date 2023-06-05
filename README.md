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
   ```git clone https://github.com/jserva90/Erply-test-task.git```
2. Change into the project directory:   
   ```cd Erply-test-task```
3. Install dependencies:
    ```go mod download```
4. To use go-sqlite3 package you need to install gcc-5

    ### For Linux
   ```
   sudo apt-get update
   sudo apt-get install -y build-essential libsqlite3-dev
   export CC=gcc
   ```
   Alternatively you can install brew [https://brew.sh/](https://brew.sh/)
   Once brew is installed, run command: ```Brew install gcc@5```
   ### For Windows
   To install the go-sqlite3 package on Windows, you can follow these steps:
    1. Install SQLite:
        - Download the precompiled SQLite DLL from the official website [https://www.sqlite.org/download.html](https://www.sqlite.org/download.html).
        - Extract the downloaded ZIP file and note the path to the extracted DLL file.
    2. Install GCC:
        - Download and install the TDM-GCC compiler [https://jmeubank.github.io/tdm-gcc/download/](https://jmeubank.github.io/tdm-gcc/download/).
        - During the installation, make sure to select the option to add the compiler to the system's PATH.
    3. Set environment variables:
        - Open the Start menu and search for "Environment Variables."
        - Select "Edit the system environment variables."
        - In the System Properties window, click the "Environment Variables" button.
        - In the "System variables" section, click "New" to add a new variable.
        - Set the variable name as CGO_CFLAGS and the value as -I{path_to_sqlite_headers}. Replace {path_to_sqlite_headers} with the path to the folder containing the SQLite header files. This is typically the folder where you extracted the SQLite DLL in step 1.
        - Click "OK" to save the variable.
        - Click "New" again to add another variable.
        - Set the variable name as CGO_LDFLAGS and the value as -L{path_to_sqlite_dll} -lsqlite3. Replace {path_to_sqlite_dll} with the path to the folder containing the SQLite DLL file.
    4. Install the go-sqlite3 package:
        - Open a command prompt.
        - Run the following command to install the package:
            ```
            goCopy code
            go get github.com/mattn/go-sqlite3
            ```
    5. This command will download and install the go-sqlite3 package and its dependencies.
   
If there are any complications with installing gcc-5, kindly refer to the go-sqlite3 documentation at [https://github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)

## Configuration
Setup a 32-byte Secret Key for encryption in the .env file (For example):
    ```0123456789abcdef0123456789abcdef```

## Usage
1. Start the application:
    ```make``` or  ```go run ./cmd```
2. The application will start at: [http://localhost:8080](http://localhost:8080)

## API Documentation
The Swagger API documentation for the endpoints can be accessed using the following methods:
- Open http://localhost:8080/swagger/index.html in your browser.
- Click on the **Swagger Docs** button in the UI.

## Authentication
Authentication in this project requires providing the Erply client code, username (mail address), and password. You can get the credentials by registering for at [https://login.erply.com/sign-up](https://login.erply.com/sign-up)
Upon successful authentication, the login credentials and session key are encrypted and stored in the SQLite3 database. For subsequent requests to get or add customers, the encrypted session key and client code are retrieved from the database, decrypted, and used to authenticate the request to the Erply API. This ensures secure storage and retrieval of authentication information, facilitating seamless and secure API interactions.

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
