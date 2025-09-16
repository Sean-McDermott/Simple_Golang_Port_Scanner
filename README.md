This program is not completely fitted with input sanitization so you may encounter input errors


DISPLAY GUIDE:

|| | UDP| = UDP CLOSED \n
||X| UDP| = UDP OPEN
|| | TCP| = TCP CLOSED
||X| TCP| = TCP OPEN

Usage of GoNmap:
 -A    Scan all 65,536 ports
  -T    Scan top 1,000 ports
  -h string
        Host(s) to scan: single IP, comma-delimited, or hyphen-delimited range
  -l    Reverse lookup for domain names
  -p string
        Port(s) to scan: single port, comma-delimited, or hyphen-delimited range
  -t    TCP scan only
  -u    UDP scan only

If no port range is specified, the program will default to 1-100

Example Usage:

.\goNmap.exe -h 192.168.1.1 -p 1-200
.\goNmap.exe -h 192.168.1.1 -p 1,2,3
.\goNmap.exe -h 192.168.1.1-3 -p 1-200
.\goNmap.exe -h 192.168.1.1
.\goNmap.exe -h 192.168.1.1 -p 1-200 -t

good luck and godspeed
