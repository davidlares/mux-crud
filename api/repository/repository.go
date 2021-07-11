package repository

import (
	"log"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// format
type TZConvertion struct {
	TimeZone string `bson:"timeZone" json:"timeZone"`
	TimeDifference string `bson:"TimeDifference" json:"TimeDifference"`
}

// string connection Object
type Repository struct {
	dbSession *mgo.Session 
	dbServer string 
	dbDatabase string 
	dbCollection string
}

// generating the new Repository
func NewRepository(dbServer string, dbDatabase string, dbCollection string) *Repository {
	// instance
	repo := new(Repository)
	repo.dbServer = dbServer
	repo.dbDatabase = dbDatabase
	repo.dbCollection = dbCollection
	// checking connection
	dbSession, err := mgo.Dial(repo.dbServer)
	if err != nil {
		log.Fatal(err)
	}
	// assigning session
	repo.dbSession = dbSession
	return repo
}

// closing
func (repo *Repository) Close() {
	repo.dbSession.Close()
}

// generating new Session (Cloning session)
func (repo *Repository) newSession() *mgo.Session {
	return repo.dbSession.Clone()
}

// Find all
func (repo *Repository) FindAll() ([]TZConvertion, error) {
	dbSession := repo.newSession()
	defer dbSession.Close()

	coll := dbSession.DB(repo.dbDatabase).C(repo.dbCollection)
	var tzcs []TZConvertion
	err := coll.Find(bson.M{}).All(&tzcs)
	return tzcs, err
}

// Find By Timezone
func (repo *Repository) FindByTimezone(tz string) (TZConvertion, error) {
	dbSession := repo.newSession()
	defer dbSession.Close()
	coll := dbSession.DB(repo.dbDatabase).C(repo.dbCollection)
	var tzc TZConvertion
	err := coll.Find(bson.M{"timeZone": tz}).One(&tzc)
	return tzc, err
}

// Insert
func (repo *Repository) Insert(tzc TZConvertion) error {
	dbSession := repo.newSession()
	defer dbSession.Close()
	coll := dbSession.DB(repo.dbDatabase).C(repo.dbCollection)
	err := coll.Insert(&tzc)
	return err 
}

func (repo *Repository) Delete(tzc TZConvertion) error {
	dbSession := repo.newSession()
	defer dbSession.Close()
	coll := dbSession.DB(repo.dbDatabase).C(repo.dbCollection)
	err := coll.Remove(bson.M{"timeZone": tzc.TimeZone})
	return err
}

func (repo *Repository) Update(tz string, tzc TZConvertion) error {
	dbSession := repo.newSession()
	defer dbSession.Close()
	coll := dbSession.DB(repo.dbDatabase).C(repo.dbCollection)
	err := coll.Update(bson.M{"timeZone": tz}, &tzc)
	return err
}