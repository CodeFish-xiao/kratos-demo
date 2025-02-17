package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"kratos-demo/api/file_upload"
)

var _ file_upload.FileUploadServiceServer = (*FileUploadService)(nil)

type FileUploadService struct {
	file_upload.UnimplementedFileUploadServiceServer
	log *log.Helper
}

func (f FileUploadService) UploadFile(ctx context.Context, request *file_upload.UploadFileRequest) (*file_upload.UploadFileReply, error) {
	req, ok := http.RequestFromServerContext(ctx)
	if !ok {
		return nil, errors.New("http.RequestFromContext error")
	}
	// 设置内存大小
	err := req.ParseMultipartForm(32 << 20)
	if err != nil {
		return nil, err
	}
	fil := req.MultipartForm.File["file_content"]
	if len(fil) == 0 {
		return nil, errors.New("file_content is empty")
	}
	for _, header := range fil {
		// 打开文件
		file, err := header.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}
	return &file_upload.UploadFileReply{}, nil
}

func NewFileUploadService() *FileUploadService {
	return &FileUploadService{}
}
