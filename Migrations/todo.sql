CREATE TABLE todos (
    id          BIGINT AUTO_INCREMENT,
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO todos (title, description)
VALUES ('go-workshop', 'gopher');
