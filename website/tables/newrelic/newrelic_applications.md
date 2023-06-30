# Table: newrelic_applications

This table shows data for NewRelic Applications.

The composite primary key for this table is ( **id**).

## Columns

| Name                  | Type       |
|-----------------------| ---------- |
| _cq_source_name       |`utf8`|
| _cq_sync_time         |`timestamp[us, tz=UTC]`|
| _cq_id                |`uuid`|
| _cq_parent_id         |`uuid`|
| id (PK)               |`int64`|
| name                  |`utf8`|
| language              |`utf8`|
| health_status         |`utf8`|
| reporting             |`utf8`|
| last_reported_status  |`timestamp[us, tz=UTC]`|
| application_summary   |`json`|
| end_user_summary      |`json`|
| settings              |`json`|
| links                 |`json`|