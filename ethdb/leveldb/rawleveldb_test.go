// Copyright 2019 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package leveldb

import (
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

func TestRawLevelDB(t *testing.T) {
	t.Run("RawLevelDbB", func(t *testing.T) {

		db, err := leveldb.OpenFile("./ddd", nil)
		if err != nil {
			t.Fatal(err)
		}
		if err := db.Put([]byte("key"), []byte("value"), nil); err != nil {
			t.Fatalf("test: failed to insert item into database: %v", err)
		}

	})
}
