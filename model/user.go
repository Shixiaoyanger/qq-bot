package model

import (
	"fmt"
	qqbotapi "github.com/catsworld/qq-bot-api"
	"github.com/go-redis/redis/v7"
	"qq-bot/constant"
)

type UserInfo qqbotapi.User
type User struct {
	UserID int64 `json:"user_id"`
}

func GetUser(userID int64) (*User, error) {
	user := &User{
		userID,
	}
	err := user.Init()

	return user, err
}
func (u *User) Init() error {
	exists := RedisClient.Exists(u.RedisKey()).Val() // 如果已经初始化过，就不进行操作
	if exists != 0 {
		return nil
	}
	err := u.SetState(constant.StateLess)
	return err
}

func (u *User) RedisKey() string {
	return fmt.Sprintf("user:%d", u.GetUserId())
}

func (u *User) hSet(key string, value interface{}) error {
	return RedisClient.HSet(u.RedisKey(), key, value).Err()
}

func (u *User) hGet(filed string) *redis.StringCmd {
	return RedisClient.HGet(u.RedisKey(), filed)
}

func (u *User) GetUserId() int64 {
	return u.UserID
}

func (u *User) SetState(state int64) error {
	return u.hSet(constant.UserKeyState, state)
}

func (u *User) GetState() (int64, error) {
	return RedisClient.HGet(u.RedisKey(), constant.UserKeyState).Int64()
}

func (u *User) SetRegion(region string) error {
	return u.hSet(constant.UserKeyRegion, region)
}

func (u *User) GetRegion() (string, error) {
	return u.hGet(constant.UserKeyRegion).Result()
}
func (u *User) SetDeliveryPoint(deliveryPoint string) error {
	return u.hSet(constant.UserKeyDeliveryPoint, deliveryPoint)
}

func (u *User) SetDestination(destination string) error {
	return u.hSet(constant.UserKeyDestination, destination)
}
func (u *User) GetDestination() (string, error) {
	return u.hGet(constant.UserKeyDestination).Result()
}
func (u *User) SetDeliverNum(deliverNum string) error {
	return u.hSet(constant.UserKeyDeliveryNum, deliverNum)
}

func (u *User) SetDeadline(deadline string) error {
	return u.hSet(constant.UserKeyDeadline, deadline)
}

func (u *User) SetType(typ int64) error {
	return u.hSet(constant.UserKeyType, typ)
}

func (u *User) GetType() (int64, error) {
	return u.hGet(constant.UserKeyType).Int64()
}
func (u *User) SetConfirmData(confirmData string) error {
	return u.hSet(constant.UserKeyConfirmData, confirmData)
}
func (u *User) GetConfirmData() (string, error) {
	return u.hGet(constant.UserKeyConfirmData).Result()
}

func (u *User) GetDispatchInfo() ([]interface{}, error) {
	return RedisClient.HMGet(u.RedisKey(),
		constant.UserKeyRegion,
		constant.UserKeyDestination,
		constant.UserKeyDeliveryPoint,
		constant.UserKeyDeliveryNum,
		constant.UserKeyDeadline,
	).Result()
}
