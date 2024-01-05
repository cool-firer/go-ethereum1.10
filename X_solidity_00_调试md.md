1、找到internal/ethapi/api.go，找到ChainId()方法（参照X_06_inprocHandler），打上断点：

![debug_01](img/debug_01.png)

<br />

2、正常开启debug console，会进入调试控制台：

![](img/debug_02.png)

<br />

3、再打开一个终端，用ipc连：

![](img/debug_03.png)

<br />

4、再输入web3.eth.chainId()就会调到断点处。很方便。

