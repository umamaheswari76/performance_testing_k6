package interfaces

import "k6_performance_test/grpc_test/models"

type IProfile interface{
	CreateProfile(*models.Profile)(*models.ProfileResponse,error)
}