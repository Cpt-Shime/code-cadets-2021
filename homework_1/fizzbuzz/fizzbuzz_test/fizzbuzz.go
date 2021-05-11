package fizzbuzz

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)



func Fizzbuzzgame(start int , end int) ([]string, error){

	if start < 0 {
		return nil, errors.New("Start is bigger then end")

	}
	if end < 0 {
		return nil, errors.New("End is less then 0")

	}
	if end < start {
		return nil, errors.New("End is lesser then start")
	}
	var fizbuzzlist []string
	for i:=start; i<= end; i++ {
		var fizbuzz = ""

		if i%3==0 {
			fizbuzz+="Fizz"
		}
		if  i%5==0 {
			fizbuzz+="Buzz"
		}
		if fizbuzz!="" {

		} else {
			fizbuzz=strconv.Itoa(i)
		}
		fizbuzzlist=append(fizbuzzlist,fizbuzz)


	}
	return fizbuzzlist, nil


}


func Printgameresults(slice []string) {
	fmt.Println(strings.Join(slice, " "))
}





