CREATE TABLE languages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code VARCHAR(5) NOT NULL,
    name VARCHAR(100) NOT NULL,
    flag_emojis VARCHAR(100),
    level VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE lessons (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(250),
    level VARCHAR(100),
    completed bool,
    languages_code VARCHAR(5) NOT NULL,
    content JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_lessons (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID,
    lesson_id UUID REFERENCES lessons(id),
    completed_at TIMESTAMP WITH TIME ZONE,
    UNIQUE (user_id, lesson_id)
);
CREATE TABLE vocabulary (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    language_code VARCHAR(5) NOT NULL,
    word VARCHAR(100) NOT NULL,
    translation VARCHAR(100) NOT NULL,
    example_sentence TEXT,
    audio_url VARCHAR(250),
    part_of_speech VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_vocabulary (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID,
    vocabulary_id UUID REFERENCES vocabulary(id),
    learned_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    last_reviewed_at TIMESTAMP WITH TIME ZONE,
    mastery_level INTEGER DEFAULT 0,
    UNIQUE (user_id, vocabulary_id)
);



CREATE TABLE exercises (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    type VARCHAR(50),
    question VARCHAR(250),
    correet_answer VARCHAR(250),
    lesson_id UUID REFERENCES lessons(id)
);








