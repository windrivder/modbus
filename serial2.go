package modbus

import (
	"time"

	"go.bug.st/serial"
)

func GetPortsList() ([]string, error) {
	return serial.GetPortsList()
}

// serialPortWrapper2 wraps a serial.Port (i.e. physical port) to
type serialPortWrapper2 struct {
	conf     *serialPortConfig
	port     serial.Port
	deadline time.Time
}

func newSerialPortWrapper2(conf *serialPortConfig) (spw *serialPortWrapper2) {
	spw = &serialPortWrapper2{
		conf: conf,
	}

	return
}

func (spw *serialPortWrapper2) Open() (err error) {
	spw.port, err = serial.Open(spw.conf.Device, &serial.Mode{
		BaudRate: int(spw.conf.Speed),
		DataBits: int(spw.conf.DataBits),
		Parity:   serial.Parity(spw.conf.Parity),
		StopBits: serial.StopBits(spw.conf.StopBits),
	})
	return
}

// Closes the serial port.
func (spw *serialPortWrapper2) Close() (err error) {
	err = spw.port.Close()

	return
}

// Reads bytes from the underlying serial port.
func (spw *serialPortWrapper2) Read(rxbuf []byte) (cnt int, err error) {
	// return a timeout error if the deadline has passed
	if time.Now().After(spw.deadline) {
		err = ErrRequestTimedOut
		return
	}

	cnt, err = spw.port.Read(rxbuf)
	// mask serial.ErrTimeout errors from the serial port
	if err != nil && cnt == 0 {
		err = nil
	}

	return
}

// Sends the bytes over the wire.
func (spw *serialPortWrapper2) Write(txbuf []byte) (cnt int, err error) {
	cnt, err = spw.port.Write(txbuf)

	return
}

// Saves the i/o deadline (only used by Read).
func (spw *serialPortWrapper2) SetDeadline(deadline time.Time) (err error) {
	spw.deadline = deadline

	return
}
