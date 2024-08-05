-- Users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Countries table
CREATE TABLE countries (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    flag VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Events table
CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    sport_type VARCHAR(50) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Athletes table
CREATE TABLE athletes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    country_id INT REFERENCES countries(id) ON DELETE CASCADE,
    sport_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Medals table
CREATE TABLE medals (
    id SERIAL PRIMARY KEY,
    country_id INT REFERENCES countries(id) ON DELETE CASCADE,
    type VARCHAR(10) CHECK (type IN ('Gold', 'Silver', 'Bronze')) NOT NULL,
    event_id INT REFERENCES events(id) ON DELETE CASCADE,
    athlete_id INT REFERENCES athletes(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- mock data
-- Users
INSERT INTO users (username, password, role) VALUES
('user1', '$2a$12$2bQKuzshRz7eF8obx7xPbuhgV/63pNCvMQd8XXGxBhUIxMGX8JJq2', 'user'),
('user2', '$2a$12$2bQKuzshRz7eF8obx7xPbuhgV/63pNCvMQd8XXGxBhUIxMGX8JJq2', 'user'),
('user3', '$2a$12$2bQKuzshRz7eF8obx7xPbuhgV/63pNCvMQd8XXGxBhUIxMGX8JJq2', 'user'),
('user4', '$2a$12$2bQKuzshRz7eF8obx7xPbuhgV/63pNCvMQd8XXGxBhUIxMGX8JJq2', 'user'),
('user5', '$2a$12$2bQKuzshRz7eF8obx7xPbuhgV/63pNCvMQd8XXGxBhUIxMGX8JJq2', 'user'),
('admin1', '$2a$12$2bQKuzshRz7eF8obx7xPbuhgV/63pNCvMQd8XXGxBhUIxMGX8JJq2', 'admin'),
('admin2', '$2a$12$2bQKuzshRz7eF8obx7xPbuhgV/63pNCvMQd8XXGxBhUIxMGX8JJq2', 'admin'),
('user6', '$2a$12$2bQKuzshRz7eF8obx7xPbuhgV/63pNCvMQd8XXGxBhUIxMGX8JJq2', 'user'),
('user7', '$2a$12$2bQKuzshRz7eF8obx7xPbuhgV/63pNCvMQd8XXGxBhUIxMGX8JJq2', 'user'),
('user8', '$2a$12$2bQKuzshRz7eF8obx7xPbuhgV/63pNCvMQd8XXGxBhUIxMGX8JJq2', 'user');

-- Countries
INSERT INTO countries (name, flag) VALUES
('USA', 'usa_flag.png'),
('China', 'china_flag.png'),
('Russia', 'russia_flag.png'),
('Japan', 'japan_flag.png'),
('Germany', 'germany_flag.png'),
('UK', 'uk_flag.png'),
('France', 'france_flag.png'),
('Italy', 'italy_flag.png'),
('Canada', 'canada_flag.png'),
('Australia', 'australia_flag.png');

-- Events
INSERT INTO events (name, sport_type, start_time, end_time) VALUES
('100m Sprint', 'Athletics', '2024-07-26 10:00:00', '2024-07-26 12:00:00'),
('200m Sprint', 'Athletics', '2024-07-27 10:00:00', '2024-07-27 12:00:00'),
('Marathon', 'Athletics', '2024-07-28 06:00:00', '2024-07-28 12:00:00'),
('100m Backstroke', 'Swimming', '2024-07-29 10:00:00', '2024-07-29 12:00:00'),
('200m Freestyle', 'Swimming', '2024-07-30 10:00:00', '2024-07-30 12:00:00'),
('Basketball Final', 'Basketball', '2024-07-31 18:00:00', '2024-07-31 20:00:00'),
('Football Final', 'Football', '2024-08-01 18:00:00', '2024-08-01 20:00:00'),
('Judo Final', 'Judo', '2024-08-02 14:00:00', '2024-08-02 16:00:00'),
('Tennis Final', 'Tennis', '2024-08-03 16:00:00', '2024-08-03 18:00:00'),
('Boxing Final', 'Boxing', '2024-08-04 20:00:00', '2024-08-04 22:00:00');

-- Athletes
INSERT INTO athletes (name, country_id, sport_type) VALUES
('Athlete1', 1, 'Athletics'),
('Athlete2', 2, 'Athletics'),
('Athlete3', 3, 'Athletics'),
('Athlete4', 4, 'Swimming'),
('Athlete5', 5, 'Swimming'),
('Athlete6', 6, 'Basketball'),
('Athlete7', 7, 'Football'),
('Athlete8', 8, 'Judo'),
('Athlete9', 9, 'Tennis'),
('Athlete10', 10, 'Boxing');

-- Medals
INSERT INTO medals (country_id, type, event_id, athlete_id) VALUES
(1, 'Gold', 1, 1),
(2, 'Silver', 1, 2),
(3, 'Bronze', 1, 3),
(4, 'Gold', 2, 4),
(5, 'Silver', 2, 5),
(6, 'Bronze', 2, 6),
(7, 'Gold', 3, 7),
(8, 'Silver', 3, 8),
(9, 'Bronze', 3, 9),
(10, 'Gold', 4, 10);


-- 20240804_create_events_table.sql
CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    sport_type VARCHAR(100) NOT NULL,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL
);

