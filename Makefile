build:
	cd cmd/teslapid/ && go build -o ../../teslapid .
docker-logs:
	docker-compose logs -f teslapi
uploader:
	go run cmd/teslapid/teslapid.go
teslapid:
	go run cmd/teslapid/teslapid.go
tailwind:
	npx tailwind build styles.css -o assets/css/teslapi.css
