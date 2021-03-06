package bot

const (
	_ = iota
	PASS_CMD
	NICK_CMD
	USER_CMD
	SERVER_CMD
	OPER_CMD
	QUIT_CMD
	SQUIT_CMD
	JOIN_CMD
	PART_CMD
	MODE_CMD
	TOPIC_CMD
	NAMES_CMD
	LIST_CMD
	INVITE_CMD
	KICK_CMD
	VERSION_CMD
	STATS_CMD
	LINKS_CMD
	TIME_CMD
	CONNECT_CMD
	TRACE_CMD
	ADMIN_CMD
	INFO_CMD
	PRIVMSG_CMD
	NOTICE_CMD
	WHO_CMD
	WHOIS_CMD
	WHOWAS_CMD
	KILL_CMD
	PING_CMD
	PONG_CMD
	ERROR_CMD
	AWAY_CMD
	REHASH_CMD
	RESTART_CMD
	SUMMON_CMD
	USERS_CMD
	WALLOPS_CMD
	USERHOST_CMD
	ISON_CMD
)

var CMD map[string]int = map[string]int{
	"PASS":     PASS_CMD,
	"NICK":     NICK_CMD,
	"USER":     USER_CMD,
	"SERVER":   SERVER_CMD,
	"OPER":     OPER_CMD,
	"QUIT":     QUIT_CMD,
	"SQUIT":    SQUIT_CMD,
	"JOIN":     JOIN_CMD,
	"PART":     PART_CMD,
	"MODE":     MODE_CMD,
	"TOPIC":    TOPIC_CMD,
	"NAMES":    NAMES_CMD,
	"LIST":     LIST_CMD,
	"INVITE":   INVITE_CMD,
	"KICK":     KICK_CMD,
	"VERSION":  VERSION_CMD,
	"STATS":    STATS_CMD,
	"LINKS":    LINKS_CMD,
	"TIME":     TIME_CMD,
	"CONNECT":  CONNECT_CMD,
	"TRACE":    TRACE_CMD,
	"ADMIN":    ADMIN_CMD,
	"INFO":     INFO_CMD,
	"PRIVMSG":  PRIVMSG_CMD,
	"NOTICE":   NOTICE_CMD,
	"WHO":      WHO_CMD,
	"WHOIS":    WHOIS_CMD,
	"WHOWAS":   WHOWAS_CMD,
	"KILL":     KILL_CMD,
	"PING":     PING_CMD,
	"PONG":     PONG_CMD,
	"ERROR":    ERROR_CMD,
	"AWAY":     AWAY_CMD,
	"REHASH":   REHASH_CMD,
	"RESTART":  RESTART_CMD,
	"SUMMON":   SUMMON_CMD,
	"USERS":    USERS_CMD,
	"WALLOPS":  WALLOPS_CMD,
	"USERHOST": USERHOST_CMD,
	"ISON":     ISON_CMD,
}

var DMC map[int]string = map[int]string{
	PASS_CMD:     "PASS",
	NICK_CMD:     "NICK",
	USER_CMD:     "USER",
	SERVER_CMD:   "SERVER",
	OPER_CMD:     "OPER",
	QUIT_CMD:     "QUIT",
	SQUIT_CMD:    "SQUIT",
	JOIN_CMD:     "JOIN",
	PART_CMD:     "PART",
	MODE_CMD:     "MODE",
	TOPIC_CMD:    "TOPIC",
	NAMES_CMD:    "NAMES",
	LIST_CMD:     "LIST",
	INVITE_CMD:   "INVITE",
	KICK_CMD:     "KICK",
	VERSION_CMD:  "VERSION",
	STATS_CMD:    "STATS",
	LINKS_CMD:    "LINKS",
	TIME_CMD:     "TIME",
	CONNECT_CMD:  "CONNECT",
	TRACE_CMD:    "TRACE",
	ADMIN_CMD:    "ADMIN",
	INFO_CMD:     "INFO",
	PRIVMSG_CMD:  "PRIVMSG",
	NOTICE_CMD:   "NOTICE",
	WHO_CMD:      "WHO",
	WHOIS_CMD:    "WHOIS",
	WHOWAS_CMD:   "WHOWAS",
	KILL_CMD:     "KILL",
	PING_CMD:     "PING",
	PONG_CMD:     "PONG",
	ERROR_CMD:    "ERROR",
	AWAY_CMD:     "AWAY",
	REHASH_CMD:   "REHASH",
	RESTART_CMD:  "RESTART",
	SUMMON_CMD:   "SUMMON",
	USERS_CMD:    "USERS",
	WALLOPS_CMD:  "WALLOPS",
	USERHOST_CMD: "USERHOST",
	ISON_CMD:     "ISON",
}

// subcommand is just a fine-grained division of command
// there may be no standard or rfc for it
const (
	_ = iota
	CTCP_ACTION_SUB
)
