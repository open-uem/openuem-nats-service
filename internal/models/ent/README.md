# OpenUEM - Ent

Repository containing the OpenUEM project's entities created using [Ent](https://entgo.io/)

This repository will be used as a git submodule from other OpenUEM's repositories

New changes to the schema files requires to run `go generate .`

## SQL Script to delete all tables

```
drop table "certificates" cascade;
drop table "agents" cascade;
drop table "releases" cascade;
drop table "agent_tags" cascade;
drop table "antiviri" cascade;
drop table "apps" cascade;
drop table "computers" cascade;
drop table "deployments" cascade;
drop table "logical_disks" cascade;
drop table "metadata" cascade;
drop table "org_metadata" cascade;
drop table "monitors" cascade;
drop table "network_adapters" cascade;
drop table "operating_systems" cascade;
drop table "printers" cascade;
drop table "revocations" cascade;
drop table "settings" cascade;
drop table "tags" cascade;
drop table "users" cascade;
drop table "sessions" cascade;
drop table "shares" cascade;
drop table "system_updates" cascade;
drop table "updates" cascade;
```
