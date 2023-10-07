# 再来几个console例子加深下印象

## clientVersion

| namespace | 所属      | name          | 对应方法               | receiver/路径           |
| --------- | --------- | ------------- | ---------------------- | ----------------------- |
| web3      | callbacks | clientVersion | ClientVersion() string | &web3API{n} node/api.go |
|           |           |               |                        |                         |

clientVersion作为web3实例的一个prop

```javascript
new Property({
  name: 'version.node',
  getter: 'web3_clientVersion'
})

// Property定义
var Property = function (options) {
    this.name = options.name; // 'version.node'
    this.getter = options.getter; // 'web3_clientVersion'
    this.setter = options.setter;
    this.outputFormatter = options.outputFormatter;
    this.inputFormatter = options.inputFormatter;
    this.requestManager = null;
};

function Web3 (provider) {
    this._requestManager = new RequestManager(provider);
  	...
    this._extend({
        properties: properties()
    });
}

// 最终形成这样的:
<< web3实例 >>
  {
  	'version': {
      'node': {
        get: GetFunc,
        enumerable: true,
      },
      ‘getNode': GetFuncAsync,
    },
  }

// 所以, 在console应该这样: web3.version.node 或者 web3.getNode()
```

<br />

\> web3.version.node

\> web3.version.getNode( function(err, res) { console.log(res); })

```javascript
键入时, 调用send:
requestManager.send({
	method: property.getter, // 'web3_clientVersion'
}));


RequestManager.prototype.send = function (data) {
/** payload:
	{
        jsonrpc: '2.0',
        id: 自增,
        method: 'web3_clientVersion',
        params: []
	}
*/
    var payload = Jsonrpc.toPayload(data.method, data.params);
    // 调用go的代码
    var result = this.provider.send(payload);
    return result.result;
};

// 到console/console.go# bridge.Send
c.jsre.Do(func(vm *goja.Runtime) {
	transport := vm.NewObject()
	transport.Set("send", jsre.MakeCallback(vm, bridge.Send))
	transport.Set("sendAsync", jsre.MakeCallback(vm, bridge.Send))
	vm.Set("_consoleWeb3Transport", transport)
	_, err = vm.RunString("var web3 = new Web3(_consoleWeb3Transport)")
})

// 正式进入go区域代码
```

<br />

receiver比较简单

```go
&web3API{
	stack *Node
}

// 看来是返回p2p server实例 "geth/v版本信息/os系统/go版本"
func (s *web3API) ClientVersion() string {
	return s.stack.Server().Name
}
```

<br />

## miner.start()

开始挖矿，挖矿奖励的币会默认保存到第一个创建的账户中。

| namespace | 所属      | name  | 对应方法                  | receiver/路径/时机                       |
| --------- | --------- | ----- | ------------------------- | ---------------------------------------- |
| miner     | callbacks | start | Start(threads *int) error | &MinerAPI{e} eth/api.go # eth/backend.go |
|           |           |       |                           |                                          |

\> miner.start()

```javascript
// miner没有在web3.js里定义.
// 而是在 go-ethereum1.10/internal/web3ext/web3ext.go 扩展.

const MinerJs = `
web3._extend({
	property: 'miner',
	methods: [
		new web3._extend.Method({
			name: 'start',
			call: 'miner_start',
			params: 1,
			inputFormatter: [null]
		}),
	],
	properties: []
});
`

// 在web3实例中会形成:
<< web3实例 >>
{
  'miner': {
    'start': SendFunc
  },
}

// method.toPayload(), 产生payload
{
  method: 'miner_start',
  params: [],
  callback: undef,
}

// method.requestManager.send(payload), 发送
// 其他的一样了
```

<br />

Go区域代码

涉及的类有：ethereum、ethereum.engine、ethereum.txPool、ethereum.gasPrice、

ethereum.accountManager、ethereum.handler.acceptTxs、ethereum.miner。

[主网ChainConfig](./X_主网ChainConfig.md)

receiver比较简单:

```go
&MinerAPI{ e }
```

<br />





## miner.stop()

停止挖矿