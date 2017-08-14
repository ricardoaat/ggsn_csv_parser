package util

import log "github.com/sirupsen/logrus"

/*CheckErr Panics if gets error
 */
func CheckErr(err error, message string) {
	if err != nil {
		log.Panic(message, err)
	}
}
