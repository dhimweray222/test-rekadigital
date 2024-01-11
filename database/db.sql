CREATE DATABASE IF NOT EXISTS test_be_rekadigital;


CREATE TABLE IF NOT EXISTS transactions (
    id VARCHAR(255) NOT NULL,
    customer_id VARCHAR(255) NOT NULL,
    menu VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    qty INT NOT NULL,
    payment VARCHAR(255) NOT NULL,
    total INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);