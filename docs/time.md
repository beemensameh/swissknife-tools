# Time

You have some command that you can use them related to time
* Now `time:now`

## Time Now
You have a many flags to run the command.
- `--update` for make the time continuously update
- `--interval` for how many seconds should update the display time (Work with `--update` tag and should be positive number)
- `--format` for show the time with a specific format. this format should be integer follow this table:

Value|Format name|Time format|Note
-----|-----------|-----------|----
1|ANSIC|Mon Jan _2 15:04:05 2006
2|UnixDate|Mon Jan _2 15:04:05 MST 2006
3|RubyDate|Mon Jan 02 15:04:05 -0700 2006
4|RFC822|02 Jan 06 15:04 MST
5|RFC822Z|02 Jan 06 15:04 -0700|RFC822 with numeric zone
6|RFC850|Monday, 02-Jan-06 15:04:05 MST
7|RFC1123|Mon, 02 Jan 2006 15:04:05 MST
8|RFC1123Z|Mon, 02 Jan 2006 15:04:05 -0700|RFC1123 with numeric zone
9|RFC3339|2006-01-02T15:04:05Z07:00
10|RFC3339Nano|2006-01-02T15:04:05.999999999Z07:00
11|Kitchen|3:04PM
12|Stamp|Jan _2 15:04:05
13|StampMilli|Jan _2 15:04:05.000
14|StampMicro|Jan _2 15:04:05.000000
15|StampNano |Jan _2 15:04:05.000000000
