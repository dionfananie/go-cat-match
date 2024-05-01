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
    ageInMonth SMALLINT NOT NULL,
    description VARCHAR(200) NOT NULL,
    imageUrls TEXT [] NOT NULL,
	hasMatched BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	ownerId SERIAL,
    CONSTRAINT fk_users
      FOREIGN KEY(ownerId) 
        REFERENCES users(id)
);
