package backend

import (
	"omega/auth"
)

type Requester struct {
	auth.Requester
}


func (request *Requester)Prepare() {

	//fmt.Println("Prepreperper")
	if request.HasLogin {
		request.ProcessRequest()
		request.AdminLoginRequired()
	}
}