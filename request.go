package httprouter

import (
    "net/http"
    "strings"
)

const (
    defaultMaxMemory = 32 << 20 // 32 MB
)

// Experemental function
func PostFormValues(r *http.Request, name string) (list map[string]string) {
    if r.PostForm == nil {
        r.ParseMultipartForm(defaultMaxMemory)
    }

    list = make(map[string]string)

    var left_i int = len(name+"[")

    for key, value := range r.PostForm {
        if strings.HasPrefix(key, name+"[") {
            key = key[left_i:len(key)-1]

            if len(value) > 0 {
                list[key] = value[0]
            } else {
                list[key] = ""
            }
        }
    }
    return
}