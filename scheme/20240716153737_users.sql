-- +goose Up
-- +goose StatementBegin

CREATE TABLE users
(
 id serial primary key,
 name varchar(255) not null,
 surname varchar(255) not null,
 patronymic varchar(255),
 address text not null,
 pasport_serie INTEGER not null,
 passport_number INTEGER not null
);

CREATE TABLE users_tasks
(
 id serial primary key,
 user_id int references users(id) on delete cascade not null,
 start_at timestamp not null default current_timestamp, 
 end_at timestamp
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE users_tasks;
-- +goose StatementEnd
