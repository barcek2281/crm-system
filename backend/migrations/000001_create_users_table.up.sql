-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create ENUM type for task status
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status') THEN
        CREATE TYPE status AS ENUM ('OPEN', 'IN_PROGRESS', 'CLOSED');
    END IF;
END$$;

-- Company table
CREATE TABLE company (
    company_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT
);

-- Job title table
CREATE TABLE job_title (
    job_title_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL
);

-- User table
CREATE TABLE "user" (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    job_id UUID REFERENCES job_title(job_title_id),
    company_id UUID REFERENCES company(company_id),
    is_admin BOOLEAN DEFAULT FALSE,
    password_hash TEXT,
    login VARCHAR(255) UNIQUE NOT NULL
);

-- Task table
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

-- Files table
CREATE TABLE files (
    file_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    task_id UUID REFERENCES task(task_id) ON DELETE CASCADE,
    filename VARCHAR(255) UNIQUE NOT NULL,
    uri TEXT UNIQUE NOT NULL
);
