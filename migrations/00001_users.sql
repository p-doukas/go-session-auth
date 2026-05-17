-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  locked_until TIMESTAMPTZ,
  email_verified BOOLEAN NOT NULL DEFAULT FALSE,
  verification_token_hash TEXT,
  verification_token_expires_at TIMESTAMPTZ
);

-- Index for performance on locked accounts
CREATE INDEX idx_users_locked_until ON users(locked_until) 
WHERE locked_until IS NOT NULL;

CREATE INDEX idx_users_verification_token ON users(verification_token_hash)
WHERE verification_token_hash IS NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
