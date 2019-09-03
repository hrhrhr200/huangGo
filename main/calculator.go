// stack_struct.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const LIMIT = 10

type Stack struct {
	ix   int // first free position, so data[ix] == 0
	data [LIMIT]int
}

func (st *Stack) Push(n int) {
	if st.ix+1 > LIMIT {
		return // stack is full!
	}
	st.data[st.ix] = n
	st.ix++
}

func (st *Stack) Pop() int {
	st.ix--
	return st.data[st.ix]
}

func (st Stack) String() string {
	str := ""
	for ix := 0; ix < st.ix; ix++ {
		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(st.data[ix]) + "] "
	}
	return str
}

func main() {
	buf := bufio.NewReader(os.Stdin)
	calc1 := new(Stack)
	fmt.Println("Give a number, an operator (+, -, *, /), or q to stop:")
	for {
		token, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println("Input error !")
			os.Exit(1)
		}
		token = token[0 : len(token)-1]
		switch {
		case token == "q":
			fmt.Println("Calculator stopped")
			return
		case token >= "0" && token <= "999999":
			i, _ := strconv.Atoi(token)
			calc1.Push(i)
		case token == "+":
			q := calc1.Pop()
			p := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p+q)
		case token == "-":
			q := calc1.Pop()
			p := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p-q)

		case token == "*":
			q := calc1.Pop()
			p := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p*q)

		case token == "/":
			q := calc1.Pop()
			p := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p/q)
		default:
			fmt.Println("No valid input")
		}
	}
}
