drop table if exists blah;

create table blah (
    id serial primary key,
    is_deleted bool not null
)
