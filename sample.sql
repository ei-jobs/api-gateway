use ei_jobs;

-- Insert first vacancy
INSERT INTO vacancies (
    user_id,
    specialization_id,
    title,
    country,
    city,
    salary_from,
    salary_to,
    salary_period,
    work_format,
    work_schedule
) VALUES (
    1,  -- assuming user_id 1 exists
    2,  -- assuming specialization_id 2 exists
    'Senior Software Engineer',
    'United States',
    'New York',
    120000,
    180000,
    'year',
    'hybrid',
    'full-time'
);

-- Get the id of the first inserted vacancy
SET @vacancy1_id = LAST_INSERT_ID();

-- Insert conditions for first vacancy
INSERT INTO vacancy_conditions (vacancy_id, icon, condition_text) VALUES
(@vacancy1_id, 'medical', 'Medical insurance for you and your family'),
(@vacancy1_id, 'laptop', 'Latest MacBook Pro'),
(@vacancy1_id, 'vacation', '25 days paid vacation'),
(@vacancy1_id, 'gym', 'Gym membership');

-- Insert requirements for first vacancy
INSERT INTO vacancy_requirements (vacancy_id, requirement) VALUES
(@vacancy1_id, '5+ years of experience in software development'),
(@vacancy1_id, 'Strong knowledge of Go, Python, or Java'),
(@vacancy1_id, 'Experience with distributed systems'),
(@vacancy1_id, 'Excellent problem-solving skills');

-- Insert second vacancy
INSERT INTO vacancies (
    user_id,
    specialization_id,
    title,
    country,
    city,
    salary_from,
    salary_to,
    salary_period,
    work_format,
    work_schedule
) VALUES (
    2,  -- assuming user_id 2 exists
    3,  -- assuming specialization_id 3 exists
    'UX/UI Designer',
    'Canada',
    'Toronto',
    75000,
    95000,
    'year',
    'remote',
    'full-time'
);

-- Get the id of the second inserted vacancy
SET @vacancy2_id = LAST_INSERT_ID();

-- Insert conditions for second vacancy
INSERT INTO vacancy_conditions (vacancy_id, icon, condition_text) VALUES
(@vacancy2_id, 'medical', 'Full health and dental coverage'),
(@vacancy2_id, 'education', 'Learning budget $2000/year'),
(@vacancy2_id, 'flexible', 'Flexible working hours'),
(@vacancy2_id, 'stock', 'Stock options');

-- Insert requirements for second vacancy
INSERT INTO vacancy_requirements (vacancy_id, requirement) VALUES
(@vacancy2_id, '3+ years of experience in UX/UI design'),
(@vacancy2_id, 'Proficiency in Figma and Adobe Creative Suite'),
(@vacancy2_id, 'Portfolio demonstrating user-centered design process'),
(@vacancy2_id, 'Experience with design systems');
