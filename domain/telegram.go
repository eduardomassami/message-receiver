package telegram

type TelegramPayload struct {
	Headers                         Headers
	MultiValueHeaders               MultiValueHeaders
	Path                            string `json:"path" validate:"required"`
	PathParameters                  string `json:"pathParameters" validate:"required"`
	RequestContext                  RequestContext
	Resource                        string `json:"resource" validate:"required"`
	HttpMethod                      string `json:"httpMethod" validate:"required"`
	QueryStringParameters           QueryStringParameters
	MultiValueQueryStringParameters MultiValueQueryStringParameters
	StageVariables                  string `json:"stageVariables" validate:"required"`
	Body                            Body
	IsOffline                       bool `json:"isOffline" validate:"required"`
}

type Headers struct {
	Host            string `json:"Host" validate:"required"`
	ContentLength   string `json:"Content-Length" validate:"required"`
	AcceptEncoding  string `json:"Accept-Encoding" validate:"required"`
	ContentType     string `json:"Content-Type" validate:"required"`
	XForwardedFor   string `json:"X-Forwarded-For" validate:"required"`
	XForwardedProto string `json:"X-Forwarded-Proto" validate:"required"`
}

type MultiValueHeaders struct {
	Host            []string `json:"Host" validate:"required"`
	ContentLength   []string `json:"Content-Length" validate:"required"`
	AcceptEncoding  []string `json:"Accept-Encoding" validate:"required"`
	ContentType     []string `json:"Content-Type" validate:"required"`
	XForwardedFor   []string `json:"X-Forwarded-For" validate:"required"`
	XForwardedProto []string `json:"X-Forwarded-Proto" validate:"required"`
}

type RequestContext struct {
	AccountId    string `json:"accountId" validate:"required"`
	ResourceId   string `json:"resourceId" validate:"required"`
	ApiId        string `json:"apiId" validate:"required"`
	RequestId    string `json:"requestId" validate:"required"`
	Identity     Identity
	Authorizer   Authorizer
	Protocol     string `json:"protocol" validate:"required"`
	ResourcePath string `json:"resourcePath" validate:"required"`
	HttpMethod   string `json:"httpMethod" validate:"required"`
}

type Identity struct {
	CognitoIdentityPoolId         string `json:"cognitoIdentityPoolId" validate:"required"`
	AccountId                     string `json:"accountId" validate:"required"`
	CognitoIdentityId             string `json:"cognitoIdentityId" validate:"required"`
	Caller                        string `json:"caller" validate:"required"`
	ApiKey                        string `json:"apiKey" validate:"required"`
	SourceIp                      string `json:"sourceIp" validate:"required"`
	CognitoAuthenticationType     string `json:"cognitoAuthenticationType" validate:"required"`
	CognitoAuthenticationProvider string `json:"cognitoAuthenticationProvider" validate:"required"`
	UserArn                       string `json:"userArn" validate:"required"`
	UserAgent                     string `json:"userAgent" validate:"required"`
	User                          string `json:"user" validate:"required"`
}

type Authorizer struct {
	PrincipalId string `json:"principalId" validate:"required"`
}

type QueryStringParameters struct {
	OriginChannel string `json:"originChannel" validate:"required"`
}

type MultiValueQueryStringParameters struct {
	OriginChannel []string `json:"originChannel" validate:"required"`
}

type Body struct {
	UpdateId int `json:"update_id" validate:"required"`
	Message  Message
}

type Message struct {
	MessageId      int `json:"message_id" validate:"required"`
	From           From
	Chat           Chat
	Date           int            `json:"date" validate:"required"`
	ReplyToMessage ReplyToMessage `json:"reply_to_message" validate:"required"`
	Text           string         `json:"text" validate:"required"`
}

type From struct {
	Id           int    `json:"id" validate:"required"`
	IsBot        bool   `json:"is_bot" validate:"required"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	UserName     string `json:"username" validate:"required"`
	LanguageCode string `json:"language_code" validate:"required"`
}

type Chat struct {
	Id        int    `json:"id" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	UserName  string `json:"username" validate:"required"`
	Type      string `json:"type" validate:"required"`
}

type ReplyToMessage struct {
	MessageId int `json:"message_id" validate:"required"`
	From      From
	Chat      Chat
	Date      int    `json:"date" validate:"required"`
	Text      string `json:"text" validate:"required"`
	Entities  []Entities
}

type Entities struct {
	Offset int    `json:"offset" validate:"required"`
	Length int    `json:"length" validate:"required"`
	Type   string `json:"type" validate:"required"`
}
