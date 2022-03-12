# modbus-plus
it seems that [goburrow/modbus](<https://github.com/goburrow/modbus>) is no longger maintained any more.So this repository is the supplement of goburrow/modbus

# What Different?

1. we can change the slaveId on every call

   for example:

   goburrow/modbus:

   ```go
   ReadHoldingRegisters(address, quantity uint16) (results []byte, err error)
   ```

   lishuangquan1987/modbusplus:

   ```go
   ReadHoldingRegisters(address uint16, count uint16, slaveId byte) ([]byte, error)
   ```

   

2. Auto convert `[]byte` to `uint16`,and the `quantity` is the length of `value`

   for example:

   goburrow/modbus:

   ```go
   WriteMultipleRegisters(address, quantity uint16, value []byte) (results []byte, err error)
   ```

   lishuangquan1987/modbusplus:

   ```go
   WriteMultipleRegisters(address uint16, value []uint16, slaveId byte) ([]byte, error)
   ```

   

# Usage



```go
import (
	"fmt"
	"time"

	"github.com/goburrow/modbus"
	"github.com/lishuangquan1987/modbusplus"
)

// Modbus RTUClientHandler/ASCIIClientHandler/TCPClientHandler, it is the same as goburrow/modbus
handler := modbus.NewRTUClientHandler("COM9")
handler.BaudRate = 9600
handler.DataBits = 8
handler.Parity = "N"
handler.StopBits = 1
handler.SlaveId = 1
handler.Timeout = 5 * time.Second

err := handler.Connect()
if err != nil {
    fmt.Printf("connect fail:%v", err)
}
defer handler.Close()

//use the modbusplus to create client
client := modbusplus.NewClient(handler)

//we can change slave id every call
for i := 0; i < 100; i++ {
    results, err := client.ReadHoldingRegisters(0x2050, 1, byte(i))
    if err != nil {
        fmt.Printf("%v", results)
    }
}
```

