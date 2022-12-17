package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"time"
)

const (
	MetadataCollection = "metadata"
)

type MongoWrapperClient struct {
	mongoClient *mongo.Client
	logger      *zap.Logger
}

type DicomMetadata struct {
	ID                string    `bson:"_id"`
	StudyInstanceUID  string    `bson:"study_instance_uid"`
	SeriesInstanceUID string    `bson:"series_instance_uid"`
	SOPInstanceUID    string    `bson:"sop_instance_uid"`
	CreatedAt         time.Time `bson:"created_at"`
	//Metadata          go2com.MappedTag `bson:"metadata"`
}

// NewMongoWrapperClient returns the pointer to a MongoWrapperClient
func NewMongoWrapperClient(logger *zap.Logger, opts ...*options.ClientOptions) (*MongoWrapperClient, error) {
	client, err := mongo.Connect(context.TODO(), opts...)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &MongoWrapperClient{mongoClient: client}, nil
}

func (c *MongoWrapperClient) InsertMetadata() {

}

func (c *MongoWrapperClient) DeleteMetadata() {

}

func (c *MongoWrapperClient) DeleteAllMetadata() {

}

//package main
//
//import (
//"bytes"
//"context"
//"encoding/json"
//"fmt"
//"github.com/klauspost/compress/zstd"
//"github.com/minio/minio-go/v7"
//"github.com/minio/minio-go/v7/pkg/credentials"
//"github.com/okieraised/go2com"
//"github.com/spf13/viper"
//"go.mongodb.org/mongo-driver/bson"
//"go.mongodb.org/mongo-driver/mongo"
//"go.mongodb.org/mongo-driver/mongo/options"
//"go.mongodb.org/mongo-driver/mongo/readpref"
//"io"
//"strings"
//"time"
//"vindr-golang-common/commonitor"
//)
//
//const uri = "mongodb://mongo:qwerty@localhost:27017/"
//
//type MinIOCredential struct {
//	Endpoint string
//	Username string
//	Password string
//}
//
//type MetadataRaw struct {
//	ID                string           `bson:"_id"`
//	StudyInstanceUID  string           `bson:"study_instance_uid"`
//	SeriesInstanceUID string           `bson:"series_instance_uid"`
//	SOPInstanceUID    string           `bson:"sop_instance_uid"`
//	CreatedAt         time.Time        `bson:"created_at"`
//	Metadata          go2com.MappedTag `bson:"metadata"`
//}
//
//type MetadataCompressed struct {
//	ID                string    `bson:"_id"`
//	StudyInstanceUID  string    `bson:"study_instance_uid"`
//	SeriesInstanceUID string    `bson:"series_instance_uid"`
//	SOPInstanceUID    string    `bson:"sop_instance_uid"`
//	CreatedAt         time.Time `bson:"created_at"`
//	Metadata          []byte    `bson:"metadata"`
//}
//
//func main() {
//	CollectionNameRaw := "metadata"
//	CollectionNameCompressed := "metadata_compress"
//
//	serviceName := "mongo_metadata"
//
//	viper.Set("otel.service_name", serviceName)
//	viper.Set("otel.service_env", "dev")
//	viper.Set("otel.service_id", "2509")
//	viper.Set("otel.agent_host", "10.124.68.181")
//	viper.Set("otel.agent_port", "6831")
//
//	sandboxMinIO := &MinIOCredential{
//		Endpoint: "localhost:11014",
//		Username: "orthanc",
//		Password: "av1l4JEx8qHuy2Dg",
//	}
//
//	minioClient, err := minio.New(
//		sandboxMinIO.Endpoint,
//		&minio.Options{
//			Creds: credentials.NewStaticV4(sandboxMinIO.Username, sandboxMinIO.Password, ""),
//		})
//	if err != nil {
//		fmt.Println("Error creating MinIO client:", err)
//		return
//	}
//
//	buckets, err := minioClient.ListBuckets(context.Background())
//	if err != nil {
//		fmt.Println("Error listing bucket:", err)
//		return
//	}
//	fmt.Println("Got buckets", buckets)
//
//	//------------------------------------------------------------------------------------------------------------------
//	// Create a new client and connect to the server
//	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
//	if err != nil {
//		fmt.Println("error connecting to mongodb", err)
//		return
//	}
//	// Ping the primary
//	err = client.Ping(context.TODO(), readpref.Primary())
//	if err != nil {
//		fmt.Println("error pinging mongodb", err)
//		return
//	}
//	fmt.Println("Successfully connected and pinged.")
//
//	metaRawCollection := client.Database("vinlab_metadata").Collection(CollectionNameRaw)
//	metaCompressedCollection := client.Database("vinlab_metadata").Collection(CollectionNameCompressed)
//
//	var encoder, _ = zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedFastest))
//
//	//------------------------------------------------------------------------------------------------------------------
//	bucketName := buckets[0].Name
//
//	prefixes := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
//	//for i, j := 0, len(prefixes)-1; i < j; i, j = i+1, j-1 {
//	//	prefixes[i], prefixes[j] = prefixes[j], prefixes[i]
//	//}
//
//	for _, first := range prefixes {
//		for _, second := range prefixes {
//			objName := first + second + "/"
//			if strings.HasPrefix(objName, "00/") {
//				continue
//			}
//			for obj := range minioClient.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{Recursive: true, Prefix: objName}) {
//
//				//--------------------------------------------------------------------------------------------------------------
//				// Download and process object
//				//_, span = tr.Start(ctx, "createMetadata.DownloadObject")
//				fmt.Println("Processing object", obj.Key)
//				info, err := minioClient.StatObject(context.Background(), bucketName, obj.Key, minio.StatObjectOptions{})
//				if err != nil {
//					fmt.Println("error getting object stat", err)
//					continue
//				}
//
//				downloaded, err := minioClient.GetObject(context.Background(), bucketName, obj.Key, minio.GetObjectOptions{})
//				if err != nil {
//					fmt.Println("error getting object", err)
//					continue
//				}
//
//				b := make([]byte, info.Size)
//				_, err = downloaded.Read(b)
//				if err != nil {
//					if err != io.EOF {
//						fmt.Println("err read", err)
//						continue
//					}
//				}
//				//span.End()
//
//				tp := commonitor.Tracer()
//				tr := tp.Tracer(commonitor.DEFAULT_TRACER_NAME)
//
//				ctx, span := tr.Start(context.Background(), "createMetadata")
//				defer span.End()
//				//--------------------------------------------------------------------------------------------------------------
//				// Parse metadata
//				_, span = tr.Start(ctx, "createMetadata.ParseDICOM")
//				dParser, err := go2com.NewParser(bytes.NewReader(b), int64(len(b)), true, false)
//				err = dParser.Parse()
//				if err != nil {
//					fmt.Println(err)
//					continue
//				}
//				span.End()
//
//				//--------------------------------------------------------------------------------------------------------------
//				// Export metadata
//				_, span = tr.Start(ctx, "createMetadata.ExportMetadataTags")
//				ds := dParser.GetDataset()
//				dUIDs, err := ds.RetrieveFileUID()
//				if err != nil {
//					fmt.Println(err)
//					continue
//				}
//				metadata := dParser.Export(false)
//				span.End()
//
//				bMeta, err := json.Marshal(metadata)
//				compressedMeta := encoder.EncodeAll(bMeta, make([]byte, 0, len(bMeta)))
//
//				//--------------------------------------------------------------------------------------------------------------
//				// Generate raw metadata mongo struct
//				_, span = tr.Start(ctx, "createMetadata.GenerateMongoStructRawJson")
//				documentRaw := MetadataRaw{
//					ID:                strings.Join([]string{dUIDs.StudyInstanceUID, dUIDs.SeriesInstanceUID, dUIDs.SOPInstanceUID}, "_"),
//					StudyInstanceUID:  dUIDs.StudyInstanceUID,
//					SeriesInstanceUID: dUIDs.SeriesInstanceUID,
//					SOPInstanceUID:    dUIDs.SOPInstanceUID,
//					Metadata:          metadata,
//					CreatedAt:         time.Now(),
//				}
//				span.End()
//
//				//--------------------------------------------------------------------------------------------------------------
//				// Insert metadata raw to mongodb
//				_, span = tr.Start(ctx, "createMetadata.InsertToMongoDBRaw")
//				_, err = metaRawCollection.InsertOne(context.TODO(), documentRaw)
//				if err != nil {
//					fmt.Println("error inserting to db", err)
//					continue
//				}
//				span.End()
//
//				//--------------------------------------------------------------------------------------------------------------
//				// Generate compressed metadata mongo struct
//				_, span = tr.Start(ctx, "createMetadata.GenerateMongoStructCompressedJson")
//				documentCompressed := MetadataCompressed{
//					ID:                strings.Join([]string{dUIDs.StudyInstanceUID, dUIDs.SeriesInstanceUID, dUIDs.SOPInstanceUID}, "_"),
//					StudyInstanceUID:  dUIDs.StudyInstanceUID,
//					SeriesInstanceUID: dUIDs.SeriesInstanceUID,
//					SOPInstanceUID:    dUIDs.SOPInstanceUID,
//					Metadata:          compressedMeta,
//					CreatedAt:         time.Now(),
//				}
//				span.End()
//
//				//--------------------------------------------------------------------------------------------------------------
//				// Insert metadata to mongodb
//				_, span = tr.Start(ctx, "createMetadata.InsertToMongoDBCompressed")
//				_, err = metaCompressedCollection.InsertOne(context.TODO(), documentCompressed)
//				if err != nil {
//					fmt.Println("error inserting to db", err)
//					continue
//				}
//				span.End()
//
//			}
//
//		}
//	}
//
//	err = client.Disconnect(context.TODO())
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//}
//
//func queryMetaData(collection *mongo.Collection, StudyInstanceUID, SeriesInstanceUID string) ([]MetadataRaw, error) {
//	filter := bson.D{
//		{
//			"$and",
//			bson.D{
//				{
//					"study_instance_uid",
//					StudyInstanceUID,
//				},
//				{
//					"series_instance_uid",
//					SeriesInstanceUID,
//				},
//			},
//		},
//	}
//	cursor, err := collection.Find(context.TODO(), filter)
//	if err != nil {
//		return nil, err
//	}
//
//	var results []MetadataRaw
//	err = cursor.All(context.TODO(), &results)
//	if err != nil {
//		return nil, err
//	}
//
//	return results, nil
//}

//package main
//
//import (
//"context"
//"fmt"
//"github.com/okieraised/go2com"
//"github.com/spf13/viper"
//"go.mongodb.org/mongo-driver/bson"
//"go.mongodb.org/mongo-driver/mongo"
//"go.mongodb.org/mongo-driver/mongo/options"
//"go.mongodb.org/mongo-driver/mongo/readpref"
//db "labcommon/labrepo"
//"time"
//)
//
//const (
//	PostgresSandbox = "postgresql://postgres:DC6vEyzuXLe7RCg@localhost:54360/backend__vinlab__io?sslmode=disable"
//	uri             = "mongodb://mongo:qwerty@localhost:27017/"
//)
//
//type MinIOCredential struct {
//	Endpoint string
//	Username string
//	Password string
//}
//
//type MetadataRaw struct {
//	ID                string           `bson:"_id"`
//	StudyInstanceUID  string           `bson:"study_instance_uid"`
//	SeriesInstanceUID string           `bson:"series_instance_uid"`
//	SOPInstanceUID    string           `bson:"sop_instance_uid"`
//	CreatedAt         time.Time        `bson:"created_at"`
//	Metadata          go2com.MappedTag `bson:"metadata"`
//}
//
//type MetadataCompressed struct {
//	ID                string    `bson:"_id"`
//	StudyInstanceUID  string    `bson:"study_instance_uid"`
//	SeriesInstanceUID string    `bson:"series_instance_uid"`
//	SOPInstanceUID    string    `bson:"sop_instance_uid"`
//	CreatedAt         time.Time `bson:"created_at"`
//	Metadata          []byte    `bson:"metadata"`
//}
//
//type MetaAggregate struct {
//	StudyInstanceUID  string   `json:"study_instance_uid"`
//	SeriesInstanceUID string   `json:"series_instance_uid"`
//	SOPInstanceUIDs   []string `pg:"sop_instance_uids,array" json:"sop_instance_uids"`
//}
//
//func main() {
//	db.InitLabellingDB(PostgresSandbox)
//	defer db.Labelling().Close()
//
//	CollectionNameRaw := "metadata"
//	//CollectionNameCompressed := "metadata_compress"
//
//	serviceName := "mongo_metadata"
//
//	viper.Set("otel.service_name", serviceName)
//	viper.Set("otel.service_env", "dev")
//	viper.Set("otel.service_id", "2509")
//	viper.Set("otel.agent_host", "10.124.68.181")
//	viper.Set("otel.agent_port", "6831")
//
//	//------------------------------------------------------------------------------------------------------------------
//	projectIDs := []string{}
//	err := db.Labelling().Model(new(db.Project)).Column("id").Select(&projectIDs)
//	if err != nil {
//		fmt.Println("error getting project_id", err)
//		return
//	}
//
//	for _, projectID := range projectIDs {
//		fmt.Println("Processing project", projectID)
//		metadata := []MetaAggregate{}
//		err := db.Labelling().Model(new(db.Object)).ColumnExpr("study_instance_uid, series_instance_uid, array_agg(sop_instance_uid) as sop_instance_uids").
//			Where("project_id = ? and type = 'IMAGE'", projectID).
//			GroupExpr("project_id, study_instance_uid, series_instance_uid").
//			Select(&metadata)
//		if err != nil {
//			fmt.Println("error aggregating objects", err)
//			continue
//		}
//
//		return
//
//	}
//	return
//
//	//------------------------------------------------------------------------------------------------------------------
//	// Create a new client and connect to the server
//	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri), options.Client().SetCompressors([]string{"zstd"}), options.Client().SetZstdLevel(2)) // options.Client().ApplyURI(uri)
//	if err != nil {
//		fmt.Println("error connecting to mongodb", err)
//		return
//	}
//	// Ping the primary
//	err = client.Ping(context.TODO(), readpref.Primary())
//	if err != nil {
//		fmt.Println("error pinging mongodb", err)
//		return
//	}
//	fmt.Println("Successfully connected and pinged.")
//
//	//------------------------------------------------------------------------------------------------------------------
//	StudyInstanceUID := "1.2.840.113619.2.278.3.717616.668.1598576641.217"
//	SeriesInstanceUID := "1.2.840.113619.2.278.3.717616.668.1598576641.315.4198401"
//
//	metaRawCollection := client.Database("vinlab_metadata").Collection(CollectionNameRaw)
//	//indexModel := mongo.IndexModel{
//	//	Keys: bson.D{
//	//		{
//	//			"study_instance_uid",
//	//			1,
//	//		},
//	//		{
//	//			"series_instance_uid",
//	//			1,
//	//		},
//	//	},
//	//}
//	//name, err := metaRawCollection.Indexes().CreateOne(context.TODO(), indexModel)
//	//if err != nil {
//	//	fmt.Println("error creating index", err)
//	//	return
//	//}
//	//fmt.Println("Name of Index Created: " + name)
//
//	//------------------------------------------------------------------------------------------------------------------
//	filter := bson.D{
//		{
//			"study_instance_uid",
//			StudyInstanceUID,
//		},
//		{
//			"series_instance_uid",
//			SeriesInstanceUID,
//		},
//		//{
//		//	"sop_instance_uid",
//		//	bson.M{
//		//		"$in": []string{"1.2.840.113619.2.278.3.717616.668.1598576641.500.251", "1.2.840.113619.2.278.3.717616.668.1598576641.500.154"},
//		//	},
//		//},
//	}
//
//	//1.2.840.113619.2.278.3.717616.668.1598576641.500.251
//	//1.2.840.113619.2.278.3.717616.668.1598576641.500.154
//	//1.2.840.113619.2.278.3.717616.668.1598576641.500.264
//	queryOpts := options.Find().SetProjection(bson.D{{"metadata", 1}, {"sop_instance_uid", 1}})
//	cursor, err := metaRawCollection.Find(context.TODO(), filter, queryOpts)
//	if err != nil {
//		fmt.Println("error finding metadata", err)
//		return
//	}
//
//	var results []MetadataRaw
//	err = cursor.All(context.TODO(), &results)
//	if err != nil {
//		fmt.Println("error retrieving metadata", err)
//		return
//	}
//
//	//fmt.Println("results", results)
//
//	fmt.Println("Got", len(results))
//
//	for _, result := range results {
//		fmt.Println(result.SOPInstanceUID)
//	}
//
//	//------------------------------------------------------------------------------------------------------------------
//	command := bson.D{{"dbStats", 1}}
//	var result bson.M
//	err = client.Database("vinlab_metadata").RunCommand(context.TODO(), command).Decode(&result)
//	if err != nil {
//		fmt.Println("error retrieving db stat", err)
//		return
//	}
//
//	fmt.Println("result", result)
//
//	err = client.Disconnect(context.TODO())
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//}
//
////type MongoClient struct {
////	client     *mongo.Client
////	database   *mongo.Database
////	collection *mongo.Collection
////	cursor     *mongo.Cursor
////}
////
////func NewMongoClient(opts ...*options.ClientOptions) (*MongoClient, error) {
////	ctx := context.Background()
////	client, err := mongo.Connect(ctx, opts...)
////	if err != nil {
////		return nil, err
////	}
////
////	mClient := &MongoClient{
////		client: client,
////	}
////
////	return mClient, nil
////}
////
////func (c *MongoClient) Database(dbName string, opts ...*options.DatabaseOptions) {
////	c.database = c.client.Database(dbName, opts...)
////}
////
////func (c *MongoClient) Collection(name string, opts ...*options.CollectionOptions) {
////	c.collection = c.database.Collection(name, opts...)
////}
////
////func (c *MongoClient) Find(filter interface{}, opts ...*options.FindOptions) error {
////	ctx := context.Background()
////	cursor, err := c.collection.Find(ctx, filter, opts...)
////	if err != nil {
////		return err
////	}
////
////	c.cursor = cursor
////	return nil
////}
//package main
//
//import (
//"context"
//"fmt"
//"github.com/okieraised/go2com"
//"github.com/spf13/viper"
//"go.mongodb.org/mongo-driver/bson"
//"go.mongodb.org/mongo-driver/mongo"
//"go.mongodb.org/mongo-driver/mongo/options"
//"go.mongodb.org/mongo-driver/mongo/readpref"
//db "labcommon/labrepo"
//"time"
//)
//
//const (
//	PostgresSandbox = "postgresql://postgres:DC6vEyzuXLe7RCg@localhost:54360/backend__vinlab__io?sslmode=disable"
//	uri             = "mongodb://mongo:qwerty@localhost:27017/"
//)
//
//type MinIOCredential struct {
//	Endpoint string
//	Username string
//	Password string
//}
//
//type MetadataRaw struct {
//	ID                string           `bson:"_id"`
//	StudyInstanceUID  string           `bson:"study_instance_uid"`
//	SeriesInstanceUID string           `bson:"series_instance_uid"`
//	SOPInstanceUID    string           `bson:"sop_instance_uid"`
//	CreatedAt         time.Time        `bson:"created_at"`
//	Metadata          go2com.MappedTag `bson:"metadata"`
//}
//
//type MetadataCompressed struct {
//	ID                string    `bson:"_id"`
//	StudyInstanceUID  string    `bson:"study_instance_uid"`
//	SeriesInstanceUID string    `bson:"series_instance_uid"`
//	SOPInstanceUID    string    `bson:"sop_instance_uid"`
//	CreatedAt         time.Time `bson:"created_at"`
//	Metadata          []byte    `bson:"metadata"`
//}
//
//type MetaAggregate struct {
//	StudyInstanceUID  string   `json:"study_instance_uid"`
//	SeriesInstanceUID string   `json:"series_instance_uid"`
//	SOPInstanceUIDs   []string `pg:"sop_instance_uids,array" json:"sop_instance_uids"`
//}
//
//func main() {
//	db.InitLabellingDB(PostgresSandbox)
//	defer db.Labelling().Close()
//
//	CollectionNameRaw := "metadata"
//	//CollectionNameCompressed := "metadata_compress"
//
//	serviceName := "mongo_metadata"
//
//	viper.Set("otel.service_name", serviceName)
//	viper.Set("otel.service_env", "dev")
//	viper.Set("otel.service_id", "2509")
//	viper.Set("otel.agent_host", "10.124.68.181")
//	viper.Set("otel.agent_port", "6831")
//
//	//------------------------------------------------------------------------------------------------------------------
//	projectIDs := []string{}
//	err := db.Labelling().Model(new(db.Project)).Column("id").Select(&projectIDs)
//	if err != nil {
//		fmt.Println("error getting project_id", err)
//		return
//	}
//
//	for _, projectID := range projectIDs {
//		fmt.Println("Processing project", projectID)
//		metadata := []MetaAggregate{}
//		err := db.Labelling().Model(new(db.Object)).ColumnExpr("study_instance_uid, series_instance_uid, array_agg(sop_instance_uid) as sop_instance_uids").
//			Where("project_id = ? and type = 'IMAGE'", projectID).
//			GroupExpr("project_id, study_instance_uid, series_instance_uid").
//			Select(&metadata)
//		if err != nil {
//			fmt.Println("error aggregating objects", err)
//			continue
//		}
//
//		return
//
//	}
//	return
//
//	//------------------------------------------------------------------------------------------------------------------
//	// Create a new client and connect to the server
//	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri), options.Client().SetCompressors([]string{"zstd"}), options.Client().SetZstdLevel(2)) // options.Client().ApplyURI(uri)
//	if err != nil {
//		fmt.Println("error connecting to mongodb", err)
//		return
//	}
//	// Ping the primary
//	err = client.Ping(context.TODO(), readpref.Primary())
//	if err != nil {
//		fmt.Println("error pinging mongodb", err)
//		return
//	}
//	fmt.Println("Successfully connected and pinged.")
//
//	//------------------------------------------------------------------------------------------------------------------
//	StudyInstanceUID := "1.2.840.113619.2.278.3.717616.668.1598576641.217"
//	SeriesInstanceUID := "1.2.840.113619.2.278.3.717616.668.1598576641.315.4198401"
//
//	metaRawCollection := client.Database("vinlab_metadata").Collection(CollectionNameRaw)
//	//indexModel := mongo.IndexModel{
//	//	Keys: bson.D{
//	//		{
//	//			"study_instance_uid",
//	//			1,
//	//		},
//	//		{
//	//			"series_instance_uid",
//	//			1,
//	//		},
//	//	},
//	//}
//	//name, err := metaRawCollection.Indexes().CreateOne(context.TODO(), indexModel)
//	//if err != nil {
//	//	fmt.Println("error creating index", err)
//	//	return
//	//}
//	//fmt.Println("Name of Index Created: " + name)
//
//	//------------------------------------------------------------------------------------------------------------------
//	filter := bson.D{
//		{
//			"study_instance_uid",
//			StudyInstanceUID,
//		},
//		{
//			"series_instance_uid",
//			SeriesInstanceUID,
//		},
//		//{
//		//	"sop_instance_uid",
//		//	bson.M{
//		//		"$in": []string{"1.2.840.113619.2.278.3.717616.668.1598576641.500.251", "1.2.840.113619.2.278.3.717616.668.1598576641.500.154"},
//		//	},
//		//},
//	}
//
//	//1.2.840.113619.2.278.3.717616.668.1598576641.500.251
//	//1.2.840.113619.2.278.3.717616.668.1598576641.500.154
//	//1.2.840.113619.2.278.3.717616.668.1598576641.500.264
//	queryOpts := options.Find().SetProjection(bson.D{{"metadata", 1}, {"sop_instance_uid", 1}})
//	cursor, err := metaRawCollection.Find(context.TODO(), filter, queryOpts)
//	if err != nil {
//		fmt.Println("error finding metadata", err)
//		return
//	}
//
//	var results []MetadataRaw
//	err = cursor.All(context.TODO(), &results)
//	if err != nil {
//		fmt.Println("error retrieving metadata", err)
//		return
//	}
//
//	//fmt.Println("results", results)
//
//	fmt.Println("Got", len(results))
//
//	for _, result := range results {
//		fmt.Println(result.SOPInstanceUID)
//	}
//
//	//------------------------------------------------------------------------------------------------------------------
//	command := bson.D{{"dbStats", 1}}
//	var result bson.M
//	err = client.Database("vinlab_metadata").RunCommand(context.TODO(), command).Decode(&result)
//	if err != nil {
//		fmt.Println("error retrieving db stat", err)
//		return
//	}
//
//	fmt.Println("result", result)
//
//	err = client.Disconnect(context.TODO())
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//}

//type MongoClient struct {
//	client     *mongo.Client
//	database   *mongo.Database
//	collection *mongo.Collection
//	cursor     *mongo.Cursor
//}
//
//func NewMongoClient(opts ...*options.ClientOptions) (*MongoClient, error) {
//	ctx := context.Background()
//	client, err := mongo.Connect(ctx, opts...)
//	if err != nil {
//		return nil, err
//	}
//
//	mClient := &MongoClient{
//		client: client,
//	}
//
//	return mClient, nil
//}
//
//func (c *MongoClient) Database(dbName string, opts ...*options.DatabaseOptions) {
//	c.database = c.client.Database(dbName, opts...)
//}
//
//func (c *MongoClient) Collection(name string, opts ...*options.CollectionOptions) {
//	c.collection = c.database.Collection(name, opts...)
//}
//
//func (c *MongoClient) Find(filter interface{}, opts ...*options.FindOptions) error {
//	ctx := context.Background()
//	cursor, err := c.collection.Find(ctx, filter, opts...)
//	if err != nil {
//		return err
//	}
//
//	c.cursor = cursor
//	return nil
//}
