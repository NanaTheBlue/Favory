CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
username VARCHAR(255),
email VARCHAR(255),
hashed_password TEXT,
refresh_token TEXT

);



CREATE TABLE IF NOT EXISTS favor(
favor_id UUID PRIMARY KEY,
creator_id UUID REFERENCES users(id),
recipient_id UUID REFERENCES users(id),
favor_text TEXT

);
