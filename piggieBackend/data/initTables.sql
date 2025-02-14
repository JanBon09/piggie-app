CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(32) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    email VARCHAR(254) UNIQUE NOT NULL,
    dateOfBirth DATE NOT NULL,
    salt TEXT UNIQUE NOT NULL,
    name VARCHAR(32),
    surname VARCHAR(32),
    country INT,
    profilePictureURL TEXT
);