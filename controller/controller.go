package controller

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"peanut/config"
	"peanut/pkg/apierrors"
	"peanut/pkg/arrays"
	"strconv"
	"strings"
	"time"
)

func bindJSON(ctx *gin.Context, obj interface{}) bool {
	err := ctx.ShouldBindJSON(obj)
	if err == nil {
		return true
	}
	_, ok := err.(validator.ValidationErrors)
	if ok {
		err = apierrors.New(apierrors.InvalidRequest, err)
	} else {
		err = apierrors.New(apierrors.BadParams, err)
	}
	ctx.Error(err).SetType(gin.ErrorTypeBind)

	return false
}

func bindQueryParams(ctx *gin.Context, obj interface{}) bool {
	err := ctx.ShouldBindQuery(obj)

	if err == nil {
		return true
	}
	_, ok := err.(validator.ValidationErrors)
	if ok {
		err = apierrors.New(apierrors.InvalidRequest, err)
	} else {
		err = apierrors.New(apierrors.BadParams, err)
	}
	ctx.Error(err).SetType(gin.ErrorTypeBind)

	return false
}
func bindForm(ctx *gin.Context, obj interface{}) bool {
	err := ctx.ShouldBind(obj)
	if err == nil {
		return true
	}
	_, ok := err.(validator.ValidationErrors)
	if ok {
		err = apierrors.New(apierrors.InvalidRequest, err)
	} else {
		err = apierrors.New(apierrors.BadParams, err)
	}
	ctx.Error(err).SetType(gin.ErrorTypeBind)

	return false
}
func checkError(ctx *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	_ = ctx.Error(err).SetType(gin.ErrorTypePublic)
	return true
}
func saveUploadedFile(ctx *gin.Context, name string, dir string) (path string, err error) {
	// Save uploaded file
	file, _ := ctx.FormFile(name)
	dst := uuid.New().String() + filepath.Ext(file.Filename)
	contentPath := dir + dst
	err = ctx.SaveUploadedFile(file, contentPath)
	if err != nil {
		err = apierrors.NewErrorf(apierrors.InternalError, err.Error())
		return
	}

	return
}
func saveUploadedFileGCS(ctx *gin.Context, name string, dir string) (path string, err error) {
	// Save uploaded file
	ct := context.Background()
	client, errGCS := storage.NewClient(ct, option.WithCredentialsFile(config.GgStorageCredential))
	if errGCS != nil {
		log.Fatalf("Failed to create client: %v", errGCS)
	}
	ct, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	f, uploadedFile, _ := ctx.Request.FormFile(name)
	dst := uuid.New().String() + filepath.Ext(uploadedFile.Filename)

	wc := client.Bucket(config.BucketName).Object(dst).NewWriter(ctx)
	if _, errCopy := io.Copy(wc, f); err != nil {
		errCopy = apierrors.NewErrorf(apierrors.InternalError, errCopy.Error())
		return
	}
	if errClose := wc.Close(); err != nil {
		errClose = apierrors.NewErrorf(apierrors.InternalError, errClose.Error())
		return
	}
	fmt.Println("-----------------------")
	url, errParse := url.Parse(config.PublicUrlGgStorage + "/" + config.BucketName + "/" + wc.Attrs().Name)
	if errParse != nil {
		err = apierrors.NewErrorf(apierrors.InternalError, errParse.Error())
		return
	}
	return url.String(), nil
	//return
}
func CheckMaxSizeUpload(size int) bool {
	maxSize, err := strconv.Atoi(os.Getenv("MAX_SIZE_UPLOAD"))
	if err != nil {
		return false
	}
	return !(size > maxSize)
}

func CheckExtensionAvailable(ctx *gin.Context, name string, listExt []string) bool {
	file, _ := ctx.FormFile(name)
	return arrays.Contains(listExt, strings.ToLower(filepath.Ext(file.Filename)))
}
