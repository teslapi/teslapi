build:
	cd cmd/teslapid/ && go build -o teslapid .
dev: tailwind
	go run cmd/teslapid/teslapid.go
tailwind:
	npx tailwind build styles.css -o assets/css/teslapi.css
test:
	go test cmd/teslapid/teslapid.go
api:
	go run cmd/teslapid/api/api.go
