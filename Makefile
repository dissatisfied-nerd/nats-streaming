db_up:
	psql -U postgres -d nats_streaming -a -f ~/projects/ns-service/migrations/up.sql > /dev/null  

db_down:
	psql -U postgres -d nats_streaming -a -f ~/projects/ns-service/migrations/down.sql  

db_clear:
	psql -U postgres -d nats_streaming -a -f ~/projects/ns-service/migrations/down.sql  
	psql -U postgres -d nats_streaming -a -f ~/projects/ns-service/migrations/up.sql > /dev/null

