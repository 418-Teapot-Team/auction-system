CREATE TABLE IF NOT EXISTS bids
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    bidderId uuid references users (id),
    auctionId uuid references auction (id),
    newValue  bigint       not null,
    createdAt timestamptz DEFAULT now(),
)