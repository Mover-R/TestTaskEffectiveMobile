CREATE SCHEMA IF NOT EXISTS mig;

CREATE TABLE IF NOT EXISTS mig.users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    patronymic VARCHAR(50),
    age INTEGER NOT NULL,
    gender VARCHAR(8),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS mig.user_country (
    pair_id SERIAL PRIMARY KEY,
    country VARCHAR(20) NOT NULL,
    probability DECIMAL(5,2) NOT NULL,
    user_id BIGINT REFERENCES mig.users(user_id) ON DELETE CASCADE
);