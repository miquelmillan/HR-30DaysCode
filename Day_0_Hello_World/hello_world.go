package main

import ("fmt"
        "bufio"
        "bytes"
        "os"
       )

func main(){
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    var buffer bytes.Buffer
    
    buffer.WriteString("Hello, World.\n")
    buffer.WriteString(text)
    
    fmt.Println(buffer.String())
}
