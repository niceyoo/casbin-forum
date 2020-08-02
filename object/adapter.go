// Copyright 2020 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package object

import (
	"runtime"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var adapter *Adapter

func InitAdapter() {
	adapter = NewAdapter("mysql", beego.AppConfig.String("dataSourceName"))
}

// Adapter represents the MySQL adapter for policy storage.
type Adapter struct {
	driverName     string
	dataSourceName string
	engine         *xorm.Engine
}

// finalizer is the destructor for Adapter.
func finalizer(a *Adapter) {
	err := a.engine.Close()
	if err != nil {
		panic(err)
	}
}

// NewAdapter is the constructor for Adapter.
func NewAdapter(driverName string, dataSourceName string) *Adapter {
	a := &Adapter{}
	a.driverName = driverName
	a.dataSourceName = dataSourceName

	// Open the DB, create it if not existed.
	a.open()

	// Call the destructor when the object is released.
	runtime.SetFinalizer(a, finalizer)

	return a
}

func (a *Adapter) createDatabase() error {
	engine, err := xorm.NewEngine(a.driverName, a.dataSourceName)
	if err != nil {
		return err
	}
	defer engine.Close()

	_, err = engine.Exec("CREATE DATABASE IF NOT EXISTS casbin_forum default charset utf8 COLLATE utf8_general_ci")
	return err
}

func (a *Adapter) open() {
	if err := a.createDatabase(); err != nil {
		panic(err)
	}

	engine, err := xorm.NewEngine(a.driverName, a.dataSourceName+"casbin_forum")
	if err != nil {
		panic(err)
	}

	a.engine = engine
	a.createTable()
}

func (a *Adapter) close() {
	a.engine.Close()
	a.engine = nil
}

func (a *Adapter) createTable() {
	err := a.engine.Sync2(new(Topic))
	if err != nil {
		panic(err)
	}

	err = a.engine.Sync2(new(Reply))
	if err != nil {
		panic(err)
	}

	err = a.engine.Sync2(new(Member))
	if err != nil {
		panic(err)
	}

	err = a.engine.Sync2(new(Node))
	if err != nil {
		panic(err)
	}

	err = a.engine.Sync2(new(Favorites))
	if err != nil {
		panic(err)
	}

	err = a.engine.Sync2(new(Tab))
	if err != nil {
		panic(err)
	}

	err = a.engine.Sync2(new(Notification))
	if err != nil {
		panic(err)
	}

	err = a.engine.Sync2(new(BasicInfo))
	if err != nil {
		panic(err)
	}

	err = a.engine.Sync2(new(Plane))
	if err != nil {
		panic(err)
	}

	err = a.engine.Sync2(new(ConsumptionRecord))
	if err != nil {
		panic(err)
	}

	err = a.engine.Sync2(new(BrowseRecord))
	if err != nil {
		panic(err)
	}

	err = a.engine.Sync2(new(ValidateCode))
	if err != nil {
		panic(err)
	}
}
