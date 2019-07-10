package pay

import (
	"fmt"
	"github.com/imokyou/wechat/mch/core"
	wechatutil "github.com/imokyou/wechat/util"
)

type TransferRequest struct {
	MchAppid       string `json:"mch_appid"`
	Mchid          string `json:"mchid"`
	DeviceInfo     string `json:"device_info"`
	NonceStr       string `json:"nonce_str"`
	PartnerTradeNo string `json:"partner_trade_no"`
	Openid         string `json:"openid"`
	CheckName      string `json:"check_name"`
	ReUserName     string `json:"re_user_name"`
	Amount         int    `json:"amount"`
	Desc           string `json:"desc"`
	SpbillCreateIp string `json:"spbill_create_ip"`
}

type TransferResponse struct {
	ReturnCode     string `json:"return_code"`
	ReturnMsg      string `json:"return_msg"`
	MchAppid       string `json:"mch_appid"`
	Mchid          string `json:"mchid"`
	DeviceInfo     string `json:"device_info"`
	NonceStr       string `json:"nonce_str"`
	ResultCode     string `json:"result_code"`
	ErrCode        string `json:"err_code"`
	ErrCodeDes     string `json:"err_code_des"`
	PartnerTradeNo string `json:"partner_trade_no"`
	PaymentNo      string `json:"payment_no"`
	PaymentTime    string `json:"payment_time"`
}

// Transfer 企业付款.
//  NOTE: 请求需要证书.
func Transfer(clt *core.Client, req *TransferRequest) (resp *TransferResponse, err error) {
	params := make(map[string]string)
	respParams := make(map[string]string)

	params["mch_appid"] = req.MchAppid
	params["mchid"] = req.Mchid

	if req.DeviceInfo != "" {
		params["DeviceInfo"] = req.DeviceInfo
	}

	if req.NonceStr != "" {
		params["nonce_str"] = req.NonceStr
	} else {
		params["nonce_str"] = wechatutil.NonceStr()
	}

	params["partner_trade_no"] = req.PartnerTradeNo
	params["openid"] = req.Openid

	if req.CheckName != "" {
		params["check_name"] = req.CheckName
	} else {
		params["check_name"] = "NO_CHECK"
	}

	if req.ReUserName != "" {
		params["re_user_name"] = req.ReUserName
	}
	params["amount"] = fmt.Sprintf("%d", req.Amount)
	params["desc"] = req.Desc

	if req.SpbillCreateIp != "" {
		params["spbill_create_ip"] = req.SpbillCreateIp
	} else {
		params["spbill_create_ip"] = "127.0.0.1"
	}

	respParams, err = transfer(clt, params)

	resp.ReturnCode = respParams["return_code"]
	resp.ReturnMsg = respParams["return_msg"]
	resp.MchAppid = respParams["mch_appid"]
	resp.Mchid = respParams["mchid"]
	resp.DeviceInfo = respParams["device_info"]
	resp.NonceStr = respParams["nonce_str"]
	resp.ResultCode = respParams["result_code"]
	resp.ErrCode = respParams["err_code"]
	resp.ErrCodeDes = respParams["err_code_res"]
	resp.PartnerTradeNo = respParams["partner_trade_no"]
	resp.PaymentNo = respParams["payment_no"]
	resp.PaymentTime = respParams["payment_time"]

	return
}

func transfer(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/mmpaymkttransfers/promotion/transfers", req)
}
