package io

import "shopping/utils/mongo"

type LoginRequest struct {
	Username string `json:"username"  binding:"required,username"`
	Password string `json:"password"  binding:"required,password"`
}

type GetMonitorDataReq struct {
	AgentName string   `json:"agent_name"    binding:"agent_name"`
	System    []string `json:"system"        binding:"max=3,dive,oneof=Linux Windows Mac"`
	Start     int64    `json:"start"         binding:"required,timestamp"`
	End       int64    `json:"end"           binding:"required,timestamp"`
}

type GetUserLogReq struct {
	PageSize  int `json:"page_size"  binding:"required,gte=10,lte=100"`
	PageIndex int `json:"page_index" binding:"gte=0"`
}

type GetUserLogResp struct {
	Total int64            `json:"total"`
	Logs  []*mongo.UserLog `json:"logs"`
}
