CREATE USER 'sammy'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';


CREATE USER 'test_user'@'localhost' IDENTIFIED WITH mysql_native_password BY 'test_password';
GRANT ALL PRIVILEGES ON test_database.* TO 'test_user'@'localhost';
FLUSH PRIVILEGES;