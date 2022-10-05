package config

type (
	Dns struct {
		username string
		password string
		hostname string
		dbname   string
	}
	Config interface {
		GetUserName(d Dns) string
		GetPassword(d Dns) string
		GetHostName(d Dns) string
		GetDbName(d Dns) string
	}
)

const (
	username = "root"
	password = "root1234"
	hostname = "127.0.0.1:3306"
	dbname   = "content"
)

func GetDnsConfig() Dns {
	return Dns{
		username,
		password,
		hostname,
		dbname,
	}
}

func GetUserName(d Dns) string {
	return d.username
}

func GetPassword(d Dns) string {
	return d.password
}

func GetHostName(d Dns) string {
	return d.hostname
}

func GetDbName(d Dns) string {
	return d.dbname
}
