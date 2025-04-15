-- Users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL
);

-- Assets table
CREATE TABLE IF NOT EXISTS assets (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(50) NOT NULL,
    owner_id INT NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES users (id)
);

-- Transactions table
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    asset_id INT NOT NULL,
    from_user_id INT NOT NULL,
    to_user_id INT NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    amount NUMERIC(20,8) NOT NULL,
    FOREIGN KEY (asset_id) REFERENCES assets (id)
);

-- SmartContracts table
CREATE TABLE IF NOT EXISTS smart_contracts (
    id SERIAL PRIMARY KEY,
    address VARCHAR(100) NOT NULL,
    asset_id INT NOT NULL,
    FOREIGN KEY (asset_id) REFERENCES assets (id)
);
