BEGIN;

ALTER TABLE medical_records
DROP COLUMN email;

DROP INDEX IF EXISTS index_on_email_on_medical_records;

COMMIT;