CREATE TABLE IF NOT EXISTS auctioncontent
(
    id          serial primary key,
    auctionId uuid references auction (id),
    downloadUrl text
)