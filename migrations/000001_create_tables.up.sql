CREATE TYPE process_status AS ENUM ('Pending','Preparing','Delivered','Cancelled');


CREATE TABLE IF NOT EXISTS process(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    product_id VARCHAR NOT NULL,
    status process_status DEFAULT 'Pending',
    amount integer NOT NULL
);


CREATE TABLE IF NOT EXISTS wishlist(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    product_id VARCHAR NOT NULL
    UNIQUE (user_id, product_id)
);


CREATE TABLE IF NOT EXISTS feedback(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    product_id VARCHAR NOT NULL,
    rating integer NOT NULL,
    description text NOT NULL
    UNIQUE (user_id, product_id)
);

CREATE TABLE IF NOT EXISTS bought(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    product_id VARCHAR NOT NULL,
    amount integer NOT NULL,
    card_number VARCHAR(32) NOT NULL,
    amount_of_money DOUBLE PRECISION NOT NULL,
    process_id UUID REFERENCES process(id),
    status VARCHAR DEFAULT 'bought'
);
