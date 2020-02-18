package constant

const (
	UserKeyType          = "type"
	UserKeyState         = "state"
	UserKeyRegion        = "region"
	UserKeyDeliveryPoint = "delivery_point"
	UserKeyDestination   = "destination"
	UserKeyDeliveryNum   = "delivery_num"
	UserKeyDeadline      = "deadline"
	UserKeyConfirmData   = "confirm_data"
)
const (
	OrderKeyIncr = "order_num"
)

const (
	OrderKey1 = iota
	OrderKey2
	OrderKey3
	OrderKey4
	OrderKey5
	OrderKey6
	OrderKey7
	OrderKey8
	OrderKey9
	OrderKey10
	OrderKey11
	OrderKey12
	OrderKey13
	OrderKey14
	OrderKey15
	OrderKey16
)
const (
	TableOrder = "order"
)

/*
	东边：
	韵苑 1、2、3、4栋互相可接
	韵苑 5栋：可接 5、6、7栋
	韵苑 6、7栋 可接：5、6、7、8、9
	韵苑 8、9栋 可接：5、6、7、8、9、10
	韵苑 10、11、栋可接5、6、7、8、9、10、11、
	韵苑 12、24、 27、 26 互相可接
	韵苑 13、14、16、17、20、21 互相可接
	韵苑 11、22、23、25互相可接
	韵苑 15、18、19 、28互相可接
	中间：
	沁苑 东九到东十三舍超级近都可以互相接单。。
	西边：
	紫崧 1、3、4、互相可接
	紫崧  7、8、11、12、13 互相可接
	紫崧  2、5、6、互相可接
	紫崧  10、9、13 互相可接

*/
