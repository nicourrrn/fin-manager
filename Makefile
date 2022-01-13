
db_tables:
	mysqldump -u $(DB_USER) -p -d fin_manager > config/mysql/tables.sql
db_migration:
	mysqldump -u $(DB_USER) -p fin_manager > config/mysql/tables.sql

