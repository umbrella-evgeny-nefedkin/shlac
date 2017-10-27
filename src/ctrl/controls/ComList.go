package controls

import (
	"hrentabd"
	"fmt"
	"time"
	"flag"
	"errors"
	"log"
)

type ComList struct{
	Com
}

var MsgUsageComList string = "usage: \n\tlist -index <index> -ts <timestamp> \n\tlist --help\n"

func (c *ComList)Exec(Tab *hrentabd.HrenTab, args []string)  (response string, err error){

	defer func(response *string, err *error){
		if r := recover(); r!=nil{
			*err        = errors.New(fmt.Sprint(r))
			*response   = MsgUsageComAdd
		}

	}(&response, &err)

	var defaultResponse string = "empty"

	var INDEX string
	var TS int64
	var HELP bool

	Args := flag.NewFlagSet("com_list", flag.PanicOnError)
	Args.StringVar(&INDEX, "index", "", "search by index")
	Args.Int64Var(&TS, "ts", 0, "search by timestamp")
	Args.BoolVar(&HELP, "help", false, "show this help")
	Args.Parse(args)

	// Args.PrintDefaults()


	if INDEX != "" && TS != 0{
		log.Fatalln("too much argument")
	}

	if HELP{
		response = MsgUsageComList

	}else if INDEX != ""{

		if found := Tab.FindByIndex(INDEX); found != nil{
			response = fmt.Sprintln("==> ", found.TimeStart().String(), "(", found.TimeStart().Unix() ,")")
			response += fmt.Sprintln(INDEX, ":", found.Command(), "(", found.Ttl() ,")")
		}

	}else if TS != 0{

		if found := Tab.FindByTime(time.Unix(TS,0), false); found != nil{

			response = ""
			for ts, ah := range Tab.List(){

				response += fmt.Sprintln("==> ", time.Unix(ts,0).String(),"(", ts ,")")
				for index, h := range ah{
					response += fmt.Sprintln(index, ":", h.Command(), "(", h.Ttl() ,")")
				}
			}
		}

	}else{

		response = ""
		for ts, ah := range Tab.List(){

			response += fmt.Sprintln("==> ", time.Unix(ts,0).String(), "(", ts ,")")
			for index, h := range ah{
				response += fmt.Sprintln(index, ":", h.Command(), "(", h.Ttl() ,")")
			}
		}
	}


	if response == "" {
		response = defaultResponse
	}

	return response, nil
}
