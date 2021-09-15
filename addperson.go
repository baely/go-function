package gofunc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	
	"cloud.google.com/go/datastore"
)

func AddPerson(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, _ := datastore.NewClient(ctx, "portstats")
	defer client.Close()

	var person Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		fmt.Println(err)
	}

	key := datastore.IncompleteKey("Person", nil)
	key, err := client.Put(ctx, key, &person)
	if err != nil {
		fmt.Println(err)
	}
}