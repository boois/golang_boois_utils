package boois_error

import "strconv"

type BooisErr struct{
	Debug string
	Msg string
	Code int
}

func New(code int,msg string,debug string) BooisErr {
	return BooisErr{
		Msg:msg,
		Code:code,
		Debug:debug,
	}

}


func (this *BooisErr) Error() string {
	return this.Msg+"("+strconv.Itoa(this.Code)+")"
}

