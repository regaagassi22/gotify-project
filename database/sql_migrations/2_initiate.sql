-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE ratings (
    id SERIAL PRIMARY KEY,
    rating INT NOT NULL,
    user_id INT REFERENCES users(id),
    music_id INT REFERENCES musics(id)
);

-- +migrate StatementEnd