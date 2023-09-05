CREATE TABLE users (
  id SERIAL PRIMARY KEY NOT NULL,
  uuid varchar(255) NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at timestamp DEFAULT NULL,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  phone varchar(30) NOT NULL,
  password varchar(255) DEFAULT NULL,
  email_verified_at timestamp DEFAULT NULL,
  CONSTRAINT users_uuid_unique UNIQUE (uuid),
  CONSTRAINT users_email_unique UNIQUE (email)
);
