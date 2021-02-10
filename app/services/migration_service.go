package services

import (
	"fmt"
)

// Migrate all the necessary indexes
func MigrateIndex() error {
	if err := migrateArticleIndex(); err != nil {
		return err
	}
	return nil
}

// Map the article index in order to type all the fields.
func migrateArticleIndex() error {
	eventMapping :=
		`{
        "mappings": {
            "properties": {
                "AggregateID": {
                    "type": "keyword"
                },
                "CreatedAt": {
                    "type": "date"
                },
                "Index": {
                    "type": "long"
                },
                "Payload": {
                    "properties": {
                        "AuthorID": {
                            "type": "long"
                        },
                        "Content": {
                            "type": "text",
                            "fields": {
                                "keyword": {
                                    "type": "keyword",
                                    "ignore_above": 256
                                }
                            }
                        },
                        "CreatedAt": {
                            "type": "date"
                        },
                        "Title": {
                            "type": "text",
                            "fields": {
                                "keyword": {
                                    "type": "keyword",
                                    "ignore_above": 256
                                }
                            }
                        }
                    }
                },
                "Typology": {
                    "type": "text",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                }
            }
        }
	}`
	if err := CreateNewIndex("article", eventMapping); err != nil {
		return fmt.Errorf("could not migrate article index: %s", err.Error())
	}
	return nil
}
