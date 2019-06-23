package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/ivahaev/russian-time"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func getTemplate(fileName string, funcmap template.FuncMap, data interface{}) (result string, err error) {
	template, err := template.New(fileName).Funcs(funcmap).ParseFiles("templates/" + fileName)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Panic(err)
	}

	var tpl bytes.Buffer
	if err := template.Execute(&tpl, data); err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Panic(err)
		panic(err)
	}

	result = tpl.String()

	return
}

func sendToHorn(text string) {
	m := map[string]interface{}{
		"text": text,
	}
	mJson, _ := json.Marshal(m)
	contentReader := bytes.NewReader(mJson)
	req, err := http.NewRequest("POST", os.Getenv("INTEGRAM_WEBHOOK_URI"), contentReader)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Panic(err)
	}

	fmt.Println(resp)
}

func hourWithMin() string {

	timeStamp := time.Unix(time.Now().Unix(), 0)

	hr, min, _ := timeStamp.Clock()

	finalTime := "%d:%d"

	result := fmt.Sprintf(finalTime, hr, min)

	return result
}

func weekDay() rtime.Weekday {
	t := rtime.Now()
	standardTime := time.Now()
	t = rtime.Time(standardTime)

	return t.Weekday()
}

func dumpError() {
	type Info struct {
		Status     string
	}

	templateData := Info {
		Status:     "Dump error!",
	}

	funcmap := template.FuncMap{
		"weekDay":            weekDay,
		"hourWithMin":        hourWithMin,
	}

	text, err := getTemplate("unsuccessful_backup.gohtml", funcmap, templateData)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Panic(err)
	}
	sendToHorn(text)
}

func dumpSuccess() {
	type Info struct {
		Status     string
	}

	templateData := Info {
		Status:     "Dump successful!",
	}

	funcmap := template.FuncMap{
		"weekDay":            weekDay,
		"hourWithMin":        hourWithMin,
	}

	text, err := getTemplate("successful_backup.gohtml", funcmap, templateData)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Panic(err)
	}
	sendToHorn(text)
}

func mysqlDump()  {
	// mysql_dump
	cmd := exec.Command("mysqldump",
		"-P" + os.Getenv("MYSQL_PORT"),
		"-h" + os.Getenv("MYSQL_HOST"),
		"-u" + os.Getenv("MYSQL_USER"),
		"-p" + os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DB") + "| gzip > " + os.Getenv("BACKUP_DIR") + os.Getenv("MYSQL_DB") + ".$(date +%F.%H%M%S).sql.gz")

	_, err := cmd.StdoutPipe()
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Fatal(err)
		dumpError()
	}
	dumpSuccess()
	cleaner()
}

func cleaner()  {
	// find old files
	// delete old file
	// send to horn if success
	fmt.Println("cleaner done")
}

func makeBackup() {
	mysqlDump()
	cleaner()
}

func tasks() {
	gocron.Every(1).Day().At("2:00").Do(makeBackup)

	// remove, clear and next_run
	_, time := gocron.NextRun()
	fmt.Println(time)

	// function Start start all the pending jobs
	<-gocron.Start()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appEnv := os.Getenv("APP_ENV")

	if appEnv == "production" {
		raven.SetDSN(os.Getenv("SENTRY_DSN"))
	}

	tasks()
}