SET SQL_MODE='ALLOW_INVALID_DATES';

DROP TABLE Relation;
DROP TABLE Person;


CREATE TABLE IF NOT EXISTS Person (
    ID_person INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    FirstName VARCHAR (50) NOT NULL,
    LastName VARCHAR (50) NOT NULL,
    NickName VARCHAR (50),
    Gender ENUM('M','F') NOT NULL,
    Rank INT NOT NULL,
    Birthday TIMESTAMP,
    Deathday TIMESTAMP
) AUTO_INCREMENT = 1;

CREATE TABLE IF NOT EXISTS Relation (
    ID_source INT NOT NULL,
    ID_dest INT NOT NULL,
    FOREIGN KEY (ID_source) REFERENCES Person(ID_person),
    FOREIGN KEY (ID_dest) REFERENCES Person(ID_person),
    type ENUM('parental','spousal') NOT NULL,
    PRIMARY KEY(ID_source,ID_dest)    
);