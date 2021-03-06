package gvabe

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/btnguyen2k/godal"
	"github.com/btnguyen2k/henge"
	"github.com/btnguyen2k/prom"
	"main/src/gvabe/bo/libro"
	"main/src/gvabe/bo/user"
	"main/src/respicite"

	"main/src/goapi"
	"main/src/utils"
)

func _createDynamodbConnect(dbtype string) *prom.AwsDynamodbConnect {
	var adc *prom.AwsDynamodbConnect = nil
	var err error
	switch dbtype {
	case "dynamo", "dynamodb", "awsdynamo", "awsdynamodb":
		region := goapi.AppConfig.GetString("gvabe.db.dynamodb.region")
		region = strings.ReplaceAll(region, `"`, "")
		cfg := &aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewEnvCredentials(),
		}
		endpoint := goapi.AppConfig.GetString("gvabe.db.dynamodb.endpoint")
		endpoint = strings.ReplaceAll(endpoint, `"`, "")
		if endpoint != "" {
			cfg.Endpoint = aws.String(endpoint)
			if strings.HasPrefix(strings.ToLower(endpoint), "http://") {
				cfg.DisableSSL = aws.Bool(true)
			}
		}
		adc, err = prom.NewAwsDynamodbConnect(cfg, nil, nil, 10000)
	}
	if err != nil {
		panic(err)
	}
	return adc
}

func _createSqlConnect(dbtype string) *prom.SqlConnect {
	timezone := goapi.AppConfig.GetString("timezone")
	var sqlc *prom.SqlConnect = nil
	var err error
	switch dbtype {
	case "sqlite":
		dir := goapi.AppConfig.GetString("gvabe.db.sqlite.directory")
		dbname := goapi.AppConfig.GetString("gvabe.db.sqlite.dbname")
		sqlc, err = henge.NewSqliteConnection(dir, dbname, timezone, "sqlite3", 10000, nil)
	case "pg", "pgsql", "postgres", "postgresql":
		url := goapi.AppConfig.GetString("gvabe.db.pgsql.url")
		sqlc, err = henge.NewPgsqlConnection(url, timezone, "pgx", 10000, nil)
	case "mysql":
		url := goapi.AppConfig.GetString("gvabe.db.mysql.url")
		sqlc, err = henge.NewMysqlConnection(url, timezone, "mysql", 10000, nil)
	case "cosmos", "cosmosdb":
		url := goapi.AppConfig.GetString("gvabe.db.cosmosdb.url")
		sqlc, err = henge.NewCosmosdbConnection(url, timezone, "gocosmos", 10000, nil)
	}
	if err != nil {
		panic(err)
	}
	return sqlc
}

func _createMongoConnect(dbtype string) *prom.MongoConnect {
	var mc *prom.MongoConnect = nil
	var err error
	switch dbtype {
	case "mongo", "mongodb":
		db := goapi.AppConfig.GetString("gvabe.db.mongodb.db")
		url := goapi.AppConfig.GetString("gvabe.db.mongodb.url")
		// poolOpts := &prom.MongoPoolOpts{
		// 	ConnectTimeout:         15 * time.Second,
		// 	SocketTimeout:          5 * time.Second,
		// 	ServerSelectionTimeout: 10 * time.Second,
		// 	MaxPoolSize:            4,
		// 	MinPoolSize:            1,
		// }
		// poolOpts = prom.MongoPoolOptsLongDistance
		// mc, err = prom.NewMongoConnectWithPoolOptions(url, db, 10000, poolOpts)
		mc, err = prom.NewMongoConnect(url, db, 10000)
	}
	if err != nil {
		panic(err)
	}
	return mc
}

func _createM2oMappingDaoMongo(mc *prom.MongoConnect, collectionName string) respicite.M2oDao {
	url := mc.GetUrl()
	return respicite.NewM2oDaoMongo(mc, collectionName, strings.Index(url, "replicaSet=") >= 0)
}

func _createUserDaoSql(sqlc *prom.SqlConnect) user.UserDao {
	if sqlc.GetDbFlavor() == prom.FlavorCosmosDb {
		return user.NewUserDaoCosmosdb(sqlc, user.TableUser, true)
	}
	return user.NewUserDaoSql(sqlc, user.TableUser, true)
}
func _createUserDaoDynamodb(adc *prom.AwsDynamodbConnect) user.UserDao {
	return user.NewUserDaoDynamodb(adc, user.TableUser)
}
func _createUserDaoMongo(mc *prom.MongoConnect) user.UserDao {
	url := mc.GetUrl()
	return user.NewUserDaoMongo(mc, user.TableUser, strings.Index(url, "replicaSet=") >= 0)
}

func _createProductDaoSql(sqlc *prom.SqlConnect) libro.ProductDao {
	if sqlc.GetDbFlavor() == prom.FlavorCosmosDb {
		return libro.NewProductDaoCosmosdb(sqlc, libro.TableProduct, true)
	}
	return libro.NewProductDaoSql(sqlc, libro.TableProduct, true)
}
func _createProductDaoDynamodb(adc *prom.AwsDynamodbConnect) libro.ProductDao {
	return libro.NewProductDaoDynamodb(adc, libro.TableProduct)
}
func _createProductDaoMongo(mc *prom.MongoConnect) libro.ProductDao {
	url := mc.GetUrl()
	return libro.NewProductDaoMongo(mc, libro.TableProduct, strings.Index(url, "replicaSet=") >= 0)
}

func _createTopicDaoSql(sqlc *prom.SqlConnect) libro.TopicDao {
	if sqlc.GetDbFlavor() == prom.FlavorCosmosDb {
		return libro.NewTopicDaoCosmosdb(sqlc, libro.TableTopic, true)
	}
	return libro.NewTopicDaoSql(sqlc, libro.TableTopic, true)
}
func _createTopicDaoDynamodb(adc *prom.AwsDynamodbConnect) libro.TopicDao {
	return libro.NewTopicDaoDynamodb(adc, libro.TableTopic)
}
func _createTopicDaoMongo(mc *prom.MongoConnect) libro.TopicDao {
	url := mc.GetUrl()
	return libro.NewTopicDaoMongo(mc, libro.TableTopic, strings.Index(url, "replicaSet=") >= 0)
}

func _createPageDaoSql(sqlc *prom.SqlConnect) libro.PageDao {
	if sqlc.GetDbFlavor() == prom.FlavorCosmosDb {
		return libro.NewPageDaoCosmosdb(sqlc, libro.TablePage, true)
	}
	return libro.NewPageDaoSql(sqlc, libro.TablePage, true)
}
func _createPageDaoDynamodb(adc *prom.AwsDynamodbConnect) libro.PageDao {
	return libro.NewPageDaoDynamodb(adc, libro.TablePage)
}
func _createPageDaoMongo(mc *prom.MongoConnect) libro.PageDao {
	url := mc.GetUrl()
	return libro.NewPageDaoMongo(mc, libro.TablePage, strings.Index(url, "replicaSet=") >= 0)
}

// var _sqliteTableSchema = map[string]map[string]string{
// 	user.TableUser:        {user.UserColMaskUid: "VARCHAR(32)"},
// 	blog.TableBlogPost:    {blog.PostColOwnerId: "VARCHAR(32)", blog.PostColIsPublic: "INT"},
// 	blog.TableBlogComment: {blog.CommentColOwnerId: "VARCHAR(32)", blog.CommentColPostId: "VARCHAR(32)", blog.CommentColParentId: "VARCHAR(32)"},
// 	blog.TableBlogVote:    {blog.VoteColOwnerId: "VARCHAR(32)", blog.VoteColTargetId: "VARCHAR(32)", blog.VoteColValue: "INT"},
// }

// var _pgsqlTableSchema = map[string]map[string]string{
// 	user.TableUser:        {user.UserColMaskUid: "VARCHAR(32)"},
// 	blog.TableBlogPost:    {blog.PostColOwnerId: "VARCHAR(32)", blog.PostColIsPublic: "INT"},
// 	blog.TableBlogComment: {blog.CommentColOwnerId: "VARCHAR(32)", blog.CommentColPostId: "VARCHAR(32)", blog.CommentColParentId: "VARCHAR(32)"},
// 	blog.TableBlogVote:    {blog.VoteColOwnerId: "VARCHAR(32)", blog.VoteColTargetId: "VARCHAR(32)", blog.VoteColValue: "INT"},
// }

// var _cosmosdbTableSpec = map[string]*henge.CosmosdbCollectionSpec{
// 	user.TableUser:        {Pk: henge.CosmosdbColId, Uk: [][]string{{"/" + user.UserFieldMaskId}}},
// 	blog.TableBlogPost:    {Pk: henge.CosmosdbColId},
// 	blog.TableBlogComment: {Pk: henge.CosmosdbColId},
// 	blog.TableBlogVote:    {Pk: henge.CosmosdbColId, Uk: [][]string{{"/" + blog.VoteFieldOwnerId, "/" + blog.VoteFieldTargetId}}},
// }

// func _createSqlTables(sqlc *prom.SqlConnect, dbtype string) {
// 	switch sqlc.GetDbFlavor() {
// 	case prom.FlavorSqlite:
// 		for tbl, schema := range _sqliteTableSchema {
// 			if err := henge.InitSqliteTable(sqlc, tbl, schema); err != nil {
// 				log.Printf("[WARN] creating table %s (%s): %s\n", tbl, dbtype, err)
// 			}
// 		}
// 	case prom.FlavorPgSql:
// 		for tbl, schema := range _pgsqlTableSchema {
// 			if err := henge.InitSqliteTable(sqlc, tbl, schema); err != nil {
// 				log.Printf("[WARN] creating table %s (%s): %s\n", tbl, dbtype, err)
// 			}
// 		}
// 	case prom.FlavorCosmosDb:
// 		for tbl, spec := range _cosmosdbTableSpec {
// 			if err := henge.InitCosmosdbCollection(sqlc, tbl, spec); err != nil {
// 				log.Printf("[WARN] creating table %s (%s): %s\n", tbl, dbtype, err)
// 			}
// 		}
// 	}
//
// 	if sqlc.GetDbFlavor() == prom.FlavorCosmosDb {
// 		return
// 	}
//
// 	// user
// 	if err := henge.CreateIndexSql(sqlc, user.TableUser, true, []string{user.UserColMaskUid}); err != nil {
// 		log.Printf("[WARN] creating table index %s/%s (%s): %s\n", user.TableUser, user.UserColMaskUid, dbtype, err)
// 	}
//
// 	// blog post
// 	if err := henge.CreateIndexSql(sqlc, blog.TableBlogPost, false, []string{blog.PostColOwnerId}); err != nil {
// 		log.Printf("[WARN] creating table index %s/%s (%s): %s\n", blog.TableBlogPost, blog.PostColOwnerId, dbtype, err)
// 	}
// 	if err := henge.CreateIndexSql(sqlc, blog.TableBlogPost, false, []string{blog.PostColIsPublic}); err != nil {
// 		log.Printf("[WARN] creating table index %s/%s (%s): %s\n", blog.TableBlogPost, blog.PostColIsPublic, dbtype, err)
// 	}
//
// 	// blog comment
// 	if err := henge.CreateIndexSql(sqlc, blog.TableBlogComment, false, []string{blog.CommentColOwnerId}); err != nil {
// 		log.Printf("[WARN] creating table index %s/%s (%s): %s\n", blog.TableBlogComment, blog.CommentColOwnerId, dbtype, err)
// 	}
// 	if err := henge.CreateIndexSql(sqlc, blog.TableBlogComment, false, []string{blog.CommentColPostId, blog.CommentColParentId}); err != nil {
// 		log.Printf("[WARN] creating table index %s/%s (%s): %s\n", blog.TableBlogComment, blog.CommentColPostId+":"+blog.CommentColParentId, dbtype, err)
// 	}
//
// 	// blog vote
// 	if err := henge.CreateIndexSql(sqlc, blog.TableBlogVote, true, []string{blog.VoteColOwnerId, blog.VoteColTargetId}); err != nil {
// 		log.Printf("[WARN] creating table index %s/%s (%s): %s\n", blog.TableBlogVote, blog.VoteColOwnerId+":"+blog.VoteColTargetId, dbtype, err)
// 	}
// 	if err := henge.CreateIndexSql(sqlc, blog.TableBlogVote, false, []string{blog.VoteColTargetId, blog.VoteColValue}); err != nil {
// 		log.Printf("[WARN] creating table index %s/%s (%s): %s\n", blog.TableBlogVote, blog.VoteColTargetId+":"+blog.VoteColValue, dbtype, err)
// 	}
// }

// func _dynamodbWaitforGSI(adc *prom.AwsDynamodbConnect, table, gsi string, timeout time.Duration) error {
// 	t := time.Now()
// 	for status, err := adc.GetGlobalSecondaryIndexStatus(nil, table, gsi); ; {
// 		if err != nil {
// 			return err
// 		}
// 		if strings.ToUpper(status) == "ACTIVE" {
// 			return nil
// 		}
// 		if time.Now().Sub(t).Milliseconds() > timeout.Milliseconds() {
// 			return errors.New("")
// 		}
// 	}
// }

// func _createDynamodbTables(adc *prom.AwsDynamodbConnect) {
// 	spec := &henge.DynamodbTablesSpec{MainTableRcu: 2, MainTableWcu: 1, CreateUidxTable: true, UidxTableRcu: 2, UidxTableWcu: 1}
// 	if err := henge.InitDynamodbTables(adc, user.TableUser, spec); err != nil {
// 		log.Printf("[WARN] creating tableName %s (%s): %s\n", user.TableUser, "DynamoDB", err)
// 	}
// 	spec = &henge.DynamodbTablesSpec{MainTableRcu: 2, MainTableWcu: 1}
// 	if err := henge.InitDynamodbTables(adc, blog.TableBlogPost, spec); err != nil {
// 		log.Printf("[WARN] creating tableName %s (%s): %s\n", blog.TableBlogPost, "DynamoDB", err)
// 	}
// 	spec = &henge.DynamodbTablesSpec{MainTableRcu: 2, MainTableWcu: 1}
// 	if err := henge.InitDynamodbTables(adc, blog.TableBlogComment, spec); err != nil {
// 		log.Printf("[WARN] creating tableName %s (%s): %s\n", blog.TableBlogComment, "DynamoDB", err)
// 	}
// 	spec = &henge.DynamodbTablesSpec{MainTableRcu: 2, MainTableWcu: 1, CreateUidxTable: true, UidxTableRcu: 2, UidxTableWcu: 1}
// 	if err := henge.InitDynamodbTables(adc, blog.TableBlogVote, spec); err != nil {
// 		log.Printf("[WARN] creating tableName %s (%s): %s\n", blog.TableBlogVote, "DynamoDB", err)
// 	}
//
// 	var tableName, gsiName, colName string
//
// 	// user
// 	tableName, colName, gsiName = user.TableUser, user.UserFieldMaskId, "gsi_"+colName
// 	if err := adc.CreateGlobalSecondaryIndex(nil, tableName, gsiName, 2, 1,
// 		[]prom.AwsDynamodbNameAndType{{Name: colName, Type: prom.AwsAttrTypeString}},
// 		[]prom.AwsDynamodbNameAndType{{Name: colName, Type: prom.AwsKeyTypePartition}}); err != nil {
// 		log.Printf("[WARN] creating GSI %s/%s (%s): %s\n", tableName, colName, "DynamoDB", err)
// 	} else if err := _dynamodbWaitforGSI(adc, tableName, gsiName, 10*time.Second); err != nil {
// 		log.Printf("[WARN] creating GSI %s/%s (%s): %s\n", tableName, colName, "DynamoDB", err)
// 	}
//
// 	// blog post
// 	tableName, colName, gsiName = blog.TableBlogPost, blog.PostFieldOwnerId, "gsi_"+colName
// 	if err := adc.CreateGlobalSecondaryIndex(nil, tableName, gsiName, 2, 1,
// 		[]prom.AwsDynamodbNameAndType{{Name: colName, Type: prom.AwsAttrTypeString}},
// 		[]prom.AwsDynamodbNameAndType{{Name: colName, Type: prom.AwsKeyTypePartition}}); err != nil {
// 		log.Printf("[WARN] creating GSI %s/%s (%s): %s\n", tableName, colName, "DynamoDB", err)
// 	} else if err := _dynamodbWaitforGSI(adc, tableName, gsiName, 10*time.Second); err != nil {
// 		log.Printf("[WARN] creating GSI %s/%s (%s): %s\n", tableName, colName, "DynamoDB", err)
// 	}
// 	tableName, colName, gsiName = blog.TableBlogPost, blog.PostFieldIsPublic, "gsi_"+colName
// 	if err := adc.CreateGlobalSecondaryIndex(nil, tableName, gsiName, 2, 1,
// 		[]prom.AwsDynamodbNameAndType{{Name: colName, Type: prom.AwsAttrTypeNumber}},
// 		[]prom.AwsDynamodbNameAndType{{Name: colName, Type: prom.AwsKeyTypePartition}}); err != nil {
// 		log.Printf("[WARN] creating GSI %s/%s (%s): %s\n", tableName, colName, "DynamoDB", err)
// 	} else if err := _dynamodbWaitforGSI(adc, tableName, gsiName, 10*time.Second); err != nil {
// 		log.Printf("[WARN] creating GSI %s/%s (%s): %s\n", tableName, colName, "DynamoDB", err)
// 	}
//
// 	// blog comment
// 	tableName, colName, gsiName = blog.TableBlogComment, blog.CommentFieldOwnerId, "gsi_"+colName
// 	if err := adc.CreateGlobalSecondaryIndex(nil, tableName, gsiName, 2, 1,
// 		[]prom.AwsDynamodbNameAndType{{Name: colName, Type: prom.AwsAttrTypeString}},
// 		[]prom.AwsDynamodbNameAndType{{Name: colName, Type: prom.AwsKeyTypePartition}}); err != nil {
// 		log.Printf("[WARN] creating GSI %s/%s (%s): %s\n", tableName, colName, "DynamoDB", err)
// 	} else if err := _dynamodbWaitforGSI(adc, tableName, gsiName, 10*time.Second); err != nil {
// 		log.Printf("[WARN] creating GSI %s/%s (%s): %s\n", tableName, colName, "DynamoDB", err)
// 	}
// 	tableName, colName, gsiName = blog.TableBlogComment, blog.CommentFieldPostId+"_"+blog.CommentFieldParentId, "gsi_"+colName
// 	if err := adc.CreateGlobalSecondaryIndex(nil, tableName, gsiName, 2, 1,
// 		[]prom.AwsDynamodbNameAndType{{Name: blog.CommentFieldPostId, Type: prom.AwsAttrTypeString}, {Name: blog.CommentFieldParentId, Type: prom.AwsAttrTypeString}},
// 		[]prom.AwsDynamodbNameAndType{{Name: blog.CommentFieldPostId, Type: prom.AwsKeyTypePartition}, {Name: blog.CommentFieldParentId, Type: prom.AwsKeyTypeSort}}); err != nil {
// 		log.Printf("[WARN] creating GSI %s/%s (%s): %s\n", tableName, colName, "DynamoDB", err)
// 	} else if err := _dynamodbWaitforGSI(adc, tableName, gsiName, 10*time.Second); err != nil {
// 		log.Printf("[WARN] creating GSI %s/%s (%s): %s\n", tableName, colName, "DynamoDB", err)
// 	}
//
// 	// blog vote
// 	tableName, colName, gsiName = blog.TableBlogVote, blog.VoteFieldOwnerId+"_"+blog.VoteFieldTargetId, "gsi_"+colName
// 	if err := adc.CreateGlobalSecondaryIndex(nil, tableName, gsiName, 2, 1,
// 		[]prom.AwsDynamodbNameAndType{{Name: blog.VoteFieldOwnerId, Type: prom.AwsAttrTypeString}, {Name: blog.VoteFieldTargetId, Type: prom.AwsAttrTypeString}},
// 		[]prom.AwsDynamodbNameAndType{{Name: blog.VoteFieldOwnerId, Type: prom.AwsKeyTypePartition}, {Name: blog.VoteFieldTargetId, Type: prom.AwsKeyTypeSort}}); err != nil {
// 		log.Printf("[WARN] creating GSI %s/%s (%s): %s\n", tableName, colName, "DynamoDB", err)
// 	} else if err := _dynamodbWaitforGSI(adc, tableName, gsiName, 10*time.Second); err != nil {
// 		log.Printf("[WARN] creating GSI %s/%s (%s): %s\n", tableName, colName, "DynamoDB", err)
// 	}
// 	tableName, colName, gsiName = blog.TableBlogVote, blog.VoteFieldTargetId+"_"+blog.VoteFieldValue, "gsi_"+colName
// 	if err := adc.CreateGlobalSecondaryIndex(nil, tableName, gsiName, 2, 1,
// 		[]prom.AwsDynamodbNameAndType{{Name: blog.VoteFieldTargetId, Type: prom.AwsAttrTypeString}, {Name: blog.VoteFieldValue, Type: prom.AwsAttrTypeNumber}},
// 		[]prom.AwsDynamodbNameAndType{{Name: blog.VoteFieldTargetId, Type: prom.AwsKeyTypePartition}, {Name: blog.VoteFieldValue, Type: prom.AwsKeyTypeSort}}); err != nil {
// 		log.Printf("[WARN] creating GSI %s/%s (%s): %s\n", tableName, colName, "DynamoDB", err)
// 	} else if err := _dynamodbWaitforGSI(adc, tableName, gsiName, 10*time.Second); err != nil {
// 		log.Printf("[WARN] creating GSI %s/%s (%s): %s\n", tableName, colName, "DynamoDB", err)
// 	}
// }

// TODO change this function to implement application's business logic
func _createMongoCollections(mc *prom.MongoConnect) {
	if err := respicite.M2oDaoMongoInitCollection(mc, tblMappingDomainProduct); err != nil {
		log.Printf("[WARN] creating collection %s (%s): %s\n", tblMappingDomainProduct, "MongoDB", err)
	}

	user.InitUserTableMongo(mc, user.TableUser)
	libro.InitProductTableMongo(mc, libro.TableProduct)
	libro.InitTopicTableMongo(mc, libro.TableTopic)
	libro.InitPageTableMongo(mc, libro.TablePage)
}

// TODO change this function to implement application's business logic
func initDaos() {
	dbtype := strings.ToLower(goapi.AppConfig.GetString("gvabe.db.type"))

	/* TODO support only MongoDB for now */

	// create DB connect instance
	// sqlc := _createSqlConnect(dbtype)
	mc := _createMongoConnect(dbtype)
	// adc := _createDynamodbConnect(dbtype)
	// if sqlc == nil && mc == nil && adc == nil {
	// 	panic(fmt.Sprintf("unknown database type: %s", dbtype))
	// }
	if mc == nil {
		panic(fmt.Sprintf("unknown database type: %s", dbtype))
	}

	// if sqlc != nil {
	// 	// create database tables
	// 	_createSqlTables(sqlc, dbtype)
	//
	// 	// create DAO instances
	// }
	// if adc != nil {
	// 	// create AWS DynamoDB tables
	// 	_createDynamodbTables(adc)
	//
	// 	// create DAO instances
	// }
	if mc != nil {
		if DEVMODE {
			log.Printf("[DEVMODE] MongoDB database: %s", mc.GetDb())
			log.Printf("[DEVMODE] MongoDB url: %s", mc.GetUrl())

			// HACK to force database creation
			colName := "__libro"
			_, err := mc.CreateCollection(colName)
			log.Printf("[DEVMODE] Create collection %s: %s", colName, err)
			// _, err = mc.GetCollection("__libro").InsertOne(context.Background(), bson.D{
			// 	{"version", goapi.AppVersion},
			// })
			// log.Printf("[DEVMODE] Insert document: %s", err)
		}

		// create MongoDB collections
		_createMongoCollections(mc)

		// create DAOs
		userDao = _createUserDaoMongo(mc)
		productDao = _createProductDaoMongo(mc)
		topicDao = _createTopicDaoMongo(mc)
		pageDao = _createPageDaoMongo(mc)
		domainProductMappingDao = _createM2oMappingDaoMongo(mc, tblMappingDomainProduct)
	}

	_initUsers()
	_initSamples()
}

func _initUsers() {
	log.Printf("[INFO] Initializing user accounts...")

	adminUserId := goapi.AppConfig.GetString("gvabe.init.admin_user_id")
	if adminUserId == "" {
		log.Printf("[WARN] Admin user-id not found at config [gvabe.init.admin_user_id], will not create admin account")
		return
	}
	adminUser, err := userDao.Get(adminUserId)
	if err != nil {
		panic(fmt.Sprintf("error while getting user [%s]: %e", adminUserId, err))
	}
	if adminUser != nil {
		log.Printf("[INFO] Admin user [%s] already existed", adminUserId)
		return
	}

	log.Printf("[INFO] Admin user [%s] not found, creating one...", adminUserId)
	adminUserPwd := goapi.AppConfig.GetString("gvabe.init.admin_user_pwd")
	if adminUserPwd == "" {
		adminUserPwd = utils.RandomString(16)
		log.Printf("[INFO] Admin password not found at config [gvabe.init.admin_user_pwd], generating one...%s", adminUserPwd)
	}
	adminUserName := goapi.AppConfig.GetString("gvabe.init.admin_user_name")
	if adminUserName == "" {
		adminUserName = adminUserId
		log.Printf("[INFO] Admin display-name not found at config [gvabe.init.admin_user_name], use default value %s", adminUserName)
	}

	adminUser = user.NewUser(goapi.AppVersionNumber, adminUserId, utils.UniqueId())
	adminUser.SetPassword(encryptPassword(adminUserId, adminUserPwd)).SetDisplayName(adminUserName).SetAdmin(true)
	fmt.Println("[DEBUG]", adminUser.ToMap(nil))
	result, err := userDao.Create(adminUser)
	if err != nil {
		panic(fmt.Sprintf("error while creating user [%s]: %s", adminUserId, err))
	}
	if !result {
		log.Printf("[ERROR] Cannot create user [%s]", adminUserId)
	}
}

func _initSamplesUpdateStats(prod *libro.Product, topic *libro.Topic) {
	if prod != nil {
		topicList, err := topicDao.GetAll(prod, nil, nil)
		if err != nil {
			panic(fmt.Sprintf("error while getting topic list for product %s: %s", prod.GetId(), err))
		} else {
			prod.SetNumTopics(len(topicList))
			result, err := productDao.Update(prod)
			if err != nil {
				panic(fmt.Sprintf("error while updating product %s: %s", prod.GetId(), err))
			}
			if !result {
				log.Printf("[ERROR] Cannot update product [%s]", prod.GetId())
			}
		}
	}

	if topic != nil {
		pageList, err := pageDao.GetAll(topic, nil, nil)
		if err != nil {
			panic(fmt.Sprintf("error while getting page list for topic %s: %s", topic.GetId(), err))
		} else {
			topic.SetNumPages(len(pageList))
			result, err := topicDao.Update(topic)
			if err != nil {
				panic(fmt.Sprintf("error while updating topic %s: %s", topic.GetId(), err))
			}
			if !result {
				log.Printf("[ERROR] Cannot update topic [%s]", topic.GetId())
			}
		}
	}
}

func _initSamples() {
	if !DEVMODE {
		// only populate sample data on dev mode
		return
	}

	demoProdId := "demo"
	demoProdName := "Demo"
	demoProdDesc := "Sample product to demonstrate Libro"
	demoProd, err := productDao.Get(demoProdId)
	if err != nil {
		panic(fmt.Sprintf("error while getting product [%s]: %s", demoProdId, err))
	}
	if demoProd != nil {
		log.Printf("[INFO] Sample product [%s] already existed.", demoProdId)
		return
	}

	log.Printf("[INFO] Sample product [%s] not found, creating one...", demoProdId)
	demoProd = libro.NewProduct(goapi.AppVersionNumber, demoProdId, demoProdName, demoProdDesc, true)
	result, err := productDao.Create(demoProd)
	if err != nil && err != godal.ErrGdaoDuplicatedEntry {
		panic(fmt.Sprintf("error while creating sample product [%s]: %s", demoProdId, err))
	}
	if !result {
		log.Printf("[ERROR] Cannot create sample product [%s]", demoProdId)
	}

	domainCsv := goapi.AppConfig.GetString("libro.samples.domains", "localhost")
	domainList := regexp.MustCompile(`[,\s]+`).Split(domainCsv, -1)
	for _, domain := range domainList {
		log.Printf("[INFO] Creating mapping {domain:%s -> product:%s}...", domain, demoProdId)
		result, err = domainProductMappingDao.Set("localhost", demoProdId)
		if err != nil && err != respicite.ErrDuplicated {
			panic(fmt.Sprintf("error while creating mapping {domain:%s -> product:%s}: %s", domain, demoProdId, err))
		}
		if !result {
			log.Printf("[ERROR] Cannot create mapping {domain:%s -> product:%s}...", domain, demoProdId)
		}
	}

	re := regexp.MustCompile(`\W+`)
	shortLorem := "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa."
	longLorerm := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi nec imperdiet turpis. Curabitur aliquet pulvinar ultrices. Etiam at posuere leo. Proin ultrices ex et dapibus feugiat link example aenean purus leo, faucibus at elit vel, aliquet scelerisque dui. Etiam quis elit euismod, imperdiet augue sit amet, imperdiet odio. Aenean sem erat, hendrerit eu gravida id, dignissim ut ante. Nam consequat porttitor libero euismod congue."

	/*----------------------------------------------------------------------*/
	topic := libro.NewTopic(goapi.AppVersionNumber, demoProd, "Quick Start", "fa-fighter-jet",
		"Libro's quick start guide. "+shortLorem)
	topic.SetPosition(1).SetId(demoProd.GetId() + "-" + re.ReplaceAllString(strings.ToLower(topic.GetTitle()), ""))
	log.Printf("[INFO] Creating topic (%s -> %s)...", demoProdId, topic.GetTitle())
	result, err = topicDao.Create(topic)
	if err != nil && err != godal.ErrGdaoDuplicatedEntry {
		panic(fmt.Sprintf("error while creating topic (%s -> %s): %s", demoProdId, topic.GetTitle(), err))
	}
	if !result {
		log.Printf("[ERROR] Cannot create topic [%s]", topic.GetTitle())
	}
	_initSamplesUpdateStats(demoProd, nil)

	page := libro.NewPage(goapi.AppVersionNumber, topic, "Download", "fa-download",
		"Download Libro and install on your infrastructure. "+shortLorem, longLorerm)
	page.SetPosition(1).SetId(topic.GetId() + "-" + re.ReplaceAllString(strings.ToLower(page.GetTitle()), ""))
	log.Printf("[INFO] Creating page (%s:%s -> %s)...", demoProdId, topic.GetTitle(), page.GetTitle())
	result, err = pageDao.Create(page)
	if err != nil && err != godal.ErrGdaoDuplicatedEntry {
		panic(fmt.Sprintf("error while creating page (%s:%s -> %s): %s", demoProdId, topic.GetTitle(), page.GetTitle(), err))
	}
	if !result {
		log.Printf("[ERROR] Cannot create page [%s]", page.GetTitle())
	}

	page = libro.NewPage(goapi.AppVersionNumber, topic, "Installation", "fa-code",
		"Installation guide: how to install Libro on local environment. "+shortLorem, longLorerm)
	page.SetPosition(2).SetId(topic.GetId() + "-" + re.ReplaceAllString(strings.ToLower(page.GetTitle()), ""))
	log.Printf("[INFO] Creating page (%s:%s -> %s)...", demoProdId, topic.GetTitle(), page.GetTitle())
	result, err = pageDao.Create(page)
	if err != nil && err != godal.ErrGdaoDuplicatedEntry {
		panic(fmt.Sprintf("error while creating page (%s:%s -> %s): %s", demoProdId, topic.GetTitle(), page.GetTitle(), err))
	}
	if !result {
		log.Printf("[ERROR] Cannot create page [%s]", page.GetTitle())
	}

	_initSamplesUpdateStats(nil, topic)
	/*----------------------------------------------------------------------*/
	topic = libro.NewTopic(goapi.AppVersionNumber, demoProd, "Components", "fa-cogs",
		"Libro has 3 components: API backend, Admin frontend, and Public frontend. "+shortLorem)
	topic.SetPosition(2).SetId(demoProd.GetId() + "-" + re.ReplaceAllString(strings.ToLower(topic.GetTitle()), ""))
	log.Printf("[INFO] Creating topic (%s -> %s)...", demoProdId, topic.GetTitle())
	result, err = topicDao.Create(topic)
	if err != nil && err != godal.ErrGdaoDuplicatedEntry {
		panic(fmt.Sprintf("error while creating topic (%s -> %s): %s", demoProdId, topic.GetTitle(), err))
	}
	if !result {
		log.Printf("[ERROR] Cannot create topic [%s]", topic.GetTitle())
	}
	_initSamplesUpdateStats(demoProd, nil)

	page = libro.NewPage(goapi.AppVersionNumber, topic, "API Backend", "fa-microchip",
		"APIs for both Admin frontend and Public frontend. "+shortLorem, longLorerm)
	page.SetPosition(1).SetId(topic.GetId() + "-" + re.ReplaceAllString(strings.ToLower(page.GetTitle()), ""))
	log.Printf("[INFO] Creating page (%s:%s -> %s)...", demoProdId, topic.GetTitle(), page.GetTitle())
	result, err = pageDao.Create(page)
	if err != nil && err != godal.ErrGdaoDuplicatedEntry {
		panic(fmt.Sprintf("error while creating page (%s:%s -> %s): %s", demoProdId, topic.GetTitle(), page.GetTitle(), err))
	}
	if !result {
		log.Printf("[ERROR] Cannot create page [%s]", page.GetTitle())
	}

	page = libro.NewPage(goapi.AppVersionNumber, topic, "Admin Frontend", "fa-tachometer",
		"GUI for administrators to manage products, document topics and pages. "+shortLorem, longLorerm)
	page.SetPosition(2).SetId(topic.GetId() + "-" + re.ReplaceAllString(strings.ToLower(page.GetTitle()), ""))
	log.Printf("[INFO] Creating page (%s:%s -> %s)...", demoProdId, topic.GetTitle(), page.GetTitle())
	result, err = pageDao.Create(page)
	if err != nil && err != godal.ErrGdaoDuplicatedEntry {
		panic(fmt.Sprintf("error while creating page (%s:%s -> %s): %s", demoProdId, topic.GetTitle(), page.GetTitle(), err))
	}
	if !result {
		log.Printf("[ERROR] Cannot create page [%s]", page.GetTitle())
	}

	page = libro.NewPage(goapi.AppVersionNumber, topic, "Public Frontend", "fa-book",
		"End users are able to view documentation from here. "+shortLorem, longLorerm)
	page.SetPosition(3).SetId(topic.GetId() + "-" + re.ReplaceAllString(strings.ToLower(page.GetTitle()), ""))
	log.Printf("[INFO] Creating page (%s:%s -> %s)...", demoProdId, topic.GetTitle(), page.GetTitle())
	result, err = pageDao.Create(page)
	if err != nil && err != godal.ErrGdaoDuplicatedEntry {
		panic(fmt.Sprintf("error while creating page (%s:%s -> %s): %s", demoProdId, topic.GetTitle(), page.GetTitle(), err))
	}
	if !result {
		log.Printf("[ERROR] Cannot create page [%s]", page.GetTitle())
	}

	_initSamplesUpdateStats(nil, topic)
	/*----------------------------------------------------------------------*/
	topic = libro.NewTopic(goapi.AppVersionNumber, demoProd, "FAQs", "fa-question",
		"List of frequently asked questions. "+shortLorem)
	topic.SetPosition(3).SetId(demoProd.GetId() + "-" + re.ReplaceAllString(strings.ToLower(topic.GetTitle()), ""))
	log.Printf("[INFO] Creating topic (%s -> %s)...", demoProdId, topic.GetTitle())
	result, err = topicDao.Create(topic)
	if err != nil && err != godal.ErrGdaoDuplicatedEntry {
		panic(fmt.Sprintf("error while creating topic (%s -> %s): %s", demoProdId, topic.GetTitle(), err))
	}
	if !result {
		log.Printf("[ERROR] Cannot create topic [%s]", topic.GetTitle())
	}

	_initSamplesUpdateStats(demoProd, nil)

	page = libro.NewPage(goapi.AppVersionNumber, topic, "General", "fa-circle-notch", "General questions: "+shortLorem, longLorerm)
	page.SetPosition(1).SetId(topic.GetId() + "-" + re.ReplaceAllString(strings.ToLower(page.GetTitle()), ""))
	log.Printf("[INFO] Creating page (%s:%s -> %s)...", demoProdId, topic.GetTitle(), page.GetTitle())
	result, err = pageDao.Create(page)
	if err != nil && err != godal.ErrGdaoDuplicatedEntry {
		panic(fmt.Sprintf("error while creating page (%s:%s -> %s): %s", demoProdId, topic.GetTitle(), page.GetTitle(), err))
	}
	if !result {
		log.Printf("[ERROR] Cannot create page [%s]", page.GetTitle())
	}

	_initSamplesUpdateStats(nil, topic)
}
