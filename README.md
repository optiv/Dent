# THIS REPOSITORY HAS BEEN ARCHIVED

To view the latest version of Dent or to submit an issue, reference https://github.com/Tylous/Dent.


<h1 align="center">
<br>
<p align="center"> <img src=Screenshots/dent.png border="2px solid #555">
<br>
Dent
</h1>



## More Information
If you want to learn more about the techniques utlized in this framework please take a look at this [article](https://www.optiv.com/insights/source-zero/blog/breaking-wdapt-rules-com).
#

## Description
This framework generates code to exploit vulnerabilties in Microsoft Defender Advanced Threat Protection's Attack Surface Reduction (ASR) rules to execute shellcode without being detected or prevented. ASR was designed to be the first line of defense, detecting events based on actions that violate a set of rules. These rules focus on specific behavior indicators on the endpoint that are often associated with an attacker’s Tactics, Techniques, or Procedures (TTPs). These rules have a heavy focus on the Microsoft Office suite, as this is a common attack vector for establishing a remote foothold on an endpoint. A lot of the rule-based controls focus on network-based or process-based behavior indicators that stand out from the normal business operation. These rules focus on either the initial compromise of a system or a technique that can severely impact an organization (e.g., disclosure of credentials or ransomware). They cover a large amount of the common attack surface and focus on hampering known techniques used to compromise assets.

Dent takes advantage of several vulnerabilities to bypass these resticive controls to execute payloads on an endpoint without being blocked or effectively detected by Microsoft Defender Advanced Threat Protection sensors. The article above outlines this vulnerabilties that are STILL present in Microsoft Defender Advanced Threat Protection even after disclosure.  



## Install
The first step as always is to clone the repo, then build it

```
go build Dent.go
```

## Help
```
./Dent -h
 
________                 __   
\______ \   ____   _____/  |_ 
 |    |  \_/ __ \ /    \   __\
 |    |   \  ___/|   |  \  |  
/_______  /\___  >___|  /__|  
        \/     \/     \/      
                (@Tyl0us)

"Call someone a hero long enough, and they'll believe it. They'll become it. 
They have no choice. Let them call you a monster, and you become a monster."


Usage of ./Dent:
  -C string
        Name of the COM object.
  -N string
        Name of the XLL playload when it's writen to disk.
  -O string
        Name of the output file. (default "output.txt")
  -P string
        Path of the DLL for your COM object. (Either use \\ or '' around the path)
  -U string
        URL where the base64 encoded XLL payload is hosted.
  -show
        Display the script in the terminal.
```

### Weaponizing 
This framwork is intended for exploiting vulnerabilities and deficiencies in Microsoft Defender Advanced Threat Protection, because of that it does not actually generate any payloads/implants. In order to generate those, you can use a large number of tools publicly availalble, however all research, development and testing was done using [ScareCrow](https://github.com/optiv/ScareCrow). Microsoft Defender Advanced Threat Protection doesn't rely on userland hooking for telemetry rather it utilized various other mechanisms such as kernel callbacks. From testing, this framework works extremely well at bypassing Microsoft Defender Advanced Threat Protection to execute shellcode. 


## Techniques

At the time of release there are currently two techniques. I will be constantly adding different ones that utilize these vulnerabilities in different ways peroidically so please stay tuned for more. 


## Fake COM Object Mode


COM objects are often created when an application is being installed on to a system. Once created, any application or script can call them, however this isn't the only way to create them. By modifying/creating registry keys into the `HKEY_CLASSES_ROOT` section of the Windows Registry, we can create a COM object that points to our shellcode on the system. This means any application or script that can utlize COM can call it, executing the shellcode.

This works because of how CoCreatInstance API functions. CoCreateInstance is used to create and initialize COM objects based on the CLSID (a globally unique identifier used to identify a specific COM class object). This function pulls the information to execute the call using the values stored in registry keys. These CLSID values can be found in the `HKEY_CLASSES_ROOT\CLSID\` path of the registry. However, before a process can call the CLSID, it must know that CLSID's value. This is done by first performing a registry query to look for the COM object in `HKEY_CLASSES_ROOT\<COM object name>`, and if it exists, a second registry query will be made to get the CLSID value stored in the subfolder.

Further inspection of the registry's subfolders shows the permission for the CLSID values are not consistent. A large majority of COM objects stored here only allow “Full Control” permission to the Trusted Installer. The Trusted Installer is a service account that owns resources to protect them, even from Administrators. This is intended to ensure that even if an attacker obtains administrative privileges, the resources cannot be manipulated maliciously. Unfortunately, a lot of COM Objects allow anyone in the Administrators groups “Full Control” permission. In addition, the root key CLSID allows the Administrators group “Full Control” permissions rather than NT AUTHORITY\System or Trusted Installer. Because of this, under an elevated context we can create, or even modify specific COM objects values.

<p align="center"> <img src=Screenshots/CLSID_Permissions.png border="2px solid #555">


**Important**

The creation of these registry keys only works if you run it under an elevated context. Double clicking on this through a GUI will not execute the .VBS file in an elevated context even if you are an administator. It is recommened you run it from an Administative shell or command prompt. However, once the keys are created, any application can call this COM object under any context. 

### ScareCrow Weaponizing 

To utilize a ScareCrow payload with this type of bypass you can run the following command:

```
./ScareCrow -I <path to your raw stageless shellcode>  -domain <domain name> -Loader dll
```

###  Usage

Once you have your payload, use the `-N` flag for the name of the payload when it's written to disk, the `-C` flag for the name of the COM object, the `-I` flag for the location to write it to, and lastly the `-O` for the output file to store the content.



## Remote .XLL Payload Mode 

This option generates a block of code to bypass several ASR rules to download, write to disk, load and execute shellcode, bypassing ASR preventative controls. This is done using the Excel.Application COM object which represents the entire Excel Application, but in an automated form, and allows for the programmatical interaction with it. Because this is still Excel, it does not trigger the ASR rule. This because when we call Excel.Application, we can see that it spawns under a Service Host process (Svchost.exe) and not the WinWord.exe process. While Svchost.exe is a system-level process used to host multiple Windows-based services, the child process created (Excel.exe) did not gain system-level privileges.

Because we created a COM object that was an entire application, the Excel process was created under Svchost.exe so it could be handled properly to prevent any instabilities to the WinWord.exe process. Though this process is under Svchost.exe, there is another challenge to contend with: executing shellcode. As binary execution or using WinAPI inside a macro will trigger other ASR rules, this limits what we can do without triggering an ASR rule or getting caught by WDAPT's EDR component. This is where DLLs shine. If a DLL-based payload is compiled with the right export functions, it can be used as an Office plugin that, when loaded, will automatically run shellcode. To do this, we can utilize Excel's RegisterXLL function. The RegisterXLL function loads an XLL plugin into memory, automatically registering and executing it. XLL files are essentially Excel-based DLLs. 

To get the content on the system, we can use another COM object (Microsoft.XMLHTTP), gaining the ability to execute an HTTP request, in this case, an HTTP GET request to a URL. The second COM object (ADODB.stream) provides the ability to read/write bytes of a data stream. By combining the two COM objects, an attacker can request a remote resource through an HTTP GET request and write the response (in this case, the file itself) to disk. This is done using the COM object (ADODB.stream) again to handle read/write bytes of the data stream. The second COM object (Microsoft.XMLDOM) allows the reading of data stored in a file. The XMLDCOM object allows for the data type to be set (in this case, base64) and once opened and stored in a string with the proper datatype, the ADODB.stream object can write the string of code to disk using a different data type (in this case, BinaryStreamType), converting the base64 string back into a binary form.



### ScareCrow Weaponizing 

To utilize a ScareCrow payload with this type of bypass you can run the following command:

```
./ScareCrow -I <path to your raw stageless shellcode>  -domain <domain name> -Loader excel  -O <Output filename>
```
Once generated, copy line 13 and 14 from the outputted file, and merge them together, making sure to remove:

* the `var <variable name>`
* the `;` at the end of each line
* the quotes around each string

<p align="center"> <img src=Screenshots/XLL_Payload.png border="2px solid #555">


###  Usage

Once you have your encoded payload, use the `-N` flag for the name of the payload when it's written to disk, the `-U` flag for the URL the encoded payload is going be hosted on (i.e https://<evilsite>/), and `-F` flag for the name of the file being hosted by the site. The outputted code is designed to work in a macro document .

### Lack of WDAPT Sensor Recording

Through further investigation, it was observed not as a gap in WDATP's sensors, but rather that WDATP does have visibility into this activity, but it is ignored. Through WDATP endpoint's timeline of events looking for any reference to Appwiz.xll, we observed that WDAPT recorded a “created file” event when Word created the file AppWiz.xll. It is important to note that .XLL files are executable.

## Disclosure Time 

**11/20/2020** - Research development and article written.

**03/14/2021** - Provided Microsoft a preliminary disclosure document outlining identified issues.

**03/31/2021** - Microsoft recognized and acknowledged that the vulnerabilities related to spawning an Office child process and writing files to disk were real vulnerabilities and began working on remediation. However, the permissions inconsistencies in the registry were deemed not a vulnerability due to the requirement of elevated privileges.

**04/21/2021** - Microsoft informed the author that signature build 1.333.1055.0 released on 03/22/2021 and 1.335.1321.0 released on 04/21/2021 contained the detection for the Office application-based vulnerabilities and closed the case.

**04/22/2021** - The author retested the same techniques identifying that the vulnerabilities were still present. 



