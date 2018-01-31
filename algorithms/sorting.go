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

func (u *Unsorted) MergeSort(a []int, p, r int) {
        defer timeTrack(time.Now(), "Merge Sort")
        if p < r {
                q := (p + r) / 2
                MergeSort(a, p, q)
                MergeSort(a, q+1, r)
                Merge(a, p, q, r)
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
        i.Populate(12345678)
        i.InsertionReverseSort()
}

