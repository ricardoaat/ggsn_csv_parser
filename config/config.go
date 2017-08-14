package config

//Conf global configure instance
var Conf config

type config struct {
	Process process
	Path    path
	Notif   notif
	SFTP    sftp
	MVNOS   map[string]mvno
}

type process struct {
	Backup bool
}

type path struct {
	LogPath      string
	ResultReport string
	CsvResources string
	SFTPDest     string
}

type notif struct {
	Recipients []string
	FromAddrs  string
}

type sftp struct {
	Host string
	Port string
	User string
	Pass string
}

type mvno struct {
	Ranges [][]int
}
