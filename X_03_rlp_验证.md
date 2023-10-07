



# RLP uint验证

1. 进入xxx/go-ethereum1.10/rlp/ 目录, 新建一个my_encode_test.go文件，内容如下：

   ```go
   package rlp
   
   import (
   	"fmt"
   	"testing"
   	"bytes"
   	"encoding/hex"
   )
   
   func TestEcodeUint(t *testing.T) {
   	b := new(bytes.Buffer)
   	var val interface{}  = uint32(256)
   	Encode(b, val)
   	fmt.Println(hex.Dump(b.Bytes()))
   }
   ```

2. cd到go-ethereum1.10/rlp/，执行 go test -v -run ^TestEcodeUint$：

   ```shell
   ➜  go test -v -run ^TestEcodeUint$
   === RUN   TestEcodeUint
   00000000  82 01 00                                          |...|
   
   --- PASS: TestEcodeUint (0.00s)
   PASS
   ok  	github.com/ethereum/go-ethereum/rlp	0.515s
   ```





# RLP bitInt验证

```go
package rlp

import (
	"fmt"
	"testing"
	"bytes"
	"encoding/hex"
	"math/big"
)

func TestEncodeBigInt(t *testing.T) {
	s, _ := hex.DecodeString("102030405060708090A0B0C0D0E0F2")
	var val interface{}  = new(big.Int).SetBytes(s)

	b := new(bytes.Buffer)
	Encode(b, val)
	fmt.Println(hex.Dump(b.Bytes()))


	b.Reset()
	s, _ = hex.DecodeString("102030405060708090A0B0C0D0E0F2102030405060708090A0B0C0D0E0F2102030405060708090A0B0C0D0E0F2102030405060708090A0B0C0D0E0F2")
	val = new(big.Int).SetBytes(s)
	Encode(b, val)
	fmt.Println(hex.Dump(b.Bytes()))
}
```



# [RLP List](#List)



```go
package rlp_test

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/rlp"
)

func ExampleEncoderBuffer() {
	var w bytes.Buffer

	// Encode [1, 2, 3, [4, 5, 6, 7, 8, 9]]
	buf := rlp.NewEncoderBuffer(&w)
	l1 := buf.List()
	buf.WriteUint64(1)
	buf.WriteUint64(2)
	buf.WriteUint64(3)
	l2 := buf.List()
	buf.WriteUint64(4)
	buf.WriteUint64(5)
	buf.WriteUint64(6)
	buf.WriteUint64(7)
	buf.WriteUint64(8)
	buf.WriteUint64(9)
	buf.ListEnd(l2)
	buf.ListEnd(l1)

	if err := buf.Flush(); err != nil {
		panic(err)
	}
	fmt.Printf("%X\n", w.Bytes())
	// Output:
	// C404C20506
}


```

