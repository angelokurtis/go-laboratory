-- name: GetAccountAndLockForUpdates :one
SELECT *
FROM account
WHERE username = ? LIMIT 1
FOR UPDATE;

-- name: GetAccount :one
SELECT *
FROM account
WHERE username = ? LIMIT 1;

-- name: UpdateAccountBalance :execresult
UPDATE account
SET balance = ?,
    version = version + 1
WHERE id = ?;

-- name: UpdateAccountBalanceVersion :execresult
UPDATE account
SET balance = ?,
    version = version + 1
WHERE id = ?
  AND version = ?;
