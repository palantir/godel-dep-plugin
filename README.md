<p align="right">
<a href="https://autorelease.general.dmz.palantir.tech/palantir/godel-dep-plugin"><img src="https://img.shields.io/badge/Perform%20an-Autorelease-success.svg" alt="Autorelease"></a>
</p>

dep-plugin
==========
`dep-plugin` is a [godel](https://github.com/palantir/godel) plugin for [`dep`](https://github.com/golang/dep). It
packages the `dep` program and exposes a task that allows the packaged version of `dep` to be run. It also adds a
`verify` task that runs `dep ensure` when apply is true and `dep ensure -no-vendor -dry-run` when apply is false to
verify that the state of `dep` in a project is valid.

Tasks
-----
* `dep`: runs the packaged copy of `dep`. All of the arguments that are passed to this task are passed to the packaged
  copy of `dep`.

Verify
------
When run as part of the `verify` task, if `apply=true`, then the `dep ensure` task is run. If `apply=false`, the
`dep ensure -novendor -dry-run` task is run, and if the task indicates that the `Gopkg.lock` file is out of date, the
verification fails (without output). If the verification task fails for any other reason, the reason for the failure is
printed.
