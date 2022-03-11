package config

import "github.com/casbin/casbin/v2"

func LoadDefaultPolicy(enforcer *casbin.Enforcer) {
	/**
	*
	 */
	// add policy
	// subject, object, action
	// subject is a level
	// object is a path
	// action as read or write
	if hasPolicy := enforcer.HasPolicy("admin", "/api/users", "read"); !hasPolicy {
		enforcer.AddPolicy("admin", "/api/users", "read")
	}
	if hasPolicy := enforcer.HasPolicy("admin", "/api/users", "write"); !hasPolicy {
		enforcer.AddPolicy("admin", "/api/users", "write")
	}
}
