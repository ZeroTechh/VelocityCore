package services

const (
	//UsersService is used to name and locate users service
	UsersService = "users-srv"
	//JWTService is used to name and locate jwt service
	JWTService = "jwt-srv"
	//EmailVerificationSrv is used to name and locate email verification service
	EmailVerificationSrv = "email-verification-srv"
	//AuthServerService is used to name and locate auth server service
	AuthServerService = "auth-server-srv"
	//ResourceServerService is used to name and locate resource server service
	ResourceServerService = "resource-server-srv"
)

//ServiceUrls is a map of all services and their url
var ServiceUrls = map[string]string{
	UsersService:         "127.0.0.1:3000",
	JWTService:           "127.0.0.1:3001",
	EmailVerificationSrv: "127.0.0.1:3002",
}
