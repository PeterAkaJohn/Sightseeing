CREATE TABLE locations (
    id          SERIAL PRIMARY KEY,
    name varchar(255),
    position varchar(255),
    address varchar(255),
    city varchar(255),
    state varchar(255),
    postal_code integer,
    description text,
    created_at timestamp with time zone DEFAULT now()
  );

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username varchar(40) NOT NULL UNIQUE,
    firstname varchar(40),
    lastname varchar(40),
    email varchar(40),
    password varchar(255) NOT NULL,
    created_at timestamp with time zone DEFAULT now()
);

CREATE TABLE favorites (
    id SERIAL PRIMARY KEY,
    user_id integer REFERENCES users(id),
    location_id integer REFERENCES locations(id),
    created_at timestamp with time zone DEFAULT now()
);
