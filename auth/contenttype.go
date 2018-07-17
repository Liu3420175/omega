package auth
/*

record models app_label and model's name
 */
import (
	"../conf"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

 type ContentType struct {
	 Id           int64          `orm:"pk;auto"`
	 AppLabel     string         `orm:"size(100)"`
     Model        string         `orm:"size(100)"`
 }


 func (contenttype *ContentType) TableName() string {
 	// table name
 	return "begoo_content_type"
 }


 func (contenttype *ContentType) TableUnique() [][]string {
 	return [][]string{
 		[]string{"AppLabel","Model"},
	}
 }

 func (contenttype *ContentType) natural_key() []string {

 	return []string{contenttype.AppLabel,contenttype.Model}
 }



 func init() {

 	databases := conf.DATABASES
 	database := databases["default"]
 	orm.RegisterDriver("mysql",orm.DRMySQL)
 	connent := database["USER"] + ":" + database["PASSWORD"] + "@tcp(" + database["HOST"] + ":" + database["PORT"] + ")/" + database["NAME"] + "?charset=utf8"
 	orm.RegisterDataBase("DB"+database["NAME"],"mysql",connent)
 	orm.RegisterModel(
 		new(ContentType),
	)
 }