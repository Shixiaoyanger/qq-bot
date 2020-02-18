package constant

const (
	//无状态
	StateLess = iota
	//等待选择校区
	StateWaitingRegion
	//确认校区
	StateConfirmRegion
	//确认快递点
	StateConfirmDeliveryPoint
	//确认目的地
	StateConfirmDestination
	//确认件数
	StateConfirmDeliveryNum
	//确认截止时间
	StateConfirmDeadline
	//确认最终信息
	StateConfirmInformation

	/****************************/
	//停车
	StateWaitingTakeOff
	StateConfirmTakeOff

	/****************************/
	//接单
	StateWaitingAcceptOrder
	StateConfirmAcceptOrder

	/****************************/
	//发单
	StateWaitingSendOrder
	StateConfirmSendOrder
)

const (
	//顺序很重要，影响业务逻辑
	TypeTakeOn = iota

	TypeDispatch
)

const (
	OrderStateDelete  = -1
	OrderStateTakenOn = -5
	OrderStateTimeOut = -10
	OrderStateStop    = -15
)
