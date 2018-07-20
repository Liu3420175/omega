package models




type Country struct {
    Base
	Iso_3166_1_a2        string      `orm:"size(2);column(iso_3166_1_a2);null"`
	Iso_3166_1_a3        string      `orm:"size(3);column(iso_3166_1_a3);null"`
	Iso_3166_1_numeric   string      `orm:"size(3);column(Iso_3166_1_numeric);null"`

	//The commonly used name; e.g. 'United Kingdom'
    PintableName         string      `orm:"size(128)"`
	// The full official name of a country
	// e.g. 'United Kingdom of Great Britain and Northern Ireland'
	//Official name
    Name                 string     `orm:"size(128)"`
    //Chinese name
	NameCn               string     `orm:"size(128);null"`
	//Higher the number, higher the country in the list
	DisplayOrder         int8       `orm:"default(0)"`

	IsActive             bool       `orm:"default(true)"`
}


type AbstractAdress struct {
	
}
