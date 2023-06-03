CREATE TABLE IF NOT EXISTS `sessions`(
    `session_id`      INTEGER PRIMARY KEY AUTOINCREMENT,
    `client_code` TEXT NOT NULL,
    `username` TEXT NOT NULL,
    `password` TEXT NOT NULL,
    `session_token` TEXT NOT NULL,
    `session_key` TEXT NOT NULL
);