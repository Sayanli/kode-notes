CREATE TABLE  users (
    id serial PRIMARY KEY,
    username varchar(200),
    password varchar(200)
);

CREATE TABLE notes (
    id SERIAL,
    user_id integer,
    text varchar(10000),
    mistakes JSONB,
    FOREIGN KEY (user_id) REFERENCES users(id)
);