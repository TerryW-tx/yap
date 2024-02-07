/*
 * Copyright (c) 2024 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mysql

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goplus/yap/ydb"
)

// Register registers a default data source for `mysql` engine.
func Register(defaultDataSource string) {
	ydb.Register("mysql", defaultDataSource)
}

func init() {
	ydb.Register("mysql", func() string {
		dataSource := os.Getenv("YDB_MYSQL_TEST")
		if dataSource == "" {
			log.Panicln("env `YDB_MYSQL_TEST` not found, please set it before running")
		}
		return dataSource
	})
}
