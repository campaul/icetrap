CREATE TABLE sessions (
    id bigserial PRIMARY KEY,
    game_id bigint NOT NULL REFERENCES games(id)
);

CREATE TABLE selections (
    id bigserial PRIMARY KEY,
    square_id bigint NOT NULL REFERENCES squares(id),
    session_id bigint NOT NULL REFERENCES sessions(id)
);

ALTER TABLE games ADD COLUMN current_session bigint REFERENCES sessions(id);
