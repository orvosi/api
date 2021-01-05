BEGIN;

CREATE TABLE IF NOT EXISTS users (
   id          BIGSERIAL      PRIMARY KEY,
   "name"      TEXT           NOT NULL,
   email       TEXT           NOT NULL,
   google_id   TEXT           UNIQUE NOT NULL,
   created_at  TIMESTAMP,
   updated_at  TIMESTAMP,
   created_by  VARCHAR(200),
   updated_by  VARCHAR(200)
);

COMMIT;