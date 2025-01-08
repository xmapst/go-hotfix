# example

This is an example project.

## how to work with this project

```bash
make
./bin/app
```

## test hotfix

```bash
telnet localhost 3333
# input
# hotfix patch/test_go.patch patch.PatchTest()
# recovery patch/test_go.patch
# hotfix patch/test_recovery_go.patch patch.PatchTest()
```

