package constant

const (
	AutoReplyPrivateTemplate                = "若要发车，请回复“发车”\n若要上车，请回复“上车”\n若要停车，请回复“停车”\n若要接单，请回复“接单”\n\n回复“帮助”获取使用帮助"
	AutoReplyHelpTemplate                   = "使用规则：\n想要代领快递：回复小二“发车”根据引导填写快递点和目的地。发车成功，如果有相同路线的快递小二会替你筛选好，推送对方联系方式给你。\n想要找人代领：回复小二“上车”根据引导填写快递点和目的地，小二会推送给你匹配的发车人联系方式，代领需求及时被解决。"
	AutoReplyDispatchConfirmInfoTemplate    = "发车需要输入目的地、快递点、可运件数、截止时间等信息，继续请回复“是”，取消请回复“否”"
	AutoReplyTakeOnConfirmInfoTemplate      = "上车需要输入目的地、快递点、件数等信息，继续请回复“是”，取消请回复“否”"
	AutoReplyTakeOffConfirmInfoTemplate     = "停车需要输入订单编号，继续请回复“是”，取消请回复“否”"
	AutoReplyAcceptOrderConfirmInfoTemplate = "接单需要输入你要帮助的顾客的订单编号，继续请回复“是”，取消请回复“否”"
	AutoReplySendOrderConfirmInfoTemplate   = "发单需要输入你上车的订单编号，继续请回复“是”，取消请回复“否”"
	AutoReplyDispatchInfoTemplate           = "发车成功，您的车辆编号是%s。如果快递太多车辆提前满员，请回复“停车”，我们将不再给您推送想上车的盆友啦"
	AutoReplyTakeOnSuccessInfoTemplate      = "小二查询到有这些车你可以搭：\n%s\n如果没有上⻋车成功 回复“发单”⼩⼆会将你的上车需求发到群⾥,很快就会有人认领啦."
	AutoReplyTakeOnFailedInfoTemplate       = "小二没有查询到没有合适的车车，你的上车编号是%s，已经将您的上车需求发到群里，请等待有缘人认领。"
	AutoReplySendOrderSuccessTemplate       = "你的上车编号%s发单成功，请耐心等待，有人接单小二会提醒你的哦"
	AutoSendGroupTakeOnMessage              = "有一位顾客需要你们的帮助，信息如下：\n\n编号：%s\n%s\n\n私戳小二，回复“接单”，帮他/她取快递哦"
	AutoReplyAcceptOrderInfo                = "该顾客的信息如下：%s\n请及时联系他/她，帮助他/她取到快递哦"
	AutoSendDispatchMessage                 = "%s号车发啦!\n\n%s\n\n有顺路的(%s及附近的)快递想上车让他/她代领快递的\n私戳小二回复“上车”，\n输入和他/她相同的信息就可以获取他/她的联系方式啦"
	AutoSendOrderAccepted                   = "你的上车订单已被认领，赶快联系他/她吧！\n他/她QQ是：%d\n被人领走后，上车订单会过期，如果没有和对方谈妥，请重新上车"
	ConfirmInfoTemplate                     = "确认请回复“是”完成%s，\n取消请回复“否”终止%s"
	AutoReplyGroupTemplate                  = "请私戳我发车/上车\n(发送任意信息即可开始，小二会给出引导哦)\n" //发车前要加我好友哦，不然小二不会理你的
	QQTestGroup                             = 863841443
	QQGroupHustHelp2                        = 881532638
	QQGroupHustDouble11                     = 691198215
	QQGroupXianZhi1                         = 925099164
	QQGroupXianZhi2                         = 385149574
	CQAtProdString                          = "\\[CQ:at,qq=1509370662\\]"
)

const (
	ChooseRegion = "请选择校区（回复数字）:\n1.韵苑\n2.紫菘\n3.沁苑"

	ChooseZiSongDestination  = "请选择目的地（回复数字，如紫菘5栋请回复5）"
	ChooseYunYuanDestination = "请选择目的地（回复数字，如韵苑5栋请回复5）"
	ChooseQinYuanDestination = "请选择目的地（回复数字，如东5舍请回复5）"

	ChooseZiSongDeliveryPoint  = "请选择紫菘快递点（回复快递点前的数字）：\n1.麦子烘培旁 圆通\n2.宝岛眼镜内 申通\n3.蜜雪冰城旁 申通\n4.顺丰快递点 学子便利店\n5.紫崧印务旁 百世\n6.⻄区服务中心内 菜鸟驿站\n7.百惠园对面⼆楼 韵达\n8.紫崧⽔果大卖场内 中通\n9.⻄区服务中心内 中通\n"
	ChooseYunYuanDeliveryPoint = "请选择韵苑快递点（回复快递点前的数字）：\n1.发之源 中通\n2.韵苑二栋 顺丰\n3.韵苑二栋 妈妈驿站 多为圆通\n4.韵苑11栋 菜⻦驿站 多为申通\n5.韵苑四栋 韵达\n6.东操后侧修车铺 京东\n7.韵体小竹林 天猫/唯品会\n8.生活区明达眼镜店 EMS/邮政\n9.东校区医院对面，东教工17栋 邮政/中通\n10.绝望坡 百世汇通\n11.韵苑操场南侧 韵达快递\n"
	ChooseQinYuanDeliveryPoint = "请选择沁苑快递点（回复快递点前的数字）：\n1.东一背后 菜鸟驿站\n2.东一背后澡堂旁 韵达\n3.东四北侧 中通\n4.东四对面 老东四食堂旁 妈妈驿站\n5.东四北侧 京东\n6.东四二楼明达眼镜 EMS/邮政\n7.保卫处旁 EMS/邮政\n8.博⼠生公寓 EMS/邮政\n"

	ChooseDeliveryNum = "请输入件数（回复数字/任意，如回复5或回复任意）"

	ChooseDeadline = "请选择截止时间\n(即刻起到今⽇日结束时间段内，回复格式如13:30)\n注意是英文的“:”"

	ConfirmInformation = "请确认您的发车信息，提交请回复“是”，取消请回复“否”\n%s"

	ChooseOrder = "请输入订单编号"
)

const (
	RegionZiSong  = "紫菘"
	RegionYunYuan = "韵苑"
	RegionQinYuan = "沁苑"
)
