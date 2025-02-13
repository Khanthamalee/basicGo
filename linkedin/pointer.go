package linkedin

import "fmt"

func Pointer() {
	v1 := 33
	pointer1 := &v1
	*pointer1 = *pointer1 / 11
	fmt.Println("pointer1 : ", pointer1)
	fmt.Printf("pointer1 : %T \n", pointer1)
	fmt.Println("*pointer1 : ", *pointer1)
	fmt.Println("v1 : ", v1)

	pointer2 := &v1
	*pointer2 = *pointer2 / 3

	fmt.Println("Pointer2 :", pointer2)
	fmt.Println("*Pointer2 :", *pointer2)
	fmt.Println("v1 :", v1)

}
