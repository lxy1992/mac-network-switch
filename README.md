# Mac Network Auto Switch
- Auto Switch the Ethernet and WiFi
- This project is mainly based on Albertbori's work [page](https://gist.github.com/albertbori/1798d88a93175b9da00b), 
  - I change the script and add some commands to make it work on the OSX 12.1.
  - Finally, make all this work into a single command.
- 由于公司WiFi速度差的离谱，于是开始用有线网络，虽然苹果说在配置中调整服务列表的顺序，可以优先使用有线网络。但实际上一旦打开 WiFi ，网络延迟就快速增加，也就是实际上连了有线网络后，也是永远优先使用WiFi。
- 这个问题的解决办法，就是把 WiFi 关了。但是每次连上都要手动关，断线后又要手动开，太过麻烦了。网上查到一些脚本，但大部分都不能直接用，而且在使用上需要配置不少权限。所以我写了个命令工具来实现这个的自动化，一键配置好。

## Function
- When Ethernet connected, WiFi will be closed automatically.
- When Ethernet disconnected, WiFi will be open automatically.

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
