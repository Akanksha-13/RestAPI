package models

type ApplicationInfo struct {
	ErpAcronym  string `json:"erp_acronym"`
	PathID      string `json:"path_id"`
	ContactInfo struct {
		Email string `json:"email"`
		Slack string `json:"slack"`
		Owner string `json:"owner"`
	} `json:"contact_info"`
}

type InstanceSpec struct {
	Memory     string            `json:"memory"`
	Type       string            `json:"type"`
	CloudType  string            `json:"cloud_type"`
	Region     string            `json:"region"`
	Tier       int               `json:"tier"`
	ValkeyName string            `json:"valkey_name"`
	Tags       map[string]string `json:"tags"`
}

type SecurityGroups struct {
	CacheInstance struct {
		Admin string `json:"admin"`
		Read  string `json:"read"`
	} `json:"cache_instance"`
	CacheData struct {
		Read        string `json:"read"`
		Contributor string `json:"contributor"`
	} `json:"cache_data"`
}

type ValkeyRequest struct {
	ApplicationInfo ApplicationInfo `json:"application_info"`
	InstanceSpec    InstanceSpec    `json:"instance_spec"`
	SecurityGroups  SecurityGroups  `json:"security_groups"`
}

type ValkeyResponse struct {
	RequestID string                 `json:"request-id"`
	Error     map[string]interface{} `json:"error"`
}
