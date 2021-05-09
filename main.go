package main

/*
#include<stdio.h>
#include <stdlib.h>
extern int suma(int n1,int n2);
*/
import "C"
import (
	"fmt"
)

func main() {

	s := C.suma(2, 3)

	fmt.Println(s)
}
