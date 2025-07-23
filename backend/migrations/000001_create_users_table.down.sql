-- Drop tables in reverse order of creation to respect foreign key dependencies

DROP TABLE IF EXISTS files;
DROP TABLE IF EXISTS task;
DROP TABLE IF EXISTS admin;
DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS job_title;
DROP TABLE IF EXISTS company;

-- Drop enum type
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status') THEN
        DROP TYPE status;
    END IF;
END$$;

-- Drop UUID extension only if desired (optional and safe to leave)
-- DROP EXTENSION IF EXISTS "uuid-ossp";
