package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(code int, data interface{}, message string) {
	g.C.JSON(code, gin.H{
		"code":    code,
		"data":    data,
		"message": message,
	})
	return
}
func SliceStrToSliceInt(old_slice []string) ([]int, error) {
	new_slice := make([]int, 0)
	for _, val := range old_slice {
		v, err := strconv.Atoi(val)
		if err != nil {
			return new_slice, err
		}
		new_slice = append(new_slice, v)
	}
	return new_slice, nil
}
