package clients

import(
	pb ""
)

type AuthClient struct{
	*pb.NewAuthClient
}

func NewAuthClient()