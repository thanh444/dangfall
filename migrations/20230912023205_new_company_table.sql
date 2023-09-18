CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "company" (
    id uuid not null default uuid_generate_v4(),
    name varchar(255) not null,
    created_at timestamp without time zone default now(),
    updated_at timestamp without time zone default now(),
    primary key (id)
)