package handler

import "fmt"

type CmdHandler struct {
	Response string
	Error    error
	Finished bool
}

func (C *CmdHandler) GetResponse() {
	if C.Error != nil {
		fmt.Printf("Encountered an error: %v \n", C.Error)
	} else {
		fmt.Printf("%s", C.Response)
	}

	C.Finished = true
}

func (C *CmdHandler) HandlePodder() {
	C.Response = "\t P O D D E R \n\n \t\t written by N.Salong"

	C.GetResponse()
}

func (C *CmdHandler) HandleVerify() {
	C.Response, C.Error = HandleVerify()

	C.GetResponse()
}

func (C *CmdHandler) HandlePods(context, path string) {
	C.Response, C.Error = HandlePods(context, path)

	C.GetResponse()
}
