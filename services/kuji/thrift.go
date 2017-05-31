package services

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	Kuji "github.com/katsew/kuji"
	ThKuji "github.com/katsew/kuji-thrift/thrift/dist/go/kuji"
	"strconv"
)

func NewThriftGachaService(
	transport thrift.TServerTransport,
	transportFactory thrift.TTransportFactory,
	protocolFactory thrift.TProtocolFactory,
	strategy Kuji.KujiStrategy,
) thrift.TProcessor {

	handler := kujiService{
		Kuji.NewKuji(strategy),
	}
	kujiHandler = &handler
	processor := ThKuji.NewKujiServiceProcessor(kujiHandler)

	return processor
}

func (k kujiService) ThRegisterCandidatesWithKey(req *ThKuji.ReqCandidates) (*ThKuji.Response, error) {

	candidates := []Kuji.KujiCandidate{}
	fmt.Println("Candidates", req.Candidates)
	for _, v := range req.Candidates.Candidates {
		k := Kuji.KujiCandidate{
			Id:     v.GetID(),
			Weight: v.GetWeight(),
		}
		candidates = append(candidates, k)
	}

	_, err := kujiHandler.Kuji.RegisterCandidatesWithKey(req.Key, candidates)
	if err != nil {
		return nil, err
	}
	response := &ThKuji.Response{
		Code:    200,
		Message: "Success register candidates",
		Data:    nil,
	}
	return response, nil
}

func (k kujiService) ThPickOneByKey(req *ThKuji.ReqPickOneByKey) (*ThKuji.Response, error) {

	picked, err := kujiHandler.Kuji.PickOneByKey(req.Key)
	if err != nil {
		return nil, err
	}
	i, err := strconv.ParseInt(picked, 10, 64)
	response := &ThKuji.Response{
		Code:    200,
		Message: "Success pick one",
		Data: &ThKuji.Data{
			ID:    i,
			IDStr: picked,
		},
	}
	return response, nil
}

func (k kujiService) ThPickOneByKeyAndIndex(req *ThKuji.ReqPickOneByKeyAndIndex) (*ThKuji.Response, error) {

	picked, err := kujiHandler.Kuji.PickOneByKeyAndIndex(req.Key, req.Index)
	if err != nil {
		return nil, err
	}
	i, err := strconv.ParseInt(picked, 10, 64)
	response := &ThKuji.Response{
		Code:    200,
		Message: "Success pick one by index",
		Data: &ThKuji.Data{
			ID:    i,
			IDStr: picked,
		},
	}
	return response, nil
}

func (k kujiService) ThPickAndDeleteOneByKey(req *ThKuji.ReqPickOneByKey) (*ThKuji.Response, error) {

	picked, err := kujiHandler.Kuji.PickOneByKey(req.Key)
	if err != nil {
		return nil, err
	}
	i, err := strconv.ParseInt(picked, 10, 64)
	response := &ThKuji.Response{
		Code:    200,
		Message: "Success pick one",
		Data: &ThKuji.Data{
			ID:    i,
			IDStr: picked,
		},
	}
	return response, nil
}
