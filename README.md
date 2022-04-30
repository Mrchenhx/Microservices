> Microservices study notes
>
> create by Mrchenhx 04/30/2022

#  计算机网络

## 1. OSI 七层网络协议

![image-20220430164118267](README.assets/image-20220430164118267.png)

OSI 模型（开放系统互连网络模型）定义了互连的七层框架，即图中显示的结构

TCP/IP 模型借鉴了 OSI 模型。OSI 现有模型再有协议，而 TCP 先有协议再有模型。

OSI 模型是理论上的模型，而 TCP 是广泛使用的标准。

## 2. 经典协议与数据包

### TCP 协议

![image-20220430164902866](README.assets/image-20220430164902866.png)

![image-20220430164929884](README.assets/image-20220430164929884.png)

TCP 段是抓包工具常用分析的层级，即三次握手挥手分析的层级。

### HTTP 协议

![image-20220430165317331](README.assets/image-20220430165317331.png)

###  Websocket 协议

![image-20220430165405556](README.assets/image-20220430165405556.png)

底层是使用 HTTP 协议

![image-20220430165610348](README.assets/image-20220430165610348.png)

### TCP 握手与挥手分析

![image-20220430165730972](README.assets/image-20220430165730972.png)

三次连接示意图：

![image-20220430165740053](README.assets/image-20220430165740053.png)

四次挥手示意图：

![image-20220430165806971](README.assets/image-20220430165806971.png)

![image-20220430170344700](README.assets/image-20220430170344700.png)

![image-20220430170440908](README.assets/image-20220430170440908.png)

样例目录：demo/base/close_wait_test/server

