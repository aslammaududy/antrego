-- +goose Up
create table queues(
    id integer not null primary key autoincrement,
    booking_code varchar(10),
    clinic_code char(3),
    number varchar(10),
    estimated_time timestamp,
    status varchar(10), --(booked, called, done, cancelled)
    created_at timestamp,
    updated_At timestamp,
    deleted_at timestamp
);

-- +goose Down
drop table if exists queues;
