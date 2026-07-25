# UniApp Admission Portal - API & Database Documentation

## 1. Overview & Architecture

This document describes the database schema and REST API endpoints for the **UniApp Admission Portal**.

- **Base URL**: `http://localhost:8080` (Direct) or `/api` (Proxied via SvelteKit Frontend)
- **Authentication**: JWT Bearer Tokens in HTTP Header (`Authorization: Bearer <token>`)
- **Data Format**: JSON (`Content-Type: application/json`)

---

## 2. Database Schema

The database consists of 14 core tables in MySQL/MariaDB:

### 2.1 Tables Summary

| Table | Primary Key | Description |
|-------|-------------|-------------|
| `University` | `u_id` | Public universities offering admission |
| `University_Transport` | `(u_id, transport_route)` | Transport routes and travel times to universities |
| `Admission_Test` | `test_id` | Admission exam units, centers, and prerequisites |
| `Conducts` | `(u_id, test_id)` | Mapping of which universities conduct which admission tests |
| `Program` | `program_id` | Degree programs offered by universities |
| `Program_Eligibility_Rules` | `(program_id, rule_type)` | Qualification criteria (GPA cutoffs, HSC subject marks, group) |
| `Program_Required_Fields` | `(program_id, field_name)` | Mandatory profile attributes needed to apply to a program |
| `Student` | `student_id` | Registered student accounts |
| `Student_Mobile` | `(student_id, mobile_no)` | Multivalued student phone numbers |
| `Student_Academics` | `(student_id, exam_level)` | SSC/HSC exam level details, board, year, GPA, education group |
| `Student_Subject_Marks` | `(student_id, exam_level, subject_name)` | Individual HSC subject marks (Physics, Mathematics, Chemistry) |
| `Student_Profile_Info` | `(student_id, field_name)` | Key-value store for student profile attributes and uploaded documents |
| `Application` | `app_id` | Submitted student program applications |
| `Payment` | `payment_id` | Payment transactions tied to applications |

---

## 3. API Endpoints Reference

---

### 3.1 System & Health

#### GET `/health`
Check if backend service and database connection are healthy.

- **Auth Required**: No
- **Request Headers**: None
- **Response 200 OK**:
  ```text
  Backend is healthy and connected to DB!
  ```

---

### 3.2 Student Authentication

#### POST `/register`
Register a new student account.

- **Auth Required**: No
- **Request Body**:
  ```json
  {
    "first_name": "Rahim",
    "last_name": "Uddin",
    "email": "student@example.com",
    "password": "SecretPassword123",
    "dob": "2004-05-15"
  }
  ```
- **Responses**:
  - **201 Created**:
    ```json
    {
      "message": "Registration successful"
    }
    ```
  - **400 Bad Request**: Date format invalid or payload malformed.
  - **499 Conflict**: Email already registered.

#### POST `/login`
Authenticate student and receive JWT token.

- **Auth Required**: No
- **Request Body**:
  ```json
  {
    "email": "student@example.com",
    "password": "SecretPassword123"
  }
  ```
- **Responses**:
  - **200 OK**:
    ```json
    {
      "message": "Login successful",
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
    ```
  - **401 Unauthorized**: Invalid email or password.

---

### 3.3 Student Profile & Academics

#### GET `/profile`
Retrieve protected profile token context.

- **Auth Required**: Yes (Student JWT)
- **Headers**: `Authorization: Bearer <token>`
- **Response 200 OK**:
  ```json
  {
    "message": "Access granted to protected profile!",
    "student_id": 1
  }
  ```

#### POST / PUT `/student/profile`
Upsert dynamic student profile info fields (e.g. addresses, guardian names, blood group, document URLs).

- **Auth Required**: Yes (Student JWT)
- **Headers**: `Authorization: Bearer <token>`
- **Request Body**:
  ```json
  {
    "FATHERS_NAME": "Md. Anowar Hossain",
    "MOTHERS_NAME": "Sultana Begum",
    "BLOOD_GROUP": "O+",
    "QUOTA": "GENERAL",
    "PRESENT_ADDRESS": "Dhaka, Bangladesh",
    "PERMANENT_ADDRESS": "Dhaka, Bangladesh",
    "PHOTO_URL": "https://example.com/photo.jpg",
    "SIGNATURE_URL": "https://example.com/signature.png"
  }
  ```
- **Response 200 OK**:
  ```json
  {
    "status": "success",
    "message": "Profile updated successfully!"
  }
  ```

---

### 3.4 Programs & Eligibility Engine

#### GET `/programs`
List available degree programs with optional search and unit filters.

- **Auth Required**: No
- **Query Parameters**:
  - `search` (optional): Filter by program name or university name (e.g., `search=Computer`)
  - `unit` (optional): Filter by exam unit (e.g., `unit=A`)
- **Response 200 OK**:
  ```json
  [
    {
      "program_id": 1,
      "p_name": "Computer Science & Engineering",
      "p_unit": "A",
      "total_seats": 120,
      "prev_cutmarks": "85.50",
      "deadline": "2026-12-31T00:00:00Z",
      "u_id": 1,
      "university_name": "BUET",
      "website": "https://buet.ac.bd",
      "location": "Dhaka",
      "logo_url": "https://example.com/buet-logo.png"
    }
  ]
  ```

#### GET `/programs/detail`
Retrieve details of a single program by ID.

- **Auth Required**: No
- **Query Parameters**:
  - `id` (required): Program ID (e.g., `id=1`)
- **Response 200 OK**:
  ```json
  {
    "program_id": 1,
    "p_name": "Computer Science & Engineering",
    "p_unit": "A",
    "total_seats": 120,
    "prev_cutmarks": "85.50",
    "deadline": "2026-12-31T00:00:00Z",
    "u_id": 1,
    "university_name": "BUET",
    "website": "https://buet.ac.bd",
    "location": "Dhaka",
    "logo_url": "https://example.com/buet-logo.png"
  }
  ```

#### GET `/programs/eligible`
Run the eligibility engine against logged-in student's SSC/HSC GPAs, education group, and subject marks (Physics, Math, Chemistry).

- **Auth Required**: Yes (Student JWT)
- **Headers**: `Authorization: Bearer <token>`
- **Response 200 OK**:
  ```json
  [
    {
      "program_id": 1,
      "program_name": "Computer Science & Engineering",
      "university_name": "BUET"
    },
    {
      "program_id": 2,
      "program_name": "Unit A (Science)",
      "university_name": "Dhaka University"
    }
  ]
  ```

---

### 3.5 Applications & Submissions

#### GET `/program/requirements`
Check requirements status for a specific program to see if student has provided all mandatory profile fields.

- **Auth Required**: Yes (Student JWT)
- **Headers**: `Authorization: Bearer <token>`
- **Query Parameters**:
  - `program_id` (required): Program ID (e.g., `program_id=1`)
- **Response 200 OK**:
  ```json
  {
    "program_id": 1,
    "is_ready_to_apply": false,
    "required_fields": [
      {
        "field_name": "PRESENT_ADDRESS",
        "value": "Dhaka, Bangladesh",
        "is_provided": true
      },
      {
        "field_name": "PHOTO_URL",
        "value": null,
        "is_provided": false
      }
    ],
    "missing_fields": [
      "PHOTO_URL"
    ]
  }
  ```

#### POST `/applications/apply`
Submit an application for a degree program.

- **Auth Required**: Yes (Student JWT)
- **Headers**: `Authorization: Bearer <token>`
- **Request Body**:
  ```json
  {
    "program_id": 1
  }
  ```
- **Responses**:
  - **201 Created**:
    ```json
    {
      "status": "success",
      "message": "Application submitted successfully! Please proceed to payment.",
      "app_id": 15
    }
    ```
  - **400 Bad Request (Incomplete Profile)**:
    ```json
    {
      "status": "incomplete_profile",
      "message": "You must provide additional information to apply to this program.",
      "missing_fields": [
        "PHOTO_URL",
        "SIGNATURE_URL"
      ]
    }
    ```

#### GET `/applications`
Retrieve all applications submitted by the logged-in student.

- **Auth Required**: Yes (Student JWT)
- **Headers**: `Authorization: Bearer <token>`
- **Response 200 OK**:
  ```json
  [
    {
      "app_id": 15,
      "sub_date": "2026-07-26T01:00:00Z",
      "status": "Pending",
      "program_id": 1,
      "program_name": "Computer Science & Engineering",
      "university_name": "BUET"
    }
  ]
  ```

---

### 3.6 Payment Processing

#### POST `/payments/process`
Process payment transaction for a program application.

- **Auth Required**: Yes (Student JWT)
- **Headers**: `Authorization: Bearer <token>`
- **Request Body**:
  ```json
  {
    "application_id": 15,
    "amount": "500.00",
    "payment_method": "bKash",
    "transaction_id": "TRX987654321"
  }
  ```
- **Responses**:
  - **200 OK**:
    ```json
    {
      "status": "success",
      "message": "Payment received successfully! Application is now complete.",
      "application_id": 15,
      "transaction_id": "TRX987654321"
    }
    ```
  - **400 Bad Request**: Application already paid for.
  - **404 Not Found**: Application not found.

---

### 3.7 Universities & Public Info

#### GET `/universities`
List all registered universities.

- **Auth Required**: No
- **Response 200 OK**:
  ```json
  [
    {
      "u_id": 1,
      "u_name": "BUET",
      "website": "https://buet.ac.bd",
      "location": "Dhaka",
      "logo_url": "https://example.com/buet-logo.png"
    }
  ]
  ```

#### GET `/universities/detail`
Get details for a specific university.

- **Auth Required**: No
- **Query Parameters**:
  - `u_id` (required): University ID (e.g., `u_id=1`)
- **Response 200 OK**:
  ```json
  {
    "u_id": 1,
    "u_name": "BUET",
    "website": "https://buet.ac.bd",
    "location": "Dhaka",
    "logo_url": "https://example.com/buet-logo.png"
  }
  ```

---

### 3.8 Admin Management

#### POST `/admin/login`
Authenticate as system administrator.

- **Auth Required**: No
- **Request Body**:
  ```json
  {
    "email": "admin@system.com",
    "password": "admin123secret"
  }
  ```
- **Response 200 OK**:
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "role": "ADMIN"
  }
  ```

#### GET `/admin/applications`
Fetch all student applications for a university.

- **Auth Required**: Yes (Admin JWT)
- **Headers**: `Authorization: Bearer <token>`
- **Query Parameters**:
  - `u_id` (required): University ID (e.g., `u_id=1`)
- **Response 200 OK**:
  ```json
  [
    {
      "app_id": 15,
      "sub_date": "2026-07-26T01:00:00Z",
      "status": "PAID",
      "program_id": 1,
      "student_id": 5,
      "first_name": "Rahim",
      "last_name": "Uddin",
      "email": "rahim@example.com"
    }
  ]
  ```

#### PUT `/admin/applications/status`
Update status of a student application.

- **Auth Required**: Yes (Admin JWT)
- **Headers**: `Authorization: Bearer <token>`
- **Request Body**:
  ```json
  {
    "app_id": 15,
    "status": "APPROVED"
  }
  ```
- **Response 200 OK**:
  ```json
  {
    "status": "success",
    "message": "Application status updated to APPROVED"
  }
  ```

#### POST `/admin/university`
Create a new university entry.

- **Auth Required**: Yes (Admin JWT)
- **Headers**: `Authorization: Bearer <token>`
- **Request Body**:
  ```json
  {
    "name": "Dhaka University",
    "website": "https://du.ac.bd",
    "location": "Dhaka",
    "logo_url": "https://example.com/du-logo.png"
  }
  ```
- **Response 201 Created**:
  ```json
  {
    "message": "University created successfully",
    "u_id": 2
  }
  ```

#### PUT `/admin/university`
Update an existing university entry.

- **Auth Required**: Yes (Admin JWT)
- **Headers**: `Authorization: Bearer <token>`
- **Query Parameters**: `u_id=2`
- **Request Body**:
  ```json
  {
    "name": "University of Dhaka",
    "website": "https://du.ac.bd",
    "location": "Dhaka, Bangladesh",
    "logo_url": "https://example.com/du-logo.png"
  }
  ```
- **Response 200 OK**:
  ```json
  {
    "message": "University updated successfully"
  }
  ```

#### DELETE `/admin/university`
Delete a university entry.

- **Auth Required**: Yes (Admin JWT)
- **Headers**: `Authorization: Bearer <token>`
- **Query Parameters**: `u_id=2`
- **Response 200 OK**:
  ```json
  {
    "message": "University deleted successfully"
  }
  ```

#### POST `/admin/program`
Create a new degree program under a university.

- **Auth Required**: Yes (Admin JWT)
- **Headers**: `Authorization: Bearer <token>`
- **Request Body**:
  ```json
  {
    "p_name": "Electrical & Electronic Engineering",
    "p_unit": "A",
    "total_seats": 120,
    "prev_cutmarks": 82.50,
    "deadline": "2026-12-31",
    "u_id": 1
  }
  ```
- **Response 201 Created**:
  ```json
  {
    "message": "Program created successfully",
    "program_id": 3
  }
  ```

#### PUT `/admin/program`
Update a degree program.

- **Auth Required**: Yes (Admin JWT)
- **Headers**: `Authorization: Bearer <token>`
- **Query Parameters**: `program_id=3`
- **Request Body**:
  ```json
  {
    "p_name": "Electrical & Electronic Engineering",
    "p_unit": "A",
    "total_seats": 130,
    "prev_cutmarks": 83.00,
    "deadline": "2026-12-31",
    "u_id": 1
  }
  ```
- **Response 200 OK**:
  ```json
  {
    "message": "Program updated successfully"
  }
  ```

#### DELETE `/admin/program`
Delete a degree program.

- **Auth Required**: Yes (Admin JWT)
- **Headers**: `Authorization: Bearer <token>`
- **Query Parameters**: `program_id=3`
- **Response 200 OK**:
  ```json
  {
    "message": "Program deleted successfully"
  }
  ```
