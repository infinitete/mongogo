package mongogo

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://fffftest:fffftestpwd@127.0.0.1:27017/"))
}

func TestIns(t *testing.T) {
	db := client.Database("compamy_statistics")

	err := db.CreateCollection(context.TODO(), "helllo")
	t.Logf("Error: %s", err)
	db.Collection("helllo").Drop(context.TODO())
}
