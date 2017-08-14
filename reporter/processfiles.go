package reporter

import (
	"path/filepath"
	"strconv"
	"time"

	"fmt"

	"encoding/csv"
	"os"

	"github.com/ricardoaat/ggsn_csv_parser/config"
	"github.com/ricardoaat/ggsn_csv_parser/util"
	log "github.com/sirupsen/logrus"
)

func processCDRfiles(d int) {

	t := time.Now()

	for i := 0; i < d; i++ {
		log.Info("Attempting to load date " + t.Format("2006/01/02"))
		fl := getFileListByDate("TIGO-PGW*" + t.Format("20060102") + "*.csv")
		if fl != nil {
			log.Info(fmt.Sprintf("Files found for %s", t.Format("2006/01/02")))
			processFileList(fl)

		} else {
			log.Warning(fmt.Sprintf("No files found for %s", t.Format("2006/01/02")))
		}

		t = t.AddDate(0, 0, -1)
	}

}

func getFileListByDate(pattern string) []string {
	p := config.Conf.Path.CsvResources + pattern
	log.Info(p)
	fl, err := filepath.Glob(p)
	util.CheckErr(err, fmt.Sprintf("Error with path %s ", p))

	if len(fl) > 0 {
		return fl
	}
	return nil
}

func processFileList(fl []string) {
	for _, fn := range fl {

		filesProcessed = append(filesProcessed, fn)
		log.Info("Processing " + fn)

		f, err := os.Open(fn)
		util.CheckErr(err, "Error reading file "+fn)
		defer f.Close()

		r := csv.NewReader(f)
		d, err := r.ReadAll()
		util.CheckErr(err, "Error reading CVS fields ")

		if len(d) <= 1 {
			log.Warning("Empty file Skipping")
		} else {
			for i, c := range d {
				if i == 0 {
					log.Debug("Skipping header")
					continue
				}
				imsi, err := strconv.Atoi(c[3][:len(c[3])-1])
				util.CheckErr(err, "Error getting download field ")
				rg, err := strconv.Atoi(c[11])
				if err != nil {
					rg = 0
				}
				do, err := strconv.Atoi(c[15])
				util.CheckErr(err, "Error getting download field ")
				up, err := strconv.Atoi(c[16])
				util.CheckErr(err, "Error getting upload field ")
				su, err := strconv.Atoi(c[17])
				util.CheckErr(err, "Error getting sumbytes field ")

				tr := c[5]
				if len(tr) == 0 {
					continue
				}
				t, err := time.Parse("2006-01-02 15:04:05", tr[:19])
				util.CheckErr(err, "Failed to parse time from column ")

				m := fetchMvnoByImsi(imsi)
				if m == "NoMVNO" {
					saveNoMvnoIMSI(imsi)
				}

				cn := countByDay[t.Format("20060102")]
				cn.Download += do
				cn.Upload += up
				cn.Sumbytes += su

				if cn.RatingGroup == nil {
					cn.RatingGroup = make(map[int]int)
				}

				cn.RatingGroup[rg] += su

				if cn.Mvno == nil {
					cn.Mvno = make(map[string]int)
				}

				cn.Mvno[m] += su
				countByDay[t.Format("20060102")] = cn
			}
		}
	}

}

func fetchMvnoByImsi(i int) string {
	for m, n := range config.Conf.MVNOS {
		for _, r := range n.Ranges {
			if r[0] < i && i < r[1] {
				return m
			}
		}
	}
	return "NoMVNO"
}

func saveNoMvnoIMSI(im int) {
	for _, i := range noMvnoIMSIS {
		if i == im {
			return
		}
	}
	noMvnoIMSIS = append(noMvnoIMSIS, im)
}
