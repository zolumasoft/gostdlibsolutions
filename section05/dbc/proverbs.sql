/* Create database */
CREATE DATABASE gostdessentials CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

/* Add user and grant privileges */
CREATE USER IF NOT EXISTS 'username'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON gostdessentials.* TO 'username'@'%';
FLUSH PRIVILEGES;

/* Use database */
USE gostdessentials;

/* Create table */
CREATE TABLE gostdessentials.proverbs (
  id INT NOT NULL AUTO_INCREMENT,
  text VARCHAR(255),
  philosopher VARCHAR(30),
  PRIMARY KEY (ID)
);

/* Add data into database */
INSERT INTO gostdessentials.proverbs (text, philosopher) 
  VALUES('Waste no more time arguing what a good man should be. Be one.', 'Marcus Aurelius');
INSERT INTO gostdessentials.proverbs (text, philosopher) 
  VALUES('We suffer more in imagination than in reality.', 'Seneca');
INSERT INTO gostdessentials.proverbs (text, philosopher) 
  VALUES('The best revenge is not to be like your enemy.', 'Marcus Aurelius');
INSERT INTO gostdessentials.proverbs (text, philosopher) 
  VALUES('I begin to speak only when I''m certain what I''ll say isn''t better left unsaid.', 'Cato');
INSERT INTO gostdessentials.proverbs (text, philosopher) 
  VALUES('Seek not for events to happen as you wish but rather wish for events to happen as they do and your life will go smoothly.', 'Epictetus');
INSERT INTO gostdessentials.proverbs (text, philosopher) 
  VALUES('Do not wait for your ship to come in, swim out to it.', 'Unknown');
