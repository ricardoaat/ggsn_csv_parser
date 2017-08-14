package reporter

import (
	"fmt"
	"os"
	"time"

	"github.com/ricardoaat/ggsn_csv_parser/config"
	"github.com/ricardoaat/ggsn_csv_parser/util"
	log "github.com/sirupsen/logrus"
)

func createResultFile() {
	t := time.Now()

	for ct, d := range countByDay {

		fn := fmt.Sprintf("cdr_ggsn_report_%s.csv", ct)
		fo, err := os.Create(config.Conf.Path.ResultReport + fn)
		util.CheckErr(err, "Error creating file "+fn)
		defer fo.Close()

		fo.WriteString("day|Download|Upload|Sumbytes\n")

		l := fmt.Sprintf("%s|%d|%d|%d\n", ct, d.Download, d.Upload, d.Sumbytes)
		_, err = fo.WriteString(l)
		log.Debug(fmt.Sprintf("Writing line: %s", l))
		util.CheckErr(err, "Error writing line "+l)

		for k, d := range d.RatingGroup {
			_, err := fo.WriteString(fmt.Sprintf("RG:%d %d|", k, d))
			util.CheckErr(err, "Error writing line "+l)
		}
		fo.WriteString("\n")

		for k, d := range d.Mvno {
			_, err := fo.WriteString(fmt.Sprintf("%s %d|", k, d))
			util.CheckErr(err, "Error writing line "+l)
		}
		fo.WriteString("\n")
	}

	if len(noMvnoIMSIS) > 0 {

		log.Info("Creating imsis's without MVNO file")
		fn := fmt.Sprintf("imsis_without_mvno_%s.csv", t.Format("20060102T150405"))
		foi, err := os.Create(config.Conf.Path.ResultReport + fn)
		util.CheckErr(err, "Error creating file "+fn)
		defer foi.Close()

		for _, i := range noMvnoIMSIS {
			_, err := foi.WriteString(fmt.Sprintf("%d\n", i))
			util.CheckErr(err, "Error writing line ")
		}
	}
}
