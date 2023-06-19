package migrations

const (
	CREATE_USER_MODEL = `CREATE TABLE IF NOT EXISTS users ( 
	id bigint NOT NULL PRIMARY KEY AUTO_INCREMENT,
	name varchar(255),
	phone_number varchar(255), 
	email varchar(255), 
	hash varchar(1000),
	salt varchar(20),
	role varchar(20),
	country_code varchar(3) DEFAULT 'IN',
	email_verified BOOLEAN DEFAULT false,
	phone_number_verified BOOLEAN DEFAULT false,
	active_account BOOLEAN DEFAULT false,
	created_at bigint,
	updated_at bigint)`
)
