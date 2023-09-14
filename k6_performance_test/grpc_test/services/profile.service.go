package services

import (
	"context"
	"k6_performance_test/grpc_test/interfaces"
	"k6_performance_test/grpc_test/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type InitCreateProfile struct{
	collecion *mongo.Collection
	ctx context.Context
}

func NewCreateProfile(collection *mongo.Collection, ctx context.Context) interfaces.IProfile{
	return &InitCreateProfile{collection, ctx}
}

func (c *InitCreateProfile)CreateProfile(profile *models.Profile)(*models.ProfileResponse,error){
	_,err := c.collecion.InsertOne(c.ctx, &profile)
	if err !=nil{
		return nil,err
	}	
	return &models.ProfileResponse{
		StringMessage: "Profile Inserted",
	},nil
}