package mongogo

import (
	"context"
	"fmt"
	"time"

	"github.com/infinitete/mongogo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dailyCountStr = "日服务统计"

type ByDateRange model.DateRange[model.ServiceStatics]

func (br ByDateRange) FilterKey() bson.D {
	return bson.D{{"key", br.Key}}
}

// 合作伙伴维度统计的服务
type ByCompany struct {
	collection *mongo.Collection
}

func (bc *ByCompany) FindByKey(key string) *ByDateRange {
	ext := &ByDateRange{}
	filter := bson.D{{"key", key}}
	err := bc.collection.FindOne(context.TODO(), filter, options.FindOne()).Decode(ext)

	if err != nil {
		ext.Key = key
		ext.Title = dailyCountStr
		ext.From = time.Now().Unix()
		ext.End = time.Now().Unix()
		ext.Value = &model.ServiceStatics{
			Total: 0,
			Mode:  make(map[uint8]int64),
			Type:  make(map[uint8]int64),
		}
	}

	return ext
}

func (bc *ByCompany) Update(document *ByDateRange) error {
	_, err := bc.collection.ReplaceOne(context.TODO(), document.FilterKey(), document)
	if err != nil {
		return err
	}

	return nil
}

// 为今天添加一次统计
func (bc *ByCompany) IncCount(partnerId int, partnerName string, serveMode, serveType uint8) error {
	key := fmt.Sprintf("%d-%d-%s", partnerId, partnerId, partnerName, time.Now().Format("2006-01-02"))
	ext := bc.FindByKey(key)
	ext.Value.Inc(serveMode, serveType)
}

// 取消或者删除一次服务后，减计数
func (bc *ByCompany) DecCount(partnerId int, serveMode, serveType uint8) {
}
