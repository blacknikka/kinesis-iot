package mongo

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	client *mongo.Client
}

const (
	// Path to the AWS CA file
	caFilePath = "rds-combined-ca-bundle.pem"

	queryTimeout    = 30
	username        = "root"
	password        = "root"
	clusterEndpoint = "mongo:27017"

	// Which instances to read from
	readPreference = "secondaryPreferred"
)

func (m *Mongo) Connect() error {
	// connectionURI := fmt.Sprintf(
	// 	"mongodb://%s:%s@%s/sample-database?ssl=true&replicaSet=rs0&readpreference=%s",
	// 	username,
	// 	password,
	// 	clusterEndpoint,
	// 	readPreference,
	// )

	connectionURI := fmt.Sprintf(
		"mongodb://%s:%s@%s",
		username,
		password,
		clusterEndpoint,
	)

	// tlsConfig, err := getCustomTLSConfig(caFilePath)
	// if err != nil {
	// 	log.Fatalf("Failed getting TLS configuration: %v", err)
	// }

	// client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI).SetTLSConfig(tlsConfig))
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	fmt.Println("Connected to DocumentDB!")
	m.client = client
	return nil

}

func getCustomTLSConfig(caFile string) (*tls.Config, error) {
	tlsConfig := new(tls.Config)
	certs, err := ioutil.ReadFile(caFile)

	if err != nil {
		return tlsConfig, err
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

	if !ok {
		return tlsConfig, errors.New("Failed parsing pem file")
	}

	return tlsConfig, nil
}

func (m *Mongo) CountAll(db string, collection string, where bson.D) (int64, error) {
	col := m.client.Database(db).Collection(collection)
	count, err := col.CountDocuments(context.TODO(), where)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (m *Mongo) GetLastOne(db string, collection string, opt bson.D) (map[string]interface{}, error) {
	col := m.client.Database(db).Collection(collection)

	// 最新１件を取得する
	findOptions := options.FindOne().SetSort(bson.D{{"_id", -1}})
	var result bson.M
	err := col.FindOne(
		context.Background(),
		opt,
		findOptions,
	).Decode(&result)

	if err != nil {
		return nil, err
	}

	bytesBuf, err := bson.MarshalExtJSON(result, true, true)
	if err != nil {
		return nil, err
	}

	var jsonDocument map[string]interface{}
	json.Unmarshal(bytesBuf, &jsonDocument)

	return jsonDocument, nil
}

func (m *Mongo) InsertOne(db string, collection string, document bson.D) error {
	col := m.client.Database(db).Collection(collection)
	_, err := col.InsertOne(context.Background(), document)
	if err != nil {
		return err
	}

	return nil
}
