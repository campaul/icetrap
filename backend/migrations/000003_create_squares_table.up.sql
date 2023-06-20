CREATE TABLE squares (
    id bigserial PRIMARY KEY,
    game_id bigint NOT NULL,
    title VARCHAR NOT NULL,

    CONSTRAINT FK_game_id FOREIGN KEY (game_id) REFERENCES games(id)
);
