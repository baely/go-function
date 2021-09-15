package gofunc

import (
	"context"
	"fmt"
	"net/http"
	
	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

func GetPerson(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, _ := datastore.NewClient(ctx, "portstats")
	defer client.Close()

	query := datastore.NewQuery("Person")
	it := client.Run(ctx, query)

	for {
		var person Person
		_, err := it.Next(&person)
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, "Person: %v\n", person)
	}
}

