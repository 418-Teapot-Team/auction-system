CREATE TABLE IF NOT EXISTS auction
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    creatorId uuid references users (id),
    title       varchar(255) not null,
    description text         not null,
    startBit    bigint       not null,
    currentBit  bigint       not null,
    createdAt timestamptz DEFAULT now(),
    updatedAt timestamptz
)