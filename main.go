package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {

	var types string
	var outputs string
	var args []string
	var notargs []string
	var in_flags bool = false
	var help bool
	for i := 0; i < len(os.Args); i++ {
		if os.Args[i][0] == '-' {
			in_flags = true
		}
		if i == 0 || in_flags {
			notargs = append(notargs, os.Args[i])
		} else {
			args = append(args, os.Args[i])
		}
	}
	if len(args) == 0 {
		fmt.Println("usage: <file.log> -t <type> -o <output file>")
		fmt.Println("-o string\n \t Writes output to the file specified")
		fmt.Println("-t string\n \t Type output txt|json (default \"txt\")")
		fmt.Println("-h string\n \t Help")
		os.Exit(0)
	}
	os.Args = notargs
	flag.StringVar(&types, "t", "txt", "Type output")
	flag.StringVar(&outputs, "o", "", "Writes output to the file specified")
	flag.BoolVar(&help, "h", false, "Help")
	flag.Parse()
	switch types {
	case "txt":
		txt(args[0], outputs)
	case "json":
		jsonExt(args[0], outputs)
	default:
		fmt.Println("unknown type file " + types)
		os.Exit(0)
	}
}

func txt(file string, dir string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	if dir == "" {
		dir = ""
		file = strings.TrimSuffix(filepath.Base(file), path.Ext(filepath.Base(file)))
	} else {
		file = strings.TrimSuffix(filepath.Base(dir), path.Ext(filepath.Base(dir)))
		dir, _ = filepath.Split(dir)
	}
	if err := ioutil.WriteFile(dir+file+".txt", data, os.ModePerm); err != nil {
		fmt.Println(err)
		return
	}

}

func jsonExt(file string, dir string) {
	data, err := ioutil.ReadFile(file)
	dataReplace := strings.Replace(string(data), "\n", "\r\n", -1)
	dataArr := strings.Split(strings.Replace(dataReplace, "\r\n", "\n", -1), "\n")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	datas := []map[string]interface{}{}
	for i := 0; i < len(dataArr); i++ {
		data := map[string]interface{}{
			"message": dataArr[i],
		}
		datas = append(datas, data)
	}
	jsonString, _ := json.Marshal(datas)

	if dir == "" {
		dir = ""
		file = strings.TrimSuffix(filepath.Base(file), path.Ext(filepath.Base(file)))
	} else {
		file = strings.TrimSuffix(filepath.Base(dir), path.Ext(filepath.Base(dir)))
		dir, _ = filepath.Split(dir)
	}
	if err := ioutil.WriteFile(dir+file+".json", jsonString, os.ModePerm); err != nil {
		fmt.Println(err)
		return
	}
}
