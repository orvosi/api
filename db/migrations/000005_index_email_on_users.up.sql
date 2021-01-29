BEGIN;

CREATE INDEX IF NOT EXISTS index_on_email_on_users
ON users USING btree (email);

COMMIT;