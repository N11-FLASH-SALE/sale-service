CREATE TYPE process_status AS ENUM ('Pending','Preparing','Delivered','Cancelled');


CREATE TABLE IF NOT EXISTS process(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    product_id VARCHAR DEFAULT 'Pending',
    status process_status NOT NULL,
    amount integer NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS wishlist(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    product_id VARCHAR NOT NULL
);


CREATE TABLE IF NOT EXISTS feedback(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    product_id VARCHAR NOT NULL,
    rating integer NOT NULL,
    description text NOT NULL
);

CREATE TABLE IF NOT EXISTS bought(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    product_id VARCHAR NOT NULL,
    amount integer NOT NULL,
    card_number VARCHAR(16) NOT NULL,
    amount_of_money DOUBLE PRECISION NOT NULL
);
