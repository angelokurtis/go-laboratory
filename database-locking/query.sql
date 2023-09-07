-- name: GetAccountAndLockForUpdates :one
SELECT *
FROM account
WHERE username = ?
LIMIT 1 FOR UPDATE;

-- name: UpdateAccountBalance :execresult
UPDATE account
SET balance = ?,
    version = version + 1
WHERE id = ?;
