package adapter

// USB USB接口 被适配的目标接口
type USB interface {
	USBCharge() string
}

// USBPhone 被适配的目标类
type USBPhone struct {
}

// NewUSB 被适配接口的工厂函数
func NewUSB() USB {
	return &USBPhone{}
}

// USBCharge 被适配的目标类实现的方法
func (u *USBPhone) USBCharge() string {
	return "usb phone charge"
}

// TypeC typ-c接口 适配的目标接口
type TypeC interface {
	TypeCCharge() string
}

// NewTypeCAdapter 适配的目标接口的工厂函数
func NewTypeCAdapter(usb USB) TypeC {
	return &TypeCAdapter{
		USB: usb,
	}
}

// TypeCAdapter TypeC转USB的适配器
type TypeCAdapter struct {
	USB // 组合需要被适配的目标接口
}

// TypeCCharge 实现适配的目标接口的方法
func (t *TypeCAdapter) TypeCCharge() string {
	return t.USBCharge() // 将usb充电适配成type-c充电
}
