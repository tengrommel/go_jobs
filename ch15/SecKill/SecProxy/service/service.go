package service

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"time"
	"crypto/md5"
)

var (
	secKillConf *SecSkillConf
)

func InitService(serviceConf *SecSkillConf) {
	secKillConf = serviceConf
	logs.Debug("init service succ, config:%v", secKillConf)
}

func SecInfoList() (data []map[string]interface{}, code int, err error)  {
	secKillConf.RWSecProductLock.RLock()
	defer secKillConf.RWSecProductLock.RUnlock()

	for _, v := range secKillConf.SecProductInfoMap{
		item, _, err := SecInfoById(v.ProductId)
		if err != nil{
			logs.Error("get product_id[%d] failed, err:%v", v.ProductId, err)
			continue
		}
		data = append(data, item)
	}

	return
}


func SecInfo(productId int) (data []map[string]interface{}, code int, err error) {

	secKillConf.RWSecProductLock.RLock()
	defer secKillConf.RWSecProductLock.RUnlock()

	item, code, err := SecInfoById(productId)
	if err != nil{
		return
	}
	data = append(data, item)

	return
}

func SecInfoById(productId int) (data map[string]interface{}, code int, err error) {

	secKillConf.RWSecProductLock.RLock()
	defer secKillConf.RWSecProductLock.RUnlock()

	v, ok := secKillConf.SecProductInfoMap[productId]
	if !ok {
		code = ErrNotFoundProductId
		err = fmt.Errorf("not found product_id:%d", productId)
		return
	}

	start := false
	end := false
	status := "success"

	now := time.Now().Unix()
	if now - v.StartTime < 0 {
		start = false
		end = false
		status = "sec kill is not start"
	}

	if now - v.StartTime > 0 {
		start = true
	}

	if now-v.EndTime > 0{
		start = false
		end = true
	}

	if v.Status == ProductStatusForceSaleOut || v.Status == ProductStatusSaleOut{
		start = false
		end = true
		status = "Product is sale out"
	}

	data = make(map[string]interface{})
	data["product_id"] = productId
	data["start"] = start
	data["end"] = end
	data["status"] = status
	return
}

func UserCheck(req *SecRequest) (err error) {
	authData := fmt.Sprintf("%d:%s", req.UserId, secKillConf.CookieSecretKey)
	authSign := fmt.Sprintf("%x", md5.Sum([]byte(authData)))
	if authSign != req.UserAuthSign{
		err = fmt.Errorf("invalid user cookie")
		return
	}
	return
}



func SecKill(req *SecRequest) (data []map[string]interface{}, code int, err error)  {
	secKillConf.RWSecProductLock.RLock()
	defer secKillConf.RWSecProductLock.RUnlock()

	err = UserCheck(req)
	if err != nil{
		code = ErrUserCheckAuthFailed
		logs.Warn("userId[%d] invalid, check failed, req[%v]", req.UserId, req)
	}

	err = antiSpam(req)
	if err != nil{

	}
	return
}