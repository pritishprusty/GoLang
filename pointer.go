package main

func main(){
	i:=10
	j:=&i
	*j=100 //this means it'll update the value of i by going into the address of i
	fmt.Println(j) //this will print the address of i
	fmt.println(i)//but this will print 100
}