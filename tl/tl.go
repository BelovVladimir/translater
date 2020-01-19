package tl

import "strings"

// создаем функцию перевода строки
func Rus(text string) string {
    // разбиваем строку на слова 
    eng := strings.Split(text, " ")
    buffer := []string{}
    // переводим
    for _, value := range eng {
        low := strings.ToLower(value)
        if dict[low] != "" {
            buffer = append(buffer, dict[low])
        } else {
            buffer = append(buffer, value)
        }
    }
    // соединяем переведенные слова в новую строку
    rus := strings.Join(buffer, " ")
    return rus
}
