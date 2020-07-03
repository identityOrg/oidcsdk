package memdbstore

import "github.com/hashicorp/go-memdb"

var (
	schema = &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"request": {
				Name: "request",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "RequestID"},
					},
					"code-sign": {
						Name:         "code-sign",
						Unique:       true,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "AuthorizationCodeSignature"},
					},
					"at-sign": {
						Name:         "at-sign",
						Unique:       true,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "AccessTokenSignature"},
					},
					"rt-sign": {
						Name:         "rt-sign",
						Unique:       true,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "RefreshTokenSignature"},
					},
				},
			},
			"client": {
				Name: "client",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
		},
	}
)
