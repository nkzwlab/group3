CREATE DATABASE IF NOT EXISTS db;
USE db;


CREATE TABLE IF NOT EXISTS users(
    id INT(11) NOT NULL AUTO_INCREMENT ,
    login_name VARCHAR(30) NOT NULL,
    PRIMARY KEY(id, login_name)
);

INSERT INTO users(login_name) VALUES ("kino-ma");


CREATE TABLE IF NOT EXISTS kadai(
    id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT(11) NOT NULL,
    done BOOLEAN NOT NULL DEFAULT false,
    title VARCHAR(40) NOT NULL,
    content VARCHAR(200) NOT NULL,
    draft VARCHAR(1000) NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id)
);

INSERT INTO kadai(user_id, title, content, draft) VALUES (1, "jn_lecture", "グループでWebサービスを作る", "課題管理サービスを開発中");
