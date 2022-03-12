package modbusplus

import (
	"encoding/binary"

	"github.com/goburrow/modbus"
)

type ClientWrapper struct {
	handler modbus.ClientHandler
	client  modbus.Client
}

func NewClient(handler modbus.ClientHandler) ClientWrapper {
	client := modbus.NewClient(handler)

	return ClientWrapper{
		handler: handler,
		client:  client,
	}
}
func setSlaveId(handler modbus.ClientHandler, slaveId byte) {
	if h, ok := handler.(*modbus.RTUClientHandler); ok {
		h.SlaveId = slaveId
		return
	}
	if h, ok := handler.(*modbus.TCPClientHandler); ok {
		h.SlaveId = slaveId
		return
	}
	if h, ok := handler.(*modbus.ASCIIClientHandler); ok {
		h.SlaveId = slaveId
		return
	}
}

func (c *ClientWrapper) ReadCoils(address, count uint16, slaveId byte) ([]byte, error) {
	setSlaveId(c.handler, slaveId)
	return c.client.ReadCoils(address, count)
}
func (c *ClientWrapper) ReadDiscreteInputs(address, count uint16, slaveId byte) ([]byte, error) {
	setSlaveId(c.handler, slaveId)
	return c.client.ReadDiscreteInputs(address, count)
}
func (c *ClientWrapper) WriteSingleCoil(address, value uint16, slaveId byte) ([]byte, error) {
	setSlaveId(c.handler, slaveId)
	return c.client.WriteSingleCoil(address, value)
}
func (c *ClientWrapper) WriteMultipleCoils(address uint16, value []uint16, slaveId byte) ([]byte, error) {
	setSlaveId(c.handler, slaveId)
	count := len(value)
	data := make([]byte, 2*count)
	for i, v := range value {
		binary.BigEndian.PutUint16(data[i*2:], v)
	}
	return c.client.WriteMultipleCoils(address, uint16(count), data)
}

func (c *ClientWrapper) ReadInputRegisters(address, count uint16, slaveId byte) ([]byte, error) {
	setSlaveId(c.handler, slaveId)
	return c.client.ReadInputRegisters(address, count)
}
func (c *ClientWrapper) ReadHoldingRegisters(address uint16, count uint16, slaveId byte) ([]byte, error) {
	setSlaveId(c.handler, slaveId)
	return c.client.ReadHoldingRegisters(address, count)
}
func (c *ClientWrapper) WriteSingleRegister(address, value uint16, slaveId byte) ([]byte, error) {
	setSlaveId(c.handler, slaveId)
	return c.client.WriteSingleRegister(address, value)
}
func (c *ClientWrapper) WriteMultipleRegisters(address uint16, value []uint16, slaveId byte) ([]byte, error) {
	setSlaveId(c.handler, slaveId)
	count := len(value)
	data := make([]byte, 2*count)
	for i, v := range value {
		binary.BigEndian.PutUint16(data[i*2:], v)
	}
	return c.client.WriteMultipleRegisters(address, uint16(count), data)
}
func (c *ClientWrapper) ReadWriteMultipleRegisters(readAddress, readCount, writeAddress uint16, value []uint16, slaveId byte) ([]byte, error) {
	setSlaveId(c.handler, slaveId)
	writeCount := len(value)
	data := make([]byte, 2*writeCount)
	for i, v := range value {
		binary.BigEndian.PutUint16(data[i*2:], v)
	}
	return c.client.ReadWriteMultipleRegisters(readAddress, readCount, writeAddress, uint16(writeCount), data)
}
func (c *ClientWrapper) MaskWriteRegister(address, andMask, orMask uint16, slaveId byte) ([]byte, error) {
	setSlaveId(c.handler, slaveId)
	return c.client.MaskWriteRegister(address, andMask, orMask)
}
func (c *ClientWrapper) ReadFIFOQueue(address uint16, slaveId byte) ([]byte, error) {
	setSlaveId(c.handler, slaveId)
	return c.client.ReadFIFOQueue(address)
}
