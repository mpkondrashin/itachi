package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type sampleData struct {
	Title    string
	Filename string
}

type homepageData struct {
	Host      string
	HTTPPort  string
	HTTPSPort string
	Random    string
	Samples   []sampleData
}

var homepageTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>Itachi</title>
	<style>
		* {
			font-family: sans-serif;
		}
		.center {
			margin-left: auto;
			margin-right: auto;
		}
		th, td {
			padding: 5px;
		}
		h1 {
			text-align: center;
		}
	</style>
</head>
<body>
	<h1>Itachi: Malware Samples Generator</h1>
	{{$random := .Random}}
	{{$host := .Host}}
	{{$port := .HTTPPort}}
	{{$sport := .HTTPSPort}}
	<div style="border: 1px solid #ccc; padding: 10px; margin: 10px;">
	<table class="center">
	{{range .Samples}}
	<tr>
		<td>{{.Title}}</td>
		<td>
			<a href="https://{{$host}}:{{$sport}}/{{.Filename}}?r={{$random}}">HTTPS</a>
		</td>
		<td>
			<a href="http://{{$host}}:{{$port}}/{{.Filename}}?r={{$random}}">HTTP</a>
		</td>
	</tr>
	{{end}}
	</table>
	</div>
	<p>
		Notes:
		<ul>
		<li>Samples (beside EICAR.COM) are designed to showcase sandbox detection. They are not supposed to be detected by static antimalware engines. Some of the samples may be detected by sophisticated ML engines.</li>
		<li>For each sample it can be downloaded though HTTPS (with encryption) and HTTP (without encryption).</li>
		<li>HTTPS is using hardcoded self-signed certificate.</li>
		<li>Each time you click on the link, a new unique file will be generated.</li>
		<li>Only for eicar.com the same file will be generated.</li>
		<li>The script adds random data to URLs to avoid caching.</li>
		<li>Dropper: upon execution it saves eicar.com to the current directory.</li>
		<li>Spyware: upon execution it connects to remote server that considered by Trend Micro to be spyware related.</li>
		<li>Downloader: upon execution it downloads eicar.com file from eicar.com server.</li>
		</ul>
	</p>
</body>
</html>
`

var Samples = []sampleData{
	{Title: "Dropper", Filename: "dropper.exe"},
	//	{Title: "Encryptor", Filename: "encryptor.exe"},
	{Title: "Spyware", Filename: "spyware.exe"},
	{Title: "Downloader", Filename: "downloader.exe"},
	{Title: "EICAR.COM", Filename: "eicar.com"},
	//	{Title: "Autorun", Filename: "autorun.exe"},
	//	{Title: "AntiAV", Filename: "antiav.exe"},
	//{Title: "NoVirus", Filename: "novirus.exe"},
}

var tmpl = template.Must(template.New("homepage").Parse(homepageTemplate))

func homepage(w http.ResponseWriter, r *http.Request) {
	data := homepageData{
		Random:    uuid.New().String(),
		Host:      strings.Split(r.Host, ":")[0],
		HTTPPort:  strconv.Itoa(*httpPort),
		HTTPSPort: strconv.Itoa(*httpsPort),
		Samples:   Samples,
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
