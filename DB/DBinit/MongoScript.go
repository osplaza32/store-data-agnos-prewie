package DBinit

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (bd *BaseDato)initmongo()  {
	ctx := context.Background()
	// Options to the database.
	clientOpts := options.Client().ApplyURI(bd.Direccion)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		fmt.Println(err)
	}

	bd.mongo = client
}
