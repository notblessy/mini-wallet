-- migrate:up
create table wallets (
    id varchar(255) primary key not null,
    owned_by varchar(255) not null,
        FOREIGN KEY (owned_by) REFERENCES users(customer_xid) ON DELETE CASCADE,
    status int not null,
    balance int,
    enabled_at timestamp,
    disabled_at timestamp
);

-- migrate:down
drop table wallets;

