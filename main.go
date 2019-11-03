package main

import "sort"

func main()  {
	aaa := "aaa,bbb,ccc,ddd,eee,1111,2222,3333"
	orderSlice(aaa)
}


func orderSlice (sourceSlice []string) []string {
	sort.Strings(sourceSlice)
	return sourceSlice
	
}
