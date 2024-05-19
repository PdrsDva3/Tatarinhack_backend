-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS users
(
    id              SERIAL NOT NULL PRIMARY KEY,
    nick            varchar UNIQUE,
    email           varchar UNIQUE,
    hashed_password VARCHAR,
    goal            VARCHAR,
    sex             varchar,
    exp             int default 0,
    cnt             int default 0,
    grammar         int default 0,
    vocabulary      int default 0,
    speaking        int default 0,
    rating          int default 0,
    lvl             int default 0,
    days            int default 0,
    achievement     int default 0
);

CREATE TABLE IF NOT EXISTS friends_link
(
    id_first  INTEGER REFERENCES users (id),
    id_second INTEGER REFERENCES users (id),
    nick      varchar references users (nick),
    sex       varchar
);

CREATE TABLE IF NOT EXISTS teachers
(
    id              SERIAL NOT NULL PRIMARY KEY,
    nick            varchar UNIQUE,
    email           varchar UNIQUE,
    hashed_password VARCHAR
);

CREATE TABLE IF NOT EXISTS courses
(
    id          SERIAL NOT NULL PRIMARY KEY,
    name        VARCHAR,
    level       int,
    description varchar
);

CREATE TABLE IF NOT EXISTS tests
(
    id      SERIAL NOT NULL PRIMARY KEY,
    name    VARCHAR,
    type    VARCHAR default '',
    level   int default 0,
    speed   VARCHAR,
    count_q int
);

CREATE TABLE IF NOT EXISTS theory
(
    id          SERIAL NOT NULL PRIMARY KEY,
    name        VARCHAR,
    description VARCHAR
);

CREATE TABLE IF NOT EXISTS questions
(
    id          SERIAL NOT NULL PRIMARY KEY,
    name        VARCHAR,
    description VARCHAR
);

CREATE TABLE IF NOT EXISTS answers
(
    id         SERIAL NOT NULL PRIMARY KEY,
    name       VARCHAR,
    is_correct bool
);


CREATE TABLE IF NOT EXISTS courses_tests
(
    id_tests   INTEGER REFERENCES tests (id),
    id_courses INTEGER REFERENCES courses (id)
);

CREATE TABLE IF NOT EXISTS theory_tests
(
    id_tests  INTEGER REFERENCES tests (id),
    id_theory INTEGER REFERENCES theory (id)
);

CREATE TABLE IF NOT EXISTS questions_tests
(
    id_tests     INTEGER REFERENCES tests (id),
    id_questions INTEGER REFERENCES questions (id)
);

CREATE TABLE IF NOT EXISTS audios
(
    id             SERIAL NOT NULL PRIMARY KEY,
    correct_answer VARCHAR
);

CREATE TABLE IF NOT EXISTS audio_tests
(
    id      SERIAL NOT NULL PRIMARY KEY,
    name    VARCHAR,
    type    VARCHAR,
    level   int,
    speed   VARCHAR,
    count_q int
);

CREATE TABLE IF NOT EXISTS audios_tests
(
    id_tests INTEGER REFERENCES audio_tests (id),
    id_audio INTEGER REFERENCES audios (id)
);

CREATE TABLE IF NOT EXISTS answers_questions
(
    id_answer   INTEGER REFERENCES answers (id),
    id_question INTEGER REFERENCES questions (id)
);


CREATE TABLE IF NOT EXISTS audios_tests
(
    id_tests INTEGER REFERENCES tests (id),
    id_audio INTEGER REFERENCES audios (id)
);

CREATE TABLE IF NOT EXISTS fight
(
    session int default 100,
    test int default 100,
    id_1 int default 100,
    id_2 int default 101,
    res_1 int default 2,
    res_2 int default 0
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table friends_link;
drop table users;
drop table teachers;
drop table audios_tests;
drop table audios;
drop table courses_tests;
drop table courses;
drop table theory_tests;
drop table theory;
drop table questions_tests;
drop table tests;
drop table answers_questions;
drop table questions;
drop table answers;
-- +goose StatementEnd
