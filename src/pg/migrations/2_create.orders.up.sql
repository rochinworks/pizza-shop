CREATE TABLE "orders" (
  id varchar(36),
  style text,
  status text,
  userId varchar(36),
  PRIMARY KEY(id),
  	CONSTRAINT fk_users
  		FOREIGN KEY(userId)
  			REFERENCES users(id)
);
