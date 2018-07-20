package auth
/*

record models app_label and model's name
 */


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


 func (contenttype *ContentType) String() string {
 	return "ContentType:" + contenttype.AppLabel + "_" + contenttype.Model
 }


