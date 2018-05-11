package main

type Node struct {
        value int
        next *Node
}


func main() {
        //Allocate pointers
        var x *Node
        var y *Node
        var z *Node
        //Create pointees
        a:= Node{}
        b:= Node{}
        c:= Node{}
        //Point pointers to pointees
        x = &a
        y = &b
        z = &c
        //Dereference and set value
        x.value = 1
        y.value = 2
        z.value = 3
        //Dereference and set next pointer
        x.next = y
        y.next = z
        z.next = x

}
