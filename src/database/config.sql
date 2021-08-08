/* for start if not have database */
CREATE DATABASE URLShortener;
CREATE TABLE shortenedUrl (
    ShortenedLink CHAR (4) UNIQUE,
	NotShortenedLink TEXT,
	NumberOfClick SMALLINT,
	ShowNumberOfClickLink CHAR (6) UNIQUE,
	IP TEXT,
    UserID VARCHAR (4),
);
