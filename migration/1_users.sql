CREATE TABLE users
(
  id         VARCHAR(45) PRIMARY KEY ,
  email      VARCHAR(255) NOT NULL UNIQUE,
  first_name VARCHAR(200),
  last_name VARCHAR(200),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);