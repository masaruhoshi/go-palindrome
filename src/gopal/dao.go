package main

import "gopkg.in/mgo.v2"

type Dao struct {
	Instance	*mgo.Session
	Settings	Settings
}

func NewDao(settings Settings) *Dao {
	dao := new(Dao)

	session, err := mgo.Dial(settings.HostName)
	if err != nil {
		panic(err) // db not responding is a good reason to panic
	}
	session.SetMode(mgo.Monotonic, true)

	dao.Instance = session
	dao.Settings = settings

	return dao
}

func (dao *Dao) Close() {
	dao.Instance.Close()
}

func (dao *Dao) GetInstance() *Dao {
	return &Dao{dao.Instance.Copy(), dao.Settings}
}

func (dao *Dao) Database() *mgo.Database {
	return dao.Instance.DB(dao.Settings.DbName)
}

func (dao *Dao) EnsureIndex() {
	c := dao.Database().C("palindromes")

	index := mgo.Index{
		Key:		[]string{"phrase"},
		Unique:	 true,
		DropDups:   true,
		Background: true,
		Sparse:	 true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

