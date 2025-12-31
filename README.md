An exploration and hack of pipewire configuration.

This little program sets my default audio output to a device with a particular name.

You could also use `jq` for this, but to me, a Go program feels a little less fragile.

1. `pw-dump` provides a list of devices,
2. we pick out the device with a specific description (hard-coded),
3. we set the device with `wpctl set-default <id>`.

Shared to the Public Domain.
