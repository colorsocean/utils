package utils

import "regexp"

/****************************************************************
** Regexp
********/

const (
	//RegexpUUID  = `([a-f\d]{8}(-[a-f\d]{4}){3}-[a-f\d]{12}?)`
	RegexpUUID         = `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`
	RegexpUUIDNoDashes = `^[a-f0-9]{32}$`
	RegexpMongoId      = `^[a-f0-9]{24}$`
	RegexpEmail        = `^[-0-9a-zA-Z.+_]+@[-0-9a-zA-Z.+_]+\.[a-zA-Z]{2,4}$`
	RegexpPhone        = `^.*$` // Todo: Find adequate regex
)

var (
	MatchUUID         = regexp.MustCompile(RegexpUUID)
	MatchUUIDNoDashes = regexp.MustCompile(RegexpUUIDNoDashes)
	MatchMongoId      = regexp.MustCompile(RegexpMongoId)
	MatchEmail        = regexp.MustCompile(RegexpEmail)
	MatchPhone        = regexp.MustCompile(RegexpPhone)
)
