
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE account (
    id          VARCHAR(255),
    email       VARCHAR(255),
    password    VARCHAR(255),
    PRIMARY KEY(id)
);

-- test users
insert into account (id, email, password) values('65b1ece8-4ab9-4be5-b433-15494faf4743','cristianchaparroa@gmail.com','12345');
insert into account (id, email, password) values('65b1ece8-4ab9-4be5-b433-15494faf4742','mauriciolopez@gmail.com','12345');
insert into account (id, email, password) values('65b1ece8-4ab9-4be5-b433-15494faf4741','santiagocastro@gmail.com','12345');
insert into account (id, email, password) values('65b1ece8-4ab9-4be5-b433-15494faf4740','merwinponce@gmail.com','12345');


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE account;
