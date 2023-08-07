CREATE TABLE users
(
    id            serial PRIMARY KEY,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE brands
(
    id serial PRIMARY KEY,
    name varchar(255) not null,
    imageBrand int not null
);

CREATE TABLE scores
(
    id serial PRIMARY KEY,
    player1Score int not null,
    player2Score int not null
);

CREATE TABLE battles
(
    id serial PRIMARY KEY,
    player1Id serial REFERENCES users(id) ON DELETE CASCADE not null,
    player2Id serial REFERENCES users(id) ON DELETE CASCADE not null,
    scoreId serial REFERENCES scores(id) ON DELETE CASCADE not null,
    currentBrandId serial REFERENCES brands(id) ON DELETE CASCADE not null
);
