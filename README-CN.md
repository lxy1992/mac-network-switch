# Mac网络自动切换
- 目前OSX 14.3已经修复这个问题，因此在14.3的系统上不需要安装这个工具，但13.3之前的系统仍旧有这个问题。
- 自动切换以太网和WiFi
    - 由于公司的WiFi环境较差，我转而使用有线网络。然而，一旦打开WiFi，网络延迟迅速增加，这意味着即使连接到有线网络，WiFi始终被优先考虑。即使你按照苹果的指示改变了网络列表中的顺序。为了解决这个问题，我创建了一个命令工具，只需一键就可以自动关闭WiFi并配置网络。
- 这个项目主要基于Albertbori的工作[页面](https://gist.github.com/albertbori/1798d88a93175b9da00b)，
    - 我改变了脚本并添加了一些命令，使其能在OSX 12.1上工作。
    - 最后，将所有这些工作整合成一个单一的命令。

## 功能
- 当以太网连接时，WiFi将自动关闭。
- 当以太网断开时，WiFi将自动打开。

## 如何使用
### 安装
```shell
go install github.com/lxy1992/mac-network-switch@latest
```
### 使用

```shell
sudo mac-network-switch toggleAirport --switch start
sudo mac-network-switch toggleAirport --switch stop
```
