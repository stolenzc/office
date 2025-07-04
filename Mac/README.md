# Office Running On MacOS

## Automatically Start Client When Boot

1. copy and rename `com.stolenzc.office.client.plist.example` to `com.stolenzc.office.client.plist`
2. edit `com.stolenzc.office.client.plist` file
   - line 10 to set the correct path to your client binary
   - line 26 to set the correct path to this project root directory
   - line 20 and line 23 to set the correct path of log files
3. move `com.stolenzc.office.client.plist` to `~/Library/LaunchAgents/`
4. `Settings > General > Login Items` to allow `client` running in background

**Warning**

- Don't forget to clean log files circularly, or it will take up a lot of disk space.
- If you want to stop the client, start `Activity Monitor`, find `client` process, and click `X` to stop it.