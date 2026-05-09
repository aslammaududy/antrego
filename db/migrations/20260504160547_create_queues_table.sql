-- +goose Up
create table queues(
    id integer not null primary key autoincrement,
    booking_code varchar(10),
    clinic_code char(3),
    number varchar(10),
    created_at timestamp,
    updated_At timestamp,
    deleted_at timestamp
);

-- +goose Down
drop table if exists queues;
