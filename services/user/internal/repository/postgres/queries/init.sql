--name: InitDB :exec
CREATE DATABASE IF NOT EXISTS user_service;


--name: InitUserTable :exec
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password TEXT NOT NULL,
    created_at INT NOT NULL,
    status INT NOT NULL DEFAULT 0
);
