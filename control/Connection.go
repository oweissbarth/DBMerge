package control

import (
	"strconv"

	_ "github.com/go-sql-driver/mysql" // the blank import is required
	logging "github.com/op/go-logging"
	"github.com/oweissbarth/DBMerge/model"
	"github.com/oweissbarth/DBMerge/utils"
)

var log = logging.MustGetLogger("DBMergeMain")

var server model.Server

/*ConnectToServer takes the
 *username, password, hostname and port
 *and authentificates with the mysql server
 */
func ConnectToServer(username string, password string, hostname string, port string) {
	portNumbder, err := strconv.Atoi(port)
	utils.CheckError(err)
	server = model.Server{hostname, portNumbder, username, password, nil}

	server.Connect()
}

/*GetDatabases retrieves
 *all available databases from the mysql server
 */
func GetDatabases() []string {
	dbs := server.GetDatabases()

	names := []string{}
	for _, db := range dbs {
		names = append(names, db.Name)
	}

	return names
}
