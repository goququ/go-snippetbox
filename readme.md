dev run command

`npx nodemon -e go,tmpl --exec go run "cmd/web/*.go" --signal SIGTERM`

or

`go run cmd/web/*.go`

**DEV:**

Create snippet

`curl -iL -X POST https://go-sbox.herokuapp.com/snippet/create`
