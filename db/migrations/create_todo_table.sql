CREATE TABLE IF NOT EXISTS todos (
  TodoId int NOT NULL AUTO_INCREMENT,
  Name char(30),
  Status char(30),
  PRIMARY KEY (TodoId)
);