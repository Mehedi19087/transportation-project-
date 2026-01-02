ALTER TABLE outside_trips 
ADD COLUMN product_id BIGINT NOT NULL;

ALTER TABLE outside_trips 
ADD COLUMN deleted_at TIMESTAMPTZ;

CREATE INDEX idx_outside_trips_product_id ON outside_trips(product_id);
CREATE INDEX idx_outside_trips_deleted_at ON outside_trips(deleted_at);