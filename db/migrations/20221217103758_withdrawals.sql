-- migrate:up
create table deposits (
    id varchar(255) primary key not null,
    withdrawn_by varchar(255) not null,
        FOREIGN KEY (withdrawn_by) REFERENCES users(id) ON DELETE CASCADE,
    reference_id varchar(255),
    status int not null,
    amount int,
    withdrawn_at timestamp,
);

-- migrate:down
drop table deposits;