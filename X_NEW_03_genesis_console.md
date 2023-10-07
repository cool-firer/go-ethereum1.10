写入了genesis block，看怎么启动它。

# 启动console

```shell
cd XXX_DATA
../build/bin/geth --datadir "./chain_data" --nodiscover console 2>>eth_output.log
```

<br />

经过一系列检查，取出配置ChainConfig：

```go
// Get: ethereum-config- + blockHash
storedcfg := rawdb.ReadChainConfig(db, stored)

// EIP150Block等都有指向, 只是是零值。
&ChainConfig {
  ChainID: *math/big.Int {10}, 
  HomesteadBlock: *math/big.Int {nil}, 
  DAOForkBlock: nil, 
  DAOForkSupport: false, 
  EIP150Block: *math/big.Int {nil}, 
  EIP150Hash: common.Hash{}, 
  EIP155Block: *math/big.Int {nil}, 
  EIP158Block: *math/big.Int {nil}, 
  ByzantiumBlock: nil, 
  ConstantinopleBlock: nil, 
  PetersburgBlock: nil, 
  IstanbulBlock: nil, 
  MuirGlacierBlock: nil, 
  BerlinBlock: nil, 
  LondonBlock: nil, 
  ArrowGlacierBlock: nil, 
  GrayGlacierBlock: nil, 
  MergeNetsplitBlock: nil, 
  ShanghaiBlock: nil, 
  CancunBlock: nil, 
  TerminalTotalDifficulty: nil, 
  TerminalTotalDifficultyPassed: false, 
  Ethash: nil, 
  Clique: nil
}
```

<br />

写db versoin：

```go
// rawdb.WriteDatabaseVersion(chainDb, core.BlockChainVersion)
DatabaseVersion	-->	rlp(8)

```

