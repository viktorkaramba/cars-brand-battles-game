CREATE TABLE users
(
    id            serial PRIMARY KEY,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE tokens
(
    id serial PRIMARY KEY,
    tokenValue varchar not null unique ,
    revoked varchar(255) not null,
    userId serial REFERENCES users(id) ON DELETE CASCADE not null
);

CREATE TABLE brands
(
    id serial PRIMARY KEY,
    name varchar(255) not null,
    imageBrand varchar not null
);

CREATE TABLE battles
(
    id serial PRIMARY KEY,
    player1Id serial REFERENCES users(id) ON DELETE CASCADE not null,
    player2Id serial REFERENCES users(id) ON DELETE CASCADE not null,
    punishment varchar not null,
    isFinished varchar(255) not null,
    currentBrandId serial REFERENCES brands(id) ON DELETE CASCADE not null
);


CREATE TABLE scores
(
    id serial PRIMARY KEY,
    userId serial REFERENCES users(id) ON DELETE CASCADE not null,
    battleId serial REFERENCES battles(id) ON DELETE CASCADE not null,
    playerScore int not null
);

