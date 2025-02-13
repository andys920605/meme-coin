package logging

type Option interface {
	Apply(*Settings)
}

func processOpts(opts []Option) *Settings {
	var o Settings
	for _, opt := range opts {
		opt.Apply(&o)
	}
	return &o
}

func WithServiceName(name string) Option {
	return withServiceName(name)
}

type withServiceName string

func (w withServiceName) Apply(s *Settings) {
	s.ServiceName = string(w)
}

func WithLevel(level Level) Option {
	return withLevel(level)
}

type withLevel Level

func (w withLevel) Apply(s *Settings) {
	s.Level = Level(w)
}

func WithShowCaller() Option {
	return withShowCaller(true)
}

type withShowCaller bool

func (w withShowCaller) Apply(s *Settings) {
	s.IsShowCaller = bool(w)
}
