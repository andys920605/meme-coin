package snowflake

import (
	"sync"
	"testing"

	"github.com/andys920605/meme-coin/pkg/conf"
	"github.com/andys920605/meme-coin/pkg/logging"
)

func TestSnowflake_ConcurrentIDs(t *testing.T) {
	logger := initLogger()
	Init(logger)

	var wg sync.WaitGroup
	idMap := sync.Map{}
	workerCount := 100
	idsPerWorker := 100

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < idsPerWorker; j++ {
				id := New()
				if _, exists := idMap.LoadOrStore(id, struct{}{}); exists {
					t.Errorf("Duplicate ID found: %d", id)
				}
			}
		}()
	}

	wg.Wait()
}

func initLogger() *logging.Logging {
	config := conf.Config{}
	config.Log.Level = "debug"
	config.Server.Name = "test"
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
