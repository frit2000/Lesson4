package main

import (
	"fmt"
	"time"
)

type SingleArray struct {
	Arr []int
}

func (a *SingleArray) Append(count int) {

	for i := 0; i < count; i++ {
		a.ResizeAdd()
		a.Arr[len(a.Arr)-1] = i
	}
	return

}

func (a *SingleArray) ResizeAdd() {
	result := make([]int, (len(a.Arr) + 1))
	for i := 0; i < len(a.Arr); i++ {
		result[i] = a.Arr[i]
	}
	a.Arr = result
}

func (a *SingleArray) ResizeSub() {
	result := make([]int, (len(a.Arr) - 1))
	for i := 0; i < len(result); i++ {
		result[i] = a.Arr[i]
	}
	a.Arr = result
}

func (a *SingleArray) Add(item int, index int) {
	a.ResizeAdd()
	for i := len(a.Arr) - 1; i >= index; i-- {
		a.Arr[i] = a.Arr[i-1]
	}
	a.Arr[index] = item
}

func (a *SingleArray) Remove(index int) (item int) {
	item = a.Arr[index]
	for i := index; i < len(a.Arr)-1; i++ {
		a.Arr[i] = a.Arr[i+1]
	}
	a.ResizeSub()
	return
}

// /////////////////////////////////////Vector Arrays
type VectorArray struct {
	Arr []int
}

func (a *VectorArray) Append(count int) {

	for i := 0; i < count; i++ {
		if len(a.Arr) == cap(a.Arr) {
			a.ResizeAdd()
		}
		a.Arr = append(a.Arr, i)
	}
	return

}

func (a *VectorArray) ResizeAdd() {
	result := make([]int, len(a.Arr), len(a.Arr)+10)
	for i := 0; i < len(a.Arr); i++ {
		result[i] = a.Arr[i]
	}
	a.Arr = result
}

func (a *VectorArray) ResizeSub() {
	result := make([]int, (len(a.Arr) - 1), cap(a.Arr))
	for i := 0; i < len(result); i++ {
		result[i] = a.Arr[i]
	}
	a.Arr = result
}

func (a *VectorArray) Add(item int, index int) {
	a.ResizeAdd()
	for i := len(a.Arr) - 1; i >= index; i-- {
		a.Arr[i] = a.Arr[i-1]
	}
	a.Arr[index] = item
}

func (a *VectorArray) Remove(index int) (item int) {
	item = a.Arr[index]
	for i := index; i < len(a.Arr)-1; i++ {
		a.Arr[i] = a.Arr[i+1]
	}
	a.ResizeSub()
	return
}

// /////////////////////////////////////Factor Arrays
type FactorArray struct {
	Arr []int
}

func (a *FactorArray) Append(count int) {

	for i := 0; i < count; i++ {
		if len(a.Arr) == cap(a.Arr) {
			a.ResizeAdd()
		}
		a.Arr = append(a.Arr, i)
	}
	return

}

func (a *FactorArray) ResizeAdd() {
	result := make([]int, len(a.Arr), len(a.Arr)*2+1)
	for i := 0; i < len(a.Arr); i++ {
		result[i] = a.Arr[i]
	}
	a.Arr = result
}

func (a *FactorArray) ResizeSub() {
	result := make([]int, (len(a.Arr) - 1), cap(a.Arr))
	for i := 0; i < len(result); i++ {
		result[i] = a.Arr[i]
	}
	a.Arr = result
}

func (a *FactorArray) Add(item int, index int) {
	a.ResizeAdd()
	for i := len(a.Arr) - 1; i >= index; i-- {
		a.Arr[i] = a.Arr[i-1]
	}
	a.Arr[index] = item
}

func (a *FactorArray) Remove(index int) (item int) {
	item = a.Arr[index]
	for i := index; i < len(a.Arr)-1; i++ {
		a.Arr[i] = a.Arr[i+1]
	}
	a.ResizeSub()
	return
}

func main() {

	aSingle := SingleArray{[]int{}}
	aVector := VectorArray{[]int{}}
	aFactor := FactorArray{[]int{}}

	for i := 10; i <= 100_000; i *= 10 {
		start := time.Now()
		aSingle.Append(i)
		duration := time.Since(start)
		fmt.Println("aSingle, count=", i, " duration=", duration)

	}
	fmt.Println("---------------")
	for i := 10; i <= 1000_000; i *= 10 {
		start := time.Now()
		aVector.Append(i)
		duration := time.Since(start)
		fmt.Println("aVector, count=", i, " duration=", duration)

	}
	fmt.Println("---------------")
	for i := 10; i <= 1000_000; i *= 10 {
		start := time.Now()
		aFactor.Append(i)
		duration := time.Since(start)
		fmt.Println("aFactor, count=", i, " duration=", duration)
	}

	//	a.Append(count)
	//	a.Add(15, 5)
	//	fmt.Println("num1=", a.Arr)
	//	item := a.Remove(5)
	//	fmt.Println("num2=", a.Arr, "item=", item)
}
