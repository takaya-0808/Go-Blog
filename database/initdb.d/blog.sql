DROP TABLE IF EXISTS `blogs`;

create table IF not exists `blogs` (
    id INT NOT NULL AUTO_INCREMENT,
    author VARCHAR(32) NOT NULL,
    title VARCHAR(32) NOT NULL,
    context VARCHAR(1024) NOT NULL,
    created_at datetime  default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY (id)
);

INSERT INTO blogs (id,author,title,context) VALUES (1, 'TOM','xxxx@mail.co.jp', 'hogehoge');

DROP TABLE IF EXISTS `titles`;

create table IF not exists `titles` (
    id INT,
    title VARCHAR(32) NOT NULL,
    urls VARCHAR(64) NOT NULL
);


INSERT INTO titles (id, title, urls) VALUES (1, 'xxxx@mail.co.jp', 'http://hogehoge');

