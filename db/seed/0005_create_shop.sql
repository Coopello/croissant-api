DROP TABLE IF EXISTS shops;

CREATE TABLE IF NOT EXISTS shops (
  ID              BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
	Name            VARCHAR(40) NOT NULL,
	Url             VARCHAR(40) NOT NULL
);
