package repos

import (
	"context"
	"fmt"
	"log"

	"github.com/levidurfee/ham/models"
	"google.golang.org/appengine/datastore"
)

// SaveQSO will store a new QSO entry
func SaveQSO(ctx context.Context, entry *models.QSO) {
	key := datastore.NewIncompleteKey(ctx, models.QSOEntityType, nil)
	r, err := datastore.Put(ctx, key, entry)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(r)
}
