# Go Vite React - Example

Embed your vite application into golang, with hot module reloading and live reload for css / tsx changes!

I recently tried to embed a vite react-ts application into a golang binary to have a single binary fullstack application. I stumbled across [this great video](https://www.youtube.com/watch?v=w_Rv_3-FF0g) which helped me to get it working. But I was not 100% satisified. One of the most valuable things about using vite, is the dev server and it's hot reloading capability. I really like that if I make a change to context in tsx files, or css files the changes are instantly available on in browser and my state does not automatically refresh.

## Project Setup

We start with a basic go application usng the Echo framework where we a single APi endpoint to return some text. We also have a vite application created using `yarn create vite` in the frontend folder.

The frontend calls the API endpoint to return the text.

We have a [frontend.go](frontend/frontend.go) file which uses embed and echo to serve the content from the dist folder.

## Run in development mode

Start using `make dev`, will run vite build in watch mode, as well as the golang server using [air](https://github.com/cosmtrek/air)

### Build a binary

Using `make build`, will build frontend assets and then compile the golang binary embeding the frontend/dist folder

### Build a dockerfile

Using `docker build -t go-vite .`, will use a multistage dockerfile to build the vite frontend code using a node image, then the golang using a golang image, the put the single binary into an alpine image.
