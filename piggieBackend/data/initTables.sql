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

CREATE TABLE wallet(
    userId INTEGER REFERENCES users(id),
    balance NUMERIC(12, 2) NOT NULL,
    currency INTEGER NOT NULL
);

CREATE TABLE transactions(
    id SERIAL PRIMARY KEY,
    userId INTEGER REFERENCES users(id),
    transactionType INTEGER NOT NULL,
    amount NUMERIC(12, 2) NOT NULL,
    transactionDate DATE NOT NULL 
);

CREATE TABLE transactionsHistory(
    userId INTEGER REFERENCES users(id),
    numberOfTransactions INTEGER NOT NULL,
    totalEarnings NUMERIC(12, 2) NOT NULL,
    totalSpendings NUMERIC(12, 2)
);
