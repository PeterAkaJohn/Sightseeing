CREATE TABLE location (
    id          SERIAL PRIMARY KEY,
    name text,
    position text,
    address text,
    city text,
    state text,
    postal_code text,
    description text,
    open_hours text,
    close_hours text
  );

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username text NOT NULL,
    firstname text,
    lastname text,
    email text,
    password text NOT NULL
);

CREATE TABLE favorite (
    id SERIAL PRIMARY KEY,
    user_id integer REFERENCES users(id),
    location_id integer REFERENCES location(id)
);
