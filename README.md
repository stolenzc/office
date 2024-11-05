
# Is the user currently working?

Check if the user is currently working

## How to use

server side:

1. change directory to `server`
2. copy `config_example.json` to `config.json`, and edit it by yourself.
3. using `go run server.go` to start server.

client side:

1. change directory to `client`
2. using `go run client.go` to start client in your mac.

## Why this?

Client is a local script running in your working device. It will send a request to server every 5 second. If you leave, make your device sleep and disconnect from the network. so update status can't sending successfully.

## Thanks

Some Idea from [changkun/office](https://github.com/changkun/office)

## License

License: MIT
