-- migrate:up
create table users (
    id varchar(255) primary key not null,
    token text,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

-- migrate:down
drop table users;

