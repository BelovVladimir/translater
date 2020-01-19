package ui

import(
    "github.com/BelovVladimir/translater/tl"
    "bufio"
    "fmt"
    "os"
)

// создаем функцию ввода-вывода в терминале
func Terminal() {
    in := bufio.NewScanner(os.Stdin)
    for in.Scan() {
        fmt.Println("\n")
        txt := in.Text()
        rus := tl.Rus(txt)
        fmt.Println(rus, "\a\n")
        if txt == "exit" {
            break
        }
    }
}
