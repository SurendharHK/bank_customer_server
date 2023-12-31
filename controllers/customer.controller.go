package controllers

import (
	"context"
	"time"

	"github.com/Grpc_bank/bank_customer_service/interfaces"
	"github.com/Grpc_bank/bank_customer_service/models"
	pro "github.com/Grpc_bank/bank_customer_proto/netxd_customer"
)

type RPCServer struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	CustomerService interfaces.ICustomer
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *pro.Customer) (*pro.CustomerResponse, error) {
	dbProfile := &models.Customer{
		CustomerId: req.CustomerId,
		FirstName:  "",
		LastName:   "",
		BankId:     "",
		Balance:    0,
		CreatedAt:  time.Time{},
		UpadtedAt:  time.Time{},
		IsActive:   false,
	}
	result, err := CustomerService.CreateCustomer(dbProfile)
	if err != nil {
		return nil, err
	} else {
		responseProfile := &pro.CustomerResponse{
			CustomerId: result.CustomerId,
		}
		return responseProfile, nil
	}
}
