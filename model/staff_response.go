package model

// ResponseListStaffInfo ...
type ResponseListStaffInfo struct {
	List []ResponseStaffInfo `json:"list"`
}

// ResponseStaffInfo ...
type ResponseStaffInfo struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}
