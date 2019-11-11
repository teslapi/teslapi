build:
	cd cmd/teslapid/ && go build -o ../../teslapid .
docker-logs:
	docker-compose logs -f teslapi
teslapid:
	go run cmd/teslapid/teslapid.go
teslapweb:
	go run cmd/teslapweb/teslapweb.go
tailwind:
	npx tailwind build styles.css -o assets/css/teslapi.css
