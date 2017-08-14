package reporter

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/sftp"
	"github.com/ricardoaat/ggsn_parsed_csv_reporter/config"
	"github.com/ricardoaat/ggsn_parsed_csv_reporter/util"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

var compressedFiles []string

func compressAllFiles() {
	for _, f := range filesProcessed {
		log.Info("Compressing file " + f)
		b, err := ioutil.ReadFile(f)
		util.CheckErr(err, "Error reading FILE ")

		fb := filepath.Base(f)
		cf := config.Conf.Path.CsvResources + fb + ".gz"
		fo, _ := os.Create(cf)
		compressedFiles = append(compressedFiles, cf)
		w := gzip.NewWriter(fo)
		w.Write(b)
		w.Close()
		log.Info("Created " + cf)
	}
}

func copyFilesToBackup() {

	g := &ssh.ClientConfig{
		User: config.Conf.SFTP.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.Conf.SFTP.Pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	u := fmt.Sprintf("%s:%s", config.Conf.SFTP.Host, config.Conf.SFTP.Port)
	c, err := ssh.Dial("tcp", u, g)
	util.CheckErr(err, "Failed to dial ")

	sf, err := sftp.NewClient(c)
	util.CheckErr(err, "Failed creating SFTP client ")
	defer sf.Close()

	copyFilesByDay(compressedFiles, sf)

}

func copyFilesByDay(fl []string, sf *sftp.Client) {
	for _, f := range fl {
		log.Info("Moving File " + f)
		b, err := ioutil.ReadFile(f)
		util.CheckErr(err, "Failed opening FILE ")

		of := config.Conf.Path.SFTPDest + filepath.Base(f)
		log.Debug(of)
		rf, err := sf.Create(of)
		util.CheckErr(err, "Failed creating file ")

		_, err = rf.Write(b)
		util.CheckErr(err, "Failed writting file SFTP ")
		log.Info("Moved to: " + of)
	}
}

func removeFiles() {
	for _, f := range filesProcessed {
		log.Info("Removing file " + f)
		err := os.Remove(f)
		util.CheckErr(err, "Failed erasing file ")
	}
	for _, f := range compressedFiles {
		log.Info("Removing file " + f)
		err := os.Remove(f)
		util.CheckErr(err, "Failed erasing file ")
	}
}
