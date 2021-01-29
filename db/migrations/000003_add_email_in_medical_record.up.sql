BEGIN;

ALTER TABLE medical_records
ADD COLUMN email VARCHAR(255) NOT NULL;

CREATE INDEX IF NOT EXISTS index_on_email
ON medical_records USING btree (email);

COMMIT;