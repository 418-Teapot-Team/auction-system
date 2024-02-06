CREATE TABLE IF NOT EXISTS auctioncontent
(
    id          serial primary key,
    auctionId   uuid references auction (id),
    title       varchar(255),
    downloadUrl varchar(255) unique
)