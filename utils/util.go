package utils

import (
	"crypto/sha256"
	"encoding/hex"
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
	/*
		type:[]string[1 2 3] --> type:[]int[1 2 3]
	*/
	new_slice := make([]int, len(old_slice))
	for index, val := range old_slice {
		v, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		//new_slice = append(new_slice, v)
		new_slice[index] = v
	}
	return new_slice, nil
}

// Sha256 Sha256加密
func Sha256(src string) string {
	m := sha256.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
