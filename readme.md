## Himanity


## Deployment (local)

### Requirements

The following are the requirements to deploy in a local enviroment
  - go (at least 1.12)
  - npm
  - postgres database
  - goose
  - Docker
  - Docker-compose

### Install

```
git clone https://github.com/cristianchaparroa/humanity
```

The backend
```
cd backend
export GO111MODULE=on
go get -d -v ./...
```

The fronted

```
cd fronted
npm install
```

Database

To be able to run the application with  users to login, you should run the script located in `backend/db/migrations` to do that you have two options. The fisrt is run the migrations `*.sql` the part inside that is just in up section. The second way is use the tool goose, then in the root of the backend directory run the command `goose up`. For this method check the credentials for your postgres connection in the file located in `backend/db/dbconfig.yml`

#### Run

The backend
```
go run main.go
```


The fronted
```
npm start
```

## Deployment (local with Docker)

If you have docker and docker compose in your system, you should be able to run the following command.
```
docker-compose up
```

If all is ok, you should be able to enter to the localhost and use the application. 
