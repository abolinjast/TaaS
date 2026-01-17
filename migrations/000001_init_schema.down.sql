-- Drop indexes 
DROP INDEX IF EXISTS idx_sessions_quiz_performance;
DROP INDEX IF EXISTS idx_sessions_course_id;
DROP INDEX IF EXISTS idx_sessions_start_time;
DROP INDEX IF EXISTS idx_sessions_user_status;

-- Drop tables in reverse dependency order
DROP TABLE IF EXISTS study_sessions;
DROP TABLE IF EXISTS courses;
DROP TABLE IF EXISTS api_keys;
DROP TABLE IF EXISTS users;

-- Drop extension
DROP EXTENSION IF EXISTS "uuid-ossp";
