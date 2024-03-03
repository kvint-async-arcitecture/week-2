-- +goose Up
-- +goose StatementBegin
alter default privileges in schema app grant select, insert, delete on tables to auth_user_app;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
