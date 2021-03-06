/*
 * File: option.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Tuesday, 30th October 2018 1:01:49 am
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package server

import "path/filepath"

type Option struct {
	env       string
	exe       string
	ip        string
	port      string
	address   string
	logpath   string
	dbType    string
	mongoIP   string
	mongoPort string
	datapath  string
	es        string
	webdir    string
	embedES   bool
	index     bool
	reindex   bool
	rmindex   bool
	debug     bool
}

var (
	PROD = "production"
	DEV  = "développement"
	TEST = "test"
)

func (o *Option) Hydrate(env, dir, ip, port, logpath, dbType, mongoIP, mongoPort, datapath, es, webdir string, embedES, index, reindex, rmindex, debug bool) {
	o.exe = dir
	o.env = env
	o.embedES = embedES
	o.port = port
	o.ip = ip
	o.datapath = datapath
	o.dbType = dbType
	o.logpath = logpath
	o.es = es
	o.webdir = webdir
	o.debug = debug
	o.index = index
	o.reindex = reindex
	o.rmindex = rmindex
	o.mongoIP = mongoIP
	o.mongoPort = mongoPort
}

// GetAddress concat ip and port and affect to address if needed
// else default address is define to 127.0.0.1:80
func (o *Option) GetAddress() string {
	if o.address == "" {
		o.address = o.ip + ":" + o.port
	}
	return o.address
}

/*
context.serverOption interface:

type IServerOption interface {
		GetExeFolder() string
		SetEnv(string)
		SetLogpath(string)
		SetDatapath(string)
		GetEnv() string
		GetLogpath() string
		GetDatapath() string
		GetAddress() string
		GetPort() string
		GetIp() string
	}
*/

// GetIP return ip
func (o *Option) GetIP() string {
	return o.ip
}

// GetPort return port
func (o *Option) GetPort() string {
	return o.port
}

// GetExeFolder return ./app.exe folder
func (o *Option) GetExeFolder() string {
	return o.exe
}

// GetLogpath return the path to logs folder
func (o *Option) GetLogpath() string {
	return o.logpath
}

// GetDatabaseType return the kind of database the user want
func (o *Option) GetDatabaseType() string {
	return o.dbType
}

// GetDatapath return the path to data folder
func (o *Option) GetDatapath() string {
	return o.datapath
}

// GetEnv return environnement description default on DEV: "développement"
func (o *Option) GetEnv() string {
	if o.env == "" {
		return DEV
	}
	return o.env
}

// GetEmbedES return the boolean indicating if
// user want the API to connect to elastic search
func (o *Option) GetEmbedES() bool {
	return o.embedES
}

// GetEs return the es serv
func (o *Option) GetEs() string {
	return o.es
}

// GetIndex return if indexation is needed default false
func (o *Option) GetIndex() bool {
	return o.index
}

// GetReindex return if reindexation is needed default false
func (o *Option) GetReindex() bool {
	return o.reindex
}

// GetRmindex return if index removal is needed default false
func (o *Option) GetRmindex() bool {
	return o.rmindex
}

func (o *Option) GetMongoIP() string {
	return o.mongoIP
}

func (o *Option) GetMongoPort() string {
	return o.mongoPort
}

func (o *Option) GetWebDir() string {
	return filepath.Join(o.exe, o.webdir)
}

func (o *Option) SetMongoPort(port string) {
	o.mongoPort = port
}

func (o *Option) SetDatabaseType(dbType string) {
	o.dbType = dbType
}

// SetEnv allow a IServerOption to change the environnement
// "prod" => "production"
// "test" => "test"
// (default) "dev" => développement
func (o *Option) SetEnv(env string) {
	if env == "prod" {
		o.env = PROD
	} else if env == "test" {
		o.env = TEST
	} else {
		o.env = DEV
	}
}

// SetLogpath allow IServerOption to change the path to logs
func (o *Option) SetLogpath(logpath string) {
	o.logpath = logpath
}

// SetDatapath allow IServerOption to change the path to data
func (o *Option) SetDatapath(datapath string) {
	o.datapath = datapath
}

func (o *Option) SetMongoIP(ip string) {
	o.mongoIP = ip
}
