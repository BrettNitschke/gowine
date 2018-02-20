DROP DATABASE IF EXISTS wine;
CREATE DATABASE wine;
DROP TABLE IF EXISTS wines;

\c wine;

CREATE TABLE wines (
  id SERIAL PRIMARY KEY,
  winery_name VARCHAR(45) NOT NULL,
  wine_name VARCHAR(45) NOT NULL,
  vintage VARCHAR(45) NOT NULL
);

INSERT INTO wines (winery_name, wine_name, vintage) values ('Hall', 'Diamond Mountain', '2013');
INSERT INTO wines (winery_name, wine_name, vintage) values ('Frank Family', 'Pinot Noir', '2015');
INSERT INTO wines (winery_name, wine_name, vintage) values ('Silver Oak', 'Napa Valley Cabernet', '2013');
