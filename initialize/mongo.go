package initialize

import (
	"context"
	"fmt"
	"ginserver/core/internal"
	"ginserver/global"
	"ginserver/utils"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	option "go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo = new(mongo)

type (
	mongo struct{}
	Index struct {
		V    any      `bson:"v"`
		Ns   any      `bson:"ns"`
		Key  []bson.E `bson:"key"`
		Name string   `bson:"name"`
	}
)

func (m *mongo) Indexes(ctx context.Context) error {
	indexMap := map[string][][]string{}
	for collection, Indexes := range indexMap {
		err := m.CreateIndexes(ctx, collection, Indexes)
		if err != nil {
			return err
		}

	}
	return nil
}

func (m *mongo) Initialization() error {
	var opts []options.ClientOptions
	if global.GVA_CONFIG.Mongo.IsZap {
		OPTS = internal.Mongo.GetCl
	}
}

func (m *mongo) CreateIndexes(ctx context.Context, name string, indexes [][]string) error {
	collection, err := global.GVA_MONGO.Database.Collection(name).CloneCollection()
	if err != nil {
		return errors.Wrapf(err, "获取[%s]的表对象失败!", name)
	}

	list, err := collection.Indexes().List(ctx)
	if err != nil {
		return errors.Wrapf(err, "获取[%s]的索引对象失败!", name)

	}
	var entities []Index
	err = list.All(ctx, &entities)
	if err != nil {
		return errors.Wrapf(err, "获取[%s]的索引列表失败!", name)
	}
	length := len(indexes)
	indexMap1 := make(map[string][]string, length)
	for i := 0; i < length; i++ {
		sort.Strings(indexes[i])
		length1 := len(indexes[i])
		keys := make([]string, 0, length1)
		for j := 0; j < length1; j++ {
			if indexes[i][i][0] == '-' {
				keys = append(keys, indexes[i][j], "-1")
				continue
			}
			keys = append(keys, indexes[i][j], "1")

		}
		key := strings.Join(keys, "-")
		_, o1 := indexMap1[key]
		if o1 {
			return errors.Errorf("索引[%s]重复!", key)
		}
		indexMap1[key] = indexes[i]

	}
	length = len(entities)
	indexMap2 := make(map[string]map[string]string, length)
	for i := 0; i < length; i++ {
		v1, o1 := indexMap2[entities[i].Name]
		if !o1 {
			KeyLength := len(entities[i].Key)
			v1 = make(map[string]string, KeyLength)
			for j := 0; j < KeyLength; j++ {
				v2, o2 := v1[entities[i].Key[j].Key]
				if !o2 {
					v1 = make(map[string]string)
				}
				v2 = entities[i].Key[j].Key
				v1[entities[i].Key[j].Key] = v2
				indexMap2[entities[i].Name] = v1
			}
		}
	}
	for k1, v1 := range indexMap1 {
		_, o2 := indexMap2[k1]
		if o2 {
			continue
		}
		if len(fmt.Sprintf("%s.%s.$%s", collection.Name(), name, v1)) > 127 {
			err = global.GVA_MONGO.Database.Collection(name).CreateOneIndex(ctx, options.IndexModel{
				Key:          v1,
				IndexOptions: option.Index().SetName(utils.MD5V([]byte(k1))),
			})
			if err != nil {
				return errors.Wrapf(err, "创建索引[%s]失败!", k1)
			}
			return nil
		}
		err = global.GVA_MONGO.Database.Collection(name).CreateOneIndex(ctx, options.IndexModel{
			Key:          v1,
			IndexOptions: option.Index().SetExpireAfterSeconds(86400),
		})
		if err != nil {
			return errors.Wrapf(err, "创建索引[%s]失败!", k1)
		}
	}
	return nil

}
