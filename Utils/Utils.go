package Utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Writefile(outFile string, result string) {
	cf, err := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer cf.Close()
	_, err = cf.Write([]byte(result))
	check(err)
}

func Readfile(inputFile string) []string {
	output, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(output), "\n")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Parser(VarNum string, shellcodez []string) string {
	var scpayload []string
	shellcode := strings.Join(shellcodez, " ")
	const MAX_LENGTH int = 850
	x := 0
	shellcodeLength := len(shellcode)
	scpayload = append(scpayload, fmt.Sprintf("\r"))
	for x < shellcodeLength {
		if x+MAX_LENGTH <= shellcodeLength {
			scpayload = append(scpayload, fmt.Sprintf(VarNum+" = "+VarNum+" & \"%s\"\n", shellcode[0+x:x+MAX_LENGTH]))
			x += MAX_LENGTH
		} else {
			finalLength := shellcodeLength - x
			scpayload = append(scpayload, fmt.Sprintf(VarNum+" = "+VarNum+" & \"%s\"\n", shellcode[0+x:x+finalLength]))
			x += finalLength
		}
	}
	brokenupstrings := strings.Join(scpayload, "")
	return brokenupstrings
}
