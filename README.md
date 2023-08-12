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

## Hot Reloading

Instead of serving static assets from go when we running in dev mode, we will setup a proxy from echo that will route the requests to a running vite dev server, unless the path is prefixed with `/api`, this will allow for the HMR and live reloading to happen just as if you were running `vite dev` but it will also allow for api paths to be served.

All the changes required to take the initial project and support hot module reloading can be found in the [pull request](https://github.com/danhawkins/go-vite-react-example/pull/1)

### Step 1

Change the [package.json](frontend/package.json) to run the standard `vite` instead of the `tsc && vite build --watch`

## Step 2

Update the [frontend.go](frontend/frontend.go) so that when we are running in dev mode we proxy requests to the vite dev server

Import a `.env` file using dotenv which just has `ENV=dev` inside

```golang
import(
_ "github.com/joho/godotenv/autoload"
)
```

If we are running in dev mode, setup the dev proxy

```golang
func RegisterHandlers(e *echo.Echo) {
  if os.Getenv("ENV") == "dev" {
    log.Println("Running in dev mode")
    setupDevProxy(e)
    return
  }
  // Use the static assets from the dist directory
  e.FileFS("/", "index.html", distIndexHTML)
  e.StaticFS("/", distDirFS)
}

func setupDevProxy(e *echo.Echo) {
  url, err := url.Parse("http://localhost:5173")
  if err != nil {
    log.Fatal(err)
  }
  // Setep a proxy to the vite dev server on localhost:5173
  balancer := middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
    {
      URL: url,
    },
  })
  e.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
    Balancer: balancer,
    Skipper: func(c echo.Context) bool {
      // Skip the proxy if the prefix is /api
      return len(c.Path()) >= 4 && c.Path()[:4] == "/api"
    },
  }))
}
```

## Step 3

Run `make dev` to start the vite dev server and air for the golang server, changes in the frontend app will now be reflected immediatly.

**IMPORTANT: The go build will faile if frontend/dist/index.html is not available, so even if you are running in dev mode, make sure to run `make build` initially to populate the folder**
