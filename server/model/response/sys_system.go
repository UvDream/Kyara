package response

import "server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
