package backend

import (
	"strconv"
	"strings"
	"omega/auth"
	"github.com/astaxie/beego/orm"
	"omega/common"
)

func (request *Requester) CustomerAccountList() {
    page := request.GetString("page","1")
    limit := request.GetString("limit","10")
	q := request.GetString("q","")

	page = strings.Trim(page," ")
	q = strings.Trim(q," ")
	//state = strings.Trim(state," ")
	limit = strings.Trim(limit," ")

	Page,_ := strconv.Atoi(page)
	Limit,_ := strconv.Atoi(limit)
	if Page <= 0{
		Page = 1
	}
	if Limit <= 0{
		Limit = 10
	}
	offset := (Page - 1) * Limit

	table_name := new(auth.User)
	query := orm.NewOrm().QueryTable(table_name).Filter("is_admin",true)
	query = query.Limit(Limit,offset)

	if len(q) > 0{
        query = common.GetSearchResults(query,[]string{"user_name","email","phone"},q)
	}
    var  users  []auth.User
    query.All(&users,"Id","UserName","Email","FirstName","LastName","Phone","DateJoined")
	result := map[string]interface{}{
		"page" : Page,
		"limit" : Limit,
		"infos" : users,
	}
	request.CommonResponse(0,result)
	return

}
