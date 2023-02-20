# Mac Network Auto Switch
- Auto Switch the Ethernet and WiFi
- This project is mainly based on Albertbori's work [page](https://gist.github.com/albertbori/1798d88a93175b9da00b), 
  - I change the script and add some commands to make it work on the OSX 12.1.
  - Finally, make all this work into a single command.

## Function
- When Ethernet connected, WiFi will be closed automatically.
- When Ethernet disconnected, WiFi will be open automatically.

## How to Use
### Install
- Download the package of this project
### Use
```shell
go install github.com/lxy1992/mac-network-switch@latest
```

```shell
sudo mac-network-switch toggleAirport --switch start
sudo mac-network-switch toggleAirport --switch stop
```