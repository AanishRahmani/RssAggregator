-- +goose Up 
ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT(
encode(sha256(random()::TEXT::bytea),'hex')
);


-- +goose Down
ALTER TABLE users drop COLUMN api_key;