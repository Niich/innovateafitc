docker rm -f InnovateAFITC_postgres

docker run -d -p 5555:5432 -e POSTGRES_HOST_AUTH_METHOD=trust --name InnovateAFITC_postgres postgres:latest
docker cp schema.sql InnovateAFITC_postgres:/schema.sql

Start-Sleep -Seconds 5

docker exec InnovateAFITC_postgres psql -U postgres -f /schema.sql
go generate

# sqlboiler --wipe psql