drop table if exists book;
drop table if exists author;

create table author
(
    id   serial primary key,
    books jsonb not null default '[]'
);
