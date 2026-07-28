-- 1. Alter Student_Mobile Table
ALTER TABLE Student_Mobile
ADD COLUMN owner_type ENUM('self', 'mother', 'father') NOT NULL DEFAULT 'self';

-- 2. Alter University Table
ALTER TABLE University
ADD COLUMN university_description TEXT,
ADD COLUMN university_history TEXT;

-- 3. Create University_Album Table
CREATE TABLE University_Album (
    album_id INT AUTO_INCREMENT PRIMARY KEY,
    u_id INT NOT NULL,
    picture_title VARCHAR(100) NOT NULL,
    picture_url VARCHAR(255) NOT NULL,
    FOREIGN KEY (u_id) REFERENCES University (u_id) ON DELETE CASCADE
);

-- 4. Create University_Department Table
CREATE TABLE University_Department (
    dept_id INT AUTO_INCREMENT PRIMARY KEY,
    u_id INT NOT NULL,
    dept_name VARCHAR(100) NOT NULL,
    dept_description TEXT,
    total_seats INT NOT NULL,
    FOREIGN KEY (u_id) REFERENCES University (u_id) ON DELETE CASCADE
);
