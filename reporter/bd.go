package reporter

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ricardoaat/ggsn_csv_parser/util"
	log "github.com/sirupsen/logrus"
)

func loadCountFile() map[string]countBytes {
	log.Info("Loading previous counters (./countbyday.json)")
	r, err := ioutil.ReadFile("./countbyday.json")
	util.CheckErr(err, "Failed reading countbyday.json file ")

	var c map[string]countBytes
	err = json.Unmarshal(r, &c)
	util.CheckErr(err, "Failed Unmarshal countbyday.json content ")
	return c
}

func toCountToFile() {
	log.Info("Saving counters (./countbyday.json)")
	b, err := json.Marshal(countByDay)
	util.CheckErr(err, "Failed marshalling JSON")

	err = ioutil.WriteFile("./countbyday.json", b, 0644)
	util.CheckErr(err, "Failed writing countbyday.json file ")
}
