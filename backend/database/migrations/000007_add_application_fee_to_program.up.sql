-- Add application fee column to Program table
ALTER TABLE Program
ADD COLUMN application_fee NUMERIC(10, 2) NOT NULL DEFAULT 0.00;
