-- Rollback: Remove application fee column
ALTER TABLE Program
DROP COLUMN application_fee;
