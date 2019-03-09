package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	genOption := flag.String("genOption", "", "option: list|map")
	typeName := flag.String("typeName", "", "name of the list/map")
	keyType := flag.String("keyType", "", "type of the key in the map")
	valueType := flag.String("valueType", "", "type of the values hold in the list/map")
	packageName := flag.String("package", "", "package name of the list/map")
	addMutableFunctions := flag.Bool("mutable", false, "add mutable functions to the list")
	flag.Parse()

	if len(os.Args) <= 1 {
		flag.Usage()
		return
	}

	type value struct {
		Option      string
		PackageName string
		TypeName    string
		KeyType     string
		ValueType   string
		Mutable     bool
	}
	v := value{
		Option:      strings.TrimSpace(strings.ToLower(*genOption)),
		PackageName: *packageName,
		TypeName:    *typeName,
		KeyType:     *keyType,
		ValueType:   *valueType,
		Mutable:     *addMutableFunctions,
	}
	//primitiveTypes := []string{"uint", "uint8,", "uint16", "uint32", "uint64", "int", "int8,", "int16", "int32", "int64", "float32", "float64", "{{.ListType}}", "bool", "string"}

	fileName := fmt.Sprintf("%s.go", *typeName)
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Unable to create file: %s - %s", fileName, err)
	}
	w := bufio.NewWriter(file)

	switch v.Option {
	case "list":
		immutableListTemplate := template.Must(template.New("immutableListTemplate").Parse(immutableListTemplate))
		if err = immutableListTemplate.Execute(w, v); err != nil {
			log.Fatalf("Failed to create[%s]: %s", fileName, err)
		}
		if v.Mutable {
			mutableListTemplate := template.Must(template.New("mutableListTemplate").Parse(mutableListTemplate))
			if err = mutableListTemplate.Execute(w, v); err != nil {
				log.Fatalf("Failed to add mutable functions[%s]: %s", fileName, err)
			}
		}
	case "map":
		immutableMapTemplate := template.Must(template.New("immutableMapTemplate").Parse(immutableMapTemplate))
		if err = immutableMapTemplate.Execute(w, v); err != nil {
			log.Fatalf("Failed to create[%s]: %s", fileName, err)
		}
		if v.Mutable {
			mutableListTemplate := template.Must(template.New("mutableListTemplate").Parse(mutableListTemplate))
			if err = mutableListTemplate.Execute(w, v); err != nil {
				log.Fatalf("Failed to add mutable functions[%s]: %s", fileName, err)
			}
		}
	}
	err1 := w.Flush()
	if err1 != nil {
		log.Fatalf("Unable to create file: %s", err1)
	}
	err2 := file.Close()
	if err2 != nil {
		log.Fatalf("Unable to create file: %s", err2)
	}

}
