CREATE TABLE persons
(
    id serial PRIMARY KEY,
    email varchar(255) not null,
    phone varchar(255) not null,
    first_name varchar(255) not null unique
)