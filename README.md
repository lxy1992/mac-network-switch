# Mac Network Auto Switch
- [中文文档](https://github.com/lxy1992/mac-network-switch/blob/master/README-CN.md)
- Auto Switch the Ethernet and WiFi
  - Due to the company's poor WiFi env, I switched to using wired network. However, once WiFi is turned on, the network latency quickly increases, meaning that even when connected to the wired network, WiFi is always given priority. Even you change the order in the network list like Apple told us. To solve this problem, I created a command tool that can automatically turn off WiFi and configure the network with just one click.
- This project is mainly based on Albertbori's work [page](https://gist.github.com/albertbori/1798d88a93175b9da00b), 
  - I change the script and add some commands to make it work on the OSX 12.1.
  - Finally, make all this work into a single command.

## Function
- When Ethernet connected, WiFi will be closed automatically.
- When Ethernet disconnected, WiFi will be open automatically.

## Support
- OSX 13
- Device
  - USB10
  - AX88179A
  - Thunderbolt Ethernet Slot

## How to Use
### Install
```shell
go install github.com/lxy1992/mac-network-switch@latest
```
### Use

```shell
sudo mac-network-switch toggleAirport --switch start
sudo mac-network-switch toggleAirport --switch stop
```
