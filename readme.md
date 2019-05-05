## Humanity


### Stack

**Backend**:

  - gin-gonic (https://github.com/gin-gonic/gin)
  - gorilla websockets:  https://github.com/gorilla/websocket

**Frontend**:
  - Reactjs (without flux)

**database**:
  - postgresql


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

**Database**

To be able to run the application with  users toÂ login, you should run the script located in `backend/db/migrations` to do that you have two options. The fisrt is run the migrations `*.sql` the part inside that is just in up section. The second way is use the tool goose, then in the root of the backend directory run the command `goose up`. For this method check the credentials for your postgres connection in the file located in `backend/db/dbconfig.yml`

Note: you must be sure that you are using in the `.env` file with the same config that is in `.env.dev.local`.

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

Note: you must be sure that `.env` file is using the same configuration  like in `.env.dev.docker`
```
docker-compose up
```

If all is ok, you should be able to enter to the localhost and use the application.


### Database population


#### Option 1.

You must to make a connection to
```
host:localhost
database:humanity
user:humanity
password: humanity
```

After that you have the connection you should run the following script
```sql
CREATE TABLE account (
    id          VARCHAR(255),
    email       VARCHAR(255),
    password    VARCHAR(255),
    nickname    VARCHAR(255),
    PRIMARY KEY(id)
);

-- test users
insert into account (id, email, password, nickname) values('65b1ece8-4ab9-4be5-b433-15494faf4743','cristianchaparroa@gmail.com','12345', 'ccchaparroa');
insert into account (id, email, password, nickname) values('65b1ece8-4ab9-4be5-b433-15494faf4742','mauriciolopez@gmail.com','12345', 'mlopez');
insert into account (id, email, password, nickname) values('65b1ece8-4ab9-4be5-b433-15494faf4741','santiagocastro@gmail.com','12345', 'scastro');
insert into account (id, email, password, nickname) values('65b1ece8-4ab9-4be5-b433-15494faf4740','merwinponce@gmail.com','12345', 'mponce');

commit;
```

#### Option 2.

If you have installed `goose` in the root of backend directory you can run the the command:

```
goose up
```


## TODO:

- Automate the database population
- Add RabbitMQ for intents not identified
- Fix the bugs reported in github
- Add more test
