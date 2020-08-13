package response

import "blog-api/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
