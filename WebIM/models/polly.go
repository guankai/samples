package models

type Polly struct {
	Username string `json:username`
	Password string `json:password`
}

type Tuling struct {
	Key    string `json:key`
	Info   string `json:info`
	Userid string `json:userid`
}

type Auth struct {
	Token string `json:token`
}

type Content struct {
	Id   string `json:id`
	Name string `json:name`
}

type AuthRes struct {
	Content []Content `json:content`
}

type Session struct {
	Token string `json:token`
}

type Answer struct {
	SessionId string `json:sessoinId`
	Lang string `json:lang`
	Body string `json:body`
	ResetContexts bool `json:resetContexts`
}

type Fulfillment struct {
	Speech string `json:speech`
}

type Status struct {
	Status int `json:status`
	Message string `json:message`
	Error string `json:error`
}

type Result struct {
	Source string `json:source`
	Fulfillment Fulfillment `json:fulfillment`
}

type PollyRes struct {
	Id string `json:id`
	Result Result `json:result`
	Status Status `json:status`
}

type TulingRes struct {
	Code int `json:code`
	Text string `json:text`
}

