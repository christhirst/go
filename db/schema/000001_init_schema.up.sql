create table accounts
(
  user_id bigserial PRIMARY KEY,
  username varchar NOT NULL,
  password varchar NOT NULL
);



create table demands
(
  id bigserial PRIMARY KEY,
  title varchar NOT NULL,
  description varchar,
  url varchar,
  price varchar,
  isDone varchar,
  account_id bigint NOT NULL
);


create table demand_transfer
(
  id bigserial PRIMARY KEY,
  from_account_id bigint NOT NULL,
  to_account_id bigint NOT NULL,
  demand_id bigint NOT NULL
);

ALTER TABLE "demands" ADD FOREIGN KEY ("account_id") REFERENCES accounts ("user_id");
ALTER TABLE "demand_transfer" ADD FOREIGN KEY ("from_account_id") REFERENCES accounts ("user_id");
ALTER TABLE "demand_transfer" ADD FOREIGN KEY ("to_account_id") REFERENCES accounts ("user_id");
ALTER TABLE "demand_transfer" ADD FOREIGN KEY ("demand_id") REFERENCES demands ("id");

