## Humanity


### 1. Stack

**Backend**:

  - gin-gonic (https://github.com/gin-gonic/gin)
  - gorilla websockets:  https://github.com/gorilla/websocket

**Frontend**:
  - Reactjs (without flux)

**database**:
  - postgresql
  - orm: gorm


## 2. Setup development environment

### 2.1. Requirements

The following are the requirements to deploy in a local environment
  - go (at least 1.12)
  - npm
  - postgres database
  - Docker
  - Docker-compose

### 2.2. Install

First of all you should download the back end and front end sources. You should clone the github repository (or use the .zip file sent).

```
git clone https://github.com/cristianchaparroa/humanity
```

**The backend:**
You should go to the backend directory after that remember set the `GO111MODULE` with `on` value to indicate to go environment that we'll use go modules.  Run the `download` command to get the direct dependencies and `tidy` to retrieve the indirect dependencies. If all is ok, you can run the `build` statement to compile the whole application.
```
cd backend
export GO111MODULE=on
go mod download
go mod tidy
go build .
```

**The fronted:**

You should be in the frontend directory and install all the dependencies with the command `npm install` after that it will be able to run.

```
cd fronted
npm install
```

### 2.3.  Database

You should modify the `.env` with` HOST_DB` variable pointing to localhost or the ip where is located your `Postgres` database.


### 2.4.  RabbitMQ

You should modify the `.env` with` RABBITMQ_HOST` variable pointing to localhost or the ip where is located your `RabbitMQ` service.


### 2.5.  Unit test

In favor of software quality we can run the unit tests  available using the following command:

```
cd backend
go test ./...
```

If you want to know the **coverage** in the  project run the command:

```
 go test ./... -coverprofile cp.out
```

You'll have the  percentage covered for unit tests by folder in the stout and you will have `cp.out` file that we'll use to view in the browser what segments of code are covered. For that you should run:

```
go tool cover -html=cp.out
```  

**Note:** The point 2.2 must be executed correctly to run this step.



### 2.6.  Run

**The backend:**
If the step 2.2 was ok, we can up the backend application with the following command:
```
go run main.go
```


**The fronted:**

If all dependencies were install correctly we can up the fronted application with the following command:
```
npm start
```

## 3. Deployment with Docker

If you have docker and docker compose in your system, you should be able to run the following command.

**Note:** you must be sure that `.env` file has setup the `HOST_DB` pointing to `postgres` service.

```
docker-compose up
```

If all is ok, you should be able to enter to the localhost and use the application.


## 4. Application test

You should test the application login  with the available profiles:

| Email | Password |
| ------------- | ------------- |
| cristianchaparroa@gmail.com| 12345 |
| mauriciolopez@gmail.com | 12345  |
| santiagocastro@gmail.com| 12345  |
| merwinponce@gmail.com | 12345  |

### TODO:

- Add RabbitMQ for intents not identified
- Fix the bugs reported in github
- Add more test
