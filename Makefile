build:
	cd frontend && yarn build
	go build -buildvcs=false -o ./bin/go-vite ./main.go

dev:
	cd frontend && yarn dev & air && fg