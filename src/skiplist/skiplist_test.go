package skiplist

import (
	"fmt"
	"math/rand"
	"testing"
)

var (
	skipList = NewSkipList(4)
	datas    = map[int64]*TestData{}
)

type TestData struct {
	Id    int64
	Value int64
}

func (this *TestData) Less(data Data) bool {
	if this.Value < data.(*TestData).Value {
		return true
	}
	if this.Value == data.(*TestData).Value && this.Id < data.(*TestData).Id {
		return true
	}
	return false
}

func (this *TestData) Equal(data Data) bool {
	return this.Id == data.(*TestData).Id
}

func init() {
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

	fmt.Println("fuck")
	skipList.Set(&TestData{
		Id:    1,
		Value: 10,
	}, datas[1])
}

func TestPrint(t *testing.T) {
	fmt.Println("print")
	skipList.Print(func(data Data) {
		fmt.Print(data.(*TestData).Id)
		fmt.Print(":")
		fmt.Print(data.(*TestData).Value)
	})
	fmt.Println("===============")
}

func TestDel(t *testing.T) {
	fmt.Println("del")
	skipList.Del(datas[1])
	fmt.Println("===============")
	TestPrint(t)
}
