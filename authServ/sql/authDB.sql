CREATE TABLE users (

	username TEXT,
	passwd TEXT

);

SELECT * FROM users;

CREATE OR REPLACE FUNCTION checkData(uName TEXT, pswd TEXT)
RETURNS INT AS $$
DECLARE
	status INT;
BEGIN
	IF (SELECT COUNT(*) FROM users WHERE username = uName AND passwd = pswd) = 1 THEN
		status = 0;
	ELSE
		status = -1;
	END IF;
	RETURN status;
END; 
$$ LANGUAGE plpgsql

SELECT * FROM checkData('admin', 'admin');


