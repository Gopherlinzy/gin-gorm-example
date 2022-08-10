package v1

import (
	"fmt"
	"net/http"

	"github.com/Gopherlinzy/go-gin-example/models"
	"github.com/Gopherlinzy/go-gin-example/pkg/e"
	"github.com/Gopherlinzy/go-gin-example/pkg/setting"
	"github.com/Gopherlinzy/go-gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetTags godoc
// @Summary      获取多个文章标签
// @Tags         tags
// @Accept       json
// @Produce      json
// @Success      200  	{string} json   "{"code":200,"data":{},"msg":"ok"}"
// @Router       /api/v1/tags [get]
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddTag godoc
// @Summary      新增文章标签
// @Tags         tags
// @Accept       json
// @Produce      json
// @Param        name   	query      	string  	true  "Name"
// @Param   	 state      query    	int    		false "State"
// @Param   	 created_by query    	string    	false "CreatedBy"
// @Success      200  	{string} json   "{"code":200,"data":{},"msg":"ok"}"
// @Router       /api/v1/tags [post]
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100个字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100个字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// EditTag godoc
// @Summary      修改文章标签
// @Tags         tags
// @Accept       json
// @Produce      json
// @Param   	 id     	path    	int    		true 	"id"
// @Param        name   	query      	string  	false  	"Name"
// @Param   	 modified_by query    	string    	false 	"CreatedBy"
// @Success      200  	{string} json   "{"code":200,"data":{},"msg":"ok"}"
// @Router       /api/v1/tags{id} [put]
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	fmt.Println(id, name, modifiedBy)
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100个字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100个字符")

	fmt.Println(valid)
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}

			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// DeleteTag godoc
// @Summary      删除文章标签
// @Tags         tags
// @Accept       json
// @Produce      json
// @Param   	 id     	path    	int    		true 	"id"
// @Success      200  	{string} json   "{"code":200,"data":{},"msg":"ok"}"
// @Router       /api/v1/tags{id} [delete]
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
