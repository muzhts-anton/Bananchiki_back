DROP TABLE IF EXISTS users          CASCADE;
DROP TABLE IF EXISTS presentation   CASCADE;
DROP TABLE IF EXISTS slideorder     CASCADE;
DROP TABLE IF EXISTS quiz           CASCADE;
DROP TABLE IF EXISTS convertedslide CASCADE;
DROP TABLE IF EXISTS vote           CASCADE;


CREATE TABLE users (
    id BIGSERIAL NOT NULL PRIMARY KEY
);

CREATE TABLE presentation (
    id                      BIGSERIAL NOT NULL PRIMARY KEY,
    creator_id              BIGINT REFERENCES users (id),
    url                     VARCHAR(128) DEFAULT '/static/presentations/' NOT NULL,
    converted_slide_num     SMALLINT NOT NULL,
    quiz_num                SMALLINT NOT NULL
);

CREATE TABLE slideorder (
    id              BIGSERIAL NOT NULL PRIMARY KEY,
    presentation_id BIGINT REFERENCES presentation (id),
    type            VARCHAR(64) NOT NULL,
    item_id         BIGINT NOT NULL,
    idx             SMALLINT NOT NULL
);

CREATE TABLE quiz (
    id          BIGSERIAL NOT NULL PRIMARY KEY,
    type        VARCHAR(64) DEFAULT 'horizontal' NOT NULL,
    question    VARCHAR(512) NOT NULL,
    background  VARCHAR(16) NOT NULL,
    font_color  VARCHAR(16) NOT NULL,
    font_size   SMALLINT NOT NULL,
    graph_color VARCHAR(16) NOT NULL
);

CREATE TABLE convertedslide (
    id      BIGSERIAL NOT NULL PRIMARY KEY,
    name    VARCHAR(64) NOT NULL,
    width   SMALLINT NOT NULL,
    height  SMALLINT NOT NULL
);

CREATE TABLE vote (
    id          BIGSERIAL NOT NULL PRIMARY KEY,
    quiz_id     BIGINT REFERENCES quiz (id),
    idx         SMALLINT NOT NULL,
    option      VARCHAR(512) NOT NULL,
    votes_num   BIGINT NOT NULL,
    color       VARCHAR(16) NOT NULL
);

