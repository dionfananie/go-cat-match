CREATE TYPE sex_type as ENUM('male', 'female');

CREATE TYPE race_type AS ENUM(
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
    race race_type NOT NULL,
    sex sex_type NOT NULL,
    ageInMonth INT NOT NULL,
    description VARCHAR(200) NOT NULL,
    imageUrls TEXT [] NOT NULL,
	hasMatched BOOLEAN NOT NULL DEFAULT false,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	ownerId INT,
  	FOREIGN KEY(ownerId) 
	REFERENCES users(id)
);
