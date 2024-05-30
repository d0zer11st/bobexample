create table example
(
    id bigint primary key generated always as identity,
    mac           macaddr not null
)