package util

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	//"samples/WebIM/models"
	"samples/WebIM/models"
	"fmt"
)

var agentId string

var token string

var sessionId string

var apiKey string

func init() {
	beego.Info("enter the init.....")
	login()
	getSession()
	Tuling("你好")

}

func login() {
	req := httplib.Post(beego.AppConfig.String("polly::auth"))
	req.Header("Connection", "keep-alive")
	polly := models.Polly{Username:beego.AppConfig.String("polly::username"), Password:beego.AppConfig.String("polly::password")}
	req.JSONBody(&polly)
	req.Body(`{"username":"yxw","password":"Hna@pwd123"}`)
	var auth models.Auth
	err := req.ToJSON(&auth)
	if err != nil {
		beego.Error("polly login error...")
	}
	token = auth.Token
	beego.Debug("the token is ", token)
	loginReq := httplib.Get(beego.AppConfig.String("polly::agent"))
	authKey := `Bearer ` + token
	loginReq.Header("Authorization", authKey)
	var authRes models.AuthRes
	loginErr := loginReq.ToJSON(&authRes)
	if loginErr != nil {
		beego.Error("polly login error...")
	}
	for _, content := range authRes.Content {
		if content.Name == beego.AppConfig.String("polly::botname") {
			agentId = content.Id
			beego.Debug("the agentId is ", agentId)
			break
		}
	}
}

func getSession() {
	apikeyUrl := beego.AppConfig.String("polly::agent") + `/` + agentId + `/apikeys`
	req := httplib.Get(apikeyUrl)
	authKey := `Bearer ` + token
	req.Header("Authorization", authKey)
	var sessions []models.Session
	err := req.ToJSON(&sessions)
	if err != nil {
		beego.Error("get session error", err)
	}
	apiKey = sessions[0].Token
	sessionId = sessions[1].Token
}

func GetApiKey() string {
	return apiKey
}

func GetAnswer(query string) (string, error) {
	req := httplib.Post(beego.AppConfig.String("polly::query"))
	req.Header("Connection", "keep-alive")
	authToken := `Bearer ` + GetApiKey()
	req.Header("Authorization", authToken)
	req.JSONBody(&models.Answer{})
	queryStr := fmt.Sprintf(`{"sessionId":"%s","lang":"zh","body":"%s","resetContexts":false}`, sessionId, query)
	beego.Debug("the querystr is ", queryStr)
	req.Body(queryStr)
	var res models.PollyRes
	err := req.ToJSON(&res)
	if err != nil {
		beego.Error("the query error is ", err)
	}
	return res.Result.Fulfillment.Speech, err
}

func Tuling(query string){
	req := httplib.Post(beego.AppConfig.String("tuling::url"))
	req.JSONBody(models.Tuling{Key:beego.AppConfig.String("tuling::apikey"),Info:query,Userid:`123123`})
	body := fmt.Sprintf(`{"key":"%s","info":"%s","userid":"123123"}`,beego.AppConfig.String("tuling::apikey"),query)
	beego.Debug("the body is ",body)
	req.Body(body)
	res,err := req.String()
	if err != nil{
		beego.Error("the err is ",err)
	}
	beego.Debug("the res is ",res)

}