-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sessions(
  id SERIAL NOT NULL PRIMARY KEY,
  user_id BIGINT NOT NULL,
  access_token TEXT NOT NULL,
  referesh_token TEXT,
  expires_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now(),

  CONSTRAINT unique_users_email UNIQUE (email),
  CONSTRAINT unique_users_username UNIQUE (username),
  CONSTRAINT unique_users_phone UNIQUE (phone_number)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sessions;
-- +goose StatementEnd
