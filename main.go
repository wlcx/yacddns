package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getIP() string {
	res, _ := http.Get("https://api.ipify.org")
	ip1, _ := ioutil.ReadAll(res.Body)

	res, _ = http.Get("https://ident.me")
	ip2, _ := ioutil.ReadAll(res.Body)

	if string(ip1) != string(ip2) {
		log.Fatalf("Got different IPs: ipify.org: %s, ident.me: %s", ip1, ip2)
	}
	return string(ip1)
}

func main() {
	var apiKey = flag.String("apikey", "", "Cloudflare API key, find it in your profile")
	var apiEmail = flag.String("email", "", "Cloudflare account email")
	var zone = flag.String("zone", "", "Name of the cloudflare zone, e.g. example.com")
	var record = flag.String("record", "", "Name of the record to update, e.g. www.example.com")
	flag.Parse()

	if *apiKey == "" || *apiEmail == "" || *zone == "" || *record == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	api, err := cloudflare.New(*apiKey, *apiEmail)
	check(err)
	zoneID, err := api.ZoneIDByName(*zone)
	check(err)

	records, err := api.DNSRecords(zoneID, cloudflare.DNSRecord{Name: *record, Type: "A"})
	check(err)
	if len(records) != 1 {
		log.Fatalf("Error, %d records found for \"%s\", expected 1!", len(records), *record)
	}

	newRecord := records[0]

	newRecord.Content = getIP()

	api.UpdateDNSRecord(zoneID, records[0].ID, newRecord)
}
