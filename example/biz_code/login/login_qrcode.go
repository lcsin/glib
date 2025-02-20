package login

import "fmt"

type QRCodeLogin struct {
	QRCode string
}

func (q *QRCodeLogin) Login() error {
	fmt.Println("qrcode: ", q.QRCode)
	return nil
}
