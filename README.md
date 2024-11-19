
# Is the user currently working?

Check if the user is currently working on MacOS.

## How to use

server side:

1. go to `server`
2. copy `config_example.json` to `config.json`, and edit it by yourself.
3. using `go run server.go` to start server.
4. you also can using `make start-server` to start server instead step 3.

client side:

1. make sure you have installed go and python
2. create a python virtual environment using `python -m venv .venv`
3. go to `client`
4. using `go run client.go` to start client in your mac.
5. you also can using `make start-client` to start client instead step 4.

## Why this?

`client/screen_lock.py` script can check macos screen status, if users macos is never locked. It means user are working now on mac. If it locked or disconnect network, it means user is offline.

## Thanks

Inspired from [changkun/office](https://github.com/changkun/office)

## License

License: MIT
