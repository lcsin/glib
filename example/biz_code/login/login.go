package login

const (
	Passwd    = 1
	PhoneCode = 2
	QRCode    = 3
)

// ILogin 登录接口
type ILogin interface {
	Login() error
}

// ILoginBuilder builder
type ILoginBuilder struct {
	email       string
	inputPasswd string

	phone     string
	inputCode string

	qRCode string
}

// NewILoginBuilder new builder
func NewILoginBuilder() *ILoginBuilder {
	return &ILoginBuilder{}
}

// WithEmail email
func (b *ILoginBuilder) WithEmail(email string) *ILoginBuilder {
	b.email = email
	return b
}

// WithInputPasswd inputPasswd
func (b *ILoginBuilder) WithInputPasswd(inputPasswd string) *ILoginBuilder {
	b.inputPasswd = inputPasswd
	return b
}

// WithPhone phone
func (b *ILoginBuilder) WithPhone(phone string) *ILoginBuilder {
	b.phone = phone
	return b
}

// WithInputCode inputCode
func (b *ILoginBuilder) WithInputCode(inputCode string) *ILoginBuilder {
	b.inputCode = inputCode
	return b
}

// Build interface build
func (b *ILoginBuilder) Build(loginType int64) ILogin {
	switch loginType {
	case Passwd: // 账号密码登录
		return &PasswdLogin{
			Email:       b.email,
			InputPasswd: b.inputPasswd,
		}
	case PhoneCode: // 手机号登录
		return &PhoneLogin{
			Phone:     b.phone,
			InputCode: b.inputCode,
		}
	case QRCode: // 扫码登录
		return &QRCodeLogin{QRCode: b.qRCode}
	}
	return nil
}
