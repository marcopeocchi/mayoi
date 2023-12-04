default:
	mkdir -p build
	GOARCH=arm64 GOOS=linux CG0_ENABLED=0 go build -o build/mayoi-arm64 cmd/web/main.go
	GOARCH=amd64 GOOS=linux CG0_ENABLED=0 go build -o build/mayoi-amd64 cmd/web/main.go

ui:
	cd cmd/web/ui && bun install && bun run build

docker:
	docker buildx build --push -t marcobaobao/mayoi --platform linux/amd64,linux/arm64 .