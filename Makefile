db_up:
	psql -U postgres -d nats_server -a -f ~/projects/wb-practice/migrations/up.sql > /dev/null  

db_down:
	psql -U postgres -d nats_server -a -f ~/projects/wb-practice/migrations/down.sql  

