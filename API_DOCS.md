# UniApp Admission Portal - Backend API Reference

This document maps the **current backend implementation** to a frontend-friendly API reference.

- **Base URL**: `http://localhost:8080`
- **Optional frontend proxy prefix**: `/api` (if your frontend proxy is configured that way)
- **Auth**: JWT Bearer token in `Authorization` header
- **Content-Type**: `application/json`

---

## Important frontend notes

### 1. Error responses are usually plain text

Most backend errors currently use Go's `http.Error(...)`, so non-2xx responses are usually **plain text**, for example:

```text
Invalid request payload
```

Do **not** assume all errors are JSON.

### 2. Some endpoints expose `sql.NullString` / `sql.NullTime` directly

A few handlers return generated DB structs directly. That means some nullable fields are serialized as objects like:

#### Nullable string

```json
{
  "String": "Dhaka",
  "Valid": true
}
```

#### Empty nullable string

```json
{
  "String": "",
  "Valid": false
}
```

#### Nullable time

```json
{
  "Time": "2026-07-28T00:00:00Z",
  "Valid": true
}
```

This affects some responses such as:

- `GET /programs`
- `GET /programs/detail`
- `GET /applications`
- `GET /admin/applications`
- `GET /universities`

`GET /universities/detail` is cleaner and returns plain strings for nullable university fields.

### 3. Protected routes

Use:

```http
Authorization: Bearer <token>
```

---

# 1. Public endpoints

## 1.1 Health

### GET `/health`

Checks backend and DB connectivity.

- **Auth required**: No

#### Response `200 OK`

```text
Backend is healthy and connected to DB!
```

---

## 1.2 Student authentication

### POST `/register`

Registers a new student account.

- **Auth required**: No

#### Request body

```json
{
  "first_name": "Rahim",
  "last_name": "Uddin",
  "email": "rahim@example.com",
  "password": "SecretPassword123",
  "dob": "2004-05-15"
}
```

#### Response `201 Created`

```json
{
  "message": "Registration successful"
}
```

#### Possible errors

- `400 Bad Request` — invalid JSON or invalid `dob` format
- `409 Conflict` — email already exists

---

### POST `/login`

Authenticates a student and returns a JWT token.

- **Auth required**: No

#### Request body

```json
{
  "email": "rahim@example.com",
  "password": "SecretPassword123"
}
```

#### Response `200 OK`

```json
{
  "message": "Login successful",
  "token": "<student-jwt-token>"
}
```

#### Possible errors

- `400 Bad Request`
- `401 Unauthorized`

---

## 1.3 Programs

### GET `/programs`

Lists programs with optional text search and unit filter.

- **Auth required**: No

#### Query parameters

- `search` _(optional)_ — matches program name or university name
- `unit` _(optional)_ — exact program unit filter, e.g. `A`

#### Example request

```http
GET /programs?search=Computer&unit=A
```

#### Response `200 OK`

> Note: nullable fields are returned as `sql.NullString` objects.

```json
[
  {
    "program_id": 3,
    "p_name": "Computer Science and Engineering",
    "p_unit": {
      "String": "A",
      "Valid": true
    },
    "total_seats": 120,
    "prev_cutmarks": {
      "String": "85.50",
      "Valid": true
    },
    "deadline": "2026-12-31T00:00:00Z",
    "u_id": 1,
    "university_name": "BUET",
    "university_location": {
      "String": "Dhaka",
      "Valid": true
    }
  }
]
```

#### Possible errors

- `500 Internal Server Error`

---

### GET `/programs/detail`

Returns a single program by ID.

- **Auth required**: No

#### Query parameters

- `id` _(required)_ — program ID

#### Example request

```http
GET /programs/detail?id=3
```

#### Response `200 OK`

> `p_unit`, `prev_cutmarks`, and `location` are nullable-wrapper objects.

```json
{
  "program_id": 3,
  "p_name": "Computer Science and Engineering",
  "p_unit": {
    "String": "A",
    "Valid": true
  },
  "total_seats": 120,
  "prev_cutmarks": {
    "String": "85.50",
    "Valid": true
  },
  "deadline": "2026-12-31T00:00:00Z",
  "u_id": 1,
  "university_name": "BUET",
  "website": "https://buet.ac.bd",
  "location": {
    "String": "Dhaka",
    "Valid": true
  }
}
```

#### Possible errors

- `400 Bad Request` — missing/invalid `id`
- `404 Not Found`
- `500 Internal Server Error`

---

## 1.4 Universities

### GET `/universities`

Returns all universities.

- **Auth required**: No

#### Response `200 OK`

> This endpoint returns the raw DB model, so nullable fields are wrappers.

```json
[
  {
    "u_id": 1,
    "u_name": "BUET",
    "website": "https://buet.ac.bd",
    "location": {
      "String": "Dhaka",
      "Valid": true
    },
    "logo_url": {
      "String": "https://example.com/buet.png",
      "Valid": true
    },
    "university_description": {
      "String": "A leading engineering university.",
      "Valid": true
    },
    "university_history": {
      "String": "Founded in ...",
      "Valid": true
    }
  }
]
```

#### Possible errors

- `500 Internal Server Error`

---

### GET `/universities/detail`

Returns a university with departments and album.

- **Auth required**: No

#### Query parameters

- `u_id` _(required)_ — university ID

#### Example request

```http
GET /universities/detail?u_id=1
```

#### Response `200 OK`

```json
{
  "u_id": 1,
  "u_name": "BUET",
  "website": "https://buet.ac.bd",
  "location": "Dhaka",
  "logo_url": "https://example.com/buet.png",
  "university_description": "A leading engineering university.",
  "university_history": "Founded in ...",
  "departments": [
    {
      "dept_id": 1,
      "dept_name": "CSE",
      "dept_description": "Computer Science and Engineering",
      "total_seats": 120
    }
  ],
  "album": [
    {
      "album_id": 1,
      "picture_title": "Campus View",
      "picture_url": "https://example.com/campus.jpg"
    }
  ]
}
```

#### Possible errors

- `400 Bad Request` — missing/invalid `u_id`
- `404 Not Found`
- `500 Internal Server Error`

---

# 2. Student-protected endpoints

## 2.1 Basic profile token check

### GET `/profile`

Currently this only returns JWT-authenticated student context.

- **Auth required**: Yes (student token)

#### Response `200 OK`

```json
{
  "message": "Access granted to protected profile!",
  "student_id": 1
}
```

---

## 2.2 Student profile fields

### POST `/student/profile`

### PUT `/student/profile`

Upserts dynamic student profile fields.

- **Auth required**: Yes (student token)

#### Request body

```json
{
  "FATHERS_NAME": "Md. Anowar Hossain",
  "MOTHERS_NAME": "Sultana Begum",
  "BLOOD_GROUP": "O+",
  "QUOTA": "GENERAL",
  "PRESENT_ADDRESS": "Dhaka, Bangladesh",
  "PERMANENT_ADDRESS": "Tangail, Bangladesh",
  "PHOTO_URL": "https://example.com/photo.jpg",
  "SIGNATURE_URL": "https://example.com/signature.png"
}
```

#### Response `200 OK`

```json
{
  "status": "success",
  "message": "Profile updated successfully!"
}
```

#### Supported field names currently used by backend logic

- `PRESENT_ADDRESS`
- `PERMANENT_ADDRESS`
- `FATHERS_NAME`
- `MOTHERS_NAME`
- `BLOOD_GROUP`
- `QUOTA`
- `PHOTO_URL`
- `SIGNATURE_URL`

#### Possible errors

- `400 Bad Request`
- `401 Unauthorized`

---

## 2.3 Student mobiles

### GET `/student/mobile`

Returns all saved mobile numbers for the logged-in student.

- **Auth required**: Yes (student token)

#### Response `200 OK`

```json
[
  {
    "student_id": 1,
    "mobile_no": "01711111111",
    "owner_type": "self"
  },
  {
    "student_id": 1,
    "mobile_no": "01822222222",
    "owner_type": "mother"
  }
]
```

---

### POST `/student/mobile`

Adds a new mobile number.

- **Auth required**: Yes (student token)

#### Request body

```json
{
  "mobile_no": "01711111111",
  "owner_type": "self"
}
```

#### Allowed `owner_type`

- `self`
- `mother`
- `father`

#### Response `201 Created`

```json
{
  "message": "Student mobile added successfully",
  "mobile": {
    "student_id": 1,
    "mobile_no": "01711111111",
    "owner_type": "self"
  }
}
```

#### Possible errors

- `400 Bad Request`
- `401 Unauthorized`
- `409 Conflict` — duplicate mobile number for the same student

---

### PUT `/student/mobile`

Updates an existing mobile row by current mobile number.

- **Auth required**: Yes (student token)

#### Request body

```json
{
  "current_mobile_no": "01711111111",
  "mobile_no": "01733333333",
  "owner_type": "father"
}
```

#### Response `200 OK`

```json
{
  "message": "Student mobile updated successfully",
  "mobile": {
    "student_id": 1,
    "mobile_no": "01733333333",
    "owner_type": "father"
  }
}
```

#### Possible errors

- `400 Bad Request`
- `401 Unauthorized`
- `404 Not Found`
- `409 Conflict`

---

### DELETE `/student/mobile`

Deletes a mobile row by query parameter.

- **Auth required**: Yes (student token)

#### Query parameters

- `mobile_no` _(required)_

#### Example request

```http
DELETE /student/mobile?mobile_no=01733333333
```

#### Response `200 OK`

```json
{
  "message": "Student mobile deleted successfully"
}
```

#### Possible errors

- `400 Bad Request`
- `401 Unauthorized`
- `404 Not Found`

---

## 2.4 Eligibility and requirements

### GET `/programs/eligible`

Runs the eligibility engine for the logged-in student.

- **Auth required**: Yes (student token)

#### Response `200 OK`

```json
[
  {
    "program_id": 3,
    "program_name": "Computer Science and Engineering",
    "university_name": "BUET"
  },
  {
    "program_id": 7,
    "program_name": "Unit A",
    "university_name": "University of Dhaka"
  }
]
```

#### Possible errors

- `401 Unauthorized`
- `500 Internal Server Error`

---

### GET `/program/requirements`

Checks whether the student has supplied all required profile fields for a program.

- **Auth required**: Yes (student token)

#### Query parameters

- `program_id` _(required)_

#### Example request

```http
GET /program/requirements?program_id=3
```

#### Response `200 OK`

```json
{
  "program_id": 3,
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
  "missing_fields": ["PHOTO_URL"]
}
```

#### Possible errors

- `400 Bad Request`
- `401 Unauthorized`
- `500 Internal Server Error`

---

## 2.5 Applications

### POST `/applications/apply`

Creates an application for a program if required profile fields are complete.

- **Auth required**: Yes (student token)

#### Request body

```json
{
  "program_id": 3
}
```

#### Success response `201 Created`

```json
{
  "status": "success",
  "message": "Application submitted successfully! Please proceed to payment.",
  "app_id": 14
}
```

#### Incomplete profile response `400 Bad Request`

```json
{
  "status": "incomplete_profile",
  "message": "You must provide additional information to apply to this program.",
  "missing_fields": ["PHOTO_URL", "SIGNATURE_URL"]
}
```

#### Other possible errors

- `401 Unauthorized`
- `500 Internal Server Error`

---

### GET `/applications`

Returns all applications of the logged-in student.

- **Auth required**: Yes (student token)

#### Response `200 OK`

> `sub_date` is returned as a nullable-wrapper object.

```json
[
  {
    "app_id": 14,
    "program_id": 3,
    "sub_date": {
      "Time": "2026-07-28T10:30:00Z",
      "Valid": true
    },
    "status": "PAID",
    "program_name": "Computer Science and Engineering",
    "university_name": "BUET"
  }
]
```

#### Possible errors

- `401 Unauthorized`
- `500 Internal Server Error`

---

## 2.6 Payments

### POST `/payments/process`

Records a payment for an existing application owned by the logged-in student and marks the application as `PAID`.

- **Auth required**: Yes (student token)

#### Request body

```json
{
  "application_id": 14,
  "amount": "500.00",
  "payment_method": "bKash",
  "transaction_id": "TRX12345678"
}
```

#### Response `200 OK`

```json
{
  "status": "success",
  "message": "Payment received successfully! Application is now complete.",
  "application_id": 14,
  "transaction_id": "TRX12345678"
}
```

#### Possible errors

- `400 Bad Request` — invalid body or already paid
- `401 Unauthorized`
- `404 Not Found` — application not found / does not belong to student
- `500 Internal Server Error`

---

# 3. Admin endpoints

## 3.1 Admin login

### POST `/admin/login`

Authenticates an admin user and returns an admin JWT.

- **Auth required**: No

#### Request body

```json
{
  "email": "admin@system.com",
  "password": "admin123secret"
}
```

#### Response `200 OK`

```json
{
  "token": "<admin-jwt-token>",
  "role": "ADMIN"
}
```

#### Possible errors

- `400 Bad Request`
- `401 Unauthorized`
- `500 Internal Server Error`

---

## 3.2 Admin application management

### GET `/admin/applications`

Returns applications for a university.

- **Auth required**: Yes (admin token)

#### Query parameters

- `u_id` _(required)_

#### Example request

```http
GET /admin/applications?u_id=1
```

#### Response `200 OK`

> `sub_date` is a nullable-wrapper object.

```json
[
  {
    "app_id": 14,
    "student_id": 1,
    "program_id": 3,
    "program_name": "Computer Science and Engineering",
    "status": "Pending",
    "sub_date": {
      "Time": "2026-07-28T10:30:00Z",
      "Valid": true
    }
  }
]
```

#### Important note

Current backend does **not** return student `first_name`, `last_name`, or `email` here.

#### Possible errors

- `400 Bad Request`
- `401 Unauthorized`
- `403 Forbidden`
- `500 Internal Server Error`

---

### PUT `/admin/applications/status`

### PATCH `/admin/applications/status`

Updates application status.

- **Auth required**: Yes (admin token)

#### Request body

```json
{
  "app_id": 14,
  "status": "APPROVED"
}
```

#### Response `200 OK`

```json
{
  "status": "success",
  "message": "Application status updated to APPROVED"
}
```

#### Possible errors

- `400 Bad Request`
- `401 Unauthorized`
- `403 Forbidden`
- `405 Method Not Allowed`
- `500 Internal Server Error`

---

## 3.3 Admin university management

### POST `/admin/university`

Creates a university with optional departments and album pictures.

- **Auth required**: Yes (admin token)

#### Request body

```json
{
  "name": "BUET",
  "website": "https://buet.ac.bd",
  "location": "Dhaka",
  "logo_url": "https://example.com/buet.png",
  "university_description": "A leading engineering university.",
  "university_history": "Founded in ...",
  "departments": [
    {
      "dept_name": "CSE",
      "dept_description": "Computer Science and Engineering",
      "total_seats": 120
    }
  ],
  "album": [
    {
      "picture_title": "Main Gate",
      "picture_url": "https://example.com/gate.jpg"
    }
  ]
}
```

#### Response `201 Created`

```json
{
  "message": "University created successfully",
  "u_id": 1
}
```

#### Possible errors

- `400 Bad Request`
- `401 Unauthorized`
- `403 Forbidden`
- `500 Internal Server Error`

---

### PUT `/admin/university`

Updates a university and replaces its departments/album with the supplied arrays.

- **Auth required**: Yes (admin token)

#### Query parameters

- `u_id` _(required)_

#### Example request

```http
PUT /admin/university?u_id=1
```

#### Request body

```json
{
  "name": "BUET",
  "website": "https://buet.ac.bd",
  "location": "Dhaka",
  "logo_url": "https://example.com/buet-new.png",
  "university_description": "Updated description",
  "university_history": "Updated history",
  "departments": [
    {
      "dept_name": "EEE",
      "dept_description": "Electrical and Electronic Engineering",
      "total_seats": 120
    }
  ],
  "album": [
    {
      "picture_title": "Updated Campus",
      "picture_url": "https://example.com/campus-new.jpg"
    }
  ]
}
```

#### Response `200 OK`

```json
{
  "message": "University updated successfully"
}
```

#### Possible errors

- `400 Bad Request`
- `401 Unauthorized`
- `403 Forbidden`
- `500 Internal Server Error`

---

### DELETE `/admin/university`

Deletes a university.

- **Auth required**: Yes (admin token)

#### Query parameters

- `u_id` _(required)_

#### Example request

```http
DELETE /admin/university?u_id=1
```

#### Response `200 OK`

```json
{
  "message": "University deleted successfully"
}
```

#### Possible errors

- `400 Bad Request`
- `401 Unauthorized`
- `403 Forbidden`
- `500 Internal Server Error`

---

## 3.4 Admin program management

### POST `/admin/program`

Creates a program.

- **Auth required**: Yes (admin token)

#### Request body

```json
{
  "p_name": "Computer Science and Engineering",
  "p_unit": "A",
  "total_seats": 120,
  "prev_cutmarks": 85.5,
  "deadline": "2026-12-31",
  "u_id": 1
}
```

#### Response `201 Created`

```json
{
  "message": "Program created successfully",
  "program_id": 3
}
```

#### Possible errors

- `400 Bad Request`
- `401 Unauthorized`
- `403 Forbidden`
- `500 Internal Server Error`

---

### PUT `/admin/program`

Updates a program.

- **Auth required**: Yes (admin token)

#### Query parameters

- `program_id` _(required)_

#### Example request

```http
PUT /admin/program?program_id=3
```

#### Request body

```json
{
  "p_name": "Computer Science and Engineering",
  "p_unit": "A",
  "total_seats": 130,
  "prev_cutmarks": 86.0,
  "deadline": "2026-12-31",
  "u_id": 1
}
```

#### Response `200 OK`

```json
{
  "message": "Program updated successfully"
}
```

#### Possible errors

- `400 Bad Request`
- `401 Unauthorized`
- `403 Forbidden`
- `500 Internal Server Error`

---

### DELETE `/admin/program`

Deletes a program.

- **Auth required**: Yes (admin token)

#### Query parameters

- `program_id` _(required)_

#### Example request

```http
DELETE /admin/program?program_id=3
```

#### Response `200 OK`

```json
{
  "message": "Program deleted successfully"
}
```

#### Possible errors

- `400 Bad Request`
- `401 Unauthorized`
- `403 Forbidden`
- `500 Internal Server Error`

---

## 3.5 Admin admission test management

### POST `/admin/admission-test`

Creates an admission test linked to a program.

- **Auth required**: Yes (admin token)

#### Request body

```json
{
  "exam_unit": "A",
  "exam_center": "Dhaka",
  "exam_date": "2026-11-15",
  "prereq_test_id": 2,
  "program_id": 3
}
```

#### Notes

- `prereq_test_id` is optional
- `exam_date` format must be `YYYY-MM-DD`
- `program_id` is required

#### Response `201 Created`

```json
{
  "message": "Admission test created successfully",
  "test_id": 7
}
```

#### Possible errors

- `400 Bad Request`
- `401 Unauthorized`
- `403 Forbidden`
- `500 Internal Server Error`

---

### PUT `/admin/admission-test`

Updates an admission test.

- **Auth required**: Yes (admin token)

#### Query parameters

- `test_id` _(required)_

#### Example request

```http
PUT /admin/admission-test?test_id=7
```

#### Request body

```json
{
  "exam_unit": "A",
  "exam_center": "Chattogram",
  "exam_date": "2026-11-20",
  "prereq_test_id": 2,
  "program_id": 3
}
```

#### Response `200 OK`

```json
{
  "message": "Admission test updated successfully"
}
```

#### Possible errors

- `400 Bad Request`
- `401 Unauthorized`
- `403 Forbidden`
- `404 Not Found`
- `500 Internal Server Error`

---

# 4. Route summary

## Public

- `GET /health`
- `POST /register`
- `POST /login`
- `GET /programs`
- `GET /programs/detail`
- `GET /universities`
- `GET /universities/detail`
- `POST /admin/login`

## Student protected

- `GET /profile`
- `POST /student/profile`
- `PUT /student/profile`
- `GET /student/mobile`
- `POST /student/mobile`
- `PUT /student/mobile`
- `DELETE /student/mobile`
- `GET /programs/eligible`
- `POST /applications/apply`
- `GET /applications`
- `GET /program/requirements`
- `POST /payments/process`

## Admin protected

- `GET /admin/applications`
- `PUT /admin/applications/status`
- `PATCH /admin/applications/status`
- `POST /admin/university`
- `PUT /admin/university`
- `DELETE /admin/university`
- `POST /admin/program`
- `PUT /admin/program`
- `DELETE /admin/program`
- `POST /admin/admission-test`
- `PUT /admin/admission-test`
