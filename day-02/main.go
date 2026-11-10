package main

import "fmt"

func set(store map[string]string,key string ,value string){
	store[key] = value
}
func get(store map[string]string ,key string)(string,bool){
	value,exists:=store[key]
	return value,exists
}
func del(store map[string]string,key string){
	delete(store,key)
}
func main(){
	//Arrays
	numbers:=[5] int {1,2,3,4,5}
	fmt.Println("This is an array:",numbers)
	fmt.Println("Len of the array:",len(numbers))
	fmt.Println("Index,value")
	for i,j:=range numbers{
		fmt.Println(i,j)
	}

	//Slices
	num_slices:=[]int{6,7,8,9,0}
	fmt.Println("This is a slice:",num_slices)
	fmt.Println("Size of the slice:",len(num_slices))

	//Maps
	store:=make(map[string]string)
	set(store,"name","Akhil")
	set(store,"Age","21")
	set(store,"language","GO")
	Name,exists:=get(store,"name")
	if exists{
		fmt.Println(Name)
	}
	delete(store,"Age")
	_,exists=get(store,"Age")
	if !exists{
		fmt.Println("Age deleted sucessfully")
	}
	fmt.Println(len(store))
}