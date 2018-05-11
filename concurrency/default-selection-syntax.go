package main

func main() {
        select {
        case i := <-c:
            // use i
        default:
            // receiving from c would block
        }
}
