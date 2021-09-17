CREATE TABLE IF NOT EXISTS todos (
  id          INTEGER  NOT NULL PRIMARY KEY AUTO_INCREMENT,
  subject     TEXT NOT NULL,
  description TEXT,
  created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);