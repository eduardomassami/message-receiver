package messenger

type WhatsappPayload struct {
	Headers                         WhatsappHeaders                         `json:"headers" validate:"required"`
	MultiValueHeaders               WhatsappMultiValueHeaders               `json:"multiValueHeaders" validate:"required"`
	Path                            string                                  `json:"path" validate:"required"`
	PathParameters                  string                                  `json:"pathParameters" validate:"required"`
	RequestContext                  WhatsappRequestContext                  `json:"requestContext" validate:"required"`
	Resource                        string                                  `json:"resource" validate:"required"`
	HttpMethod                      string                                  `json:"httpMethod" validate:"required"`
	QueryStringParameters           WhatsappQueryStringParameters           `json:"queryStringParameters" validate:"required"`
	MultiValueQueryStringParameters WhatsappMultiValueQueryStringParameters `json:"multiValueQueryStringParameters" validate:"required"`
	StageVariables                  string                                  `json:"stageVariables" validate:"required"`
	Body                            WhatsappBody                            `json:"body" validate:"required"`
	IsOffline                       bool                                    `json:"isOffline" validate:"required"`
}

type WhatsappHeaders struct {
	Host            string `json:"Host" validate:"required"`
	UserAgent       string `json:"User-Agent" validate:"required"`
	ContentLength   string `json:"Content-Length" validate:"required"`
	AcceptEncoding  string `json:"Accept-Encoding" validate:"required"`
	AcceptLanguage  string `json:"Accept-Language" validate:"required"`
	ContentType     string `json:"Content-Type" validate:"required"`
	XForwardedFor   string `json:"X-Forwarded-For" validate:"required"`
	XForwardedProto string `json:"X-Forwarded-Proto" validate:"required"`
	XRequestId      string `json:"X-Request-Id" validate:"required"`
	XWaAccountId    string `json:"X-Wa-Account-Id" validate:"required"`
}

type WhatsappMultiValueHeaders struct {
	Host            []string `json:"Host" validate:"required"`
	UserAgent       []string `json:"User-Agent" validate:"required"`
	ContentLength   []string `json:"Content-Length" validate:"required"`
	AcceptEncoding  []string `json:"Accept-Encoding" validate:"required"`
	AcceptLanguage  []string `json:"Accept-Language" validate:"required"`
	ContentType     []string `json:"Content-Type" validate:"required"`
	XForwardedFor   []string `json:"X-Forwarded-For" validate:"required"`
	XForwardedProto []string `json:"X-Forwarded-Proto" validate:"required"`
	XRequestId      []string `json:"X-Request-Id" validate:"required"`
	XWaAccountId    []string `json:"X-Wa-Account-Id" validate:"required"`
}

type WhatsappRequestContext struct {
	AccountId    string             `json:"accountId" validate:"required"`
	ResourceId   string             `json:"resourceId" validate:"required"`
	ApiId        string             `json:"apiId" validate:"required"`
	RequestId    string             `json:"requestId" validate:"required"`
	Identity     WhatsappIdentity   `json:"identity" validate:"required"`
	Authorizer   WhatsappAuthorizer `json:"authorizer" validate:"required"`
	Protocol     string             `json:"protocol" validate:"required"`
	ResourcePath string             `json:"resourcePath" validate:"required"`
	HttpMethod   string             `json:"httpMethod" validate:"required"`
}

type WhatsappIdentity struct {
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

type WhatsappAuthorizer struct {
	PrincipalId string `json:"principalId" validate:"required"`
}

type WhatsappQueryStringParameters struct {
	OriginChannel string `json:"originChannel" validate:"required"`
}

type WhatsappMultiValueQueryStringParameters struct {
	OriginChannel []string `json:"originChannel" validate:"required"`
}

type WhatsappBody struct {
	Contacts []WhatsappContacts `json:"contacts" validate:"required"`
	Messages []WhatsappMessage  `json:"messages" validate:"required"`
}

type WhatsappContacts struct {
	Profile WhatsappProfile `json:"profile" validate:"required"`
	WaId    string          `json:"wa_id" validate:"required"`
}

type WhatsappProfile struct {
	Name string `json:"name" validate:"required"`
}

type WhatsappMessage struct {
	From      string       `json:"from" validate:"required"`
	Id        string       `json:"id" validate:"required"`
	Text      WhatsappText `json:"text" validate:"required"`
	Timestamp string       `json:"timestamp" validate:"required"`
	Type      string       `json:"type" validate:"required"`
}

type WhatsappText struct {
	Body string `json:"body" validate:"required"`
}
