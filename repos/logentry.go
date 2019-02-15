package repos

import (
	"context"
	"fmt"
	"log"

	"github.com/levidurfee/ham/hamlog"
	"google.golang.org/appengine/datastore"
)

// QSO Entity type
var QSO = "QSOEntry"

// StoreEntry will store a new hamlog entry
func StoreEntry(ctx context.Context, entry *hamlog.Entry) {
	key := datastore.NewIncompleteKey(ctx, QSO, nil)
	r, err := datastore.Put(ctx, key, entry)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(r)
}
