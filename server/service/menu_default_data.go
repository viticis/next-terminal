package service

import "next-terminal/server/model"

var DefaultMenu = []*model.Menu{
	model.NewMenu("dashboard", "控制面板", "root",
		model.NewPermission("GET", "/overview/counter"),
		model.NewPermission("GET", "/overview/asset"),
		model.NewPermission("GET", "/overview/date-counter"),
	),

	model.NewMenu("resource", "资源管理", "root"),

	model.NewMenu("asset", "资产管理", "resource",
		model.NewPermission("GET", "/assets/paging"),
		model.NewPermission("GET", "/tags"),
	),
	model.NewMenu("asset-access", "接入", "asset",
		model.NewPermission("POST", "/sessions"),
		model.NewPermission("GET", "/term"),
		model.NewPermission("GET", "/tunnel"),
	),
	model.NewMenu("asset-add", "新建", "asset",
		model.NewPermission("POST", "/assets"),
	),
	model.NewMenu("asset-edit", "编辑", "asset",
		model.NewPermission("GET", "/assets/:id"),
		model.NewPermission("PUT", "/assets/:id"),
	),
	model.NewMenu("asset-del", "删除", "asset",
		model.NewPermission("DELETE", "/assets/:id"),
	),
	model.NewMenu("asset-copy", "复制", "asset",
		model.NewPermission("GET", "/assets/:id"),
		model.NewPermission("POST", "/assets"),
	),
	model.NewMenu("asset-conn-test", "连通性测试", "asset",
		model.NewPermission("POST", "/assets/:id/tcping"),
	),
	model.NewMenu("asset-import", "导入资产", "asset",
		model.NewPermission("POST", "/assets/import"),
	),
	model.NewMenu("asset-change-owner", "更换所有者", "asset",
		model.NewPermission("GET", "/users"),
		model.NewPermission("POST", "/assets/:id/change-owner"),
	),
	model.NewMenu("asset-detail", "详情", "asset",
		model.NewPermission("GET", "/assets/:id"),
	),
	model.NewMenu("asset-authorised-user", "资产授权用户", "asset-detail",
		model.NewPermission("GET", "/authorised/users/paging"),
		model.NewPermission("GET", "/authorised/selected"),
		model.NewPermission("GET", "/users"),
		model.NewPermission("GET", "/strategies"),
		model.NewPermission("GET", "/command-filters"),
		model.NewPermission("POST", "/authorised/users"),
	),
	model.NewMenu("asset-authorised-user-add", "增加授权", "asset-authorised-user",
		model.NewPermission("POST", "/authorised/:id/users"),
	),
	model.NewMenu("asset-authorised-user-del", "移除授权", "asset-authorised-user",
		model.NewPermission("DELETE", "/authorised/:id"),
	),
	model.NewMenu("asset-authorised-user-group", "资产授权用户组", "asset-detail",
		model.NewPermission("GET", "/authorised/user-groups/paging"),
		model.NewPermission("GET", "/authorised/selected"),
		model.NewPermission("GET", "/user-groups"),
		model.NewPermission("GET", "/strategies"),
		model.NewPermission("GET", "/command-filters"),
		model.NewPermission("POST", "/authorised/user-groups"),
	),
	model.NewMenu("asset-authorised-user-group-add", "增加授权", "asset-authorised-user-group",
		model.NewPermission("POST", "/authorised/:id/user-groups"),
	),
	model.NewMenu("asset-authorised-user-group-del", "移除授权", "asset-authorised-user-group",
		model.NewPermission("DELETE", "/authorised/:id"),
	),

	model.NewMenu("credential", "授权凭证", "resource",
		model.NewPermission("GET", "/credentials/paging"),
	),
	model.NewMenu("credential-add", "增加", "credential",
		model.NewPermission("POST", "/credentials"),
	),
	model.NewMenu("credential-del", "删除", "credential",
		model.NewPermission("DELETE", "/credentials/:id"),
	),
	model.NewMenu("credential-edit", "修改", "credential",
		model.NewPermission("POST", "/credentials/:id"),
	),

	model.NewMenu("command", "动态指令", "resource",
		model.NewPermission("GET", "/commands/paging"),
	),
	model.NewMenu("command-add", "增加", "command",
		model.NewPermission("POST", "/commands"),
	),
	model.NewMenu("command-edit", "修改", "command",
		model.NewPermission("PUT", "/commands/:id"),
	),
	model.NewMenu("command-del", "删除", "command",
		model.NewPermission("DELETE", "/commands/:id"),
	),
	model.NewMenu("command-exec", "执行", "command",
		model.NewPermission("GET", "/assets/paging"),
		model.NewPermission("GET", "/tags"),
		model.NewPermission("POST", "/sessions"),
		model.NewPermission("GET", "/term"),
	),
	model.NewMenu("command-change-owner", "更换所有者", "command",
		model.NewPermission("GET", "/users"),
		model.NewPermission("POST", "/commands/:id/change-owner"),
	),

	model.NewMenu("access-gateway", "接入网关", "resource",
		model.NewPermission("GET", "/access-gateways/paging"),
	),
	model.NewMenu("access-gateway-add", "增加", "access-gateway",
		model.NewPermission("POST", "/access-gateways"),
	),
	model.NewMenu("access-gateway-del", "删除", "access-gateway",
		model.NewPermission("DELETE", "/access-gateways/:id"),
	),
	model.NewMenu("access-gateway-edit", "修改", "access-gateway",
		model.NewPermission("PUT", "/access-gateways/:id"),
	),

	model.NewMenu("session-audit", "会话审计", "root"),

	model.NewMenu("online-session", "在线会话", "session-audit",
		model.NewPermission("GET", "/sessions/paging"),
	),
	model.NewMenu("online-session-disconnect", "断开", "online-session",
		model.NewPermission("GET", "/sessions/:id/disconnect"),
	),
	model.NewMenu("online-session-monitor", "监控", "online-session",
		model.NewPermission("GET", "/sessions/:id/tunnel-monitor"),
		model.NewPermission("GET", "/sessions/:id/ssh-monitor"),
	),

	model.NewMenu("offline-session", "历史会话", "session-audit",
		model.NewPermission("GET", "/sessions/paging"),
	),
	model.NewMenu("offline-session-playback", "回放", "offline-session",
		model.NewPermission("GET", "/sessions/:id/recording"),
	),
	model.NewMenu("offline-session-del", "删除", "offline-session",
		model.NewPermission("DELETE", "/sessions/:id"),
	),
	model.NewMenu("offline-session-clear", "清空", "offline-session",
		model.NewPermission("POST", "/sessions/clear"),
	),
	model.NewMenu("offline-session-command", "命令记录", "offline-session",
		model.NewPermission("GET", "/sessions/:id/commands/paging"),
	),
	model.NewMenu("offline-session-reviewed", "标记已读", "offline-session"),       // TODO
	model.NewMenu("offline-session-unreviewed", "标记未读", "offline-session"),     // TODO
	model.NewMenu("offline-session-reviewed-all", "全部标记已读", "offline-session"), // TODO

	model.NewMenu("log-audit", "日志审计", "root"),

	model.NewMenu("login-log", "登录日志", "log-audit",
		model.NewPermission("GET", "/login-logs/paging"),
	),
	model.NewMenu("login-log-del", "删除", "login-log",
		model.NewPermission("DELETE", "/login-logs/:id"),
	),
	model.NewMenu("login-log-clear", "清空", "login-log",
		model.NewPermission("POST", "/login-logs/clear"),
	),

	model.NewMenu("storage-log", "文件日志", "log-audit",
		model.NewPermission("GET", "/storage-logs/paging"),
	),

	model.NewMenu("storage-log-del", "删除", "storage-log",
		model.NewPermission("DELETE", "/storage-logs/:id"),
	),
	model.NewMenu("storage-log-clear", "清空", "storage-log",
		model.NewPermission("POST", "/storage-logs/clear"),
	),

	model.NewMenu("ops", "系统运维", "root"),

	model.NewMenu("job", "计划任务", "ops",
		model.NewPermission("GET", "/jobs/paging"),
	),
	model.NewMenu("job-add", "增加", "job",
		model.NewPermission("POST", "/jobs"),
		model.NewPermission("GET", "/assets/paging"),
	),
	model.NewMenu("job-del", "删除", "job",
		model.NewPermission("DELETE", "/jobs/:id"),
	),
	model.NewMenu("job-edit", "修改", "job",
		model.NewPermission("PUT", "/jobs/:id"),
		model.NewPermission("GET", "/assets/paging"),
	),
	model.NewMenu("job-run", "执行", "job",
		model.NewPermission("POST", "/jobs/:id/exec"),
	),
	model.NewMenu("job-change-status", "开启/关闭", "job"),
	model.NewMenu("job-log", "日志", "job",
		model.NewPermission("GET", "/jobs/:id/logs/paging"),
	),

	model.NewMenu("storage", "磁盘空间", "ops",
		model.NewPermission("GET", "/storages/paging"),
	),
	model.NewMenu("storage-add", "增加", "storage",
		model.NewPermission("POST", "/storages"),
	),
	model.NewMenu("storage-del", "删除", "storage",
		model.NewPermission("DELETE", "/storages/:id"),
	),
	model.NewMenu("storage-edit", "修改", "storage",
		model.NewPermission("PUT", "/storages/:id"),
	),
	model.NewMenu("storage-browse", "浏览", "storage",
		model.NewPermission("GET", "/storages/:id/ls"),
	),
	model.NewMenu("storage-browse-download", "下载", "storage-browse",
		model.NewPermission("GET", "/storages/:id/download"),
	),
	model.NewMenu("storage-browse-upload", "上传", "storage-browse",
		model.NewPermission("POST", "/storages/:id/upload"),
	),
	model.NewMenu("storage-browse-mkdir", "创建文件夹", "storage-browse",
		model.NewPermission("POST", "/storages/:id/mkdir"),
	),
	model.NewMenu("storage-browse-rm", "删除", "storage-browse",
		model.NewPermission("POST", "/storages/:id/rm"),
	),
	model.NewMenu("storage-browse-rename", "重命名", "storage-browse",
		model.NewPermission("POST", "/storages/:id/rename"),
	),
	model.NewMenu("storage-browse-edit", "编辑", "storage-browse",
		model.NewPermission("POST", "/storages/:id/edit"),
	),

	model.NewMenu("monitoring", "系统监控", "ops",
		model.NewPermission("GET", "/overview/ps"),
	),

	model.NewMenu("security", "安全策略", "root"),

	model.NewMenu("access-security", "访问安全", "security",
		model.NewPermission("GET", "/securities/paging"),
	),
	model.NewMenu("access-security-add", "增加", "access-security",
		model.NewPermission("POST", "/securities"),
	),
	model.NewMenu("access-security-del", "删除", "access-security",
		model.NewPermission("DELETE", "/securities/:id"),
	),
	model.NewMenu("access-security-edit", "修改", "access-security",
		model.NewPermission("PUT", "/securities/:id"),
	),

	model.NewMenu("login-policy", "登录策略", "security",
		model.NewPermission("GET", "/login-policies/paging"),
	),
	model.NewMenu("login-policy-add", "增加", "login-policy",
		model.NewPermission("POST", "/login-policies"),
	),
	model.NewMenu("login-policy-del", "删除", "login-policy",
		model.NewPermission("DELETE", "/login-policies/:id"),
	),
	model.NewMenu("login-policy-edit", "修改", "login-policy",
		model.NewPermission("PUT", "/login-policies/:id"),
	),
	model.NewMenu("login-policy-detail", "详情", "login-policy",
		model.NewPermission("GET", "/login-policies/:id"),
	),
	model.NewMenu("login-policy-bind-user", "绑定用户", "login-policy-detail",
		model.NewPermission("GET", "/login-policies/:id/users/paging"),
	),
	model.NewMenu("login-policy-unbind-user", "解绑", "user-login-policy",
		model.NewPermission("DELETE", "/authorised/:id"),
	),

	model.NewMenu("identity", "用户管理", "root"),

	model.NewMenu("user", "用户管理", "identity",
		model.NewPermission("GET", "/users/paging"),
		model.NewPermission("GET", "/roles"),
	),
	model.NewMenu("user-add", "增加", "user",
		model.NewPermission("POST", "/users"),
	),
	model.NewMenu("user-del", "删除", "user",
		model.NewPermission("DELETE", "/users/:id"),
	),
	model.NewMenu("user-edit", "修改", "user",
		model.NewPermission("GET", "/users/:id"),
		model.NewPermission("PUT", "/users/:id"),
	),
	model.NewMenu("user-change-password", "修改密码", "user",
		model.NewPermission("POST", "/users/:id/change-password"),
	),
	model.NewMenu("user-reset-totp", "重置双因素认证", "user",
		model.NewPermission("POST", "/users/:id/reset-totp"),
	),
	model.NewMenu("user-detail", "用户详情", "user",
		model.NewPermission("GET", "/users/:id"),
		model.NewPermission("GET", "/authorised/assets/paging"),
	),
	model.NewMenu("user-authorised-asset", "授权的资产", "user-detail",
		model.NewPermission("GET", "/authorised/assets/paging"),
	),
	model.NewMenu("user-bind-asset", "授权", "user-authorised-asset",
		model.NewPermission("GET", "/authorised/selected"),
		model.NewPermission("GET", "/assets"),
		model.NewPermission("GET", "/strategies"),
		model.NewPermission("GET", "/command-filters"),
	),
	model.NewMenu("user-unbind-asset", "移除", "user-authorised-asset",
		model.NewPermission("DELETE", "/authorised/:id"),
	),
	model.NewMenu("user-login-policy", "登录策略", "user-detail",
		model.NewPermission("GET", "/login-policies/paging", "userId"),
	),
	model.NewMenu("user-unbind-login-policy", "解绑", "user-login-policy",
		model.NewPermission("DELETE", "/authorised/:id"),
	),

	model.NewMenu("role", "角色管理", "identity",
		model.NewPermission("GET", "/roles/paging"),
	),
	model.NewMenu("role-add", "增加", "role",
		model.NewPermission("POST", "/roles"),
	),
	model.NewMenu("role-del", "删除", "role",
		model.NewPermission("DELETE", "/roles/:id"),
	),
	model.NewMenu("role-edit", "修改", "role",
		model.NewPermission("GET", "/roles/:id"),
		model.NewPermission("PUT", "/roles/:id"),
	),
	model.NewMenu("role-detail", "详情", "role",
		model.NewPermission("GET", "/roles/:id"),
		model.NewPermission("GET", "/menus"),
	),

	model.NewMenu("user-group", "用户组管理", "identity",
		model.NewPermission("GET", "/user-groups/paging"),
	),
	model.NewMenu("user-group-add", "增加", "user-group",
		model.NewPermission("POST", "/user-groups"),
	),
	model.NewMenu("user-group-del", "删除", "user-group",
		model.NewPermission("DELETE", "/user-groups:/id"),
	),
	model.NewMenu("user-group-edit", "修改", "user-group",
		model.NewPermission("GET", "/user-groups/:id"),
		model.NewPermission("PUT", "/user-groups/:id"),
	),
	model.NewMenu("user-group-detail", "详情", "user-group",
		model.NewPermission("GET", "/user-groups/:id"),
	),
	model.NewMenu("user-group-authorised-asset", "授权的资产", "user-group",
		model.NewPermission("GET", "/authorised/assets/paging"),
	),
	model.NewMenu("user-group-bind-asset", "授权", "user-group-authorised-asset",
		model.NewPermission("GET", "/authorised/selected"),
		model.NewPermission("GET", "/assets"),
		model.NewPermission("GET", "/strategies"),
		model.NewPermission("GET", "/command-filters"),
	),
	model.NewMenu("user-group-unbind-asset", "移除", "user-group-authorised-asset",
		model.NewPermission("DELETE", "/authorised/:id"),
	),

	model.NewMenu("authorised", "授权策略", "root"),

	model.NewMenu("command-filter", "命令过滤", "authorised",
		model.NewPermission("GET", "/command-filters/paging"),
	),
	model.NewMenu("command-filter-add", "增加", "command-filter",
		model.NewPermission("POST", "/command-filters"),
	),
	model.NewMenu("command-filter-del", "删除", "command-filter",
		model.NewPermission("DELETE", "/command-filters/:id"),
	),
	model.NewMenu("command-filter-edit", "编辑", "command-filter",
		model.NewPermission("GET", "/command-filters/:id"),
		model.NewPermission("PUT", "/command-filters/:id"),
	),
	model.NewMenu("command-filter-detail", "详情", "command-filter",
		model.NewPermission("GET", "/command-filters/:id"),
	),
	model.NewMenu("command-filter-rule", "规则", "command-filter-detail",
		model.NewPermission("GET", "/command-filter-rules/:id"),
	),
	model.NewMenu("command-filter-rule-add", "增加", "command-filter-rule",
		model.NewPermission("POST", "/command-filter-rules"),
	),
	model.NewMenu("command-filter-rule-put", "修改", "command-filter-rule",
		model.NewPermission("GET", "/command-filter-rules/:id"),
		model.NewPermission("PUT", "/command-filter-rules/:id"),
	),
	model.NewMenu("command-filter-rule-del", "删除", "command-filter-rule",
		model.NewPermission("DELETE", "/command-filter-rules/:id"),
	),

	model.NewMenu("strategy", "授权策略", "authorised",
		model.NewPermission("GET", "/strategies/paging"),
	),
	model.NewMenu("strategy-add", "增加", "strategy",
		model.NewPermission("POST", "/strategies"),
	),
	model.NewMenu("strategy-edit", "修改", "strategy",
		model.NewPermission("GET", "/strategies/:id"),
		model.NewPermission("PUT", "/strategies/:id"),
	),
	model.NewMenu("strategy-del", "删除", "strategy",
		model.NewPermission("DELETE", "/strategies/:id"),
	),
	model.NewMenu("strategy-detail", "详情", "strategy",
		model.NewPermission("GET", "/strategies/:id"),
	),

	model.NewMenu("setting", "系统设置", "root",
		model.NewPermission("GET", "/properties"),
		model.NewPermission("PUT", "/properties"),
		model.NewPermission("POST", "/ldap-user-sync"),
		model.NewPermission("GET", "/license"),
		model.NewPermission("POST", "/license"),
		model.NewPermission("GET", "/license/machine-id"),
	),
	model.NewMenu("info", "个人中心", "root"),
}
