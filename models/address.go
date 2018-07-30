package models

import (
	"omega/auth"
	"bytes"
	"strconv"
)

type Country struct {
    Base
	Iso_3166_1_a2        string       `orm:"size(2);column(iso_3166_1_a2);null"`
	Iso_3166_1_a3        string       `orm:"size(3);column(iso_3166_1_a3);null"`
	Iso_3166_1_numeric   string       `orm:"size(3);column(Iso_3166_1_numeric);null"`

	//The commonly used name; e.g. 'United Kingdom'
    PintableName         string       `orm:"size(128)"`
	// The full official name of a country
	// e.g. 'United Kingdom of Great Britain and Northern Ireland'
	//Official name
    Name                 string       `orm:"size(128)"`
    //Chinese name
	NameCn               string       `orm:"size(128);null"`
	//Higher the number, higher the country in the list
	DisplayOrder         int8         `orm:"default(0)"`

	IsActive             bool         `orm:"default(true)"`

	UserAdresses         []*UserAdress `orm:"reverse(many)"`
}


type AbstractAdress struct {
	Base
	Title                string       `orm:"null;size(64)"`
	FirstName            string       `orm:"null;size(255)"`
    LastName             string       `orm:"null;size(255)"`
    Line1                string       `orm:"null;size(255)"`
    Line2                string       `orm:"null;size(255)"`
    Line3                string       `orm:"null;size(255)"`
    City                 string       `orm:"null;size(255)"`
    State                string       `orm:"null;size(255)"`
    Postcode             string       `orm:"null;size(64)"`

}



type UserAdress struct {
	AbstractAdress
	Country                    *Country      `orm:"rel(fk);null;on_delete(set_null)"`
    User                       *auth.User    `orm:"rel(fk);null;on_delete(set_null)"`
    PhoneNumber                 string        `orm:"null;size(64)"`
    Notes                       string        `orm:"null;size(512)"`
	IsDefaultForShipping        bool          `orm:"default(false)"`
	IsDefaultForBilling         bool          `orm:"default(false)"`
	NumOrdersAsShippingAddress  int           `orm:"default(0)"`
	NumOrdersAsBillingAddress   int           `orm:"default(0)"`
	//Adress hash
	Hash                        string        `orm:"null;size(255)"`
}



func (useraddress *UserAdress) UserAddressHash() string {
	// jisuan address hash
	if useraddress.Country == nil || useraddress.User == nil{
		return ""
	}

	var buf bytes.Buffer
	address_id := strconv.FormatInt(useraddress.Country.Id,10)
	user_id := strconv.FormatInt(useraddress.User.Id,10)
	buf.WriteString(address_id)
	buf.WriteString(useraddress.State)
	buf.WriteString(useraddress.City)
	buf.WriteString(useraddress.Line1)
	buf.WriteString(useraddress.Line2)
	buf.WriteString(useraddress.Line3)
	buf.WriteString(useraddress.Postcode)
	buf.WriteString(user_id)
	return buf.String()


}
