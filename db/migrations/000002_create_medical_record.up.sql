BEGIN;

CREATE TABLE IF NOT EXISTS medical_records (
   id           BIGSERIAL       PRIMARY KEY,
   user_id      BIGINT          NOT NULL,
   symptom      TEXT            NOT NULL,
   diagnosis    TEXT            NOT NULL,
   therapy      TEXT            NOT NULL,
   result       TEXT            NOT NULL,
   created_at   TIMESTAMP,
   updated_at   TIMESTAMP,
   created_by   VARCHAR(200),
   updated_by   VARCHAR(200)
);

COMMIT;