create table accounts
(
  user_id varchar PRIMARY KEY,
  username varchar NOT NULL,
  password varchar NOT NULL
);



create table demands
(
  id VARCHAR NOT NULL PRIMARY KEY,
  title varchar NOT NULL,
  description varchar,
  url varchar,
  price varchar,
  isDone varchar,
  owner VARCHAR
);


ALTER TABLE "demands" ADD FOREIGN KEY ("owner") REFERENCES accounts ("user_id");

