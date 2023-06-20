CREATE TABLE games (
    id bigserial PRIMARY KEY,
    public_path uuid DEFAULT uuid_generate_v4 (),
    private_path uuid DEFAULT uuid_generate_v4 (),
    title VARCHAR NOT NULL
);
