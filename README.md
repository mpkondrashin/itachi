# Itachi &mdash; Generate Harmless Malware Samples

**At each run you will receive five unique that are detected by Sandbox, but not by static analysis**

## Samples

### &#x261E; Dropper

Upon execution dropper.exe is writing to current directory [eicar.com](https://www.eicar.com/download-anti-malware-testfile/) test malware file

### &#x261E; Encryptor

ransomware file that "encrypts" all MS Office files in C:\ directory.

**Note:** Second run will decrypt them back

### &#x261E; Spyware

Upon execution this sample attempt to connect to wrs21.winshipway.com web site that is harmlessm but assumed by sandbox to be spyware-related

### &#x261E; Downloader

This sample downloads [eicar.com](https://www.eicar.com/download-anti-malware-testfile/) from web site and saves to currecnt directory

### &#x261E; AntiAV

This sample kill all antimalware related processes

 ### &#x261E; NoVirus

Harmless file that should not be detected by any securiyt solution.

**Note:** If it is detected by some static analysis product, it is False Positive

## Installation
Not required. Just download Itachi for your OS from [releases](https://github.com/mpkondrashin/itachi/releases) 

