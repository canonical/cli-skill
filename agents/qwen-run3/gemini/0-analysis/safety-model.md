# Safety Model

- **Configuration Safety**: Using `set` allows blind overrides. There are no built-in rollbacks explicitly exposed.
- **Boot Safety**: To prevent starting intensive inference applications when core components (model weights, binaries) aren't present up-front via Snap Components, internal daemon scripts implement a heavy-duty staggered spinlock (waiting up to an hour) before letting the daemon start sequence proceed. Thus `show-engine` output acts as a safety gate for component integrity.