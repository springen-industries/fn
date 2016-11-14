## Quick Example for a Go Function (3 minutes)

This example will show you how to test and deploy Go (Golang) code to IronFunctions.

### 1. Prepare the `functions.yaml` file:

At functions.yaml you will find:
```yml
app: goapp
route: /hello
image: USERNAME/hello:0.0.1
build:
- docker run --rm -v "$PWD":/go/src/ -w /go/src/ -e "GOPATH=/go/src/vendor:/go" iron/go:dev go build -o hello
```

The important step here is to ensure you replace `USERNAME` with your Docker Hub account name. Some points of note:
the application name is `goapp` and the route for incoming requests is `/hello`. These informations are relevant for
the moment you try to test this function.

### 2. Build:

```sh
fnctl publish
```

`-v` is optional, but it allows you to see how this function is being built.

### 3. Queue jobs for your function

Now you can start jobs on your function. Let's quickly queue up a job to try it out.

```sh
cat hello.payload.json | fnctl run goapp /hello
```

Here's a curl example to show how easy it is to do in any language:

```sh
curl -H "Content-Type: application/json" -X POST -d @hello.payload.json http://localhost:8080/r/goapp/hello
```