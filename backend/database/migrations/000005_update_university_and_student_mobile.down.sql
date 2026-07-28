-- 1. Drop the new tables first to prevent foreign key constraint errors
DROP TABLE IF EXISTS University_Department;
DROP TABLE IF EXISTS University_Album;

-- 2. Remove the added columns from University
ALTER TABLE University
DROP COLUMN university_description,
DROP COLUMN university_history;

-- 3. Remove the added column from Student_Mobile
ALTER TABLE Student_Mobile
DROP COLUMN owner_type;
