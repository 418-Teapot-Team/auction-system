CREATE TABLE IF NOT EXISTS users
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username    VARCHAR(255) not null unique,
    fullName    varchar(255) not null,
    password    varchar(255) not null,
    createdAt   timestamptz DEFAULT now()
)