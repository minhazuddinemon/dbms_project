-- 1. Create Universities Table
CREATE TABLE universities (
  university_id SERIAL PRIMARY KEY,
  university_name VARCHAR(255) NOT NULL,
  website VARCHAR(255),
  location_address TEXT,
  transport_routes TEXT,
  est_travel_time VARCHAR(100)
);

-- 2. Create Students Table
CREATE TABLE students (
  student_id SERIAL PRIMARY KEY,
  first_name VARCHAR(100) NOT NULL,
  last_name VARCHAR(100) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  date_of_birth DATE NOT NULL,
  ssc_roll INT NOT NULL,
  ssc_reg INT NOT NULL,
  ssc_gpa DECIMAL(3, 2) NOT NULL,
  ssc_board VARCHAR(100) NOT NULL,
  ssc_year INT NOT NULL,
  hsc_roll INT NOT NULL,
  hsc_reg INT NOT NULL,
  hsc_gpa DECIMAL(3, 2) NOT NULL,
  hsc_board VARCHAR(100) NOT NULL,
  hsc_year INT NOT NULL
);

-- 3. Create Student Mobiles Table (Multivalued Attribute)
CREATE TABLE student_mobiles (
  student_id INT,
  mobile VARCHAR(20),
  PRIMARY KEY (student_id, mobile),
  FOREIGN KEY (student_id) REFERENCES students (student_id) ON DELETE CASCADE
);

-- 4. Create Programs Table
CREATE TABLE programs (
  program_id SERIAL PRIMARY KEY,
  university_id INT NOT NULL,
  program_name VARCHAR(255) NOT NULL,
  program_unit VARCHAR(50) NOT NULL,
  total_seats INT NOT NULL,
  min_ssc_gpa DECIMAL(3, 2) NOT NULL,
  min_hsc_gpa DECIMAL(3, 2) NOT NULL,
  required_hsc_group VARCHAR(50) NOT NULL,
  min_math DECIMAL(3, 2),
  min_physics DECIMAL(3, 2),
  min_chemistry DECIMAL(3, 2),
  min_english DECIMAL(3, 2),
  min_bangla DECIMAL(3, 2),
  prev_cutoff_marks DECIMAL(5, 2),
  FOREIGN KEY (university_id) REFERENCES universities (university_id) ON DELETE CASCADE
);

-- 5. Create Applications Table
CREATE TABLE applications (
  application_id SERIAL PRIMARY KEY,
  student_id INT NOT NULL,
  program_id INT NOT NULL,
  submission_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  status VARCHAR(50) DEFAULT 'Applied',
  FOREIGN KEY (student_id) REFERENCES students (student_id) ON DELETE CASCADE,
  FOREIGN KEY (program_id) REFERENCES programs (program_id) ON DELETE CASCADE
);

-- 6. Create Payments Table
CREATE TABLE payments (
  payment_id SERIAL PRIMARY KEY,
  application_id INT NOT NULL UNIQUE,
  transaction_id VARCHAR(255) NOT NULL UNIQUE,
  amount DECIMAL(10, 2) NOT NULL,
  payment_method VARCHAR(50) NOT NULL,
  payment_status VARCHAR(50) DEFAULT 'Pending',
  TIMESTAMP TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (application_id) REFERENCES applications (application_id) ON DELETE CASCADE
);

-- 7. Create Admission Tests Table
CREATE TABLE admission_tests (
  test_id SERIAL PRIMARY KEY,
  university_id INT NOT NULL,
  prereq_test_id INT NULL,
  exam_unit VARCHAR(50) NOT NULL,
  exam_center VARCHAR(255) NOT NULL,
  FOREIGN KEY (university_id) REFERENCES universities (university_id) ON DELETE CASCADE,
  FOREIGN KEY (prereq_test_id) REFERENCES admission_tests (test_id) ON DELETE SET NULL
);

-- 8. Create Student Test Results Table (M:N Associative Table)
CREATE TABLE student_test_results (
  student_id INT,
  test_id INT,
  marks_obtained DECIMAL(5, 2),
  merit_position INT,
  PRIMARY KEY (student_id, test_id),
  FOREIGN KEY (student_id) REFERENCES students (student_id) ON DELETE CASCADE,
  FOREIGN KEY (test_id) REFERENCES admission_tests (test_id) ON DELETE CASCADE
);

-- 9. Create Notifications Table (Weak Entity)
CREATE TABLE notifications (
  student_id INT,
  notification_id INT, -- Removed AUTO_INCREMENT here
  message TEXT NOT NULL,
  deadline_date TIMESTAMP,
  PRIMARY KEY (student_id, notification_id),
  FOREIGN KEY (student_id) REFERENCES students (student_id) ON DELETE CASCADE
);
