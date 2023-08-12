build:
	cd frontend && yarn build
	ENV=prod go build -buildvcs=false -o ./bin/go-vite ./main.go

dev:
	cd frontend && yarn dev & air && fg