package main

import (
  f"fmt"
  s"strconv"
  r"math/rand"
  size"golang.org/x/crypto/ssh/terminal"
  b"bufio"
  "os"
)
func printer(e string){
  f.Print(e + "\n")
}

const (
    MAX_SIZE int=50000
)

// Needed..
type Sizes struct {
  index int
  size_height [MAX_SIZE]int
  size_width [MAX_SIZE]int
}

func (setSizes* Sizes) set(width,height int) {
  setSizes.size_height[setSizes.index] = height
  setSizes.size_width[setSizes.index] = width
}
func (sizes Sizes) print_with_size(index *int) {
    /* GETTING STRING */
    reader := b.NewReader(os.Stdin)
    f.Print("TYPE SOMETHING TO PRINT: ")
    to_print,_ := reader.ReadString('\n')
    
    /* Checking length */
    var leng int = len(to_print)
    var i int = 0
    for ; i < leng; i++ {
        t := 0
        if(i==10) {
            t = i
            sizes.size_height[*index]+=i
            sizes.size_height[*index]+=i+2
        }

        if(t==leng-1) { 
            break
        }
    }

    for {
      f.Printf("\x1b[" + s.Itoa(r.Intn(sizes.size_height[*index])) + ";" + s.Itoa(r.Intn(sizes.size_width[*index])) + "H\x1b[" + s.Itoa(40 + r.Intn(7)) + "m\x1b[" + s.Itoa(30 + r.Intn(7))+"m"+to_print+"\x1b[39m\x1b[49m")
    }
    *index++
}

func main() {
    /* GETTING AND PRINTING SIZES */
    size_:=Sizes{}
    size_.index = 0
    width,height,_:=size.GetSize(0) //haha

    /* Storing sizes in struct */
    size_.set(width,height)

    /* PRINTING */
    size_.print_with_size(&size_.index)
}