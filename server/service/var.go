package service

var (
	AssetService             = new(assetService)
	BackupService            = new(backupService)
	CredentialService        = new(credentialService)
	GatewayService           = new(gatewayService)
	JobService               = new(jobService)
	MailService              = new(mailService)
	PropertyService          = new(propertyService)
	SecurityService          = new(securityService)
	SessionService           = new(sessionService)
	StorageService           = new(storageService)
	UserService              = new(userService)
	UserGroupService         = new(userGroupService)
	AccessTokenService       = new(accessTokenService)
	ShareSessionService      = new(shareSessionService)
	RoleService              = new(roleService)
	LoginPolicyService       = new(loginPolicyService)
	CommandFilterService     = new(commandFilterService)
	CommandFilterRuleService = new(commandFilterRuleService)
	StorageLogService        = new(storageLogService)
	AuthorisedService        = new(authorisedService)
)
