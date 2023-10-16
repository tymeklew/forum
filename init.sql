use forum;
CREATE TABLE IF NOT EXISTS users (
    uuid varchar(36) NOT NULL,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    PRIMARY KEY (uuid)
);
CREATE TABLE IF NOT EXISTS posts (
    uuid  varchar(36) NOT NULL,
    owner varchar(36) NOT NULL,
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    PRIMARY KEY (uuid),
    FOREIGN KEY (owner) REFERENCES users(uuid)
);
CREATE TABLE IF NOT EXISTS sessions (
    uuid varchar(36) NOT NULL,
    user varchar(36) NOT NULL,
    PRIMARY KEY (uuid),
    FOREIGN KEY (user) REFERENCES users(uuid)
)
