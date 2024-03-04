CREATE SEQUENCE order_id_seq;

CREATE TABLE IF NOT EXISTS orders (
  id integer primary key DEFAULT nextval('order_id_seq'),
  car_id integer NOT NULL,
  order_date date NOT NULL,
  pickup_date date NOT NULL,
  dropoff_date date NOT NULL,
  pickup_location varchar(50) NOT NULL,
  pickup_lat double precision NOT NULL,
  pickup_long double precision NOT NULL,
  dropoff_location varchar(50) NOT NULL,
  dropoff_lat double precision NOT NULL,
  dropoff_long double precision NOT NULL,
  created_by integer default 0 NOT NULL,
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_by integer default 0 NOT NULL,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  deleted_by integer,
  deleted_at timestamp WITH TIME ZONE
);

ALTER SEQUENCE order_id_seq OWNED BY orders.id;