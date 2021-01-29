BEGIN;

ALTER TABLE medical_records
DROP COLUMN email;

COMMIT;