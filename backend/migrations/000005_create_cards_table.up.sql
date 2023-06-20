CREATE TABLE cards (
    id bigserial PRIMARY KEY,
    public_path uuid DEFAULT uuid_generate_v4 (),
    private_path uuid DEFAULT uuid_generate_v4 (),
    session_id bigint NOT NULL REFERENCES sessions(id)
);

CREATE TABLE card_squares (
    id bigserial PRIMARY KEY,
    square_id bigint NOT NULL REFERENCES squares(id),
    card_id bigint NOT NULL REFERENCES cards(id)
);
