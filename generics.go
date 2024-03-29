package main

import (
	"fmt"
)

//sumints
type Number interface{
	int64 | float64
}

func sumInts(m map[string]int64)int64{
	var sum int64

	for _,v :=range m{
		sum +=v
	}
	return sum
}

func sumFloats(m map[string]float64)float64{
	var sum float64

	for _,v :=range m{
		sum +=v
	}
	return sum
}
//generic function
func sumIntOrFloat[k comparable, v int64 | float64](m map[k]v)v{
	var sum v
	for _,v:=range m{
		sum +=v
	}
	return sum
}

//with interface
func sumIntOrFloatWithInterface[k comparable, v Number | float64](m map[k]v)v{
	var sum v
	for _,v:=range m{
		sum +=v
	}
	return sum
}

func Generic(){
	mInt:=map[string]int64{"first":20,"second":30}
	mfloat:=map[string]float64{"first":21.5,"second":33.6}
	fmt.Println(sumInts(mInt))

	fmt.Println(sumFloats(mfloat))

}
func GenericCall(){
	mInt:=map[string]int64{"first":20,"second":30}
	mfloat:=map[string]float64{"first":21.5,"second":33.6}

	fmt.Printf("sum for int %v:",sumIntOrFloat[string,int64](mInt))
	fmt.Printf("sum for floats %v:",sumIntOrFloat[string,float64](mfloat))

	fmt.Printf("sum for int with interface %v:",sumIntOrFloatWithInterface(mInt))
	fmt.Printf("sum for floats with interface %v:",sumIntOrFloatWithInterface(mfloat))

}