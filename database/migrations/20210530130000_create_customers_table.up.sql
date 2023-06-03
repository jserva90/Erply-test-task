CREATE TABLE IF NOT EXISTS `customers`(
    `client_code` TEXT NOT NULL,
    `customer_id` INTEGER NOT NULL,
    `full_name`      TEXT,
    `email` TEXT,
    `phone` TEXT,
    `created_at` INTEGER
);