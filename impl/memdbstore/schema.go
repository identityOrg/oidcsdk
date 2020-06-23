package memdbstore

import "github.com/hashicorp/go-memdb"

var (
	schema = &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"request": &memdb.TableSchema{
				Name: "request",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "RequestID"},
					},
					"code-sign": &memdb.IndexSchema{
						Name:    "code-sign",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "AuthorizationCodeSignature"},
					},
					"at-sign": &memdb.IndexSchema{
						Name:    "at-sign",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "AccessTokenSignature"},
					},
					"rt-sign": &memdb.IndexSchema{
						Name:    "rt-sign",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "RefreshTokenSignature"},
					},
				},
			},
		},
	}
)
