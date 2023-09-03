package main

import "flag"

// SQL Config
var config_driver_name = "mysql"
var config_data_source_name = "root:root@/test?charset=utf8"

// Dev Config
var config_debug = true

func getConfig() {
	flag.StringVar(&config_driver_name, "driver", "mysql", "Database driver name")
	flag.StringVar(&config_data_source_name, "dsn", "root:root@/test?charset=utf8", "Data source name")
	flag.BoolVar(&config_debug, "debug", true, "Enable debugging")

	// 解析命令行参数
	flag.Parse()
}
