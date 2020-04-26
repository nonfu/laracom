package log

import (
    "github.com/sirupsen/logrus"
    "os"
)

func init() {
    if os.Getenv("MICRO_SERVICE_ENV") == "dev" {
        logrus.SetFormatter(&logrus.TextFormatter{
            TimestampFormat: "2006-01-02T15:04:05.000",
            FullTimestamp: true,
        })
    } else {
        logrus.SetFormatter(&logrus.JSONFormatter{})
    }
}