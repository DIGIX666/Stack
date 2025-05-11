-- Table users
CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash CHAR(255) NOT NULL,
    avatar_url VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    role VARCHAR(10) NOT NULL CHECK (role IN ('user', 'admin')),
    session VARCHAR(255),
    is_verified_email BOOLEAN NOT NULL DEFAULT FALSE,
    password_reset_token VARCHAR(255),
    password_reset_expires_at TIMESTAMPTZ,
);

