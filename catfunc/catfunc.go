package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    // "github.com/k0kubun/pp"
    "encoding/json"
)

func check_regexp(reg, str string) bool {
    return regexp.MustCompile(reg).Match([]byte(str))
}

func main() {
    var fp *os.File
    var err error

    if len(os.Args) < 2 {
        fp  = os.Stdin
    } else {
        fp, err = os.Open(os.Args[1])
        if err != nil {
            panic(err)
        }
        defer fp.Close()
    }

    scanner := bufio.NewScanner(fp)
    m := make(map[string]string)
    mm := make(map[string]map[string]string)
    for scanner.Scan() {
        funcStr := ""
        var _ = funcStr

        check_regexp(`^func`, scanner.Text())
        // Instance method
        r := regexp.MustCompile(`^func \(.*\) (\w+)\(.*`)
        // r := regexp.MustCompile(`^func\s[.*\s\(\)]?(\w+)\(.*`)
        funcName := r.FindAllStringSubmatch(scanner.Text(), -1)
        if len(funcName) > 0 {
            // fmt.Println(funcName[0][1])
            // fmt.Println(r.FindAllStringSubmatch(scanner.Text(), -1))
            if check_regexp(`^func`, scanner.Text()) == true {
                funcStr += scanner.Text() + "\n"
                for scanner.Scan() {
                    funcStr += scanner.Text() + "\n"
                    if check_regexp(`^}$`, scanner.Text()) == true {
                        // fmt.Println(funcName[0][1])
                        // pp.Println(m)
                        m[funcName[0][1]] = funcStr
                        break
                    }
                }
            }
        }
        // normarl method
        r = regexp.MustCompile(`^func\s(\w+)\(.*`)
        funcName = r.FindAllStringSubmatch(scanner.Text(), -1)
        if len(funcName) > 0 {
            // fmt.Println(funcName[0][1])
            // fmt.Println(r.FindAllStringSubmatch(scanner.Text(), -1))
            if check_regexp(`^func`, scanner.Text()) == true {
                funcStr += scanner.Text() + "\n"
                for scanner.Scan() {
                    funcStr += scanner.Text() + "\n"
                    if check_regexp(`^}$`, scanner.Text()) == true {
                        // fmt.Println(funcName[0][1])
                        // pp.Println(m)
                        m[funcName[0][1]] = funcStr
                        break
                    }
                }
            }
        }
    }
    mm[os.Args[1]] = m
    // pp.Print(m)
    if err := scanner.Err(); err != nil {
        panic(err)
    }
    data, err := json.MarshalIndent(mm, "", "  ")
    if err != nil {
        panic(err)
    }
    fmt.Println(string(data))
}
