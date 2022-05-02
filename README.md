> Microservices study notes
>
> create by Mrchenhx 04/30/2022

#  计算机网络

## 1. OSI 七层网络协议

![image-20220430180002645](README.assets/image-20220430180002645.png)

OSI 模型（开放系统互连网络模型）定义了互连的七层框架，即图中显示的结构

TCP/IP 模型借鉴了 OSI 模型。OSI 现有模型再有协议，而 TCP 先有协议再有模型。

OSI 模型是理论上的模型，而 TCP 是广泛使用的标准。

## 2. 经典协议与数据包

### TCP 协议

![image-20220430180025331](README.assets/image-20220430180025331.png)

![image-20220430180051506](README.assets/image-20220430180051506.png)

TCP 段是抓包工具常用分析的层级，即三次握手挥手分析的层级。

### HTTP 协议

![image-20220430180110505](README.assets/image-20220430180110505.png)

###  Websocket 协议

![image-20220430180123002](README.assets/image-20220430180123002.png)

底层是使用 HTTP 协议

![image-20220430180134372](README.assets/image-20220430180134372.png)

### TCP 篇

#### TCP 握手与挥手分析

![image-20220430180154689](README.assets/image-20220430180154689.png)

三次连接示意图：

![image-20220430180203485](README.assets/image-20220430180203485.png)

四次挥手示意图：

![image-20220430180211899](README.assets/image-20220430180211899.png)

为什么 time_wait 需要等待 2MSL？

![image-20220430180238858](README.assets/image-20220430180238858.png)

为什么会出现大量 close_wait ？

![image-20220430180321007](README.assets/image-20220430180321007.png)

样例关键代码：

```go

```

#### TCP - 流量控制

![image-20220430185538436](README.assets/image-20220430185538436.png)

简易示意图：

![image-20220430185557081](README.assets/image-20220430185557081.png)

#### TCP - 拥塞控制

![image-20220430185838578](README.assets/image-20220430185838578.png)

拥塞控制分为：慢开始和拥塞避免、快速重传和快速恢复

##### 慢开始和拥塞避免

![image-20220430190042458](README.assets/image-20220430190042458.png)

![image-20220430190121070](README.assets/image-20220430190121070.png)

##### 快重传和快恢复

![image-20220430190317974](README.assets/image-20220430190317974.png)

##### 粘包/拆包

![image-20220430190507358](README.assets/image-20220430190507358.png)

#### 出现的原因：

![image-20220430190711407](README.assets/image-20220430190711407.png)

##### 解决方案

![image-20220430191829302](README.assets/image-20220430191829302.png)

## 3. 网络代理

### 什么是网络代理

![image-20220501213304710](README.assets/image-20220501213304710.png)

### 网络代理与网络转发的区别

#### 网络转发

![image-20220501213550158](README.assets/image-20220501213550158.png)

#### 网络代理

![image-20220501213755337](README.assets/image-20220501213755337.png)

#### 实际区别

![image-20220501213842731](README.assets/image-20220501213842731.png)

### 网络代理类型

![image-20220501214056078](README.assets/image-20220501214056078.png)

### 正向代理

![image-20220501214217596](README.assets/image-20220501214217596.png)

#### 实现 Web 浏览器代理

![image-20220501214502817](README.assets/image-20220501214502817.png)

![image-20220501214626633](README.assets/image-20220501214626633.png)









### 反向代理