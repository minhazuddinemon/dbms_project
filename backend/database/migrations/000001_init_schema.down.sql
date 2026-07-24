-- 000001_init_schema.down.sql or 000002_init_schema.down.sql

-- Drop child tables dependent on Application and Student
DROP TABLE IF EXISTS Payment;
DROP TABLE IF EXISTS Application;

-- Drop child tables dependent on Student and Admission_Test
DROP TABLE IF EXISTS Gives;
DROP TABLE IF EXISTS Notification;
DROP TABLE IF EXISTS Student_Academics;
DROP TABLE IF EXISTS Student_Mobile;
DROP TABLE IF EXISTS Student;

-- Drop child tables dependent on Program and University
DROP TABLE IF EXISTS Program_Eligibility_Rules;
DROP TABLE IF EXISTS Program;

-- Drop child tables dependent on University and Admission_Test
DROP TABLE IF EXISTS Conducts;
DROP TABLE IF EXISTS Admission_Test;
DROP TABLE IF EXISTS University_Transport;

-- Drop root tables
DROP TABLE IF EXISTS University;
