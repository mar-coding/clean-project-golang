package config

import (
	_ "embed"
	"time"
)

//go:embed service_info.yml
var serviceInfo []byte

const (
	_defaultCacheExpireDuration = 30 * time.Minute
)

var (
	CacheExpireDuration    time.Duration
	ContactUSForwardEmails []string
	FrontBaseURL           string
)

//func NewConfig(configPath string) (*config.Config, error) {
//	cfg, err := config.New(configPath)
//	if err != nil {
//		return nil, err
//	}
//	return cfg, initiateExtraData(cfg)
//}
//
//func NewServiceInfo() (*info.ServiceInfo, error) {
//	return info.NewFromEmbed(serviceInfo)
//}
//
//func NewError(serviceInfo *info.ServiceInfo, serviceConfig *config.Config) (errHandler.Handler, error) {
//	return errHandler.NewError(
//		uint32(serviceInfo.Code),
//		serviceInfo.Name,
//		serviceInfo.Version,
//		serviceConfig.Domain,
//	)
//}
//
//func initiateExtraData(config *config.Config) error {
//	if v, ok := config.ExtraData["cache_expire_duration"]; ok {
//		CacheExpireDuration = _defaultCacheExpireDuration
//		if CacheExpireDuration != 0 {
//			CacheExpireDuration = time.Duration(v.(int)) * time.Minute
//		}
//	} else {
//		CacheExpireDuration = _defaultCacheExpireDuration
//	}
//
//	if v, ok := config.ExtraData["front_base_url"]; ok {
//		FrontBaseURL = v.(string)
//	}
//
//	if v, ok := config.ExtraData["contact_us_forward_email"]; ok {
//		contacts := v.([]interface{})
//		for _, contact := range contacts {
//			ContactUSForwardEmails = append(ContactUSForwardEmails, contact.(string))
//		}
//	} else {
//		return errors.New("config: failed to get field contact_us_forward_email in extra data")
//	}
//
//	return nil
//}
