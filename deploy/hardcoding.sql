insert into users (id, nick, email, goal, sex, hashed_password) values (100, 'Deniskas', 'denis.kalina@gmail.com', '15 новых слов', 'Мужской', '1234');
insert into users (id, nick, email, goal, sex, hashed_password) values (101, 'Angelina', 'biba.com', '10 новых слов', 'Женский', '$2a$10$R/u9KWwNi42UqHChsVClM.XaIpv1xQsK2avxlLWzDPX9NzqU8VcFG');

insert into answers (id, name, is_correct) values (100, 'Алга', true);
insert into answers (id, name, is_correct) values (111, 'Не знаю', false);
insert into answers (id, name, is_correct) values (112, 'Вперед', false);
insert into questions (id, name, description) values (101, 'Вопрос 1', 'Как будет "Вперед" на татарском?');

insert into answers_questions (id_answer, id_question) values (111, 101);
insert into answers_questions (id_answer, id_question) values (100, 101);
insert into answers_questions (id_answer, id_question) values (112, 101);
insert into questions (id, name, description) values (103, 'Вопрос 2', 'Как будет "Привет" на татарском?');

insert into answers (id, name, is_correct) values (102, 'Салам', true);
insert into answers (id, name, is_correct) values (122, 'Привет', false);
insert into answers (id, name, is_correct) values (123, 'Шалом', false);
insert into answers_questions (id_answer, id_question) values (102, 103);
insert into answers_questions (id_answer, id_question) values (122, 103);
insert into answers_questions (id_answer, id_question) values (123, 103);
insert into answers (id, name, is_correct) values (104, 'Мин сине яратам', true);
insert into answers (id, name, is_correct) values (144, 'I live you', false);
insert into answers (id, name, is_correct) values (145, 'Я тебя люблю', false);
insert into questions (id, name, description) values (105, 'Вопрос 3', 'Как будет "Я тебя люблю" на татарском?');

insert into answers_questions (id_answer, id_question) values (104, 105);
insert into answers_questions (id_answer, id_question) values (144, 105);
insert into answers_questions (id_answer, id_question) values (145, 105);

insert into answers (id, name, is_correct) values (200, 'Ничек', true);
insert into answers (id, name, is_correct) values (201, 'How are you', false);
insert into answers (id, name, is_correct) values (202, 'Я не знаю', false);
insert into questions (id, name, description) values (203, 'Вопрос 4', 'Как будет "Как дела" на татарском?');

insert into answers_questions (id_answer, id_question) values (200, 203);
insert into answers_questions (id_answer, id_question) values (201, 203);
insert into answers_questions (id_answer, id_question) values (202, 203);
insert into tests (id, name, speed, count_q) values (100, 'Борьба умов', 'fastik', 4);


insert into questions_tests (id_tests, id_questions) values (100, 105);
insert into questions_tests (id_tests, id_questions) values (100, 103);
insert into questions_tests (id_tests, id_questions) values (100, 101);
insert into questions_tests (id_tests, id_questions) values (100, 203);

CREATE TABLE IF NOT EXISTS fight
(
    session int default 100,
    test int default 100,
    id_1 int default 100,
    id_2 int default 101,
    res_1 int default 2,
    res_2 int default 0
);

insert into fight (session) values (100);