# UUID

You have some command that you can use them related to time
* Now `uuid:generate`

## `uuid:generate`
You have a many flags to run the command.
- `--version` for UUID version which are generated. It is from 1 to 5.
- `--uuid-type` It is DCE uuid types (should be one of [group, person, security] and should add value when uuid version is 2).
- `--uuid-security-type` It is DCE security uuid types (should be one of [group, person, org] and should add value when uuid version is 2 and the uuid type is security).
- `--name` It is used for uuid version 3 and 5 (maybe anything - no constrained).
- `--number` Number of UUID need to generate
- `--separated` The separated character that should separate UUIDs
