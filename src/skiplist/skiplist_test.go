package skiplist

import (
	"fmt"
	"math/rand"
	"testing"
)

var (
	skipList = NewSkipList(4)
)

type TestData struct {
	Id    int64
	Value int64
}

func (this *TestData) Less(data Data) bool {
	return this.Value < data.(*TestData).Value
}

func (this *TestData) Equal(data Data) bool {
	return this.Id == data.(*TestData).Id
}

func init() {
	datas := map[int64]*TestData{}
	for _, val := range rand.Perm(10) {
		data := &TestData{
			Id:    int64(val),
			Value: int64(val),
		}
		skipList.Set(data, nil)
		datas[data.Id] = data
	}

	for _, val := range rand.Perm(20) {
		data := &TestData{
			Id:    int64(val),
			Value: int64(val),
		}
		old, ok := datas[data.Id]
		if ok {
			skipList.Set(data, old)
		} else {
			skipList.Set(data, nil)
		}
	}
}

func TestPrint(t *testing.T) {
	fmt.Println("print")
	skipList.Print(func(data Data) {
		fmt.Print(data.(*TestData).Value)
	})
	fmt.Println("===============")
}
