package portscan

var KNOWN_TCP_PORTS = map[int]string{
	20: "ftp - filetransfer",
	21: "ftp - control",
	22: "SSH",
	23: "telnet",
	24: "LMTP",
	25: "SMTP",
	43: "whois",
	49: "TACACS",
	53: "DNS",
	69: "tftp",
	88: "kerberos",
	109: "pop2",
	110: "pop3",
	123: "ntp",
	137: "netbios",
	139: "netbios",
	445: "Samba",
	631: "cups",
}

var KNOWN_UDP_PORTS = map[int]string{
	53: "DNS",
}