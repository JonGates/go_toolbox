# GO工具箱

## authcode
##### discuz authcode for golang ，and keep the data consistent。

##### Install 
```javascript
    go get github.com/JonGates/go_toolbox
```
##### Usage
```javascript
    package main
    import (
        "fmt"
        "github.com/JonGates/go_toolbox"
    )
    func main() {
        str := "JonGates"
        encode := go_toolbox.AuthCode(str, "ENCODE", "", 0, 4)
        fmt.Println(encode)
        decode := go_toolbox.AuthCode(encode, "DECODE", "", 0, 4)
        fmt.Println(decode)
    }
```
 #### Mutual solution of PHP and GO
```javascript
    Key,ckey_length must be consistent with PHP
```

