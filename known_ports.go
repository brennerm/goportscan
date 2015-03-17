package goportscan

const KNOWN_TEST_PORT = 12345
const UNKNOWN_TEST_PORT = 12346

var KNOWN_TCP_PORTS = map[int]string{
	20: "ftp - filetransfer",
	21: "ftp - control",
	22: "ssh",
	23: "telnet",
	24: "lmtp",
	25: "smtp",
	43: "whois",
	49: "tacas",
	53: "dns",
	57: "mtp",
	69: "tftp",
	80: "http",
	88: "kerberos",
	109: "pop2",
	110: "pop3",
	115: "sftp",
	118: "sql",
	137: "netbios",
	139: "netbios",
	143: "imap",
	194: "irc",
	201: "appletalk",
	220: "imap3",
	389: "ldap",
	443: "https",
	445: "Samba",
	513: "rlogin",
	546: "dhcpv6client",
	547: "dhcpv6server",
	554: "rtsp",
	631: "ipp cups",
	636: "ldaps",
	1194: "openvpn",
	1293: "ipsec",
	1812: "radius",
	8080: "http",
	KNOWN_TEST_PORT: "test port",
}

var KNOWN_UDP_PORTS = map[int]string{
	53: "dns",
	123: "ntp",
}
