package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// for scanner.Scan() {
	// 	checkDomain(scanner.Text())
	// }
	checkDomain("yahoo.com")
	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading from stdin: %v\n", err)
	}

}

func checkDomain(domain string) {

	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("%s: %v (no MX record) \r ", domain, err) // no MX record
	}

	if len(mxRecords) > 0 {
		hasMx = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("%s: %v (no TXT record) \r ", domain, err) // no TXT record
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("%s: %v (no DMARC record) \r ", domain, err) // no DMARC record
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%s: MX: %t, SPF: %t, DMARC: %t, SPF Record: %s, DMARC Record: %s \r ", domain, hasMx, hasSPF, hasDMARC, spfRecord, dmarcRecord)

}
