CREATE TABLE flights (
  flight_id UUID PRIMARY KEY,
  plane_number VARCHAR(255) REFERENCES planes(plane_number),
  departure_point VARCHAR(255) NOT NULL,
  destination_point VARCHAR(255) NOT NULL,
  departure_time TIMESTAMP NOT NULL,
  estimated_arrival_time TIMESTAMP NOT NULL,
  available_seats INTEGER NOT NULL
  real_departure_time TIMESTAMP,
  real_arrival_time TIMESTAMP,
  status VARCHAR(255) NOT NULL CHECK (status IN ('scheduled', 'flying', 'finished'))
);

ALTER TABLE flights ADD CONSTRAINT flights_plane_number_fkey
    FOREIGN KEY (plane_number)
    REFERENCES planes (plane_number)
    ON DELETE RESTRICT;