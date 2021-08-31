-- CREATE DATABASE urlshortener;
CREATE TABLE urlshortened (
	ShortenedLink CHAR (4) UNIQUE,
	NotShortenedLink TEXT,
	NumberOfClick SMALLINT,
	ShowNumberOfClickLink CHAR (6) UNIQUE,
	IP TEXT,
    UserID VARCHAR (4),
	CreateTime TIMESTAMPTZ DEFAULT NOW()
)
