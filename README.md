
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
2. create a python virtual environment using `python -m venv .venv`, And install dependencies using `pip install -r requirements.txt`
3. copy `config_example.json` to `config.json`, and edit it by yourself.
4. go to `client`
5. using `go run client.go` to start client in your mac.
6. you also can using `make start-client` to start client instead step 5.

automatically start client when mac boot to see [Mac README](Mac/README.md).

## Configuration

### server config

- `port`: The port that server will listen on, Required.
- `user_name`: The user name will be shown in the web page, Required.
- `dingtalk_id`: The DingTalk ID of the user, click `快速联系` will jumps to profile page in DingTalk.

### client config

- `server_address`: The URL of the server, Required.
- `expected_ssid`: The expected SSID of the Wi-Fi network, if the client is not connected to this SSID, it will be considered offline. Optional.

## Why this?

`client/screen_lock.py` script can check macos screen status, if users macos is never locked. It means user are working now on mac. If it locked or disconnect network, it means user is offline.

## Thanks

Inspired from [changkun/office](https://github.com/changkun/office)

## License

License: MIT
