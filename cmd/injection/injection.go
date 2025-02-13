// nolint
package injection

import (
	_ "time/tzdata"

	_ "github.com/joho/godotenv/autoload"

	domain_service "github.com/andys920605/meme-coin/internal/domain/service"
	"github.com/andys920605/meme-coin/internal/north/local/appservice"
	meme_coin_dao "github.com/andys920605/meme-coin/internal/south/adapter/repository/dao/meme_coin/mysql"
	meme_coin_rep "github.com/andys920605/meme-coin/internal/south/adapter/repository/meme_coin"
	"github.com/andys920605/meme-coin/pkg/conf"
	"github.com/andys920605/meme-coin/pkg/logging"
	"github.com/andys920605/meme-coin/pkg/mysqlx"
	"github.com/andys920605/meme-coin/pkg/snowflake"
)

type Injection struct {
	Config             *conf.Config
	Logger             *logging.Logging
	MemeCoinAppService *appservice.MemeCoinAppService
}

func New() *Injection {
	config := initConfig()
	logger := initLogger(config)

	snowflake.Init(logger)

	mysqlxClient := initMysqlClient(config, logger)
	memeCoinDao := meme_coin_dao.NewMemeCoinDao(mysqlxClient)
	memeCoinRep := meme_coin_rep.NewMemeCoinRepository(memeCoinDao)
	memeCoinDomainSvc := domain_service.NewMemeCoinDomainService(logger, memeCoinRep)
	memeCoinAppSvc := appservice.NewMemeCoinAppService(logger, memeCoinDomainSvc)

	return &Injection{
		Config:             config,
		Logger:             logger,
		MemeCoinAppService: memeCoinAppSvc,
	}
}

func initLogger(config *conf.Config) *logging.Logging {
	loggingLevel, err := logging.ParserLevel(config.Log.Level)
	if err != nil {
		panic(err)
	}

	logger := logging.New(
		logging.WithServiceName(config.Server.Name),
		logging.WithLevel(loggingLevel),
		logging.WithShowCaller(),
	)
	return logger
}

func initConfig() *conf.Config {
	config, err := conf.NewConfig()
	if err != nil {
		panic(err)
	}
	return config
}

func initMysqlClient(config *conf.Config, logger *logging.Logging) *mysqlx.Client {
	client, err := mysqlx.NewClient(config)
	if err != nil {
		logger.Emergencyf("failed to initialize mysql client: %v", err)
	}
	logger.Infof("mysql client initialized")
	return client
}
