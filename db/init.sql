CREATE TABLE movie (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(50), NOT NULL,
  desc TEXT,
  rating SMALLINT NOT NULL DEFAULT 0,
  author VARCHAR(50), NOT NULL,
);

CREATE TABLE playing (
  id BIGSERIAL PRIMARY KEY,
  movie_id INT,
  price INT,
  start TIMESTAMPTZ,
  end TIMESTAMPTZ,
  CONSTRAINT fk_movie_playing FOREIGN KEY movie_id REFERENCES movie(id)
);

CREATE TABLE viewer (
  id BIGSERIAL PRIMARY KEY,
  playing_id INT,
  name VARCHAR(50),
  CONSTRAINT fk_playing_viewer FOREIGN KEY playing_id REFERENCES playing(id)
)