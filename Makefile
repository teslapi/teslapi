build:
	cd cmd/teslapid/ && go build -o teslapid .
dev:
	go run cmd/teslapid/teslapid.go
tailwind:
	npx tailwind build styles.css -o assets/css/teslapi.css
