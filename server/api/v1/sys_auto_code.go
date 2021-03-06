package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
	"os"
	"server/global"
	"server/global/response"
	"server/model"
	"server/service"
	"server/utils"
)

// @Tags SysApi
// @Summary 自动代码模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AutoCodeStruct true "创建自动代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/createTemp [post]
func CreateTemp(c *gin.Context) {
	var a model.AutoCodeStruct
	_ = c.ShouldBindJSON(&a)
	AutoCodeVerify := utils.Rules{
		"Abbreviation": {utils.NotEmpty()},
		"StructName":   {utils.NotEmpty()},
		"PackageName":  {utils.NotEmpty()},
		"Fields":       {utils.NotEmpty()},
	}
	WKVerifyErr := utils.Verify(a, AutoCodeVerify)
	if WKVerifyErr != nil {
		response.FailWithMessage(WKVerifyErr.Error(), c)
		return
	}
	if a.AutoCreateApiToSql {
		apiList := [6]model.SysApi{
			{
				Path:        "/" + a.Abbreviation + "/" + "create" + a.StructName,
				Description: "新增" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "POST",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName,
				Description: "删除" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "DELETE",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName + "ByIds",
				Description: "批量删除" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "DELETE",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "update" + a.StructName,
				Description: "更新" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "PUT",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "find" + a.StructName,
				Description: "根据ID获取" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "GET",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "get" + a.StructName + "List",
				Description: "获取" + a.Description + "列表",
				ApiGroup:    a.Abbreviation,
				Method:      "GET",
			},
		}
		for _, v := range apiList {
			errC := service.CreateApi(v)
			if errC != nil {
				c.Writer.Header().Add("success", "false")
				c.Writer.Header().Add("msg", url.QueryEscape(fmt.Sprintf("自动化创建失败，%v，请自行清空垃圾数据", errC)))
				return
			}
		}
	}
	err := service.CreateTemp(a)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
		os.Remove("./ginvueadmin.zip")
	} else {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "ginvueadmin.zip")) // fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.Header().Add("success", "true")
		c.File("./ginvueadmin.zip")
		os.Remove("./ginvueadmin.zip")
	}
}

// @Tags SysApi
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/getTables [get]

func GetTables(c *gin.Context) {
	dbName := c.DefaultQuery("dbName", global.GVA_CONFIG.Mysql.Dbname)
	err, tables := service.GetTables(dbName)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询table失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{
			"tables": tables,
		}, c)
	}
}

// @Tags SysApi
// @Summary 获取当前所有数据库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/getDatabase [get]
func GetDB(c *gin.Context) {
	err, dbs := service.GetDB()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询table失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{
			"dbs": dbs,
		}, c)
	}
}

// @Tags SysApi
// @Summary 获取当前表所有字段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/getDatabase [get]
func GetColume(c *gin.Context) {
	dbName := c.DefaultQuery("dbName", global.GVA_CONFIG.Mysql.Dbname)
	tableName := c.Query("tableName")
	err, columes := service.GetColume(tableName, dbName)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询table失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{
			"columes": columes,
		}, c)
	}
}
