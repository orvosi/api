BEGIN;

ALTER TABLE medical_records
DROP COLUMN user_id;

COMMIT;