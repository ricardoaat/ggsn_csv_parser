package reporter

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/ricardoaat/ggsn_parsed_csv_reporter/config"
	"github.com/ricardoaat/ggsn_parsed_csv_reporter/util"
	log "github.com/sirupsen/logrus"
)

//Init Starts report program
func Init(d int) {
	countByDay = loadCountFile()
	log.Info(fmt.Sprintf("Processing files until %d days ago", d))
	processCDRfiles(d)
	log.Info("Creating result file")
	toCountToFile()

	createResultFile()
	if len(filesProcessed) > 0 {
		//Backup Process
		if config.Conf.Process.Backup {
			compressAllFiles()
			copyFilesToBackup()
			removeFiles()
		}

		notifyResult()
	} else {
		notifyTeam("<b> No files loaded! <b>")
	}
}

func notifyResult() {
	m := "<p>Process <b>RESULT</b></p>"
	m += `
	<table border="1" cellpadding="3" cellspacing="0">
		<thead>
		<tr>
			<th align="center"><b>Date</b></td>    
			<th align="center"><b>Download</b></td>
			<th align="center"><b>Upload</b></td>
			<th align="center"><b>Sumbytes</b></td>
		</tr>
		</thead>
		<tbody>
	`
	ds := make(timeSlice, 0, len(countByDay))

	for d := range countByDay {
		t, err := time.Parse("20060102", d)
		util.CheckErr(err, "Failed to parse date ")
		ds = append(ds, t)
	}

	sort.Sort(ds)

	for _, t := range ds {
		m += `<tr>`
		tr := `
			<tr>
				<td align="center">%s</td>
				<td align="center">%d</td>
				<td align="center">%d</td>
				<td align="center">%d</td>
			</tr>		
		`
		mtf := t.Format("20060102")
		m += fmt.Sprintf(
			tr,
			t.Format("2006-01-02"),
			countByDay[mtf].Download,
			countByDay[mtf].Upload,
			countByDay[mtf].Sumbytes)
		m += `</tr><tr><td colspan="4">`
		for k, d := range countByDay[mtf].RatingGroup {
			m += fmt.Sprintf("RG:%d %d|", k, d)
		}
		m += `</td></tr>`

		m += `</tr><tr><td colspan="4">`
		for k, d := range countByDay[mtf].Mvno {
			m += fmt.Sprintf("%s %d|", k, d)
		}
		m += `</td></tr>`
	}
	m += `</tbody></table>`
	if len(noMvnoIMSIS) > 0 {
		m += `<p>IMEI's without MVNO</p>`
		for _, i := range noMvnoIMSIS {
			m += strconv.Itoa(i) + ","
		}
	}
	notifyTeam(m)
}
