CREATE TYPE sex as ENUM('male', 'female');

CREATE TYPE race AS ENUM(
    'Maine Coon',
    'Persian',
    'Ragdoll',
    'Siamese',
    'Bengal',
    'Sphynx',
    'British Shorthair',
    'Abyssinian',
    'Scottish Fold',
    'Birman' 
);

CREATE TABLE IF NOT EXISTS cats (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    race race NOT NULL,
    sex sex NOT NULL,
    ageInMonth INT NOT NULL,
    description VARCHAR(200) NOT NULL,
    imageUrls TEXT [] NOT NULL,
	hasMatched BOOLEAN NOT NULL DEFAULT false,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	ownerId INT,
  	FOREIGN KEY(ownerId) 
	REFERENCES users(id)
);
