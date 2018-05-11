//sorting 
package main

import (
        "fmt"
        "math/rand"
        "time"
)

type Unsorted struct {
        list []int
}

//function to be called using defer in any other function.
func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    fmt.Printf("%s took %s", name, elapsed)
}

func (u *Unsorted) InsertionSort() {
        defer timeTrack(time.Now(),  "Insertion Sort")
        for j :=1 ; j < len(u.list); j++ {
                key := u.list[j] //pick the next element in the list
                //Insert key into the sorted sequence, u.list[1..j-1]
                i := j - 1 //pick the element before the key
                for i >= 0 && u.list[i] > key { // whilst i is not the beginning of the list, and it's value is greater than key
                        u.list[i+1] = u.list[i] //swap key and u.list[i] (up the list)
                        i -= 1 //decrement i
                u.list[i+1] = key // 'decrement' the key (move the 'hand' down the list)
                }
        }
        //fmt.Println("Insertion sorted list: ")
        //fmt.Print(u.list)
}

func (u *Unsorted) InsertionReverseSort() {
        defer timeTrack(time.Now(), "Reverse Insertion Sort")
        for j :=1 ; j < len(u.list); j++ {
                key := u.list[j] //pick the next element in the list
                //Insert key into the sorted sequence, u.list[1..j-1]
                i := j - 1 //pick the element before the key
                for i >= 0 && u.list[i] < key { // whilst i is not the beginning of the list, and it's value is lsthn than key
                        u.list[i+1] = u.list[i] //swap key and u.list[i] (up the list)
                        i -= 1 //decrement i
                u.list[i+1] = key // 'decrement' the key (move the 'hand' down the list)
                }
        }
        //fmt.Println("Insertion reverse sorted list: ")
        //fmt.Print(u.list)
}

func (u *Unsorted) MergeSort(p, r int) {
        //defer timeTrack(time.Now(), "Merge Sort")
        fmt.Printf("Merge sort called with p: %v and r: %v \n", p,r)
        if p < r {
                fmt.Printf("In Mergesort, p: %v, r: %v \n", p,r)
                q := (p + r) / 2
                fmt.Printf("In Mergesort q: %v\n", q)
                u.MergeSort(p, q)
                u.MergeSort(q+1, r)
                u.Merge (p, q, r)
	}
}

func (u *Unsorted) Merge(p, q, r int) {
        fmt.Printf("Merge called with p: %v, q: %v, r: %v \n", p, q, r)
	n1 := q - p + 1
	n2 := r - q
	L := make([]int, n1+2, n1+2)
        R := make([]int, n2+2, n2+2)
	fmt.Println("The entire list inside merge", u.list)
	fmt.Println("The p'th element of the list outside of any loop (inside merge)", u.list[p])
	for i:=1; i<=n1; i++ {
                fmt.Println("The p+i-1'th element in i's loop", u.list[p+i-1])
		L[i] = u.list[p + i - 1]
	}
	for j:=1; j<=n2; j++ {
                fmt.Println("The q+j'th element in j's loop", u.list[q+j])
		R[j] = u.list[q + j]
	}
	L[n1+1] = 5000 //approximate infinity to 5000
	R[n2+1] = 5000
	i, j := 1, 1
	for k:=p; k < r; k++ {
		if L[i] <= R[j] {
                        fmt.Printf("setting u.list[%d] to %d\n", k, L[i])
			u.list[k] = L[i]
			i += 1
		} else {
                        fmt.Printf("setting u.list[%d] to %d\n", k, R[i])
			u.list[k] = R[j]
			j += 1
		}
	}
	return
}

func (u *Unsorted) Populate(size int) {
        u.list = make([]int, size)
        for i:= range u.list {
                u.list[i] = rand.Intn(4096)
        }
}

func main() {
        i := Unsorted{}
        i.Populate(5)
        i.MergeSort(0, len(i.list))
	fmt.Println(i.list)
}

