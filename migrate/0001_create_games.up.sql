CREATE TABLE games (
id uuid PRIMARY KEY,
game_type varchar(32) NOT NULL,
finished boolean NOT NULL,
game_state jsonb NOT NULL
);
