-- migrate:up
create table users (
    customer_xid varchar(255) primary key not null,
    created_at timestamp
);

-- migrate:down
drop table users;

