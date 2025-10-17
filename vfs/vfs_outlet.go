package vfs

type UploadService interface {
	Upload() (string, error)
}

