/*
@Time : 2020/5/13 18:29
@Author : zxr
@File : ProdService
@Software: GoLand
*/
package services

import (
	"bili/grpcpro/pbfiles"
	"context"
)

type ProdService struct {
}

func (p *ProdService) GetProdStock(ctx context.Context, req *pbfiles.ProdRequest) (*pbfiles.ProdResponse, error) {
	var stock int32 = 0
	if req.Area == pbfiles.ProdArea_A {
		stock = 30
	}
	if req.Area == pbfiles.ProdArea_B {
		stock = 40
	}
	if req.Area == pbfiles.ProdArea_C {
		stock = 50
	}
	return &pbfiles.ProdResponse{ProdStock: stock}, nil
}

func (p *ProdService) GetProdList(ctx context.Context, in *pbfiles.ProdListRequest) (*pbfiles.ProdListResponse, error) {
	list := []*pbfiles.ProdList{
		{
			ProdId:     10,
			ProdName:   "JAVA",
			IsSubsidy:  true,
			IsIntegral: 666,
			Price:      9.99,
		},
		{
			ProdId:     20,
			ProdName:   "PHP",
			IsSubsidy:  true,
			IsIntegral: 666,
			Price:      9.99,
		},
		{
			ProdId:     30,
			ProdName:   "GO",
			IsSubsidy:  true,
			IsIntegral: 666,
			Price:      9.99,
		},
		{
			ProdId:     40,
			ProdName:   "C++",
			IsSubsidy:  true,
			IsIntegral: 666,
			Price:      9.99,
		},
	}
	return &pbfiles.ProdListResponse{ProdList: list}, nil
}

func (p *ProdService) GetProdInfo(context.Context, *pbfiles.ProdRequest) (*pbfiles.ProdModel, error) {
	return &pbfiles.ProdModel{
		ProdId:    1,
		ProdName:  "test product",
		ProdPrice: 9.99,
	}, nil
}
