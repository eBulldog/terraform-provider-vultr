---
layout: "vultr"
page_title: "Vultr: vultr_startup_script"
sidebar_current: "docs-vultr-startup-script"
description: |-
  Provides a Vultr Startup Script resource. This can be used to create, read, modify, and delete Startup Scripts.
---

# vultr_startup_script

Provides a Vultr Startup Script resource. This can be used to create, read, modify, and delete Startup Scripts.

## Example Usage

Create a new Startup Script

```hcl
resource "vultr_startup_script" "my_script" {
    name = "man_run_docs"
    script = "echo $PATH"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the given script.
* `script` - (Required) Contents of the startup script.
* `type` - (Optional) Type of startup script. Default is boot.

## Attributes Reference

The following attributes are exported:

* `id` - ID of the script.
* `name` - Name of the given script.
* `date_created` - Date the script was created.
* `date_modified` - Date the script was last modified.
* `type` - The type of startup script this is.
* `script` - The contents of the startup script.

## Import

Startup Scripts can be imported using the Startup Scripts `SCRIPTID`, e.g.

```
terraform import vultr_startup_script.my_script 537932
```