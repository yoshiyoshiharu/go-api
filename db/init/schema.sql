CREATE TABLE IF NOT EXISTS users (
  id INTEGER NOT NULL auto_increment PRIMARY KEY,
  email VARCHAR(255) NOT NULL,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL
);

INSERT INTO users(email, first_name, last_name) VALUES
  ("test1@test.com", "Test1", "User1"),
  ("test2@test.com", "Test2", "User2"),
  ("test3@test.com", "Test3", "User3");
