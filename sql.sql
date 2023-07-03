CREATE TABLE IF NOT EXISTS conversions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    amount FLOAT,
    from_currency VARCHAR(3),
    to_currency VARCHAR(3),
    exchange_rate FLOAT,
    converted_value FLOAT,
    currency_symbol VARCHAR(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
);