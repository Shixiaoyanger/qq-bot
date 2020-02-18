package bot

import (
	"encoding/json"
	"errors"
	"fmt"
	"qq-bot/constant"
	"qq-bot/model"
	"qq-bot/util"
	"regexp"
	"strconv"
)

var reTime = regexp.MustCompile("([01][0-9]|2[0-3]):[0-5][0-9]")
var reAt = regexp.MustCompile(constant.CQAtProdString)

type StateFunc func(b *BotAPI, u *model.User) error

var stateEvent = map[int64]StateFunc{
	constant.StateLess:                 Less,
	constant.StateWaitingRegion:        WaitingRegion,
	constant.StateConfirmRegion:        ConfirmRegion,
	constant.StateConfirmDeliveryPoint: ConfirmDeliveryPoint,
	constant.StateConfirmDestination:   ConfirmDestination,
	constant.StateConfirmDeliveryNum:   ConfirmDeliveryNum,
	constant.StateConfirmDeadline:      ConfirmDeadline,
	constant.StateConfirmInformation:   ConfirmInformation,

	constant.StateWaitingTakeOff: WaitingTakeOff,
	constant.StateConfirmTakeOff: ConfirmTakeOff,

	constant.StateWaitingAcceptOrder: WaitingAcceptOrder,
	constant.StateConfirmAcceptOrder: ConfirmAcceptOrder,

	constant.StateWaitingSendOrder: WaitingSendOrder,
	constant.StateConfirmSendOrder: ConfirmSendOrder,
}

var regionList = map[string]string{
	"1": constant.RegionYunYuan,
	"2": constant.RegionZiSong,
	"3": constant.RegionQinYuan,
}
var yunYuanDeliveryPointList = map[string]string{
	"1":  "发之源 中通",
	"2":  "韵苑二栋 顺丰",
	"3":  "韵苑二栋 妈妈驿站 多为圆通",
	"4":  "韵苑11栋 菜⻦驿站 多为申通",
	"5":  "韵苑四栋 韵达",
	"6":  "东操后侧修车铺 京东",
	"7":  "韵体小竹林 天猫/唯品会",
	"8":  "生活区明达眼镜店 EMS/邮政",
	"9":  "东校区医院对面，东教工17栋 邮政/中通",
	"10": "绝望坡 百世汇通",
	"11": "韵苑操场南侧 韵达快递",
}
var ziSongDeliveryPointList = map[string]string{
	"1": "麦子烘培旁 圆通",
	"2": "宝岛眼镜内 申通",
	"3": "蜜雪冰城旁 申通",
	"4": "顺丰快递点 学子便利店",
	"5": "紫崧印务旁 百世",
	"6": "⻄区服务中心内 菜鸟驿站",
	"7": "百惠园对面⼆楼 韵达",
	"8": "紫崧⽔果大卖场内 中通",
	"9": "⻄区服务中心内 中通",
}
var qinYuanDeliveryPointList = map[string]string{
	"1": "东一背后 菜鸟驿站",
	"2": "东一背后澡堂旁 韵达",
	"3": "东四北侧 中通",
	"4": "东四对面 老东四食堂旁 妈妈驿站",
	"5": "东四北侧 京东",
	"6": "东四二楼明达眼镜 EMS/邮政",
	"7": "保卫处旁 EMS/邮政",
	"8": "博⼠生公寓 EMS/邮政",
}

var regionToDelivery = map[string]map[string]string{
	constant.RegionYunYuan: yunYuanDeliveryPointList,
	constant.RegionZiSong:  ziSongDeliveryPointList,
	constant.RegionQinYuan: qinYuanDeliveryPointList,
}
var regionToDestination = map[string]string{
	constant.RegionYunYuan: constant.ChooseYunYuanDestination,
	constant.RegionZiSong:  constant.ChooseZiSongDestination,
	constant.RegionQinYuan: constant.ChooseQinYuanDestination,
}

func ReceivePrivateMsg(b *BotAPI) {
	userID := b.Update.UserID
	user, err := model.GetUser(userID)
	if err != nil || user == nil {
		b.ReplyMsg("系统错误，请稍后再试")
	}
	if b.Update.Message.Text == "重置" {
		err := user.SetState(constant.StateLess)
		if err != nil {
			b.ReplyMsg("系统错误，请稍后再试")
		}
		b.ReplyMsg("状态重置成功")
		return
	}
	state, err := user.GetState()
	if err != nil {
		b.ReplyMsg("系统错误，请稍后再试")
	}

	if f, ok := stateEvent[state]; ok {
		err := f(b, user)
		if err != nil {
			b.ReplyMsg("系统错误，请稍后再试 stateEvent[state]")
			fmt.Println("stateEvent[state]", err)
		}
	}

}
func ReceiveGroupMsg(b *BotAPI) {

	if ok := CheckAt(b.Update.Message.Message.CQString()); ok {
		userID := b.Update.Message.From.ID

		b.ReplyMsg(fmt.Sprintf("[CQ:at,qq=%d] \n%s", userID, constant.AutoReplyGroupTemplate))
	}
}

func Less(b *BotAPI, u *model.User) error {
	switch b.Update.Message.Text {
	case "发车":
		b.ReplyMsg(constant.AutoReplyDispatchConfirmInfoTemplate)
		u.SetType(constant.TypeDispatch)
		return u.SetState(constant.StateWaitingRegion)
	case "上车":
		b.ReplyMsg(constant.AutoReplyTakeOnConfirmInfoTemplate)
		u.SetType(constant.TypeTakeOn)
		return u.SetState(constant.StateWaitingRegion)
	case "停车":
		b.ReplyMsg(constant.AutoReplyTakeOffConfirmInfoTemplate)
		return u.SetState(constant.StateWaitingTakeOff)
	case "接单":
		b.ReplyMsg(constant.AutoReplyAcceptOrderConfirmInfoTemplate)
		return u.SetState(constant.StateWaitingAcceptOrder)
	case "发单":
		b.ReplyMsg(constant.AutoReplySendOrderConfirmInfoTemplate)
		return u.SetState(constant.StateWaitingSendOrder)
	case "帮助":
		b.ReplyMsg(constant.AutoReplyHelpTemplate)
	default:
		b.ReplyMsg(constant.AutoReplyPrivateTemplate)
	}

	return nil
}

//等待选择校区
func WaitingRegion(b *BotAPI, u *model.User) error {
	switch b.Update.Message.Text {
	case "是":
		b.ReplyMsg(constant.ChooseRegion)
		return u.SetState(constant.StateConfirmRegion)
	case "否":
		b.ReplyMsg("服务已结束")
		return u.SetState(constant.StateLess)
	default:
		b.ReplyMsg(constant.ErrorInput)
		return nil
	}
}

//确认校区
func ConfirmRegion(b *BotAPI, u *model.User) error {
	msg := b.Update.Message.Text

	region, ok := regionList[msg]
	if !ok {
		b.ReplyMsg(constant.ErrorInput)
		return nil
	}

	err := u.SetRegion(region)
	if err != nil {
		return err
	}

	switch msg {
	case "1":
		b.ReplyMsg(constant.ChooseYunYuanDeliveryPoint)
	case "2":
		b.ReplyMsg(constant.ChooseZiSongDeliveryPoint)
	case "3":
		b.ReplyMsg(constant.ChooseQinYuanDeliveryPoint)
	default:
		b.ReplyMsg(constant.ErrorInput)
	}
	return u.SetState(constant.StateConfirmDeliveryPoint)
}

//确认快递点
func ConfirmDeliveryPoint(b *BotAPI, u *model.User) error {
	msg := b.Update.Message.Text

	region, err := u.GetRegion()
	if err != nil {
		return err
	}

	deliveryPoint, ok := regionToDelivery[region][msg]
	if !ok {

		b.ReplyMsg(constant.ErrorInput)
		return nil
	}

	chooseDestination, ok := regionToDestination[region]
	if !ok {
		return errors.New("系统错误")
	}

	b.ReplyMsg(chooseDestination)

	err = u.SetDeliveryPoint(deliveryPoint)
	if err != nil {

		//TODO
		return err
	}
	return u.SetState(constant.StateConfirmDestination)

}

//确认目的地 如韵苑5栋
func ConfirmDestination(b *BotAPI, u *model.User) error {
	msg := b.Update.Message.Text
	region, err := u.GetRegion()
	if err != nil {
		return errors.New("region not exist")
	}

	var destination string
	if msg != "任意" {
		num, err := StrToNum(msg)
		if err != nil {
			b.ReplyMsg(constant.ErrorInput)
			return nil
		}
		switch region {
		case constant.RegionYunYuan:
			destination = fmt.Sprintf("韵苑%d栋", num)
		case constant.RegionZiSong:
			destination = fmt.Sprintf("紫菘%d栋", num)
		case constant.RegionQinYuan:
			destination = fmt.Sprintf("东%d舍", num)
		default:
			return errors.New("region not exist")
		}
	} else {
		switch region {
		case constant.RegionYunYuan:
			destination = "韵苑任意栋"
		case constant.RegionZiSong:
			destination = "紫菘任意栋"
		case constant.RegionQinYuan:
			destination = "东任意舍"
		default:
			return errors.New("region not exist")
		}
	}

	b.ReplyMsg(constant.ChooseDeliveryNum)
	err = u.SetDestination(destination)
	if err != nil {
		return err
	}
	return u.SetState(constant.StateConfirmDeliveryNum)

}

//确认件数
func ConfirmDeliveryNum(b *BotAPI, u *model.User) error {
	msg := b.Update.Message.Text
	if msg != "任意" {
		_, err := StrToNum(msg)
		if err != nil {
			b.ReplyMsg(constant.ErrorInput)
			return nil
		}
	}

	err := u.SetDeliverNum(msg)
	if err != nil {
		return err
	}

	if typ, err := u.GetType(); typ == constant.TypeTakeOn {
		if err != nil {
			return err
		}
		confirmData, err := GetConfirmData(u)
		if err != nil {
			return err
		}
		err = u.SetConfirmData(confirmData)
		if err != nil {
			return err
		}

		b.ReplyMsg(fmt.Sprintf("请确认您的上车信息\n%s", confirmData))
		b.ReplyMsg(fmt.Sprintf(constant.ConfirmInfoTemplate, "上车", "上车"))
		return u.SetState(constant.StateConfirmInformation)
	}

	b.ReplyMsg(constant.ChooseDeadline)

	return u.SetState(constant.StateConfirmDeadline)
}

//确认截止时间
func ConfirmDeadline(b *BotAPI, u *model.User) error {
	msg := b.Update.Message.Text
	if !CheckTime(msg) {
		b.ReplyMsg("时间已过或" + constant.ErrorInput)
		return nil
	}

	err := u.SetDeadline(msg)
	if err != nil {
		return err
	}

	confirmData, err := GetConfirmData(u)
	if err != nil {
		return err
	}
	err = u.SetConfirmData(confirmData)
	if err != nil {
		return err
	}

	b.ReplyMsg(fmt.Sprintf("请确认您的信息\n%s", confirmData))
	b.ReplyMsg(fmt.Sprintf(constant.ConfirmInfoTemplate, "发车", "发车"))

	return u.SetState(constant.StateConfirmInformation)
}

//确认最终信息
func ConfirmInformation(b *BotAPI, u *model.User) error {
	msg := b.Update.Message.Text
	typ, err := u.GetType()
	if err != nil {
		return err
	}

	event := "发车"
	if typ == constant.TypeTakeOn {
		event = "上车"
	}

	switch msg {
	case "是":
		return commitInfo(b, u, typ)
	case "否":
		b.ReplyMsg("已终止" + event)
	default:
		b.ReplyMsg(constant.ErrorInput)
		return nil
	}
	return u.SetState(constant.StateLess)
}

/****************************/
//接单
func WaitingAcceptOrder(b *BotAPI, u *model.User) error {
	switch b.Update.Message.Text {
	case "是":
		b.ReplyMsg(constant.ChooseOrder)
		return u.SetState(constant.StateConfirmAcceptOrder)
	case "否":
		b.ReplyMsg("服务已结束")
		return u.SetState(constant.StateLess)
	default:
		b.ReplyMsg(constant.ErrorInput)
		return nil
	}
}
func ConfirmAcceptOrder(b *BotAPI, u *model.User) error {
	msg := b.Update.Message.Text
	num, err := StrToNum(msg)
	if err != nil {
		b.ReplyMsg(constant.ErrorInput)
		return nil
	}

	order, err := model.SearchTakeOnOrder(num)
	if err != nil || order.ConfirmData == "" {
		b.ReplyMsg(constant.ErrorInputTakeOnNum)
		return u.SetState(constant.StateLess)
	}
	model.DeleteOrder(num)
	info, id := fmt.Sprintf("\n上车人QQ：%d\n上车人信息：%s\n\n", order.UserID, order.ConfirmData), order.UserID

	b.ReplyMsg(fmt.Sprintf(constant.AutoReplyAcceptOrderInfo, info))
	b.SendMsgToPrivate(id, fmt.Sprintf(constant.AutoSendOrderAccepted, u.GetUserId()))
	return u.SetState(constant.StateLess)
}

/****************************/
//停车
func WaitingTakeOff(b *BotAPI, u *model.User) error {
	switch b.Update.Message.Text {
	case "是":
		b.ReplyMsg(constant.ChooseOrder)
		return u.SetState(constant.StateConfirmTakeOff)
	case "否":
		b.ReplyMsg("服务已结束")
		return u.SetState(constant.StateLess)
	default:
		b.ReplyMsg(constant.ErrorInput)
		return nil
	}
}
func ConfirmTakeOff(b *BotAPI, u *model.User) error {
	msg := b.Update.Message.Text
	num, err := StrToNum(msg)
	if err != nil {
		b.ReplyMsg(constant.ErrorInput)
		return nil
	}

	model.StopOrder(u.GetUserId(), num)
	b.ReplyMsg("停车成功，欢迎下次光临")
	return u.SetState(constant.StateLess)
}

/****************************/
//发单
func WaitingSendOrder(b *BotAPI, u *model.User) error {
	switch b.Update.Message.Text {
	case "是":
		b.ReplyMsg(constant.ChooseOrder)
		return u.SetState(constant.StateConfirmSendOrder)
	case "否":
		b.ReplyMsg("服务已结束")
		return u.SetState(constant.StateLess)
	default:
		b.ReplyMsg(constant.ErrorInput)
		return nil
	}
}
func ConfirmSendOrder(b *BotAPI, u *model.User) error {
	msg := b.Update.Message.Text
	num, err := StrToNum(msg)
	if err != nil {
		fmt.Println("error input", err)
		b.ReplyMsg(constant.ErrorInput)
		return nil
	}
	order, err := model.SearchTakeOnOrder(num)
	if err != nil || order.ConfirmData == "" {
		b.ReplyMsg(constant.ErrorInputTakeOnNum)
		return u.SetState(constant.StateLess)
	}

	if order.ConfirmData == "" {
		b.ReplyMsg(constant.ErrorInputTakeOnNum)
		return u.SetState(constant.StateLess)
	}

	b.ReplyMsg(fmt.Sprintf(constant.AutoReplySendOrderSuccessTemplate, msg))
	groupInfo := fmt.Sprintf(constant.AutoSendGroupTakeOnMessage, msg, order.ConfirmData)
	b.SendMsgToGroups(groupInfo,
		constant.QQGroupHustDouble11,
		constant.QQGroupHustHelp2,
		//	constant.QQGroupXianZhi1,
		//	constant.QQGroupXianZhi2
		constant.QQTestGroup,
	)

	return u.SetState(constant.StateLess)
}

func GetConfirmData(u *model.User) (string, error) {
	info, err := u.GetDispatchInfo()
	if err != nil {
		return "", errors.New("系统错误")
	}
	typ, err := u.GetType()
	if err != nil {
		return "", errors.New("系统错误")
	}

	var confirmData string

	for i := 0; i < int(typ)+4; i++ {
		switch i + 1 {
		case 1:
			confirmData += fmt.Sprintf("校区：%s\n", info[i])
		case 2:
			confirmData += fmt.Sprintf("目的地：%s\n", info[i])
		case 3:
			confirmData += fmt.Sprintf("快递点：%s\n", info[i])
		case 4:
			confirmData += fmt.Sprintf("件数：%s", info[i])
		case 5:
			confirmData += fmt.Sprintf("\n截止时间：%s", info[i])
		default:
		}
	}
	return confirmData, nil
}
func StrToNum(numStr string) (int64, error) {
	return strconv.ParseInt(numStr, 10, 16)
}
func CheckTime(timeStr string) bool {
	return reTime.MatchString(timeStr) && (util.GetExpireTime(timeStr) > 0)
}
func CheckAt(text string) bool {
	return reAt.MatchString(text)
}
func commitInfo(b *BotAPI, u *model.User, typ int64) error {

	switch typ {
	case constant.TypeTakeOn:
		num, ok, err := u.GetTakeOnNum()
		if err != nil {
			return err
		}
		//没找到车
		if !ok {
			failedInfo := fmt.Sprintf(constant.AutoReplyTakeOnFailedInfoTemplate, num)
			b.ReplyMsg(failedInfo)
			confirmData, _ := u.GetConfirmData()
			if _, err := StrToNum(num); err != nil {
				break
			}
			groupInfo := fmt.Sprintf(constant.AutoSendGroupTakeOnMessage, num, confirmData)
			b.SendMsgToGroups(groupInfo,
				constant.QQGroupHustDouble11,
				constant.QQGroupHustHelp2,
				//	constant.QQGroupXianZhi1,
				//	constant.QQGroupXianZhi2
				constant.QQTestGroup,
			)
			return u.SetState(constant.StateLess)
		}
		//找到车
		carInfo := fmt.Sprintf(constant.AutoReplyTakeOnSuccessInfoTemplate, num)
		b.ReplyMsg(carInfo)
		return u.SetState(constant.StateLess)

	case constant.TypeDispatch:
		num := u.GetDispatchNum()
		confirmData, _ := u.GetConfirmData()
		destination, _ := u.GetDestination()
		b.ReplyMsg(fmt.Sprintf(constant.AutoReplyDispatchInfoTemplate, num))
		groupInfo := fmt.Sprintf(constant.AutoSendDispatchMessage, num, confirmData, destination)
		b.SendMsgToGroups(groupInfo,
			constant.QQGroupHustDouble11,
			constant.QQGroupHustHelp2,
			//	constant.QQGroupXianZhi1,
			//	constant.QQGroupXianZhi2
			constant.QQTestGroup,
		)
	default:

	}
	return u.SetState(constant.StateLess)
}

func PrintJson(data interface{}) string {
	datas, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("json.marshal failed, err:%v", err)
		return ""
	}

	fmt.Printf("%s\n", string(datas))
	return fmt.Sprintf("%s\n", string(datas))
}
