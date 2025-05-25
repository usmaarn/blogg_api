-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
  id SERIAL NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  username VARCHAR(255),
  email VARCHAR(255) NOT NULL,
  phone_number VARCHAR(20) NOT NULL,
  type VARCHAR(10) NOT NULL,
  status VARCHAR(10) NOT NULL,
  password TEXT NOT NULL,
  email_verified_at TIMESTAMP,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now(),

  CONSTRAINT unique_users_email UNIQUE (email),
  CONSTRAINT unique_users_username UNIQUE (username),
  CONSTRAINT unique_users_phone UNIQUE (phone_number)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
