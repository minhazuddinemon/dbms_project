-- 1. Table to define what a specific program REQUIRES
CREATE TABLE Program_Required_Fields (
    program_id INT NOT NULL,
    field_name ENUM(
        'PRESENT_ADDRESS',
        'PERMANENT_ADDRESS',
        'FATHERS_NAME',
        'MOTHERS_NAME',
        'BLOOD_GROUP',
        'QUOTA',
        'PHOTO_URL',
        'SIGNATURE_URL'
    ) NOT NULL,
    PRIMARY KEY (program_id, field_name),
    FOREIGN KEY (program_id) REFERENCES Program (program_id) ON DELETE CASCADE
);

-- 2. Table to store what the student HAS PROVIDED
CREATE TABLE Student_Profile_Info (
    student_id INT NOT NULL,
    field_name ENUM(
        'PRESENT_ADDRESS',
        'PERMANENT_ADDRESS',
        'FATHERS_NAME',
        'MOTHERS_NAME',
        'BLOOD_GROUP',
        'QUOTA',
        'PHOTO_URL',
        'SIGNATURE_URL'
    ) NOT NULL,
    field_value VARCHAR(255) NOT NULL,
    PRIMARY KEY (student_id, field_name),
    FOREIGN KEY (student_id) REFERENCES Student (student_id) ON DELETE CASCADE
);
