CREATE DATABASE IF NOT EXISTS app;

USE app;

DROP TABLE if EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(120) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;