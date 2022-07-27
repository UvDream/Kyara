package system

import (
	"github.com/google/uuid"
	"server/model/common"
)

type SysOperatingHistory struct {
	common.Model
	UUID         uuid.UUID `json:"uuid" gorm:"comment:操作的UUID"`
	UserID       int       `json:"user_id" gorm:"comment:用户id"`
	User         SysUser   `json:"user"`
	IP           string    `json:"ip" gorm:"column:ip;comment:请求ip"`                       // 请求ip
	Method       string    `json:"method" gorm:"column:method;comment:请求方法"`               // 请求方法
	Path         string    `json:"path" gorm:"column:path;comment:请求路径"`                   // 请求路径
	Status       int       `json:"status" gorm:"column:status;comment:请求状态"`               // 请求状态
	Latency      int       `json:"latency" gorm:"column:latency;comment:延迟"`               // 延迟
	Agent        string    `json:"agent" gorm:"column:agent;comment:代理"`                   // 代理
	ErrorMessage string    `json:"error_message" gorm:"column:error_message;comment:错误信息"` // 错误信息
	Body         string    `json:"body" gorm:"column:body;comment:请求Body"`                 // 请求Body
	Resp         string    `json:"resp" gorm:"column:resp;comment:响应Body"`                 // 响应Body
}
