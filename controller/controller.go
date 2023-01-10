package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"peanut/pkg/apierrors"
	"peanut/pkg/arrays"
	"strconv"
	"strings"
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
