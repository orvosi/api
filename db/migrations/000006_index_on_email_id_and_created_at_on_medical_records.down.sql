BEGIN;

DROP INDEX IF EXISTS index_on_email_id_created_at_on_medical_records;

CREATE INDEX IF NOT EXISTS index_on_email_on_medical_records
ON medical_records USING btree (email);

COMMIT;