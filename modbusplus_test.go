package modbusplus

import (
	"fmt"
	"testing"
	"time"

	"github.com/goburrow/modbus"
)

func TestClient(t *testing.T) {
	// Modbus RTU/ASCII, it is the same as goburrow/modbus
	handler := modbus.NewRTUClientHandler("COM9")
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 1
	handler.Timeout = 5 * time.Second

	err := handler.Connect()
	if err != nil {
		t.Errorf("connect fail:%v", err)
	}
	defer handler.Close()

	//use the modbusplus to create client
	client := NewClient(handler)

	//we can change slave id every call
	for i := 0; i < 100; i++ {
		results, err := client.ReadHoldingRegisters(0x2050, 1, byte(i))
		if err != nil {
			fmt.Printf("%v", results)
		}
	}

}
