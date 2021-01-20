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
  owner varchar
);


create table demand_transfer
(
  id VARCHAR NOT NULL PRIMARY KEY,
  from_account_id varchar NOT NULL,
  to_account_id varchar NOT NULL,
  demand_id varchar NOT NULL
);

ALTER TABLE "demands" ADD FOREIGN KEY ("owner") REFERENCES accounts ("user_id");
ALTER TABLE "demand_transfer" ADD FOREIGN KEY ("from_account_id") REFERENCES accounts ("user_id");
ALTER TABLE "demand_transfer" ADD FOREIGN KEY ("to_account_id") REFERENCES accounts ("user_id");
ALTER TABLE "demand_transfer" ADD FOREIGN KEY ("demand_id") REFERENCES demands ("id");

