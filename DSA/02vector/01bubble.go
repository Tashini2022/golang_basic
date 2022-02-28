package main

func Bubblesort(x []int) {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x)-i; j++ {
			if x[j] > x[j+1] {
				tmp := x[j]
				x[j] = x[j+1]
				x[j+1] = tmp
			}
		}
	}
}

/*
b[i],b[j]=b[j],b[i]
或者
b[i]=b[i]+b[j]																													
b[j]=b[i]-b[j]
b[i]=b[i]-b[j]
*/

func Bubblesort2(x []int){
	
}
