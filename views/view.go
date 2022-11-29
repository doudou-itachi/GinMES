package views

import (
	"GinMES/database"
	"GinMES/models"
	"GinMES/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// 产品CRUD 存在外键的示例 正向 one to one
func ProductCreate(c *gin.Context) {
	util_response := utils.Gin{C: c}
	code := c.PostForm("code")
	product_name := c.PostForm("product_name")
	unit_id := c.PostForm("unit_id")
	specification := c.PostForm("specification")
	remark := c.PostForm("remark")
	if len(product_name) == 0 || len(unit_id) == 0 {
		message := "产品名称或单位不能为空"
		util_response.Response(200, "", message)
		return
	}
	unit_id_int, _ := strconv.Atoi(unit_id)
	//-----------------------------------------------正向创建关联关系的数据-----one to one------------------------------------------
	product_obj := models.ProductInfo{Code: code, ProductName: product_name, Specification: specification, Remark: remark, ProductUnitInfoID: unit_id_int}
	result := database.Db.Create(&product_obj)
	if result.Error != nil {
		message := "创建失败"
		util_response.Response(http.StatusOK, message, "")

		return
	} else {
		util_response.Response(http.StatusOK, "", "")
	}
}

// 未做分页
func GetProduct(c *gin.Context) {
	util_response := utils.Gin{C: c}
	product_name := c.Query("product_name")
	var product_objects []models.ProductInfo
	where := database.Db.Where(&models.BaseModel{IsValid: 1})
	if len(product_name) != 0 {
		where = where.Where(&models.ProductInfo{ProductName: product_name})
	}
	where.Find(&product_objects)
	util_response.Response(http.StatusOK, product_objects, "")

}
func ProductUpdate(c *gin.Context) {
	util_response := utils.Gin{C: c}
	product_id, er := strconv.Atoi(c.PostForm("product_id"))
	code := c.PostForm("code")
	product_name := c.PostForm("product_name")
	unit_id, e := strconv.Atoi(c.PostForm("unit_id"))
	specification := c.PostForm("specification")
	remark := c.PostForm("remark")
	if e != nil {
		util_response.Response(http.StatusBadRequest, "", "单位ID不正确")
		return
	}
	if er != nil {
		util_response.Response(http.StatusBadRequest, "", "产品ID不正确")
		return
	}
	var product_obj models.ProductInfo
	result := database.Db.Where(&models.BaseModel{ID: product_id}).First(&product_obj)
	if result.RowsAffected == 0 {
		util_response.Response(http.StatusOK, "", "未查询到数据")
		return
	}
	product_obj.Code = code
	product_obj.ProductName = product_name
	product_obj.ProductUnitInfoID = unit_id
	product_obj.Specification = specification
	product_obj.Remark = remark
	save := database.Db.Save(&product_obj)
	if save.Error != nil {
		util_response.Response(http.StatusBadRequest, "", "更新失败")
		return
	} else {
		util_response.Response(http.StatusOK, product_obj, "")
		return
	}

}
func ProductDelete(c *gin.Context) {
	util_response := utils.Gin{C: c}
	product_id, er := strconv.Atoi(c.PostForm("product_id"))
	var product_object models.ProductInfo
	if er != nil {
		util_response.Response(http.StatusBadRequest, "", "产品ID错误")
		return
	}
	res := database.Db.Where(&models.BaseModel{ID: product_id, IsValid: 1}).First(&product_object)
	if res.RowsAffected == 0 {
		util_response.Response(http.StatusOK, "", "未查询到相关数据")
		return
	}
	product_object.IsValid = 0
	save := database.Db.Save(product_object)
	if save.Error != nil {
		util_response.Response(http.StatusBadRequest, "", "删除失败")
		return
	} else {
		util_response.Response(http.StatusOK, product_object, "")
		return
	}
}
func ProductDetail(c *gin.Context) {
	var product_object models.ProductInfo
	util_response := utils.Gin{C: c}
	product_id, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		util_response.Response(http.StatusBadRequest, "", "产品ID错误")
		return
	}
	res := database.Db.Where(&models.BaseModel{ID: product_id, IsValid: 1}).First(&product_object)
	if res.RowsAffected == 0 {
		util_response.Response(http.StatusOK, "", "未查询到数据")
		return
	} else {
		util_response.Response(http.StatusOK, product_object, "")
		return
	}

}

// 单位CRUD 存在外键的示例 反向
func UnitCreate(c *gin.Context) {
	util_response := utils.Gin{C: c}
	unit_code := c.PostForm("unit_code")
	unit_name := c.PostForm("unit_name")
	remark := c.PostForm("remark")
	if len(unit_code) == 0 || len(unit_name) == 0 {
		util_response.Response(http.StatusOK, "", "单位代码或单位名称不能为空")
		return
	}
	unit_object := models.ProductUnitInfo{Code: unit_code, UnitName: unit_name, Remark: remark}
	res := database.Db.Create(&unit_object)
	if res.Error != nil {
		util_response.Response(http.StatusOK, "", "创建失败")
		return
	} else {
		util_response.Response(http.StatusOK, unit_object, "")
		return
	}
}
func UnitUpdate(c *gin.Context) {
	util_response := utils.Gin{C: c}
	unit_id, e := strconv.Atoi(c.PostForm("unit_id"))
	unit_code := c.PostForm("unit_code")
	unit_name := c.PostForm("unit_name")
	remark := c.PostForm("remark")
	var unit_object models.ProductUnitInfo
	if e != nil {
		util_response.Response(http.StatusBadRequest, "", "ID有误")
		return
	}
	res := database.Db.Where(&models.BaseModel{ID: unit_id, IsValid: 1}).First(&unit_object)
	if res.RowsAffected == 0 {
		util_response.Response(http.StatusBadRequest, "", "未查询到数据")
	} else {
		unit_object.Code = unit_code
		unit_object.UnitName = unit_name
		unit_object.Remark = remark
		save := database.Db.Save(&unit_object)
		if save.Error != nil {
			util_response.Response(http.StatusBadRequest, "", "更新失败")
		} else {
			util_response.Response(http.StatusOK, unit_object, "")
		}
	}
}
func UnitGet(c *gin.Context) {
	unit_name := c.Query("unit_name")
	util_response := utils.Gin{C: c}
	var unit_object []models.ProductUnitInfo
	where := database.Db.Where(&models.BaseModel{IsValid: 1})
	if len(unit_name) != 0 {
		where = where.Where(&models.ProductUnitInfo{UnitName: unit_name})
	}
	where.Preload("ProductInfo").Find(&unit_object)
	if where.Error != nil {
		util_response.Response(http.StatusOK, "", "查询失败")
	}
	util_response.Response(http.StatusOK, unit_object, "")
}
func UnitDelete(c *gin.Context) {
	util_response := utils.Gin{C: c}
	unit_id, e := strconv.Atoi(c.PostForm("unit_id"))
	if e != nil {
		util_response.Response(http.StatusOK, "", "单位ID有误")
	}
	var unit_object models.ProductUnitInfo
	// Preload加载关联数据 预加载
	res := database.Db.Where(&models.BaseModel{ID: unit_id, IsValid: 1}).Preload("ProductInfo").First(&unit_object)

	if res.Error != nil {
		util_response.Response(http.StatusOK, "", "查询失败")
	}
	if res.RowsAffected == 0 {
		util_response.Response(http.StatusOK, "", "未查询到数据")
	}

	//	物理删除
	//database.Db.Delete(&unit_object) 存在引用关系无法删除
	// 先clear清楚外键引用，在物理删除数据
	database.Db.Model(&unit_object).Association("ProductInfo").Clear()
	database.Db.Delete(&unit_object)
	util_response.Response(http.StatusOK, "", "")
}

// 工序CRUD
func WorkProcessCreate(c *gin.Context) {
	util_response := utils.Gin{C: c}
	process_code := c.PostForm("code")
	process_name := c.PostForm("process_name")
	remark := c.PostForm("remark")
	if len(process_code) == 0 || len(process_name) == 0 {
		util_response.Response(http.StatusOK, "", "工序代码或工序名称必须输入")
		return
	}
	process_obj := models.WorkProcessInfo{Code: process_code, WorkProcessName: process_name, Remark: remark}
	res := database.Db.Create(&process_obj)
	if res.Error != nil {
		util_response.Response(http.StatusOK, "", "创建失败")
		return
	}
	util_response.Response(http.StatusOK, process_obj, "")
	return
}
func WorkProcessupdate(c *gin.Context) {
	util_response := utils.Gin{C: c}
	process_id, e := strconv.Atoi(c.PostForm("process_id"))
	process_name := c.PostForm("process_name")
	remark := c.PostForm("remark")
	var process_object models.WorkProcessInfo
	if len(process_name) == 0 {
		util_response.Response(http.StatusBadRequest, "", "工序名称必须输入")
		return
	}
	if e != nil {
		util_response.Response(http.StatusBadRequest, "", "工序ID错误")
		return
	}
	res := database.Db.Where(&models.BaseModel{ID: process_id, IsValid: 1}).First(&process_object)
	if res.RowsAffected == 0 {
		util_response.Response(http.StatusOK, "", "未查询到数据")
		return
	} else if res.Error != nil {
		util_response.Response(http.StatusBadRequest, "", "查询失败")
		return
	}
	process_object.WorkProcessName = process_name
	process_object.Remark = remark
	save := database.Db.Save(&process_object)
	if save.Error != nil {
		util_response.Response(http.StatusInternalServerError, "", "保存失败")
		return
	}
	if save.RowsAffected == 0 {
		util_response.Response(http.StatusInternalServerError, "", "保存失败")
		return
	}
	util_response.Response(http.StatusOK, process_object, "")
	return
}
func WorkProcessGet(c *gin.Context) {
	util_response := utils.Gin{C: c}
	process_code := c.Query("code")
	process_name := c.Query("process_name")
	var process_object []models.WorkProcessInfo
	where := database.Db.Where(&models.BaseModel{IsValid: 1})
	if len(process_code) != 0 {
		where = where.Where(&models.WorkProcessInfo{Code: process_code})
	}
	if len(process_name) != 0 {
		where = where.Where(&models.WorkProcessInfo{WorkProcessName: process_name})
	}
	res := where.Find(&process_object)
	if res.Error != nil {
		util_response.Response(http.StatusBadRequest, "", "查询异常")
		return
	} else if res.RowsAffected == 0 {
		util_response.Response(http.StatusOK, "", "未查询到数据")
		return
	}
	util_response.Response(http.StatusOK, process_object, "")
	return
}
func WorkProcessDelete(c *gin.Context) {
	util_response := utils.Gin{C: c}
	var process_object models.WorkProcessInfo
	process_id, e := strconv.Atoi(c.PostForm("process_id"))
	if e != nil {
		util_response.Response(http.StatusBadRequest, "", "工序id错误")
		return
	}
	res := database.Db.Where(&models.BaseModel{ID: process_id, IsValid: 1}).First(&process_object)
	if res.Error != nil {
		util_response.Response(http.StatusInternalServerError, "", "查询失败")
	}
	if res.RowsAffected == 0 {
		util_response.Response(http.StatusOK, "", "未查询到数据")
	}
	database.Db.Model(&process_object).Association("WorkStationInfo").Clear()
	database.Db.Delete(&process_object)
	util_response.Response(http.StatusOK, process_object, "")
}

// 工艺CRUD many to many
func WorkCraftCreate(c *gin.Context) {
	util_response := utils.Gin{C: c}
	craft_code := c.PostForm("code")
	craft_name := c.PostForm("craft_name")
	remark := c.PostForm("remark")
	process_id_slice := strings.Split(c.PostForm("process_ids"), ",")
	process_id, err := utils.SliceStrToSliceInt(process_id_slice)
	if err != nil {
		util_response.Response(http.StatusOK, "", "工序信息有误")
		return
	}
	if len(craft_code) == 0 {
		util_response.Response(http.StatusOK, "", "工艺代码未传入")
		return
	}
	if len(craft_name) == 0 {
		util_response.Response(http.StatusOK, "", "工艺名称未传入")
		return
	}
	//-----------------------------------------------------------------------------------------------------------------
	// 多对多创建数据
	// by条件查询需要关联的对象实例
	//var work_craft_object models.WorkCraftInfo
	var process_objects []models.WorkProcessInfo
	database.Db.Where("ID IN ?", process_id).Find(&process_objects)
	// 创建主表的实例对象
	work_craft_object := models.WorkCraftInfo{WorkCraftName: craft_name, Code: craft_code, Remark: remark}
	//database.Db.Create(&work_craft_object)
	// 使用Association根据字段Append字表的实例对象，会自动向中间表创建关系
	//  INSERT INTO `work_craft_process` (`work_craft_info_id`,`work_process_info_id`) VALUES (8,1) ON DUPLICATE KEY UPDATE `work_craft_info_id`=`work_craft_info_id`
	//database.Db.Model(&work_craft_object).Association("WorkProcessInfos").Append(process_objects)
	database.Db.Create(&work_craft_object).Association("WorkProcessInfos").Append(process_objects)
	util_response.Response(http.StatusOK, work_craft_object, "")
}
func WorkCraftGET(c *gin.Context) {
	util_response := utils.Gin{C: c}
	craft_code := c.Query("code")
	craft_name := c.Query("craft_name")
	var craft_object []models.WorkCraftInfo
	res := database.Db.Where(&models.BaseModel{IsValid: 1})
	if len(craft_code) != 0 {
		res = res.Where(&models.WorkCraftInfo{Code: craft_code})
	}
	if len(craft_name) != 0 {
		res = res.Where(&models.WorkCraftInfo{WorkCraftName: craft_name})
	}
	res = res.Preload("WorkProcessInfos").Find(&craft_object)

	if res.Error != nil {
		util_response.Response(http.StatusBadRequest, "", "查询错误")
		return
	} else if res.RowsAffected == 0 {
		util_response.Response(http.StatusOK, "", "未查询到数据")
		return
	}
	util_response.Response(http.StatusOK, craft_object, "")
	return
}
func WorkCraftupdate(c *gin.Context) {
	work_craft_id, e := strconv.Atoi(c.PostForm("work_craft_id"))
	util_response := utils.Gin{C: c}
	craft_name := c.PostForm("craft_name")
	remark := c.PostForm("remark")
	process_id_slice := strings.Split(c.PostForm("process_ids"), ",")
	process_id, err := utils.SliceStrToSliceInt(process_id_slice)
	if err != nil {
		util_response.Response(http.StatusBadRequest, "", "参数格式错误")
	}
	if e != nil {
		util_response.Response(http.StatusBadRequest, "", "参数id错误")
	}
	var craft_object models.WorkCraftInfo
	database.Db.Where(&models.BaseModel{ID: work_craft_id, IsValid: 1}).Preload("WorkProcessInfos").First(&craft_object)
	craft_object.WorkCraftName = craft_name
	craft_object.Remark = remark
	var process_object []models.WorkProcessInfo
	database.Db.Where("ID IN ?", process_id).Find(&process_object)
	//多对多更新 1.查询到关联数据的实例对象和主表的实例对象，使用Model(&craft_object).Association("WorkProcessInfos").Replace(process_object)更新
	database.Db.Model(&craft_object).Association("WorkProcessInfos").Replace(process_object)
	util_response.Response(http.StatusOK, craft_object, "")

}

func WorkCraftDelete(c *gin.Context) {
	util_response := utils.Gin{C: c}
	work_craft_id, e := strconv.Atoi(c.PostForm("work_craft_id"))
	if e != nil {
		util_response.Response(http.StatusOK, "", "参数错误")
		return
	}
	var craft_object models.WorkCraftInfo
	res := database.Db.Where(&models.BaseModel{ID: work_craft_id, IsValid: 1}).First(&craft_object)
	if res.Error != nil {
		util_response.Response(http.StatusOK, "", "查询失败")
		return
	} else if res.RowsAffected == 0 {
		util_response.Response(http.StatusOK, "", "未查询到数据")
		return
	}
	craft_object.IsValid = 0
	// 伪删除后清空多对多中间表与此数据相关的关联关系数据
	database.Db.Model(&craft_object).Association("WorkProcessInfos").Clear()
	r := database.Db.Save(&craft_object)
	if r.Error != nil || r.RowsAffected == 0 {
		util_response.Response(http.StatusOK, "", "保存失败")
		return
	}

	util_response.Response(http.StatusOK, craft_object, "")
	return
}

func Login(c *gin.Context) {
	utils_response := utils.Gin{C: c}
	username := c.PostForm("username")
	password := c.PostForm("password")
	var user_object models.Users
	res := database.Db.Where(&models.Users{Username: username}).First(&user_object)
	if res.RowsAffected != 0 {
		if user_object.Username == username && user_object.Password == utils.Sha256(password) {
			token, er := utils.GenToken(user_object)
			if er != nil {
				utils_response.Response(-1, "", "token签发错误")
				return
			} else {
				token_map := make(map[string]interface{})
				token_map["token"] = token
				utils_response.Response(http.StatusOK, token_map, "")
				return
			}
		} else {
			utils_response.Response(http.StatusOK, "", "账号或密码错误")
			return
		}
	} else {
		utils_response.Response(http.StatusOK, "", "不存在该用户")
		return
	}
}

func LineCreate(c *gin.Context) {
	utils_response := utils.Gin{C: c}
	var linecreatebind utils.LineCreateBind
	var line_object models.LineInfo
	err := c.ShouldBind(&linecreatebind)
	if err != nil {
		utils_response.Response(-1, "", "获取参数失败")
		return
	}
	line_object.Code = linecreatebind.LineCode
	line_object.LineName = linecreatebind.LineName
	line_object.Reamrk = linecreatebind.Remark
	res := database.Db.Create(&line_object)
	if res.Error != nil && res.RowsAffected == 0 {
		utils_response.Response(-1, "", "创建失败")
		return
	}
	utils_response.Response(0, line_object, "")
	return
}
func LineGet(c *gin.Context) {
	utils_response := utils.Gin{C: c}
	var linegetbind utils.LineGetBind
	var line_object []models.LineInfo
	err := c.ShouldBind(&linegetbind)
	if err != nil {
		utils_response.Response(-1, "", "获取参数失败")
		return
	}
	qs := database.Db.Model(&line_object).Where(&models.BaseModel{IsValid: 1})
	if len(linegetbind.LineCode) != 0 {
		qs = qs.Where(&models.LineInfo{Code: linegetbind.LineCode})
	}
	if len(linegetbind.LineName) != 0 {
		qs = qs.Where(&models.LineInfo{LineName: linegetbind.LineName})
	}
	res := qs.Preload("WorkStationInfo").Find(&line_object)
	if res.Error != nil || res.RowsAffected == 0 {
		utils_response.Response(-1, "", "查询失败")
		return
	}
	utils_response.Response(0, line_object, "")
	return
}

func LineUpdate(c *gin.Context) {
	utils_response := utils.Gin{C: c}
	var linebind utils.LineUpdateBind
	var line_object models.LineInfo
	err := c.ShouldBind(&linebind)
	if err != nil {
		utils_response.Response(-1, "", "获取参数失败")
		return
	}
	res := database.Db.Model(&line_object).Where(&models.BaseModel{IsValid: 1, ID: linebind.LineId}).First(&line_object)
	if res.Error != nil || res.RowsAffected == 0 {
		utils_response.Response(-1, "", "查询失败或未查询到相关数据")
		return
	}
	if len(linebind.LineCode) != 0 {
		line_object.Code = linebind.LineCode
	}
	if len(linebind.LineName) != 0 {
		line_object.LineName = linebind.LineName
	}
	if len(linebind.Remark) != 0 {
		line_object.Reamrk = linebind.Remark
	}
	res2 := database.Db.Save(&line_object)
	if res2.Error != nil || res2.RowsAffected == 0 {
		utils_response.Response(-1, "", "更新失败")
		return
	}
	utils_response.Response(0, line_object, "")
	return
}
func LineDelete(c *gin.Context) {
	utils_response := utils.Gin{C: c}
	var linebind utils.LineUpdateBind
	var line_object models.LineInfo
	err := c.ShouldBind(&linebind)
	if err != nil {
		utils_response.Response(-1, "", "获取参数失败")
		return
	}
	res := database.Db.Model(&line_object).Where(&models.BaseModel{ID: linebind.LineId, IsValid: 1}).First(&line_object)
	if res.Error != nil || res.RowsAffected == 0 {
		utils_response.Response(-1, "", "查询失败或未查询到相关数据")
		return
	}
	line_object.IsValid = 0
	res2 := database.Db.Save(&line_object)
	if res2.Error != nil || res2.RowsAffected == 0 {
		utils_response.Response(-1, "", "删除失败")
		return
	}
	utils_response.Response(0, line_object, "")
	return
}

// 工位
func StationCreate(c *gin.Context) {
	utils_response := utils.Gin{C: c}
	var stationbind utils.StationBind
	var station_object models.WorkStationInfo
	err := c.ShouldBind(&stationbind)
	if err != nil {
		utils_response.Response(-1, "", "获取参数失败")
		return
	} else if len(stationbind.StationCode) == 0 || len(stationbind.StationName) == 0 {
		utils_response.Response(-1, "", "工位代码或工位名称不能为空")
		return
	} else if stationbind.LineId == 0 {
		utils_response.Response(-1, "", "产线id不能为空")
		return
	}
	station_object.Code = stationbind.StationCode
	station_object.WorkStationName = stationbind.StationName
	station_object.Remark = stationbind.Remark
	station_object.LineInfoID = stationbind.LineId
	station_object.WorkProcessInfoID = stationbind.ProcessId

	res := database.Db.Create(&station_object)
	if res.Error != nil || res.RowsAffected == 0 {
		utils_response.Response(-1, "", "创建失败")
		return
	}
	utils_response.Response(0, station_object, "")
	return
}

func StationGet(c *gin.Context) {
	utils_response := utils.Gin{C: c}
	var stationbind utils.StationBind
	var station_object []models.WorkStationInfo
	err := c.ShouldBind(&stationbind)
	if err != nil {
		utils_response.Response(-1, "", "获取参数失败")
		return
	}
	res := database.Db.Model(&station_object).Where(&models.BaseModel{IsValid: 1})
	if len(stationbind.StationCode) != 0 {
		res = res.Where(&models.WorkStationInfo{Code: stationbind.StationCode})
	}
	if len(stationbind.StationName) != 0 {
		res = res.Where(&models.WorkStationInfo{WorkStationName: stationbind.StationName})
	}
	res = res.Find(&station_object)
	if res.Error != nil || res.RowsAffected == 0 {
		utils_response.Response(-1, "", "未查询导数据或查询失败")
		return
	}
	utils_response.Response(0, station_object, "")
	return
}

func StationPut(c *gin.Context) {
	utils_response := utils.Gin{C: c}
	var station_object models.WorkStationInfo
	var stationbind utils.StationBind
	res := c.ShouldBind(&stationbind)
	if res != nil {
		utils_response.Response(-1, "", "获取参数失败")
		return
	} else if stationbind.StationId == 0 {
		utils_response.Response(-1, "", "id参数未传入")
		return
	} else if stationbind.LineId == 0 {
		utils_response.Response(-1, "", "产线id参数未传入")
		return
	}
	res2 := database.Db.Model(&station_object).Where(&models.BaseModel{ID: stationbind.StationId, IsValid: 1}).First(&station_object)
	if res2.Error != nil || res2.RowsAffected == 0 {
		utils_response.Response(-1, "", "查询失败或未查询到数据")
		return
	}
	if len(stationbind.StationCode) != 0 {
		station_object.Code = stationbind.StationCode
	}
	if len(stationbind.StationName) != 0 {
		station_object.WorkStationName = stationbind.StationName
	}
	if len(stationbind.Remark) != 0 {
		station_object.Remark = stationbind.Remark
	}
	station_object.LineInfoID = stationbind.LineId
	//database.Db.Model(&station_object).Association("LineInfoID").Clear()
	//database.Db.Model(&station_object).Association("LineInfoID").Replace(stationbind.LineId)
	if stationbind.ProcessId == 0 {
		database.Db.Model(&station_object).Association("WorkProcessInfoID").Clear()
	} else {
		station_object.WorkProcessInfoID = stationbind.ProcessId
	}

	//database.Db.Model(&station_object).Association("WorkProcessInfoID").Clear()
	//database.Db.Model(&station_object).Association("WorkProcessInfoID").Replace(stationbind.ProcessId)
	res3 := database.Db.Save(&station_object)
	if res3.Error != nil || res3.RowsAffected == 0 {
		utils_response.Response(-1, "", "保存失败")
		return
	}
	utils_response.Response(0, station_object, "")
	return
}
func StationDelete(c *gin.Context) {
	utils_response := utils.Gin{C: c}
	var station_object models.WorkStationInfo
	var stationbind utils.StationBind
	err := c.ShouldBind(&stationbind)
	if err != nil {
		utils_response.Response(-1, "", "获取参数失败")
		return
	}
	if stationbind.StationId == 0 {
		utils_response.Response(-1, "", "id未传入")
		return
	}
	res := database.Db.Model(&station_object).Where(&models.BaseModel{ID: stationbind.StationId, IsValid: 1}).First(&station_object)
	if res.Error != nil || res.RowsAffected == 0 {
		utils_response.Response(-1, "", "未查询到数据")
		return
	}
	station_object.IsValid = 0
	res2 := database.Db.Save(&station_object)
	if res2.Error != nil || res2.RowsAffected == 0 {
		utils_response.Response(-1, "", "删除失败")
		return
	}
	utils_response.Response(0, station_object, "")
	return
}
