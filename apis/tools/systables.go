package tools

import (
	"Web_Api/models/tools"
	"Web_Api/pkg"
	"Web_Api/pkg/app"
	"Web_Api/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 分页列表数据
// @Description 生成表分页列表
// @Tags 工具 - 生成表
// @Param tableName query string false "tableName / 数据表名称"
// @Param pageSize query int false "pageSize / 页条数"
// @Param pageIndex query int false "pageIndex / 页码"
// @Success 200 {object} models.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/tables/page [get]
func GetSysTableList(c *gin.Context)  {
	var data tools.SysTables
	var err error
	var pageSize = 10
	var pageIndex = 1
	if size := c.Request.FormValue("pageSize");size!=""{
		pageSize = pkg.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex");index!="" {
		pageIndex = pkg.StrToInt(err, index)
	}
	data.TableName = c.Request.FormValue("tableName")
	result, count, err := data.GetPage(pageSize, pageIndex)
	pkg.HasError(err, "", -1)
	var mp = make(map[string]interface{}, 3)
	mp["list"] = result
	mp["count"] = count
	mp["pageIndex"] = pageIndex
	mp["pageSize"] = pageSize
	var res app.Response
	res.Data = mp
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 获取配置
// @Description 获取JSON
// @Tags 工具 - 生成表
// @Param configKey path int true "configKey"
// @Success 200 {object} models.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/tables/info/{tableId} [get]
// @Security
func GetSysTables(c *gin.Context)  {
	var data tools.SysTables
	data.TableId, _ = utils.StringToInt(c.Param("tableId"))
	result, err := data.Get()
	pkg.HasError(err, "未找到相关信息",-1)
	var res app.Response
	res.Data = result
	mp := make(map[string]interface{})
	mp["rows"] = result.Columns
	mp["info"] = result
	res.Data = mp
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 添加表结构
// @Description 添加表结构
// @Tags 工具 - 生成表
// @Accept  application/json
// @Product application/json
// @Param tables query string false "tableName / 数据表名称"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/sys/tables/info [post]
// @Security Bearer
