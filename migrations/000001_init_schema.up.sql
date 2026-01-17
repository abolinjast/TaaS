-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- 1. USERS
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(100),
    last_login_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 2. API KEYS
CREATE TABLE api_keys (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    label VARCHAR(50), 
    key_hash VARCHAR(64) UNIQUE NOT NULL,
    last_used_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 3. COURSES
CREATE TABLE courses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    source_url VARCHAR(500), 
    platform VARCHAR(100), 
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'archived', 'completed')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 4. STUDY SESSIONS
CREATE TABLE study_sessions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    course_id UUID REFERENCES courses(id) ON DELETE CASCADE NOT NULL,
    module VARCHAR(255),
    topic VARCHAR(255),
    -- Time Tracking
    start_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    end_time TIMESTAMP WITH TIME ZONE,
    duration_seconds INTEGER, 
    -- Activity Type Logic
    -- 'study': Standard video watching or reading
    -- 'quiz': Taking an assessment
    activity_type VARCHAR(20) DEFAULT 'study' CHECK (activity_type IN ('study', 'quiz')),
    -- Quiz Specific Columns (Nullable)
    quiz_score SMALLINT CHECK (quiz_score >= 0 AND quiz_score <= 100),
    quiz_passed BOOLEAN, 
    -- Status
    status VARCHAR(20) DEFAULT 'running' CHECK (status IN ('running', 'completed', 'abandoned')),
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- INDEXES
CREATE INDEX idx_sessions_user_status ON study_sessions(user_id, status);

CREATE INDEX idx_sessions_start_time ON study_sessions(start_time DESC);

CREATE INDEX idx_sessions_course_id ON study_sessions(course_id);

CREATE INDEX idx_sessions_quiz_performance ON study_sessions(user_id, quiz_passed) WHERE activity_type = 'quiz';
