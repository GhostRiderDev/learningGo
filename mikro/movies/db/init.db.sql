USE main;
CREATE TABLE IF NOT EXISTS Movie(id VARCHAR(255), title VARCHAR(255), description VARCHAR(500), director VARCHAR(255));
CREATE TABLE IF NOT EXISTS Rating(record_id VARCHAR(255), record_type VARCHAR(255), user_id VARCHAR(255), value INTEGER);