# Table: newrelic_applications

This table shows data for NewRelic Applications.

The composite primary key for this table is ( **id**).

## Columns

| Name                | Type       |
|---------------------| ---------- |
| _cq_source_name     |`utf8`|
| _cq_sync_time       |`timestamp[us, tz=UTC]`|
| _cq_id              |`uuid`|
| _cq_parent_id       |`uuid`|
| id (PK)             |`int64`|
| name                |`utf8`|
| type                |`utf8`|
| created_at          |`timestamp[us, tz=UTC]`|
| updated_at          |`timestamp[us, tz=UTC]`|
| incident_preference |`json`|
| configuration       |`json`|
| links               |`json`|