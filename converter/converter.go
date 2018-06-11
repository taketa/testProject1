package main

import (
	"os"
	"fmt"
	"strconv"
	"net/http"
	"strings"
	"io/ioutil"
	"encoding/json"
)


func main() {

	args := os.Args
	if len(args)!=3{
		fmt.Println("The number of the command-line arguments must be equal to 2")
		fmt.Println("Example of the correct request: ./converter 10.5 UAH")
		return
	}
	num, err := strconv.ParseFloat(args[1], 64)
	if err!=nil{
		fmt.Println("The first argument must be a number")
		fmt.Println("Example of correct request: ./converter 10.5 UAH")
		return
	}
	currency:=strings.ToLower(args[2])
	if currency!="uah" && currency!="eur" && currency!="gbp"{
		fmt.Println("The second argument must be the currency: UAH, EUR or GBP")
		fmt.Println("Example of the correct request: ./converter 10.5 UAH")
		return
	}

	fmt.Printf("%v USD = %v %v\n",num, convert(num,currency),strings.ToUpper(args[2]))

}
func convert(num float64,currency string) float64{
	type Value struct{
		Val float64
	}
	var resJson map[string]Value
	res,err:= http.Get("https://free.currencyconverterapi.com/api/v5/convert?q=USD_"+currency+"&compact=y")
	if err!=nil{
		fmt.Println(err)
	}
	bs, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err := json.Unmarshal(bs, &resJson); err != nil {
		panic(err)
	}
	return resJson["USD_"+strings.ToUpper(currency)].Val*num
}
