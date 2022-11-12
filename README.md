# Itachi &mdash; Generate Harmless Malware Samples

**At each run you will receive five unique samples that are detected by Sandbox, but not by static analysis**

## Samples

### &#x261E; Dropper

Upon execution it will write to current directory [eicar.com](https://www.eicar.com/download-anti-malware-testfile/) test malware file

### &#x261E; Encryptor

ransomware malware that "encrypts" all MS Office files in C:\ directory.

**Note:** Second run will decrypt them back

### &#x261E; Spyware

Upon execution this sample attempt to connect to wrs21.winshipway.com web site that is harmless but assumed by sandbox to be spyware-related

### &#x261E; Downloader

This sample downloads [eicar.com](https://www.eicar.com/download-anti-malware-testfile/) from web site and saves to currecnt directory

### &#x261E; AntiAV

This sample kills all antimalware related processes

### &#x261E; NoVirus

Harmless file that should not be detected by any security solution.

**Note:** If it is detected by some static analysis product, it is False Positive

## Installation
Not required. Just download Itachi for your OS from [releases](https://github.com/mpkondrashin/itachi/releases) 

## Static analysys

Plese note that some of these file are deetected by some static analysis engines. Please refer to following table that shows results faithful for Nov 12 2022.

|                   | spyware | encryptor | dropper | downloader | novirus | antiav |
| ----------------- | ------- | --------- | ------- | ---------- | ------- | ------ |
| Avast             |         |           | x       |            |         |        |
| AVG               |         |           | x       |            |         |        |
| Cybereason        |         |           |         |            | x       |        |
| Cylance           |         | x         | x       |            | x       |        |
| Cynet             | x       | x         | x       | x          | x       | x      |
| Cyren             |         |           | x       |            |         |        |
| Elastic           |         |           | x       | x          | x       |        |
| Google            | x       |           | x       | x          |         |        |
| Ikarus            | x       |           | x       | x          |         |        |
| MaxSecure         | x       | x         |         | x          | x       |        |
| Microsoft         |         | x         |         |            |         |        |
| SecureAge         | x       | x         | x       | x          | x       | x      |
| Trellix (FireEye) |         |           |         |            | x       |        |