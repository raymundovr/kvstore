# TODO's
- Storage tests.
- Close method to gracefully close the storage.
- The service can close with events still in the write buffer: events can get lost.
- Keys and values aren’t encoded in the transaction log: multiple lines or whitespace will fail to parse correctly.
- The sizes of keys and values are unbound: huge keys or values can be added, filling the disk.
- The transaction log is written in plain text: it will take up more disk space than it probably needs to.
- The log retains records of deleted values forever: it will grow indefinitely.
