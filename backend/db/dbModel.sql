CREATE DATABASE IF NOT EXISTS FUT21;

USE FUT21;

CREATE TABLE Players (
    player_id VARCHAR(40) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    position VARCHAR(100) NOT NULL,
    nationality VARCHAR(100) NOT NULL,
    club VARCHAR(100) NOT NULL,
    PRIMARY KEY (player_id)    
)


