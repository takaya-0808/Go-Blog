-- DROP SCHEMA IF EXISTS go_blog;

-- CREATE SCHEMA go_blog;

-- USE  go_blog

-- DROP TABLE IF EXISTS users;

-- CREATE TABLE users (
--     id INT NOT NULL AUTO_INCREMENT,
--     username VARCHAR(32) NOT NULL,
--     useremail VARCHAR(32) NOT NULL,
--     userpassword VARCHAR(32) NOT NULL,
--     created_at datetime  default current_timestamp,
--     updated_at timestamp default current_timestamp on update current_timestamp
--     PRIMARY KEY (id)
-- );

-- INSERT INTO users (id,username,useremail,userpassword) VALUES (1, 'TOM','xxxx@mail.co.jp');
-- INSERT INTO users (id,username,useremail,userpassword) VALUES (2, 'Thoge','hogehoge');
-- INSERT INTO users (id,username,useremail,userpassword) VALUES (3, 'fuga','fugafuga');