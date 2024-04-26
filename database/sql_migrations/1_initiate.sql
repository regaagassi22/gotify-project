-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    nama_role VARCHAR(255) NOT NULL
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role_id INT NOT NULL,
    FOREIGN KEY (role_id) REFERENCES roles(id)
);

CREATE TABLE genres (
    id SERIAL PRIMARY KEY,
    nama_genre VARCHAR(255) NOT NULL
);

CREATE TABLE musics (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    artist VARCHAR(255) NOT NULL,
    duration INT NOT NULL,
    genre_id INT NOT NULL,
    FOREIGN KEY (genre_id) REFERENCES genres(id)
);


-- +migrate StatementEnd