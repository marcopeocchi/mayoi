default:
	mkdir -p build
	GOARCH=arm64 GOOS=linux CG0_ENABLED=0 go build -o build/mayoi-arm64 cmd/api/main.go
	GOARCH=amd64 GOOS=linux CG0_ENABLED=0 go build -o build/mayoi-amd64 cmd/api/main.go

ui:
	cd cmd/api/ui && bun install && bun run build