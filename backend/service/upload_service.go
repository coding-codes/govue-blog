package service

import (
	"fmt"
	"github.com/coding-codes/utils"
	"mime/multipart"
	"path/filepath"
)

func UploadImageAvatarService(file *multipart.FileHeader, t string) (info utils.ImageAvatarInfo, err error) {

	name := filepath.Base(file.Filename)
	filename := utils.MD5ImageName(name)

	info = utils.GenImageAvatarInfo(t, filename)
	relPath := filepath.Dir(info.AvatarRelPath)
	// If the path is empty, Dir returns ".".
	if relPath == "." {
		relPath = filepath.Dir(info.ImageRelPath)
	}

	if err = utils.IsNotExistMkDir(relPath); err != nil {
		return
	}

	SetAvatarUrl(fmt.Sprintf("http://localhost:9528" + info.ImageFullUrl))

	return
}
