package model

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"qq-bot/constant"
	"qq-bot/util"
	"strconv"
	"time"
)

type Order struct {
	ID             bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	OrderID        int64         `bson:"order_id" json:"order_id"`
	UserID         int64         `bson:"user_id" json:"user_id"`
	Region         string        `bson:"region" json:"region"`
	DeliveryPoint  string        `bson:"delivery_point" json:"delivery_point"`
	Destination    string        `bson:"destination" json:"destination"`
	DestinationKey string        `bson:"destination_key" json:"destination_key"`
	DeliveryNum    string        `bson:"delivery_num" json:"delivery_num"`
	Deadline       int64         `bson:"deadline" json:"deadline"`
	Type           int           `bson:"type" json:"type"`
	ConfirmData    string        `bson:"confirm_data" json:"confirm_data"`
	OrderStatus    int           `bson:"order_status" json:"order_status"` //1 正常 -1满员

}

/*
	韵苑
	1 2 3 4
	5 6 7 8 9
	10 11 12 24 26 27
	22 23 25
	13 14 16 17 20 21
	15 18 19 28

	紫菘
	1 3 4
	7 8 11 12 13
	2 5 6
	9  10 13

	沁苑
	1 6 7
	2 3 4 5 8
	9 10 11 12 13
*/
var orderKey = map[string]string{
	"韵苑1栋":  "yunYuan1234",
	"韵苑2栋":  "yunYuan1234",
	"韵苑3栋":  "yunYuan1234",
	"韵苑4栋":  "yunYuan1234",
	"韵苑5栋":  "yunYuan56789",
	"韵苑6栋":  "yunYuan56789",
	"韵苑7栋":  "yunYuan56789",
	"韵苑8栋":  "yunYuan56789",
	"韵苑9栋":  "yunYuan56789",
	"韵苑10栋": "yunYuan101112242627",
	"韵苑11栋": "yunYuan101112242627",
	"韵苑12栋": "yunYuan101112242627",
	"韵苑13栋": "yunYuan131416172021",
	"韵苑14栋": "yunYuan131416172021",
	"韵苑15栋": "yunYuan15181928",
	"韵苑16栋": "yunYuan131416172021",
	"韵苑17栋": "yunYuan131416172021",
	"韵苑18栋": "yunYuan15181928",
	"韵苑19栋": "yunYuan15181928",
	"韵苑20栋": "yunYuan131416172021",
	"韵苑21栋": "yunYuan131416172021",
	"韵苑22栋": "yunYuan222325",
	"韵苑23栋": "yunYuan222325",
	"韵苑24栋": "yunYuan101112242627",
	"韵苑25栋": "yunYuan222325",
	"韵苑26栋": "yunYuan101112242627",
	"韵苑27栋": "yunYuan101112242627",
	"韵苑28栋": "yunYuan15181928",
	"韵苑任意栋": "yunYuanArbitrary",

	"紫菘1栋":  "ziSong134",
	"紫菘2栋":  "ziSong256",
	"紫菘3栋":  "ziSong134",
	"紫菘4栋":  "ziSong134",
	"紫菘5栋":  "ziSong256",
	"紫菘6栋":  "ziSong256",
	"紫菘7栋":  "ziSong78111213",
	"紫菘8栋":  "ziSong78111213",
	"紫菘9栋":  "ziSong910",
	"紫菘10栋": "ziSong910",
	"紫菘11栋": "ziSong78111213",
	"紫菘12栋": "ziSong78111213",
	"紫菘13栋": "ziSong78111213",
	"紫菘任意栋": "ziSongArbitrary",

	"东1舍":  "qinYuan167",
	"东2舍":  "qinYuan23458",
	"东3舍":  "qinYuan23458",
	"东4舍":  "qinYuan23458",
	"东5舍":  "qinYuan23458",
	"东6舍":  "qinYuan167",
	"东7舍":  "qinYuan167",
	"东8舍":  "qinYuan23458",
	"东9舍":  "qinYuan910111213",
	"东10舍": "qinYuan910111213",
	"东11舍": "qinYuan910111213",
	"东12舍": "qinYuan910111213",
	"东13舍": "qinYuan910111213",
	"东任意舍": "qinYuanArbitrary",
}
var orderKeyArbitrary = map[string]string{
	constant.RegionYunYuan: "yunYuanArbitrary",
	constant.RegionZiSong:  "ziSongArbitrary",
	constant.RegionQinYuan: "qinYuanArbitrary",
}

func (u *User) GetDispatchNum() string {
	info, err := RedisClient.HGetAll(u.RedisKey()).Result()
	if err != nil {
		return "error"
	}
	orderID, err := RedisClient.Incr(constant.OrderKeyIncr).Result()
	if err != nil {
		return "error"
	}
	orderKeyDestination, ok := orderKey[info["destination"]]
	if !ok {
		return "\n由于输入的楼栋不存在，获取编号失败\n"
	}

	status, _ := strconv.Atoi(info["delivery_num"])
	order := Order{
		ID:             bson.NewObjectId(),
		OrderID:        orderID,
		UserID:         u.GetUserId(),
		Region:         info["region"],
		Destination:    info["destination"],
		DestinationKey: orderKeyDestination,
		DeliveryPoint:  info["delivery_point"],
		DeliveryNum:    info["delivery_num"],
		Deadline:       util.GetTimeStampByMinute(info["deadline"]),
		Type:           constant.TypeDispatch,
		ConfirmData:    info["confirm_data"],
		OrderStatus:    status,
	}

	_ = insertDocs(constant.TableOrder, order)

	return strconv.FormatInt(orderID, 10)
}

func (u *User) GetTakeOnNum() (string, bool, error) {
	info, err := RedisClient.HGetAll(u.RedisKey()).Result()
	if err != nil {
		return "error info", false, err
	}
	orderID, err := RedisClient.Incr(constant.OrderKeyIncr).Result()
	if err != nil {
		return "error", false, err
	}

	orderKeyDestination, ok := orderKey[info["destination"]]
	if !ok {
		return "\n由于输入的楼栋不存在，获取编号失败\n", false, err
	}

	status, _ := strconv.Atoi(info["delivery_num"])
	order := Order{
		ID:             bson.NewObjectId(),
		OrderID:        orderID,
		UserID:         u.GetUserId(),
		Region:         info["region"],
		Destination:    info["destination"],
		DestinationKey: orderKeyDestination,
		DeliveryPoint:  info["delivery_point"],
		DeliveryNum:    info["delivery_num"],
		Deadline:       util.GetTimeStampByMinute(info["deadline"]),
		Type:           constant.TypeTakeOn,
		ConfirmData:    info["confirm_data"],
		OrderStatus:    status,
	}
	insertDocs(constant.TableOrder, order)

	cars, err := SearchCar(info)
	if err != nil {
		fmt.Println(err)
		return "error car", false, err
	}
	var carsInfo string
	if len(cars) > 0 {
		for _, car := range cars {
			info := fmt.Sprintf("\n发车人QQ：%d\n发车人信息：%s\n\n你的上车编号是：%d\n", car.UserID, car.ConfirmData, orderID)
			carsInfo += info
		}
		return carsInfo, true, nil
	}

	return strconv.FormatInt(orderID, 10), false, nil
}
func SearchCar(info map[string]string) ([]Order, error) {
	cntrl := NewCloneMgoDBCntlr()
	defer cntrl.Close()
	query := bson.M{
		"delivery_point": info["delivery_point"],
		"type":           constant.TypeDispatch,
		"order_status": bson.M{
			"$gte": 0,
		},
	}
	orders, err := findOrders(query, bson.M{})
	if err != nil {
		return nil, err
	}
	var data []Order
	i := 0
	for _, order := range orders {
		destinationKey := orderKey[info["destination"]]
		if order.DestinationKey == destinationKey || order.DestinationKey == orderKeyArbitrary[order.Region] {
			if order.Deadline > time.Now().Unix() {
				data = append(data, order)
				i += 1
			} else {
				order.OrderStatus = constant.OrderStateTimeOut
				updateOrder(order)
			}
		}
		if i == 4 {
			break
		}
	}
	return data, nil
}

func SearchTakeOnOrder(num int64) (Order, error) {
	query := bson.M{
		"order_id": num,
		"type":     constant.TypeTakeOn,
		"order_status": bson.M{
			"$gte": 0,
		},
	}
	order, err := findOrder(query, bson.M{})
	if err != nil || order.ConfirmData == "" {
		return order, err
	}
	return order, nil
}
func StopOrder(userID, num int64) {
	query := bson.M{
		"order_id": num,
		"order_status": bson.M{
			"$gte": 0,
		},
	}
	order, err := findOrder(query, bson.M{})
	if err != nil || order.ConfirmData == "" {
		return
	}
	if order.UserID != userID {
		return
	}
	order.OrderStatus = constant.OrderStateStop
	updateOrder(order)
	return
}
func DeleteOrder(num int64) {
	query := bson.M{
		"order_id": num,
		"order_status": bson.M{
			"$gte": 0,
		},
	}
	order, err := findOrder(query, bson.M{})
	if err != nil {
		return
	}
	order.OrderStatus = constant.OrderStateTakenOn
	updateOrder(order)
	return
}
func updateOrder(order Order) error {
	cntrl := NewCloneMgoDBCntlr()
	defer cntrl.Close()
	query := bson.M{
		"_id": order.ID,
	}
	return updateDoc(constant.TableOrder, query, order)

}
func findOrders(query bson.M, selectField bson.M) ([]Order, error) {
	var data []Order
	cntrl := NewCopyMgoDBCntlr()
	defer cntrl.Close()
	table := cntrl.GetTable(constant.TableOrder)
	err := table.Find(query).Select(selectField).All(&data)
	return data, err
}
func findOrder(query bson.M, selectField bson.M) (Order, error) {
	var data Order
	cntrl := NewCopyMgoDBCntlr()
	defer cntrl.Close()
	table := cntrl.GetTable(constant.TableOrder)
	err := table.Find(query).Select(selectField).One(&data)
	return data, err
}

func insertDocs(tableName string, docs ...interface{}) error {
	cntrl := NewCloneMgoDBCntlr()
	defer cntrl.Close()
	table := cntrl.GetTable(tableName)
	return table.Insert(docs...)
}
func updateDoc(tableName string, query, update interface{}) error {
	cntrl := NewCloneMgoDBCntlr()
	defer cntrl.Close()
	table := cntrl.GetTable(tableName)
	return table.Update(query, update)
}
