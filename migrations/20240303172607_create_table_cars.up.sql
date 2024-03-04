CREATE SEQUENCE car_id_seq;

CREATE TABLE IF NOT EXISTS cars (
  id integer primary key DEFAULT nextval('car_id_seq'),
  car_name    varchar(50) NOT NULL,
  day_rate double precision NOT NULL,
  month_rate double precision NOT NULL,
  image varchar(256) NOT NULL,
  created_by integer default 0 NOT NULL,
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_by integer default 0 NOT NULL,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  deleted_by integer,
  deleted_at timestamp WITH TIME ZONE
);

ALTER SEQUENCE car_id_seq OWNED BY cars.id;