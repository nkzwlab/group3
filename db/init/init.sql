CREATE DATABASE IF NOT EXISTS db;
USE db;
CREATE TABLE IF NOT EXISTS info(
    id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    kg VARCHAR(10) NOT NULL,
    login_name VARCHAR(15) NOT NULL
);

INSERT INTO info(name, kg, login_name) VALUES ("Seiki Makino", "ONE", "kino-ma");
