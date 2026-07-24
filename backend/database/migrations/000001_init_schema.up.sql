-- 1. Create Universities Table
CREATE TABLE University (
    u_id INT AUTO_INCREMENT PRIMARY KEY,
    u_name VARCHAR(50) NOT NULL,
    website VARCHAR(100) NOT NULL,
    location VARCHAR(50)
);

-- 2. Create Universities_Transport Table
CREATE TABLE University_Transport (
    u_id INT NOT NULL,
    transport_route VARCHAR(100) NOT NULL,
    est_travel_time VARCHAR(15) NOT NULL,

    PRIMARY KEY (u_id, transport_route),
    FOREIGN KEY (u_id) REFERENCES University (u_id) ON DELETE CASCADE
);

-- 3. Create Admission_Test Table
CREATE TABLE Admission_Test (
    test_id INT AUTO_INCREMENT PRIMARY KEY,
    exam_unit VARCHAR(5),
    exam_center VARCHAR(50),
    prereq_test_id INT,
    FOREIGN KEY (prereq_test_id) REFERENCES Admission_Test (
        test_id
    ) ON DELETE SET NULL
);

-- 4. Create Conducts Table
CREATE TABLE Conducts (
    u_id INT NOT NULL,
    test_id INT NOT NULL,
    PRIMARY KEY (u_id, test_id),
    FOREIGN KEY (u_id) REFERENCES University (u_id),
    FOREIGN KEY (test_id) REFERENCES Admission_Test (test_id) ON DELETE CASCADE
);

-- 5. Create Program Table
CREATE TABLE Program (
    program_id INT AUTO_INCREMENT PRIMARY KEY,
    p_name VARCHAR(30) NOT NULL,
    p_unit VARCHAR(5),
    total_seats INT NOT NULL,
    prev_cutmarks NUMERIC(6, 2),
    deadline DATE NOT NULL,
    u_id INT NOT NULL,
    FOREIGN KEY (u_id) REFERENCES University (u_id) ON DELETE CASCADE
);

-- 6. Create Program_Eligibility_Rules Table
CREATE TABLE Program_Eligibility_Rules (
    program_id INT NOT NULL,
    rule_type VARCHAR(15) NOT NULL,
    rule_value VARCHAR(10),
    PRIMARY KEY (program_id, rule_type),
    FOREIGN KEY (program_id) REFERENCES Program (program_id) ON DELETE CASCADE
);

-- 7. Create Student Table
CREATE TABLE Student (
    student_id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    dob DATE NOT NULL
);

-- 8. Create Student_Mobile Table (Multivalued attribute)
CREATE TABLE Student_Mobile (
    student_id INT NOT NULL,
    mobile_no VARCHAR(20) NOT NULL,
    PRIMARY KEY (student_id, mobile_no),
    FOREIGN KEY (student_id) REFERENCES Student (student_id) ON DELETE CASCADE
);

-- 9. Create Student_Academics Table
CREATE TABLE Student_Academics (
    student_id INT NOT NULL,
    exam_level VARCHAR(20) NOT NULL, -- e.g., 'SSC', 'HSC'
    year INT NOT NULL,
    roll_no VARCHAR(20) NOT NULL,
    reg_no VARCHAR(20) NOT NULL,
    gpa NUMERIC(3, 2) NOT NULL,
    board VARCHAR(50) NOT NULL,
    PRIMARY KEY (student_id, exam_level),
    FOREIGN KEY (student_id) REFERENCES Student (student_id) ON DELETE CASCADE
);

-- 10. Create Notification Table (Deadline removed as requested)
CREATE TABLE Notification (
    notif_id INT AUTO_INCREMENT PRIMARY KEY,
    student_id INT NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES Student (student_id) ON DELETE CASCADE
);

-- 11. Create Gives Table (Student taking Admission Test)
CREATE TABLE Gives (
    student_id INT NOT NULL,
    test_id INT NOT NULL,
    marks NUMERIC(6, 2),
    merit_position INT,
    PRIMARY KEY (student_id, test_id),
    FOREIGN KEY (student_id) REFERENCES Student (student_id) ON DELETE CASCADE,
    FOREIGN KEY (test_id) REFERENCES Admission_Test (test_id) ON DELETE CASCADE
);

-- 12. Create Application Table
CREATE TABLE Application (
    app_id INT AUTO_INCREMENT PRIMARY KEY,
    sub_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(20) NOT NULL DEFAULT 'Pending',
    program_id INT NOT NULL,
    student_id INT NOT NULL,
    FOREIGN KEY (program_id) REFERENCES Program (program_id),
    FOREIGN KEY (student_id) REFERENCES Student (student_id)
);

-- 13. Create Payment Table
CREATE TABLE Payment (
    payment_id INT AUTO_INCREMENT PRIMARY KEY,
    tx_id VARCHAR(100) UNIQUE NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    status VARCHAR(20) NOT NULL,
    method VARCHAR(30) NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    app_id INT NOT NULL,
    FOREIGN KEY (app_id) REFERENCES Application (app_id) ON DELETE CASCADE
);
