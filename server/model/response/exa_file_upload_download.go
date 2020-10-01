package response

import "server/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
