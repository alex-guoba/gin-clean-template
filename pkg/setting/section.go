package setting

import (
	"time"
)

type ServerSettingS struct {
	RunMode      string
	HTTPPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize       int
	MaxPageSize           int
	DefaultContextTimeout time.Duration
	UploadSavePath        string
	UploadServerURL       string
	UploadImageMaxSize    int
	UploadImageAllowExts  []string
}

type LogSettingS struct {
	LogSavePath string
	LogFileName string
	MaxSize     int // MB
	MaxBackups  int
	Compress    bool
	Level       string
}

type EmailSettingS struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type RatelimitSettingS struct {
	Enable          bool
	ConfigFile      string
	CPULoadThresh   float64
	CPULoadStrategy int
}

var sections = make(map[string]any)

func (s *Setting) ReadSection(k string, v any) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
