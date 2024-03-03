-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    uid      VARCHAR(36)  PRIMARY KEY DEFAULT public.uuid_generate_v4(),
    email    VARCHAR(100) NOT NULL UNIQUE CHECK ( email <> '' ),
    password VARCHAR(100) NOT NULL CHECK ( password <> '' ),
    role VARCHAR(100) NOT NULL CHECK ( role <> '' )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
