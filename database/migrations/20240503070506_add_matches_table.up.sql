CREATE TABLE IF NOT EXISTS matches (
    id SERIAL PRIMARY KEY,
    issued_user_id BIGINT,
    status VARCHAR(50),
    user_cat_id BIGINT,
    match_cat_id BIGINT,
    message VARCHAR(125),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
