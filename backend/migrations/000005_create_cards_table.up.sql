CREATE TABLE cards (
    id bigserial PRIMARY KEY,
    session_id bigint NOT NULL REFERENCES sessions(id)
);

CREATE TABLE card_squares (
    id bigserial PRIMARY KEY,
    square_id bigint NOT NULL REFERENCES squares(id),
    card_id bigint NOT NULL REFERENCES cards(id)
);
