DROP TABLE IF EXISTS `users`;

create table IF not exists `users` (
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(32) NOT NULL,
    useremail VARCHAR(32) NOT NULL,
    userpassword VARCHAR(32) NOT NULL,
    created_at datetime  default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY (id)
);

INSERT INTO users (id,username,useremail,userpassword) VALUES (1, 'TOM','xxxx@mail.co.jp', 'hogehoge');
INSERT INTO users (id,username,useremail,userpassword) VALUES (2, 'Thoge','hogehoge', 'hogehoge');
INSERT INTO users (id,username,useremail,userpassword) VALUES (3, 'fuga','fugafuga', 'hogehoge');