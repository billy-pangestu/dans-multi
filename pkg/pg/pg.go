package pg

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// Connection ...
type Connection struct {
	Host    string
	DB      string
	User    string
	Pass    string
	Port    int
	Loc     *time.Location
	SslMode string
}

type MySQL struct {
	*gorm.DB
}

// Connect ...
func (m Connection) Connect() (*MySQL, error) {
	connStr := fmt.Sprintf(
		`host=%s user=%s password=%s port=%d dbname=%s sslmode=%s TimeZone=UTC`,
		m.Host, m.User, m.Pass, m.Port, m.DB, m.SslMode,
	)

	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Enable Logger, show detailed log
	// db.LogMode(true)

	db.DB().SetConnMaxLifetime(5 * time.Minute)

	return &MySQL{db}, err
}
