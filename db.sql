CREATE DATABASE IF NOT EXISTS simple_todo;

CREATE TABLE IF NOT EXISTS `simple_todo`.`todo_items` (
    id BIGINT PRIMARY KEY AUTO_RANDOM,
    title VARCHAR(255),
    completed BOOLEAN,
    user_id BIGINT,
    INDEX idx_user_id(user_id)
);

CREATE TABLE IF NOT EXISTS `simple_todo`.`users` (
    id BIGINT PRIMARY KEY AUTO_RANDOM,
    username VARCHAR(255),
    source VARCHAR(255),
    token VARCHAR(255),
    INDEX idx_token(username,source, token)
);


