package repos_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/happilymarrieddad/learning-go-gRPC/repos"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	err error
	db  *xorm.Engine
	//dbSql *sql.DB
	//mock  sqlmock.Sqlmock

	gr repos.GlobalRepository

	truncateUsers = func() {
		// mock.ExpectQuery("TRUNCATE users").
		// 	WillReturnRows(sqlmock.NewRows([]string{}))

		_, err = db.Query("TRUNCATE users")
		Ω(err).To(BeNil())
	}

	clearDatabase = func() {
		if db == nil {
			Fail("unable to run test because database is missing")
		}

		truncateUsers()

		return
	}
)

var _ = BeforeSuite(func() {
	// connection string - root:pass@tcp(localhost:3306)/grpc
	db, err = xorm.NewEngine("mysql", "root:pass@tcp(localhost:3306)/grpc")
	Ω(err).To(BeNil())
	_, err = db.DBMetas()
	Ω(err).To(BeNil())
	// dbSql, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	// Ω(err).To(BeNil())
	// db.DB().DB = dbSql

	gr = repos.GlobalRepo(db)
})

var _ = AfterSuite(func() {
	// err = mock.ExpectationsWereMet()
	// Ω(err).To(BeNil())
})

func TestRepos(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repos Suite")
}
