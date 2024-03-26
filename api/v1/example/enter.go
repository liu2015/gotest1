package example

import "ginserver/service"

type ApiGroup struct {
	CustomerApi
	FileUploadAndDownloadApi
}

var (
	customerService              = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
)
