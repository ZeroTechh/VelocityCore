package services

const (
	// UserMainService is used to name and locate user main service
	UserMainService = "user-main-srv"
	// UserExtraService is used to name and locate user extra service
	UserExtraService = "user-extra-srv"
	// UserMetaService is used to name and locate user meta service
	UserMetaService = "user-meta-srv"
	// JWTService is used to name and locate jwt service
	JWTService = "jwt-srv"
	// EmailVerificationSrv is used to name and locate email verification service
	EmailVerificationSrv = "email-verification-srv"
	// AuthServerService is used to name and locate auth server service
	AuthServerService = "auth-server-srv"
	// ResourceServerService is used to name and locate resource server service
	ResourceServerService = "resource-server-srv"
	// VerificationCodeService is used to name and locate verification code service
	VerificationCodeService = "verification-code-srv"
)

// ServiceUrls is a map of all services and their url
var ServiceUrls = map[string]string{
	UserMainService:         "127.0.0.1:3000",
	JWTService:              "127.0.0.1:3001",
	EmailVerificationSrv:    "127.0.0.1:3002",
	UserExtraService:        "127.0.0.1:3003",
	UserMetaService:         "127.0.0.1:3004",
	VerificationCodeService: "127.0.0.1:3005",
}
