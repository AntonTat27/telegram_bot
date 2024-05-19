CREATE TABLE messages (
      id bigserial constraint messages_pk
          primary key,
      created_at timestamp with time zone not null,
      sending_date timestamp with time zone,
      message varchar,
      sender_id int,
      message_id int
);

CREATE TABLE filtered_messages (
       id bigserial constraint filtered_messages_pk
           primary key,
       created_at timestamp with time zone not null,
       sending_date timestamp with time zone,
       sender_id int,
       message varchar,
       message_id int,
       word_filtered varchar
);