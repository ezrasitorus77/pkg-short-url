package consts

const (
	HSTS                    string = "Strict-Transport-Security"
	CSP                     string = "Content-Security-Policy"
	ContentType             string = "Content-Type"
	ContentDisposition      string = "Content-Disposition"
	ContentTransferEncoding string = "Content-Transfer-Encoding"
	Expires                 string = "Expires"

	JSON      string = "application/json"
	XWWW      string = "application/x-www-form-urlencoded"
	Multipart string = "multipart/form-data"
	HTML      string = "text/html; charset=utf-8"
	SHEET     string = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"

	BinaryEncoding string = "binary"

	DefaultHSTS    string = "max-age=63072000; includeSubDomains; preload"
	DefaultCSP     string = "default-src 'self'; img-src https://*; script-src 'self'; object-src 'none'"
	AllowedMethods string = "GET,POST"

	AccessControlAllowOrigin string = "Access-Control-Allow-Origin"
	AccesControlAllowMethods string = "Access-Control-Allow-Methods"

	Token string = "Token"

	XAPIKey               string = "X-Api-Key"
	XForwardedHost        string = "X-Forwarded-Host"
	XForwardedFor         string = "X-Forwarded-For"
	XOriginHost           string = "X-Origin-Host"
	XRequestedWith        string = "X-Requested-With"
	XRequestID            string = "X-Request-Id"
	XRealIP               string = "X-Real-Ip"
	XForwardedProto       string = "X-Forwarded-Proto"
	XForwardedPort        string = "X-Forwarded-Port"
	XScheme               string = "X-Scheme"
	XOriginalForwardedFor string = "X-Original-Forwarded-For"

	XRequestAccess string = "X-Request-Access"
	Origin         string = "Origin"
)
