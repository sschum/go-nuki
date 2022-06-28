package command

import "encoding/binary"

func NewRequestReboot(pin uint16, nonce []byte) Command {
	payload := make([]byte, 0, len(nonce)+2)
	payload = append(payload, nonce...)

	pinAsByte := make([]byte, 2)
	binary.LittleEndian.PutUint16(pinAsByte, pin)
	payload = append(payload, pinAsByte...)

	return NewCommand(IdRequestReboot, payload)
}
