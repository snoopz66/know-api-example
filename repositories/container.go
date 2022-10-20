package repositories

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/snoopz66/know-api-example/repositories/userRepo"
	"github.com/snoopz66/know-api-example/services"

	"github.com/snoopz66/know-api-example/repositories/config"
	"go.uber.org/zap"
)

type Container struct {
	srvConf         *config.ServiceConf
	log             *zap.Logger
	userRepo        *userRepo.DynDB
	userService     *services.User
	dynDBConnection *dynamodb.DynamoDB
}

func (c *Container) getLogger() *zap.Logger {
	if c.log == nil {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		c.log = logger
		c.log.Info("Logger created")
	}
	return c.log
}

func (c *Container) getConfig() *config.ServiceConf {
	log := c.getLogger()
	log.Info("Getting config")
	if c.srvConf == nil {
		conf, err := config.Get()
		if err != nil {
			// if we cant get config from aws GTFO
			os.Exit(1)
		}
		c.srvConf = conf
	}
	return c.srvConf
}

func (c *Container) getDynamoDBSession() *dynamodb.DynamoDB {
	log := c.getLogger()
	log.Info("Getting dynamoDB session")
	if c.dynDBConnection == nil {
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
		svc := dynamodb.New(sess)
		c.dynDBConnection = svc
	}
	return c.dynDBConnection
}

func (c *Container) getUserRepo() *userRepo.DynDB {
	log := c.getLogger()
	log.Info("Getting user repo")
	if c.userRepo == nil {
		c.userRepo = userRepo.New(c.getConfig(), c.getDynamoDBSession())
	}
	return c.userRepo
}

func (c *Container) GetUserService() *services.User {
	log := c.getLogger()
	log.Info("Getting user service")
	if c.userService == nil {
		c.userService = services.New(c.getUserRepo(), c.getLogger())
	}
	return c.userService
}
