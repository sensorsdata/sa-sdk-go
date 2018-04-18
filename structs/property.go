package structs

type LibProperties struct {
	Lib        string `json:"$lib"`
	LibVersion string `json:"$lib_version"`
	LibMethod  string `json:"$lib_method"`
	AppVersion string `json:"$app_version,omitempty"`
	LibDetail  string `json:"$lib_detail"`
}
