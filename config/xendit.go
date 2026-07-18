package config

import "rent-car-project/utils"

func GetXenditSecretKey() string {
	return utils.GetEnv("XENDIT_SECRET_KEY", "")
}

func GetXenditCallbackToken() string {
	return utils.GetEnv("XENDIT_CALLBACK_TOKEN", "")
}
