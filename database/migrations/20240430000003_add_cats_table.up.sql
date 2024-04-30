CREATE TYPE sex as ENUM("male", "female")
CREATE TYPE race AS ENUM(
    "Maine Coon",
    "Persian",
    "Ragdoll",
    "Siamese",
    "Bengal",
    "Sphynx",
    "British Shorthair",
    "Abyssinian",
    "Scottish Fold",
    "Birman" 
)

CREATE TABLE cats (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    race race NOT NULL,
    sex sex NOT NULL,
    ageInMonth SMALLINT NOT NULL
    description VARCHAR(200) NOT NULL
    imageUrls TEXT [] NOT NULL
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);