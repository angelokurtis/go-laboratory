CREATE TABLE IF NOT EXISTS account
(
    id       BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username text    NOT NULL,
    balance  float   NOT NULL DEFAULT 0,
    version  integer NOT NULL DEFAULT 1
);

INSERT INTO account (username)
SELECT *
FROM (SELECT 'kurtis') AS tmp
WHERE NOT EXISTS (SELECT username
                  FROM account
                  WHERE username = 'kurtis')
LIMIT 1;
