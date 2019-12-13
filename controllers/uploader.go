package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/axetroy/go-fs"
	"github.com/gin-gonic/gin"

	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

const FIELD = "file"

// 支持的图片后缀名
var supportImageExtNames = []string{".jpg", ".jpeg", ".png", ".ico", ".svg", ".bmp", ".gif"}

/**
check a file is a image or not
*/
func isImage(extName string) bool {
	for i := 0; i < len(supportImageExtNames); i++ {
		if supportImageExtNames[i] == extName {
			return true
		}
	}
	return false
}

/**
Handler the parse error
*/
func parseFormFail(context *gin.Context) {
	context.JSON(http.StatusBadRequest, gin.H{
		"message": "Can not parse form",
	})
}

/**
Upload image handler
*/
func UploaderImage(context *gin.Context) {
	var (
		maxUploadSize = 10485760 // 最大上传大小
		distPath      string     // 最终的输出目录
		err           error
		file          *multipart.FileHeader
		src           multipart.File
		dist          *os.File
	)

	// Source
	if file, err = context.FormFile(FIELD); err != nil {
		parseFormFail(context)
		return
	}

	extname := strings.ToLower(path.Ext(file.Filename))

	if isImage(extname) == false {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Unsupport upload file type " + extname,
		})
		return
	}

	if file.Size > int64(maxUploadSize) {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Upload file too large, The max upload limit is " + strconv.Itoa(int(maxUploadSize)),
		})
		return
	}

	if src, err = file.Open(); err != nil {

	}
	defer src.Close()

	hash := md5.New()

	io.Copy(hash, src)

	md5string := hex.EncodeToString(hash.Sum([]byte("")))

	fileName := md5string + extname

	// Destination
	distPath = path.Join(Config.Path, Config.Image.Path, fileName)
	if dist, err = os.Create(distPath); err != nil {

	}
	defer dist.Close()

	// FIXME: open 2 times
	if src, err = file.Open(); err != nil {
		//
	}

	// Copy
	io.Copy(dist, src)

	context.JSON(http.StatusOK, gin.H{
		"hash":     md5string,
		"filename": fileName,
		"origin":   file.Filename,
		"size":     file.Size,
		"code":     0,
	})
}

/**
Upload file handler
*/
func UploadFile(context *gin.Context) {
	var (
		isSupportFile bool
		maxUploadSize = Config.Image.MaxSize  // 最大上传大小
		allowTypes    = Config.File.AllowType // 可上传的文件类型
		distPath      string                  // 最终的输出目录
		err           error
		file          *multipart.FileHeader
		src           multipart.File
		dist          *os.File
	)
	// Source
	if file, err = context.FormFile(FIELD); err != nil {
		parseFormFail(context)
		return
	}

	extname := path.Ext(file.Filename)

	if len(allowTypes) != 0 {
		for i := 0; i < len(allowTypes); i++ {
			if allowTypes[i] == extname {
				isSupportFile = true
				break
			}
		}

		if isSupportFile == false {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "Unsupport upload file type " + extname,
			})
			return
		}
	}

	if file.Size > int64(maxUploadSize) {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Upload file too large, The max upload limit is " + strconv.Itoa(int(maxUploadSize)),
		})
		return
	}

	if src, err = file.Open(); err != nil {
		// open the file fail...
	}
	defer src.Close()

	hash := md5.New()

	io.Copy(hash, src)

	md5string := hex.EncodeToString(hash.Sum([]byte("")))

	fileName := md5string + extname

	// Destination
	distPath = path.Join(Config.Path, Config.File.Path, fileName)
	if dist, err = os.Create(distPath); err != nil {
		// create dist file fail...
	}
	defer dist.Close()

	// FIXME: open 2 times
	if src, err = file.Open(); err != nil {
		//
	}

	// Copy
	io.Copy(dist, src)

	context.JSON(http.StatusOK, gin.H{
		"hash":     md5string,
		"filename": fileName,
		"origin":   file.Filename,
		"size":     file.Size,
		"code":     0,
	})
}

/**
Generate Upload example Template
*/
func UploaderTemplate(template string) func(context *gin.Context) {
	return func(context *gin.Context) {
		header := context.Writer.Header()
		header.Set("Content-Type", "text/html; charset=utf-8")
		context.String(200, `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Upload</title>
</head>
<body>
<form action="`+Config.UrlPrefix+`/upload/image" method="post" enctype="multipart/form-data">
  <h2>Image Upload</h2>
  <input type="file" name="file">
  <input type="submit" value="Upload">
</form>

</hr>

<form action="`+Config.UrlPrefix+`/upload/file" method="post" enctype="multipart/form-data">
  <h2>File Upload</h2>
  <input type="file" name="file">
  <input type="submit" value="Upload">
</form>

</body>
</html>
	`)
	}

}

/**
Get Origin image
*/
func GetOriginImage(context *gin.Context) {
	filename := context.Param("filename")
	originImagePath := path.Join(Config.Path, Config.Image.Path, filename)
	if fs.PathExists(originImagePath) == false {
		// if the path not found
		http.NotFound(context.Writer, context.Request)
		return
	}
	http.ServeFile(context.Writer, context.Request, originImagePath)
}

/**
Get file raw
*/
func GetFileRaw(context *gin.Context) {
	filename := context.Param("filename")
	filePath := path.Join(Config.Path, Config.File.Path, filename)
	if isExistFile := fs.PathExists(filePath); isExistFile == false {
		// if the path not found
		http.NotFound(context.Writer, context.Request)
		return
	}
	http.ServeFile(context.Writer, context.Request, filePath)
}

/**
Download a file
*/
func DownloadFile(context *gin.Context) {
	filename := context.Param("filename")
	filePath := path.Join(Config.Path, Config.File.Path, filename)
	if isExistFile := fs.PathExists(filePath); isExistFile == false {
		// if the path not found
		http.NotFound(context.Writer, context.Request)
		return
	}
	http.ServeFile(context.Writer, context.Request, filePath)
}

type FileConfig struct {
	Path      string   `valid:"required,length(1|20)"`  // 普通文件的存放目录
	MaxSize   int      `valid:"required"`               // 普通文件上传的限制大小，单位byte, 最大单位1GB
	AllowType []string `valid:"required,length(0|100)"` // 允许上传的文件后缀名
}

type ImageConfig struct {
	Path    string `valid:"required,length(1|20)"` // 图片存储路径
	MaxSize int    `valid:"required"`              // 最大图片上传限制，单位byte
}

type TConfig struct {
	Path      string `valid:"required,length(1|20)"` //文件上传的根目录
	UrlPrefix string `valid:"required,length(0|20)"` // api的url前缀
	File      FileConfig
	Image     ImageConfig
}

type Uploader struct {
	Upload   *gin.RouterGroup
	Download *gin.RouterGroup
	Config   TConfig
	Engine   *gin.Engine
}

var Config TConfig

func InitUploader(c *TConfig) (err error) {
	if err = fs.EnsureDir(c.Path); err != nil {
		return
	}
	if err = fs.EnsureDir(path.Join(c.Path, c.File.Path)); err != nil {
		return
	}
	if err = fs.EnsureDir(path.Join(c.Path, c.Image.Path)); err != nil {
		return
	}
	return
}

/**
Create Router
*/
func New(e *gin.Engine, c TConfig) (u *Uploader, err error) {
	Config = c

	var (
		isValidConfig bool
	)

	if isValidConfig, err = govalidator.ValidateStruct(c); err != nil {
		err = errors.New(`invalid uploader config: [` + err.Error() + `]`)
		return
	} else {
		if isValidConfig == false {
			err = errors.New("invalid Uploader config")
			return
		}
	}

	if err = InitUploader(&Config); err != nil {
		return
	}

	// upload all
	uploader := e.Group(Config.UrlPrefix + "/upload")
	// download all
	downloader := e.Group(Config.UrlPrefix + "/download")
	downloader.Use(func(context *gin.Context) {
		header := context.Writer.Header()
		// alone dns prefect
		header.Set("X-DNS-Prefetch-Control", "on")
		// IE No Open
		header.Set("X-Download-Options", "noopen")
		// not cache
		header.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		header.Set("Expires", "max-age=0")
		// Content Security Policy
		header.Set("Content-Security-Policy", "default-src 'self'")
		// xss protect
		// it will caught some problems is old IE
		header.Set("X-XSS-Protection", "1; mode=block")
		// Referrer Policy
		header.Set("Referrer-Header", "no-referrer")
		// cros frame, allow same origin
		header.Set("X-Frame-Options", "SAMEORIGIN")
		// HSTS
		header.Set("Strict-Transport-Security", "max-age=5184000;includeSubDomains")
		// no sniff
		header.Set("X-Content-Type-Options", "nosniff")
	})

	return &Uploader{
		Upload:   uploader,
		Download: downloader,
		Config:   c,
		Engine:   e,
	}, nil

}

/**
Resolve
*/
func (u *Uploader) Resolve() {

	// upload the file/image
	u.Upload.POST("/image", UploaderImage)
	u.Upload.POST("/file", UploadFile)

	u.Upload.GET("/example", UploaderTemplate("image"))

	// get file which upload
	uploadFile := u.Download.Group("/file")
	uploadFile.GET("/raw/:filename", GetFileRaw)
	uploadFile.GET("/download/:filename", DownloadFile)

	// get image which upload
	downloadImage := u.Download.Group("/image")
	downloadImage.GET("/origin/:filename", GetOriginImage)

	return
}
