package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"thrift_demo/predict"
	"time"
)

type PredictImpl struct {

}

func (p *PredictImpl) Predict(ctx context.Context, req *predict.Req) (r *predict.Rsp, err error){
	t1 := time.Now()
	contents := []*predict.SortContent{}
	for i := 0; i < 10; i++ {
		sortContent := predict.SortContent{ID: int64(i), Score: 0.01}
		contents = append(contents, &sortContent)
	}
	featrues := []string{""}

	time.Sleep(40 * time.Millisecond)
	rsp := predict.Rsp{Status: "", Contents: contents, Features: &featrues}
	fmt.Printf("cost time:%d\n", time.Now().Sub(t1).Milliseconds())
	return &rsp, nil
}
// Parameters:
//  - Req
func(p *PredictImpl) DropPredict(ctx context.Context, req *predict.Req) (err error){
	return nil
}

func Serve(){
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	predictImpr := &PredictImpl{}
	addr := "10.124.16.133:9999"
	processor := predict.NewPredictProcessor(predictImpr)

	tServerSocket,err := thrift.NewTServerSocket(addr)
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
	predictServer := thrift.NewTSimpleServer4(processor, tServerSocket, transportFactory, protocolFactory)
	err = predictServer.Serve()
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
}
