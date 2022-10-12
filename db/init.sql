CREATE TABLE user
(
  id           INT(10) NOT NULL PRIMARY KEY,
  name     VARCHAR(40) NOT NULL
);

INSERT INTO user (id, name) VALUES (1, "hoge");
INSERT INTO user (id, name) VALUES (2, "fuga");
