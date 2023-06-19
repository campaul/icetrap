CREATE TABLE games (
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    title VARCHAR NOT NULL
);
