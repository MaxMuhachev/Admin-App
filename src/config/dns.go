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
	username = "uf6c4ofdy7rwuroo"
	password = "KOBI3FANKlfw8E2zEEuP"
	hostname = "bvgghoxnzxbvck2zfnld-mysql.services.clever-cloud.com:3306"
	dbname   = "bvgghoxnzxbvck2zfnld"
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
