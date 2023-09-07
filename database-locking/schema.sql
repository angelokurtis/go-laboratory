CREATE TABLE IF NOT EXISTS account
(
    id       BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username text    NOT NULL,
    balance  float   NOT NULL DEFAULT 0,
    version  integer NOT NULL DEFAULT 1
);
