-- Step 1: Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Step 2: Create ENUM type for task status
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status') THEN
        CREATE TYPE status AS ENUM ('OPEN', 'IN_PROGRESS', 'CLOSED');
    END IF;
END$$;

-- Step 3: Company table
CREATE TABLE company (
    company_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT UNIQUE NOT NULL,
    description TEXT
);

-- Step 4: Job title table
CREATE TABLE job_title (
    job_title_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL
);

-- Step 5: User table
CREATE TABLE "user" (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    phone_number VARCHAR(127),
    role UUID REFERENCES job_title(job_title_id),
    is_admin BOOLEAN DEFAULT FALSE,
    session TEXT,
    password_hash TEXT
);

-- Step 6: Admin table
CREATE TABLE admin (
    admin_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID UNIQUE REFERENCES "user"(user_id) ON DELETE CASCADE
);

-- Step 7: Task table
CREATE TABLE task (
    task_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    author_id UUID REFERENCES "user"(user_id) ON DELETE SET NULL,
    header VARCHAR(255),
    body TEXT,
    start_deadline TIMESTAMP,
    end_deadline TIMESTAMP,
    company_id UUID REFERENCES company(company_id) ON DELETE SET NULL,
    status status,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Step 8: Files table
CREATE TABLE files (
    file_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    task_id UUID REFERENCES task(task_id) ON DELETE CASCADE,
    filename VARCHAR(255) UNIQUE NOT NULL,
    uri TEXT UNIQUE NOT NULL
);
