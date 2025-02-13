package logging

type Settings struct {
	ServiceName string
	Level       Level
	// whether add zap.AddCaller() in zap.Option
	IsShowCaller bool
}
