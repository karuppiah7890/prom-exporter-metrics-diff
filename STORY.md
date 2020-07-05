# Story

This file will be capturing the story behind this repo and it's features and
everything :)

I'm creating a tool in this repo, which will help in finding the difference
between the metrics of two prometheus exporter versions. This is important in
case of breaking changes in the newer versions.

In my team, we noticed this when the prometheus node exporter was not in a
stable version but we were already using it, as the whole world was and we
we were upgrading our prometheus operator chart version which in turn upgraded
our node exporter version, again to an unstable one, but with breaking changes.
At the time, 1.0 release candidates were coming out. The stable version got
released recently - https://github.com/prometheus/node_exporter/releases/tag/v1.0.0 
on May 26th 2020 according to my timezone (UTC+05:30)

This tool will help people in case their exporters have breaking changes.
At least that's the aim. I hope that it's possible. :) Of course with some
hard work.

So, I have chosen two versions of node exporter 0.15.0 and 0.18.0. From what I
rememeber, these were the concerning versions when we were doing our upgrade.
I ran the two versions of these exporters in my Mac machine and got the metrics
from the metrics endpoint

```bash
$ ./node_exporter
```

```bash
$ curl localhost:9100/metrics > <version>.metrics
```

Surely some things might have been different if it had been run in Linux.
Anyways, the aim of the tool is to find differences, and it should be able to
do it for any exporter metrics, with exporter running anywhere and for any
tool, so I just needed a sample and I got it.

The idea is to get a sample of metrics from the two exporters. I hope that this
sample is a good enough sample, with all the metrics and labels that is needed
to find the differences. If that's not the case, the user getting the sample
needs to make sure that the sample is a good representation of all the features
(metrics and labels) that the exporter can provide, so that no differences are
missed out :)

Now, with the parsing. I am going to use the PromParser to parse the metrics

https://pkg.go.dev/github.com/prometheus/prometheus/pkg/textparse?tab=doc#PromParser

I'm going to play with this to see how it works! :)

To play with it, I'm getting the module first. I was wondering if I was getting
the latest version of the module and then I checked the github releases page
to find that `v2.19.2` is the latest release but that wasn't getting downloaded
for some reason which I didn't get. So I used the commit sha of that release
and got the module. Commit sha was `c448ada63d83002e9c1d2c9f84e09f55a61f0ff7`

https://github.com/prometheus/prometheus/releases/tag/v2.19.2
https://github.com/prometheus/prometheus/commit/c448ada63d83002e9c1d2c9f84e09f55a61f0ff7

```bash
$ go get -u -v github.com/prometheus/prometheus
go: downloading github.com/prometheus/prometheus v1.8.2
go: downloading github.com/prometheus/prometheus v2.5.0+incompatible
go: github.com/prometheus/prometheus upgrade => v2.5.0+incompatible

$ go get -u -v github.com/prometheus/prometheus@2.19.2
go get github.com/prometheus/prometheus@2.19.2: github.com/prometheus/prometheus@2.19.2: invalid version: unknown revision 2.19.2

$ go get -u -v github.com/prometheus/prometheus@v2.19.2
go get github.com/prometheus/prometheus@v2.19.2: github.com/prometheus/prometheus@v2.19.2: invalid version: module contains a go.mod file, so major version must be compatible: should be v0 or v1, not v2

$ go get -u -v github.com/prometheus/prometheus@c448ada63d83002e9c1d2c9f84e09f55a61f0ff7
go: github.com/prometheus/prometheus c448ada63d83002e9c1d2c9f84e09f55a61f0ff7 => v1.8.2-0.20200626085723-c448ada63d83
go: downloading github.com/prometheus/prometheus v1.8.2-0.20200626085723-c448ada63d83
```

Now I'm going to just list down the metric names and metric types in the old
exporter version metrics and new exporter version metrics

So I finally did that part! :)

```bash
$ ./prom-exporter-metrics-diff sample/0.15.0.metrics sample/0.18.0.metrics

old metrics : 
Metric Name: go_gc_duration_seconds, Metric Type: summary
Metric Name: go_goroutines, Metric Type: gauge
Metric Name: go_info, Metric Type: gauge
Metric Name: go_memstats_alloc_bytes, Metric Type: gauge
Metric Name: go_memstats_alloc_bytes_total, Metric Type: counter
Metric Name: go_memstats_buck_hash_sys_bytes, Metric Type: gauge
Metric Name: go_memstats_frees_total, Metric Type: counter
Metric Name: go_memstats_gc_cpu_fraction, Metric Type: gauge
Metric Name: go_memstats_gc_sys_bytes, Metric Type: gauge
Metric Name: go_memstats_heap_alloc_bytes, Metric Type: gauge
Metric Name: go_memstats_heap_idle_bytes, Metric Type: gauge
Metric Name: go_memstats_heap_inuse_bytes, Metric Type: gauge
Metric Name: go_memstats_heap_objects, Metric Type: gauge
Metric Name: go_memstats_heap_released_bytes, Metric Type: gauge
Metric Name: go_memstats_heap_sys_bytes, Metric Type: gauge
Metric Name: go_memstats_last_gc_time_seconds, Metric Type: gauge
Metric Name: go_memstats_lookups_total, Metric Type: counter
Metric Name: go_memstats_mallocs_total, Metric Type: counter
Metric Name: go_memstats_mcache_inuse_bytes, Metric Type: gauge
Metric Name: go_memstats_mcache_sys_bytes, Metric Type: gauge
Metric Name: go_memstats_mspan_inuse_bytes, Metric Type: gauge
Metric Name: go_memstats_mspan_sys_bytes, Metric Type: gauge
Metric Name: go_memstats_next_gc_bytes, Metric Type: gauge
Metric Name: go_memstats_other_sys_bytes, Metric Type: gauge
Metric Name: go_memstats_stack_inuse_bytes, Metric Type: gauge
Metric Name: go_memstats_stack_sys_bytes, Metric Type: gauge
Metric Name: go_memstats_sys_bytes, Metric Type: gauge
Metric Name: go_threads, Metric Type: gauge
Metric Name: http_request_duration_microseconds, Metric Type: summary
Metric Name: http_request_size_bytes, Metric Type: summary
Metric Name: http_requests_total, Metric Type: counter
Metric Name: http_response_size_bytes, Metric Type: summary
Metric Name: node_cpu, Metric Type: counter
Metric Name: node_disk_read_bytes_total, Metric Type: counter
Metric Name: node_disk_read_seconds_total, Metric Type: counter
Metric Name: node_disk_read_sectors_total, Metric Type: counter
Metric Name: node_disk_reads_completed_total, Metric Type: counter
Metric Name: node_disk_write_seconds_total, Metric Type: counter
Metric Name: node_disk_writes_completed_total, Metric Type: counter
Metric Name: node_disk_written_bytes_total, Metric Type: counter
Metric Name: node_disk_written_sectors_total, Metric Type: counter
Metric Name: node_exporter_build_info, Metric Type: gauge
Metric Name: node_filesystem_avail, Metric Type: gauge
Metric Name: node_filesystem_device_error, Metric Type: gauge
Metric Name: node_filesystem_files, Metric Type: gauge
Metric Name: node_filesystem_files_free, Metric Type: gauge
Metric Name: node_filesystem_free, Metric Type: gauge
Metric Name: node_filesystem_readonly, Metric Type: gauge
Metric Name: node_filesystem_size, Metric Type: gauge
Metric Name: node_load1, Metric Type: gauge
Metric Name: node_load15, Metric Type: gauge
Metric Name: node_load5, Metric Type: gauge
Metric Name: node_memory_active_bytes_total, Metric Type: gauge
Metric Name: node_memory_bytes_total, Metric Type: gauge
Metric Name: node_memory_free_bytes_total, Metric Type: gauge
Metric Name: node_memory_inactive_bytes_total, Metric Type: gauge
Metric Name: node_memory_swapped_in_pages_total, Metric Type: gauge
Metric Name: node_memory_swapped_out_pages_total, Metric Type: gauge
Metric Name: node_memory_wired_bytes_total, Metric Type: gauge
Metric Name: node_network_receive_bytes, Metric Type: gauge
Metric Name: node_network_receive_errs, Metric Type: gauge
Metric Name: node_network_receive_multicast, Metric Type: gauge
Metric Name: node_network_receive_packets, Metric Type: gauge
Metric Name: node_network_transmit_bytes, Metric Type: gauge
Metric Name: node_network_transmit_errs, Metric Type: gauge
Metric Name: node_network_transmit_multicast, Metric Type: gauge
Metric Name: node_network_transmit_packets, Metric Type: gauge
Metric Name: node_scrape_collector_duration_seconds, Metric Type: gauge
Metric Name: node_scrape_collector_success, Metric Type: gauge
Metric Name: node_time, Metric Type: gauge



new metrics : 
Metric Name: go_gc_duration_seconds, Metric Type: summary
Metric Name: go_goroutines, Metric Type: gauge
Metric Name: go_info, Metric Type: gauge
Metric Name: go_memstats_alloc_bytes, Metric Type: gauge
Metric Name: go_memstats_alloc_bytes_total, Metric Type: counter
Metric Name: go_memstats_buck_hash_sys_bytes, Metric Type: gauge
Metric Name: go_memstats_frees_total, Metric Type: counter
Metric Name: go_memstats_gc_cpu_fraction, Metric Type: gauge
Metric Name: go_memstats_gc_sys_bytes, Metric Type: gauge
Metric Name: go_memstats_heap_alloc_bytes, Metric Type: gauge
Metric Name: go_memstats_heap_idle_bytes, Metric Type: gauge
Metric Name: go_memstats_heap_inuse_bytes, Metric Type: gauge
Metric Name: go_memstats_heap_objects, Metric Type: gauge
Metric Name: go_memstats_heap_released_bytes, Metric Type: gauge
Metric Name: go_memstats_heap_sys_bytes, Metric Type: gauge
Metric Name: go_memstats_last_gc_time_seconds, Metric Type: gauge
Metric Name: go_memstats_lookups_total, Metric Type: counter
Metric Name: go_memstats_mallocs_total, Metric Type: counter
Metric Name: go_memstats_mcache_inuse_bytes, Metric Type: gauge
Metric Name: go_memstats_mcache_sys_bytes, Metric Type: gauge
Metric Name: go_memstats_mspan_inuse_bytes, Metric Type: gauge
Metric Name: go_memstats_mspan_sys_bytes, Metric Type: gauge
Metric Name: go_memstats_next_gc_bytes, Metric Type: gauge
Metric Name: go_memstats_other_sys_bytes, Metric Type: gauge
Metric Name: go_memstats_stack_inuse_bytes, Metric Type: gauge
Metric Name: go_memstats_stack_sys_bytes, Metric Type: gauge
Metric Name: go_memstats_sys_bytes, Metric Type: gauge
Metric Name: go_threads, Metric Type: gauge
Metric Name: node_boot_time_seconds, Metric Type: gauge
Metric Name: node_cpu_seconds_total, Metric Type: counter
Metric Name: node_disk_read_bytes_total, Metric Type: counter
Metric Name: node_disk_read_sectors_total, Metric Type: counter
Metric Name: node_disk_read_time_seconds_total, Metric Type: counter
Metric Name: node_disk_reads_completed_total, Metric Type: counter
Metric Name: node_disk_write_time_seconds_total, Metric Type: counter
Metric Name: node_disk_writes_completed_total, Metric Type: counter
Metric Name: node_disk_written_bytes_total, Metric Type: counter
Metric Name: node_disk_written_sectors_total, Metric Type: counter
Metric Name: node_exporter_build_info, Metric Type: gauge
Metric Name: node_filesystem_avail_bytes, Metric Type: gauge
Metric Name: node_filesystem_device_error, Metric Type: gauge
Metric Name: node_filesystem_files, Metric Type: gauge
Metric Name: node_filesystem_files_free, Metric Type: gauge
Metric Name: node_filesystem_free_bytes, Metric Type: gauge
Metric Name: node_filesystem_readonly, Metric Type: gauge
Metric Name: node_filesystem_size_bytes, Metric Type: gauge
Metric Name: node_load1, Metric Type: gauge
Metric Name: node_load15, Metric Type: gauge
Metric Name: node_load5, Metric Type: gauge
Metric Name: node_memory_active_bytes, Metric Type: gauge
Metric Name: node_memory_compressed_bytes, Metric Type: gauge
Metric Name: node_memory_free_bytes, Metric Type: gauge
Metric Name: node_memory_inactive_bytes, Metric Type: gauge
Metric Name: node_memory_swapped_in_bytes_total, Metric Type: counter
Metric Name: node_memory_swapped_out_bytes_total, Metric Type: counter
Metric Name: node_memory_total_bytes, Metric Type: gauge
Metric Name: node_memory_wired_bytes, Metric Type: gauge
Metric Name: node_network_receive_bytes_total, Metric Type: counter
Metric Name: node_network_receive_errs_total, Metric Type: counter
Metric Name: node_network_receive_multicast_total, Metric Type: counter
Metric Name: node_network_receive_packets_total, Metric Type: counter
Metric Name: node_network_transmit_bytes_total, Metric Type: counter
Metric Name: node_network_transmit_errs_total, Metric Type: counter
Metric Name: node_network_transmit_multicast_total, Metric Type: counter
Metric Name: node_network_transmit_packets_total, Metric Type: counter
Metric Name: node_scrape_collector_duration_seconds, Metric Type: gauge
Metric Name: node_scrape_collector_success, Metric Type: gauge
Metric Name: node_textfile_scrape_error, Metric Type: gauge
Metric Name: node_time_seconds, Metric Type: gauge
Metric Name: promhttp_metric_handler_requests_in_flight, Metric Type: gauge
Metric Name: promhttp_metric_handler_requests_total, Metric Type: counter
```

Next is - find differences between the old and new metrics. We also have to
checkout labels next.

Something to note is - my aim is to notice breaking changes, which means
concentrating more on the metrics that are present in the old version but
missing in the new version - because they are missing, or just renamed.
Other breaking changes are - changing the type of the metrics. And then
of course there's labels. Now the thing with the word "diff" or "difference"
is that, it's also meant to show what's present in the new version but not
present in the old. For me these are possibly new features, or just some
metrics or labels renamed or whose type have been changed.

I gotta think what I need to provide. For now I think I'll go with the main
aim of this tool, and concentrate only on the breaking changes difference, not
worrying too much on the name of the tool :)

Also, once a person finds out what's there in old version but not in new
version, they will check the new version to see if there are new stuff that
replace the old stuff or something. Even if the tool shows the new stuff,
it's upto the user to check the help text of the metric and the exporter's
change log to understand what metric gives what and how the naming has changed
or if the metric has been removed and stuff.

The one downside to this tool is that, it's going to show you all the breaking
changes based on all the old and new metrics. But it's very possible that you
are not using all the metrics of the exporter, for example in your grafana
dashboards. In which case, you just need to worry about metrics that you are
using, and forget about the others.

So, the ideal steps would be to
* First get metric names from the old metrics, with or without the tool. I'll
see how the tool can help :)
* See what all metrics you use in your grafana dashboard or any place where the
metrics are referred to
* With the tool, find out what old metrics are broken (name, type, label name,
label type) and focus only on the metrics you use and chuck the others
* Then you can move on to see how to fix the broken metrics
    * Checking changelogs / docs of the exporters to see what happened to
    these broken metrics
* For fixing the usage, you might have to use another tool, and the input to
that tool will have to be manually created, and may be the fixing itself can
be done manually if you don't have automation. For this step, I'm planning to
create a separate tool to fix grafana dashboard JSONs given input on what should
be changed to what.

Next what I did was get the labels - names and values. I checked the code in
github.com/prometheus/prometheus - pkg/textparse/promparse_test.go and found
out how to get the labels

```bash
$ ./prom-exporter-metrics-diff sample/0.15.0.metrics sample/0.18.0.metrics
old metrics : 
Metric Name: go_gc_duration_seconds, Metric Type: summary
Metric : go_gc_duration_seconds{quantile="0"}, Metric Labels: {__name__="go_gc_duration_seconds", quantile="0"}
Metric : go_gc_duration_seconds{quantile="0.25"}, Metric Labels: {__name__="go_gc_duration_seconds", quantile="0.25"}
Metric : go_gc_duration_seconds{quantile="0.5"}, Metric Labels: {__name__="go_gc_duration_seconds", quantile="0.5"}
Metric : go_gc_duration_seconds{quantile="0.75"}, Metric Labels: {__name__="go_gc_duration_seconds", quantile="0.75"}
Metric : go_gc_duration_seconds{quantile="1"}, Metric Labels: {__name__="go_gc_duration_seconds", quantile="1"}
Metric : go_gc_duration_seconds_sum, Metric Labels: {__name__="go_gc_duration_seconds_sum"}
Metric : go_gc_duration_seconds_count, Metric Labels: {__name__="go_gc_duration_seconds_count"}
Metric Name: go_goroutines, Metric Type: gauge
Metric : go_goroutines, Metric Labels: {__name__="go_goroutines"}
Metric Name: go_info, Metric Type: gauge
Metric : go_info{version="go1.9.1"}, Metric Labels: {__name__="go_info", version="go1.9.1"}
Metric Name: go_memstats_alloc_bytes, Metric Type: gauge
Metric : go_memstats_alloc_bytes, Metric Labels: {__name__="go_memstats_alloc_bytes"}
Metric Name: go_memstats_alloc_bytes_total, Metric Type: counter
Metric : go_memstats_alloc_bytes_total, Metric Labels: {__name__="go_memstats_alloc_bytes_total"}
Metric Name: go_memstats_buck_hash_sys_bytes, Metric Type: gauge
Metric : go_memstats_buck_hash_sys_bytes, Metric Labels: {__name__="go_memstats_buck_hash_sys_bytes"}
Metric Name: go_memstats_frees_total, Metric Type: counter
Metric : go_memstats_frees_total, Metric Labels: {__name__="go_memstats_frees_total"}
Metric Name: go_memstats_gc_cpu_fraction, Metric Type: gauge
Metric : go_memstats_gc_cpu_fraction, Metric Labels: {__name__="go_memstats_gc_cpu_fraction"}
Metric Name: go_memstats_gc_sys_bytes, Metric Type: gauge
Metric : go_memstats_gc_sys_bytes, Metric Labels: {__name__="go_memstats_gc_sys_bytes"}
Metric Name: go_memstats_heap_alloc_bytes, Metric Type: gauge
Metric : go_memstats_heap_alloc_bytes, Metric Labels: {__name__="go_memstats_heap_alloc_bytes"}
Metric Name: go_memstats_heap_idle_bytes, Metric Type: gauge
Metric : go_memstats_heap_idle_bytes, Metric Labels: {__name__="go_memstats_heap_idle_bytes"}
Metric Name: go_memstats_heap_inuse_bytes, Metric Type: gauge
Metric : go_memstats_heap_inuse_bytes, Metric Labels: {__name__="go_memstats_heap_inuse_bytes"}
Metric Name: go_memstats_heap_objects, Metric Type: gauge
Metric : go_memstats_heap_objects, Metric Labels: {__name__="go_memstats_heap_objects"}
Metric Name: go_memstats_heap_released_bytes, Metric Type: gauge
Metric : go_memstats_heap_released_bytes, Metric Labels: {__name__="go_memstats_heap_released_bytes"}
Metric Name: go_memstats_heap_sys_bytes, Metric Type: gauge
Metric : go_memstats_heap_sys_bytes, Metric Labels: {__name__="go_memstats_heap_sys_bytes"}
Metric Name: go_memstats_last_gc_time_seconds, Metric Type: gauge
Metric : go_memstats_last_gc_time_seconds, Metric Labels: {__name__="go_memstats_last_gc_time_seconds"}
Metric Name: go_memstats_lookups_total, Metric Type: counter
Metric : go_memstats_lookups_total, Metric Labels: {__name__="go_memstats_lookups_total"}
Metric Name: go_memstats_mallocs_total, Metric Type: counter
Metric : go_memstats_mallocs_total, Metric Labels: {__name__="go_memstats_mallocs_total"}
Metric Name: go_memstats_mcache_inuse_bytes, Metric Type: gauge
Metric : go_memstats_mcache_inuse_bytes, Metric Labels: {__name__="go_memstats_mcache_inuse_bytes"}
Metric Name: go_memstats_mcache_sys_bytes, Metric Type: gauge
Metric : go_memstats_mcache_sys_bytes, Metric Labels: {__name__="go_memstats_mcache_sys_bytes"}
Metric Name: go_memstats_mspan_inuse_bytes, Metric Type: gauge
Metric : go_memstats_mspan_inuse_bytes, Metric Labels: {__name__="go_memstats_mspan_inuse_bytes"}
Metric Name: go_memstats_mspan_sys_bytes, Metric Type: gauge
Metric : go_memstats_mspan_sys_bytes, Metric Labels: {__name__="go_memstats_mspan_sys_bytes"}
Metric Name: go_memstats_next_gc_bytes, Metric Type: gauge
Metric : go_memstats_next_gc_bytes, Metric Labels: {__name__="go_memstats_next_gc_bytes"}
Metric Name: go_memstats_other_sys_bytes, Metric Type: gauge
Metric : go_memstats_other_sys_bytes, Metric Labels: {__name__="go_memstats_other_sys_bytes"}
Metric Name: go_memstats_stack_inuse_bytes, Metric Type: gauge
Metric : go_memstats_stack_inuse_bytes, Metric Labels: {__name__="go_memstats_stack_inuse_bytes"}
Metric Name: go_memstats_stack_sys_bytes, Metric Type: gauge
Metric : go_memstats_stack_sys_bytes, Metric Labels: {__name__="go_memstats_stack_sys_bytes"}
Metric Name: go_memstats_sys_bytes, Metric Type: gauge
Metric : go_memstats_sys_bytes, Metric Labels: {__name__="go_memstats_sys_bytes"}
Metric Name: go_threads, Metric Type: gauge
Metric : go_threads, Metric Labels: {__name__="go_threads"}
Metric Name: http_request_duration_microseconds, Metric Type: summary
Metric : http_request_duration_microseconds{handler="prometheus",quantile="0.5"}, Metric Labels: {__name__="http_request_duration_microseconds", handler="prometheus", quantile="0.5"}
Metric : http_request_duration_microseconds{handler="prometheus",quantile="0.9"}, Metric Labels: {__name__="http_request_duration_microseconds", handler="prometheus", quantile="0.9"}
Metric : http_request_duration_microseconds{handler="prometheus",quantile="0.99"}, Metric Labels: {__name__="http_request_duration_microseconds", handler="prometheus", quantile="0.99"}
Metric : http_request_duration_microseconds_sum{handler="prometheus"}, Metric Labels: {__name__="http_request_duration_microseconds_sum", handler="prometheus"}
Metric : http_request_duration_microseconds_count{handler="prometheus"}, Metric Labels: {__name__="http_request_duration_microseconds_count", handler="prometheus"}
Metric Name: http_request_size_bytes, Metric Type: summary
Metric : http_request_size_bytes{handler="prometheus",quantile="0.5"}, Metric Labels: {__name__="http_request_size_bytes", handler="prometheus", quantile="0.5"}
Metric : http_request_size_bytes{handler="prometheus",quantile="0.9"}, Metric Labels: {__name__="http_request_size_bytes", handler="prometheus", quantile="0.9"}
Metric : http_request_size_bytes{handler="prometheus",quantile="0.99"}, Metric Labels: {__name__="http_request_size_bytes", handler="prometheus", quantile="0.99"}
Metric : http_request_size_bytes_sum{handler="prometheus"}, Metric Labels: {__name__="http_request_size_bytes_sum", handler="prometheus"}
Metric : http_request_size_bytes_count{handler="prometheus"}, Metric Labels: {__name__="http_request_size_bytes_count", handler="prometheus"}
Metric Name: http_requests_total, Metric Type: counter
Metric : http_requests_total{code="200",handler="prometheus",method="get"}, Metric Labels: {__name__="http_requests_total", code="200", handler="prometheus", method="get"}
Metric Name: http_response_size_bytes, Metric Type: summary
Metric : http_response_size_bytes{handler="prometheus",quantile="0.5"}, Metric Labels: {__name__="http_response_size_bytes", handler="prometheus", quantile="0.5"}
Metric : http_response_size_bytes{handler="prometheus",quantile="0.9"}, Metric Labels: {__name__="http_response_size_bytes", handler="prometheus", quantile="0.9"}
Metric : http_response_size_bytes{handler="prometheus",quantile="0.99"}, Metric Labels: {__name__="http_response_size_bytes", handler="prometheus", quantile="0.99"}
Metric : http_response_size_bytes_sum{handler="prometheus"}, Metric Labels: {__name__="http_response_size_bytes_sum", handler="prometheus"}
Metric : http_response_size_bytes_count{handler="prometheus"}, Metric Labels: {__name__="http_response_size_bytes_count", handler="prometheus"}
Metric Name: node_cpu, Metric Type: counter
Metric : node_cpu{cpu="cpu0",mode="idle"}, Metric Labels: {__name__="node_cpu", cpu="cpu0", mode="idle"}
Metric : node_cpu{cpu="cpu0",mode="nice"}, Metric Labels: {__name__="node_cpu", cpu="cpu0", mode="nice"}
Metric : node_cpu{cpu="cpu0",mode="system"}, Metric Labels: {__name__="node_cpu", cpu="cpu0", mode="system"}
Metric : node_cpu{cpu="cpu0",mode="user"}, Metric Labels: {__name__="node_cpu", cpu="cpu0", mode="user"}
Metric : node_cpu{cpu="cpu1",mode="idle"}, Metric Labels: {__name__="node_cpu", cpu="cpu1", mode="idle"}
Metric : node_cpu{cpu="cpu1",mode="nice"}, Metric Labels: {__name__="node_cpu", cpu="cpu1", mode="nice"}
Metric : node_cpu{cpu="cpu1",mode="system"}, Metric Labels: {__name__="node_cpu", cpu="cpu1", mode="system"}
Metric : node_cpu{cpu="cpu1",mode="user"}, Metric Labels: {__name__="node_cpu", cpu="cpu1", mode="user"}
Metric : node_cpu{cpu="cpu2",mode="idle"}, Metric Labels: {__name__="node_cpu", cpu="cpu2", mode="idle"}
Metric : node_cpu{cpu="cpu2",mode="nice"}, Metric Labels: {__name__="node_cpu", cpu="cpu2", mode="nice"}
Metric : node_cpu{cpu="cpu2",mode="system"}, Metric Labels: {__name__="node_cpu", cpu="cpu2", mode="system"}
Metric : node_cpu{cpu="cpu2",mode="user"}, Metric Labels: {__name__="node_cpu", cpu="cpu2", mode="user"}
Metric : node_cpu{cpu="cpu3",mode="idle"}, Metric Labels: {__name__="node_cpu", cpu="cpu3", mode="idle"}
Metric : node_cpu{cpu="cpu3",mode="nice"}, Metric Labels: {__name__="node_cpu", cpu="cpu3", mode="nice"}
Metric : node_cpu{cpu="cpu3",mode="system"}, Metric Labels: {__name__="node_cpu", cpu="cpu3", mode="system"}
Metric : node_cpu{cpu="cpu3",mode="user"}, Metric Labels: {__name__="node_cpu", cpu="cpu3", mode="user"}
Metric : node_cpu{cpu="cpu4",mode="idle"}, Metric Labels: {__name__="node_cpu", cpu="cpu4", mode="idle"}
Metric : node_cpu{cpu="cpu4",mode="nice"}, Metric Labels: {__name__="node_cpu", cpu="cpu4", mode="nice"}
Metric : node_cpu{cpu="cpu4",mode="system"}, Metric Labels: {__name__="node_cpu", cpu="cpu4", mode="system"}
Metric : node_cpu{cpu="cpu4",mode="user"}, Metric Labels: {__name__="node_cpu", cpu="cpu4", mode="user"}
Metric : node_cpu{cpu="cpu5",mode="idle"}, Metric Labels: {__name__="node_cpu", cpu="cpu5", mode="idle"}
Metric : node_cpu{cpu="cpu5",mode="nice"}, Metric Labels: {__name__="node_cpu", cpu="cpu5", mode="nice"}
Metric : node_cpu{cpu="cpu5",mode="system"}, Metric Labels: {__name__="node_cpu", cpu="cpu5", mode="system"}
Metric : node_cpu{cpu="cpu5",mode="user"}, Metric Labels: {__name__="node_cpu", cpu="cpu5", mode="user"}
Metric : node_cpu{cpu="cpu6",mode="idle"}, Metric Labels: {__name__="node_cpu", cpu="cpu6", mode="idle"}
Metric : node_cpu{cpu="cpu6",mode="nice"}, Metric Labels: {__name__="node_cpu", cpu="cpu6", mode="nice"}
Metric : node_cpu{cpu="cpu6",mode="system"}, Metric Labels: {__name__="node_cpu", cpu="cpu6", mode="system"}
Metric : node_cpu{cpu="cpu6",mode="user"}, Metric Labels: {__name__="node_cpu", cpu="cpu6", mode="user"}
Metric : node_cpu{cpu="cpu7",mode="idle"}, Metric Labels: {__name__="node_cpu", cpu="cpu7", mode="idle"}
Metric : node_cpu{cpu="cpu7",mode="nice"}, Metric Labels: {__name__="node_cpu", cpu="cpu7", mode="nice"}
Metric : node_cpu{cpu="cpu7",mode="system"}, Metric Labels: {__name__="node_cpu", cpu="cpu7", mode="system"}
Metric : node_cpu{cpu="cpu7",mode="user"}, Metric Labels: {__name__="node_cpu", cpu="cpu7", mode="user"}
Metric Name: node_disk_read_bytes_total, Metric Type: counter
Metric : node_disk_read_bytes_total{device="disk0"}, Metric Labels: {__name__="node_disk_read_bytes_total", device="disk0"}
Metric Name: node_disk_read_seconds_total, Metric Type: counter
Metric : node_disk_read_seconds_total{device="disk0"}, Metric Labels: {__name__="node_disk_read_seconds_total", device="disk0"}
Metric Name: node_disk_read_sectors_total, Metric Type: counter
Metric : node_disk_read_sectors_total{device="disk0"}, Metric Labels: {__name__="node_disk_read_sectors_total", device="disk0"}
Metric Name: node_disk_reads_completed_total, Metric Type: counter
Metric : node_disk_reads_completed_total{device="disk0"}, Metric Labels: {__name__="node_disk_reads_completed_total", device="disk0"}
Metric Name: node_disk_write_seconds_total, Metric Type: counter
Metric : node_disk_write_seconds_total{device="disk0"}, Metric Labels: {__name__="node_disk_write_seconds_total", device="disk0"}
Metric Name: node_disk_writes_completed_total, Metric Type: counter
Metric : node_disk_writes_completed_total{device="disk0"}, Metric Labels: {__name__="node_disk_writes_completed_total", device="disk0"}
Metric Name: node_disk_written_bytes_total, Metric Type: counter
Metric : node_disk_written_bytes_total{device="disk0"}, Metric Labels: {__name__="node_disk_written_bytes_total", device="disk0"}
Metric Name: node_disk_written_sectors_total, Metric Type: counter
Metric : node_disk_written_sectors_total{device="disk0"}, Metric Labels: {__name__="node_disk_written_sectors_total", device="disk0"}
Metric Name: node_exporter_build_info, Metric Type: gauge
Metric : node_exporter_build_info{branch="HEAD",goversion="go1.9.1",revision="6e2053c557f96efb63aef3691f15335a70baaffd",version="0.15.0"}, Metric Labels: {__name__="node_exporter_build_info", branch="HEAD", goversion="go1.9.1", revision="6e2053c557f96efb63aef3691f15335a70baaffd", version="0.15.0"}
Metric Name: node_filesystem_avail, Metric Type: gauge
Metric : node_filesystem_avail{device="/dev/disk1s1",fstype="apfs",mountpoint="/System/Volumes/Data"}, Metric Labels: {__name__="node_filesystem_avail", device="/dev/disk1s1", fstype="apfs", mountpoint="/System/Volumes/Data"}
Metric : node_filesystem_avail{device="/dev/disk1s4",fstype="apfs",mountpoint="/private/var/vm"}, Metric Labels: {__name__="node_filesystem_avail", device="/dev/disk1s4", fstype="apfs", mountpoint="/private/var/vm"}
Metric : node_filesystem_avail{device="/dev/disk1s5",fstype="apfs",mountpoint="/"}, Metric Labels: {__name__="node_filesystem_avail", device="/dev/disk1s5", fstype="apfs", mountpoint="/"}
Metric : node_filesystem_avail{device="map auto_home",fstype="autofs",mountpoint="/System/Volumes/Data/home"}, Metric Labels: {__name__="node_filesystem_avail", device="map auto_home", fstype="autofs", mountpoint="/System/Volumes/Data/home"}
Metric Name: node_filesystem_device_error, Metric Type: gauge
Metric : node_filesystem_device_error{device="/dev/disk1s1",fstype="apfs",mountpoint="/System/Volumes/Data"}, Metric Labels: {__name__="node_filesystem_device_error", device="/dev/disk1s1", fstype="apfs", mountpoint="/System/Volumes/Data"}
Metric : node_filesystem_device_error{device="/dev/disk1s4",fstype="apfs",mountpoint="/private/var/vm"}, Metric Labels: {__name__="node_filesystem_device_error", device="/dev/disk1s4", fstype="apfs", mountpoint="/private/var/vm"}
Metric : node_filesystem_device_error{device="/dev/disk1s5",fstype="apfs",mountpoint="/"}, Metric Labels: {__name__="node_filesystem_device_error", device="/dev/disk1s5", fstype="apfs", mountpoint="/"}
Metric : node_filesystem_device_error{device="map auto_home",fstype="autofs",mountpoint="/System/Volumes/Data/home"}, Metric Labels: {__name__="node_filesystem_device_error", device="map auto_home", fstype="autofs", mountpoint="/System/Volumes/Data/home"}
Metric Name: node_filesystem_files, Metric Type: gauge
Metric : node_filesystem_files{device="/dev/disk1s1",fstype="apfs",mountpoint="/System/Volumes/Data"}, Metric Labels: {__name__="node_filesystem_files", device="/dev/disk1s1", fstype="apfs", mountpoint="/System/Volumes/Data"}
Metric : node_filesystem_files{device="/dev/disk1s4",fstype="apfs",mountpoint="/private/var/vm"}, Metric Labels: {__name__="node_filesystem_files", device="/dev/disk1s4", fstype="apfs", mountpoint="/private/var/vm"}
Metric : node_filesystem_files{device="/dev/disk1s5",fstype="apfs",mountpoint="/"}, Metric Labels: {__name__="node_filesystem_files", device="/dev/disk1s5", fstype="apfs", mountpoint="/"}
Metric : node_filesystem_files{device="map auto_home",fstype="autofs",mountpoint="/System/Volumes/Data/home"}, Metric Labels: {__name__="node_filesystem_files", device="map auto_home", fstype="autofs", mountpoint="/System/Volumes/Data/home"}
Metric Name: node_filesystem_files_free, Metric Type: gauge
Metric : node_filesystem_files_free{device="/dev/disk1s1",fstype="apfs",mountpoint="/System/Volumes/Data"}, Metric Labels: {__name__="node_filesystem_files_free", device="/dev/disk1s1", fstype="apfs", mountpoint="/System/Volumes/Data"}
Metric : node_filesystem_files_free{device="/dev/disk1s4",fstype="apfs",mountpoint="/private/var/vm"}, Metric Labels: {__name__="node_filesystem_files_free", device="/dev/disk1s4", fstype="apfs", mountpoint="/private/var/vm"}
Metric : node_filesystem_files_free{device="/dev/disk1s5",fstype="apfs",mountpoint="/"}, Metric Labels: {__name__="node_filesystem_files_free", device="/dev/disk1s5", fstype="apfs", mountpoint="/"}
Metric : node_filesystem_files_free{device="map auto_home",fstype="autofs",mountpoint="/System/Volumes/Data/home"}, Metric Labels: {__name__="node_filesystem_files_free", device="map auto_home", fstype="autofs", mountpoint="/System/Volumes/Data/home"}
Metric Name: node_filesystem_free, Metric Type: gauge
Metric : node_filesystem_free{device="/dev/disk1s1",fstype="apfs",mountpoint="/System/Volumes/Data"}, Metric Labels: {__name__="node_filesystem_free", device="/dev/disk1s1", fstype="apfs", mountpoint="/System/Volumes/Data"}
Metric : node_filesystem_free{device="/dev/disk1s4",fstype="apfs",mountpoint="/private/var/vm"}, Metric Labels: {__name__="node_filesystem_free", device="/dev/disk1s4", fstype="apfs", mountpoint="/private/var/vm"}
Metric : node_filesystem_free{device="/dev/disk1s5",fstype="apfs",mountpoint="/"}, Metric Labels: {__name__="node_filesystem_free", device="/dev/disk1s5", fstype="apfs", mountpoint="/"}
Metric : node_filesystem_free{device="map auto_home",fstype="autofs",mountpoint="/System/Volumes/Data/home"}, Metric Labels: {__name__="node_filesystem_free", device="map auto_home", fstype="autofs", mountpoint="/System/Volumes/Data/home"}
Metric Name: node_filesystem_readonly, Metric Type: gauge
Metric : node_filesystem_readonly{device="/dev/disk1s1",fstype="apfs",mountpoint="/System/Volumes/Data"}, Metric Labels: {__name__="node_filesystem_readonly", device="/dev/disk1s1", fstype="apfs", mountpoint="/System/Volumes/Data"}
Metric : node_filesystem_readonly{device="/dev/disk1s4",fstype="apfs",mountpoint="/private/var/vm"}, Metric Labels: {__name__="node_filesystem_readonly", device="/dev/disk1s4", fstype="apfs", mountpoint="/private/var/vm"}
Metric : node_filesystem_readonly{device="/dev/disk1s5",fstype="apfs",mountpoint="/"}, Metric Labels: {__name__="node_filesystem_readonly", device="/dev/disk1s5", fstype="apfs", mountpoint="/"}
Metric : node_filesystem_readonly{device="map auto_home",fstype="autofs",mountpoint="/System/Volumes/Data/home"}, Metric Labels: {__name__="node_filesystem_readonly", device="map auto_home", fstype="autofs", mountpoint="/System/Volumes/Data/home"}
Metric Name: node_filesystem_size, Metric Type: gauge
Metric : node_filesystem_size{device="/dev/disk1s1",fstype="apfs",mountpoint="/System/Volumes/Data"}, Metric Labels: {__name__="node_filesystem_size", device="/dev/disk1s1", fstype="apfs", mountpoint="/System/Volumes/Data"}
Metric : node_filesystem_size{device="/dev/disk1s4",fstype="apfs",mountpoint="/private/var/vm"}, Metric Labels: {__name__="node_filesystem_size", device="/dev/disk1s4", fstype="apfs", mountpoint="/private/var/vm"}
Metric : node_filesystem_size{device="/dev/disk1s5",fstype="apfs",mountpoint="/"}, Metric Labels: {__name__="node_filesystem_size", device="/dev/disk1s5", fstype="apfs", mountpoint="/"}
Metric : node_filesystem_size{device="map auto_home",fstype="autofs",mountpoint="/System/Volumes/Data/home"}, Metric Labels: {__name__="node_filesystem_size", device="map auto_home", fstype="autofs", mountpoint="/System/Volumes/Data/home"}
Metric Name: node_load1, Metric Type: gauge
Metric : node_load1, Metric Labels: {__name__="node_load1"}
Metric Name: node_load15, Metric Type: gauge
Metric : node_load15, Metric Labels: {__name__="node_load15"}
Metric Name: node_load5, Metric Type: gauge
Metric : node_load5, Metric Labels: {__name__="node_load5"}
Metric Name: node_memory_active_bytes_total, Metric Type: gauge
Metric : node_memory_active_bytes_total, Metric Labels: {__name__="node_memory_active_bytes_total"}
Metric Name: node_memory_bytes_total, Metric Type: gauge
Metric : node_memory_bytes_total, Metric Labels: {__name__="node_memory_bytes_total"}
Metric Name: node_memory_free_bytes_total, Metric Type: gauge
Metric : node_memory_free_bytes_total, Metric Labels: {__name__="node_memory_free_bytes_total"}
Metric Name: node_memory_inactive_bytes_total, Metric Type: gauge
Metric : node_memory_inactive_bytes_total, Metric Labels: {__name__="node_memory_inactive_bytes_total"}
Metric Name: node_memory_swapped_in_pages_total, Metric Type: gauge
Metric : node_memory_swapped_in_pages_total, Metric Labels: {__name__="node_memory_swapped_in_pages_total"}
Metric Name: node_memory_swapped_out_pages_total, Metric Type: gauge
Metric : node_memory_swapped_out_pages_total, Metric Labels: {__name__="node_memory_swapped_out_pages_total"}
Metric Name: node_memory_wired_bytes_total, Metric Type: gauge
Metric : node_memory_wired_bytes_total, Metric Labels: {__name__="node_memory_wired_bytes_total"}
Metric Name: node_network_receive_bytes, Metric Type: gauge
Metric : node_network_receive_bytes{device="awdl0"}, Metric Labels: {__name__="node_network_receive_bytes", device="awdl0"}
Metric : node_network_receive_bytes{device="bridge0"}, Metric Labels: {__name__="node_network_receive_bytes", device="bridge0"}
Metric : node_network_receive_bytes{device="en0"}, Metric Labels: {__name__="node_network_receive_bytes", device="en0"}
Metric : node_network_receive_bytes{device="en1"}, Metric Labels: {__name__="node_network_receive_bytes", device="en1"}
Metric : node_network_receive_bytes{device="en2"}, Metric Labels: {__name__="node_network_receive_bytes", device="en2"}
Metric : node_network_receive_bytes{device="gif0"}, Metric Labels: {__name__="node_network_receive_bytes", device="gif0"}
Metric : node_network_receive_bytes{device="llw0"}, Metric Labels: {__name__="node_network_receive_bytes", device="llw0"}
Metric : node_network_receive_bytes{device="lo0"}, Metric Labels: {__name__="node_network_receive_bytes", device="lo0"}
Metric : node_network_receive_bytes{device="p2p0"}, Metric Labels: {__name__="node_network_receive_bytes", device="p2p0"}
Metric : node_network_receive_bytes{device="stf0"}, Metric Labels: {__name__="node_network_receive_bytes", device="stf0"}
Metric : node_network_receive_bytes{device="utun0"}, Metric Labels: {__name__="node_network_receive_bytes", device="utun0"}
Metric : node_network_receive_bytes{device="utun1"}, Metric Labels: {__name__="node_network_receive_bytes", device="utun1"}
Metric Name: node_network_receive_errs, Metric Type: gauge
Metric : node_network_receive_errs{device="awdl0"}, Metric Labels: {__name__="node_network_receive_errs", device="awdl0"}
Metric : node_network_receive_errs{device="bridge0"}, Metric Labels: {__name__="node_network_receive_errs", device="bridge0"}
Metric : node_network_receive_errs{device="en0"}, Metric Labels: {__name__="node_network_receive_errs", device="en0"}
Metric : node_network_receive_errs{device="en1"}, Metric Labels: {__name__="node_network_receive_errs", device="en1"}
Metric : node_network_receive_errs{device="en2"}, Metric Labels: {__name__="node_network_receive_errs", device="en2"}
Metric : node_network_receive_errs{device="gif0"}, Metric Labels: {__name__="node_network_receive_errs", device="gif0"}
Metric : node_network_receive_errs{device="llw0"}, Metric Labels: {__name__="node_network_receive_errs", device="llw0"}
Metric : node_network_receive_errs{device="lo0"}, Metric Labels: {__name__="node_network_receive_errs", device="lo0"}
Metric : node_network_receive_errs{device="p2p0"}, Metric Labels: {__name__="node_network_receive_errs", device="p2p0"}
Metric : node_network_receive_errs{device="stf0"}, Metric Labels: {__name__="node_network_receive_errs", device="stf0"}
Metric : node_network_receive_errs{device="utun0"}, Metric Labels: {__name__="node_network_receive_errs", device="utun0"}
Metric : node_network_receive_errs{device="utun1"}, Metric Labels: {__name__="node_network_receive_errs", device="utun1"}
Metric Name: node_network_receive_multicast, Metric Type: gauge
Metric : node_network_receive_multicast{device="awdl0"}, Metric Labels: {__name__="node_network_receive_multicast", device="awdl0"}
Metric : node_network_receive_multicast{device="bridge0"}, Metric Labels: {__name__="node_network_receive_multicast", device="bridge0"}
Metric : node_network_receive_multicast{device="en0"}, Metric Labels: {__name__="node_network_receive_multicast", device="en0"}
Metric : node_network_receive_multicast{device="en1"}, Metric Labels: {__name__="node_network_receive_multicast", device="en1"}
Metric : node_network_receive_multicast{device="en2"}, Metric Labels: {__name__="node_network_receive_multicast", device="en2"}
Metric : node_network_receive_multicast{device="gif0"}, Metric Labels: {__name__="node_network_receive_multicast", device="gif0"}
Metric : node_network_receive_multicast{device="llw0"}, Metric Labels: {__name__="node_network_receive_multicast", device="llw0"}
Metric : node_network_receive_multicast{device="lo0"}, Metric Labels: {__name__="node_network_receive_multicast", device="lo0"}
Metric : node_network_receive_multicast{device="p2p0"}, Metric Labels: {__name__="node_network_receive_multicast", device="p2p0"}
Metric : node_network_receive_multicast{device="stf0"}, Metric Labels: {__name__="node_network_receive_multicast", device="stf0"}
Metric : node_network_receive_multicast{device="utun0"}, Metric Labels: {__name__="node_network_receive_multicast", device="utun0"}
Metric : node_network_receive_multicast{device="utun1"}, Metric Labels: {__name__="node_network_receive_multicast", device="utun1"}
Metric Name: node_network_receive_packets, Metric Type: gauge
Metric : node_network_receive_packets{device="awdl0"}, Metric Labels: {__name__="node_network_receive_packets", device="awdl0"}
Metric : node_network_receive_packets{device="bridge0"}, Metric Labels: {__name__="node_network_receive_packets", device="bridge0"}
Metric : node_network_receive_packets{device="en0"}, Metric Labels: {__name__="node_network_receive_packets", device="en0"}
Metric : node_network_receive_packets{device="en1"}, Metric Labels: {__name__="node_network_receive_packets", device="en1"}
Metric : node_network_receive_packets{device="en2"}, Metric Labels: {__name__="node_network_receive_packets", device="en2"}
Metric : node_network_receive_packets{device="gif0"}, Metric Labels: {__name__="node_network_receive_packets", device="gif0"}
Metric : node_network_receive_packets{device="llw0"}, Metric Labels: {__name__="node_network_receive_packets", device="llw0"}
Metric : node_network_receive_packets{device="lo0"}, Metric Labels: {__name__="node_network_receive_packets", device="lo0"}
Metric : node_network_receive_packets{device="p2p0"}, Metric Labels: {__name__="node_network_receive_packets", device="p2p0"}
Metric : node_network_receive_packets{device="stf0"}, Metric Labels: {__name__="node_network_receive_packets", device="stf0"}
Metric : node_network_receive_packets{device="utun0"}, Metric Labels: {__name__="node_network_receive_packets", device="utun0"}
Metric : node_network_receive_packets{device="utun1"}, Metric Labels: {__name__="node_network_receive_packets", device="utun1"}
Metric Name: node_network_transmit_bytes, Metric Type: gauge
Metric : node_network_transmit_bytes{device="awdl0"}, Metric Labels: {__name__="node_network_transmit_bytes", device="awdl0"}
Metric : node_network_transmit_bytes{device="bridge0"}, Metric Labels: {__name__="node_network_transmit_bytes", device="bridge0"}
Metric : node_network_transmit_bytes{device="en0"}, Metric Labels: {__name__="node_network_transmit_bytes", device="en0"}
Metric : node_network_transmit_bytes{device="en1"}, Metric Labels: {__name__="node_network_transmit_bytes", device="en1"}
Metric : node_network_transmit_bytes{device="en2"}, Metric Labels: {__name__="node_network_transmit_bytes", device="en2"}
Metric : node_network_transmit_bytes{device="gif0"}, Metric Labels: {__name__="node_network_transmit_bytes", device="gif0"}
Metric : node_network_transmit_bytes{device="llw0"}, Metric Labels: {__name__="node_network_transmit_bytes", device="llw0"}
Metric : node_network_transmit_bytes{device="lo0"}, Metric Labels: {__name__="node_network_transmit_bytes", device="lo0"}
Metric : node_network_transmit_bytes{device="p2p0"}, Metric Labels: {__name__="node_network_transmit_bytes", device="p2p0"}
Metric : node_network_transmit_bytes{device="stf0"}, Metric Labels: {__name__="node_network_transmit_bytes", device="stf0"}
Metric : node_network_transmit_bytes{device="utun0"}, Metric Labels: {__name__="node_network_transmit_bytes", device="utun0"}
Metric : node_network_transmit_bytes{device="utun1"}, Metric Labels: {__name__="node_network_transmit_bytes", device="utun1"}
Metric Name: node_network_transmit_errs, Metric Type: gauge
Metric : node_network_transmit_errs{device="awdl0"}, Metric Labels: {__name__="node_network_transmit_errs", device="awdl0"}
Metric : node_network_transmit_errs{device="bridge0"}, Metric Labels: {__name__="node_network_transmit_errs", device="bridge0"}
Metric : node_network_transmit_errs{device="en0"}, Metric Labels: {__name__="node_network_transmit_errs", device="en0"}
Metric : node_network_transmit_errs{device="en1"}, Metric Labels: {__name__="node_network_transmit_errs", device="en1"}
Metric : node_network_transmit_errs{device="en2"}, Metric Labels: {__name__="node_network_transmit_errs", device="en2"}
Metric : node_network_transmit_errs{device="gif0"}, Metric Labels: {__name__="node_network_transmit_errs", device="gif0"}
Metric : node_network_transmit_errs{device="llw0"}, Metric Labels: {__name__="node_network_transmit_errs", device="llw0"}
Metric : node_network_transmit_errs{device="lo0"}, Metric Labels: {__name__="node_network_transmit_errs", device="lo0"}
Metric : node_network_transmit_errs{device="p2p0"}, Metric Labels: {__name__="node_network_transmit_errs", device="p2p0"}
Metric : node_network_transmit_errs{device="stf0"}, Metric Labels: {__name__="node_network_transmit_errs", device="stf0"}
Metric : node_network_transmit_errs{device="utun0"}, Metric Labels: {__name__="node_network_transmit_errs", device="utun0"}
Metric : node_network_transmit_errs{device="utun1"}, Metric Labels: {__name__="node_network_transmit_errs", device="utun1"}
Metric Name: node_network_transmit_multicast, Metric Type: gauge
Metric : node_network_transmit_multicast{device="awdl0"}, Metric Labels: {__name__="node_network_transmit_multicast", device="awdl0"}
Metric : node_network_transmit_multicast{device="bridge0"}, Metric Labels: {__name__="node_network_transmit_multicast", device="bridge0"}
Metric : node_network_transmit_multicast{device="en0"}, Metric Labels: {__name__="node_network_transmit_multicast", device="en0"}
Metric : node_network_transmit_multicast{device="en1"}, Metric Labels: {__name__="node_network_transmit_multicast", device="en1"}
Metric : node_network_transmit_multicast{device="en2"}, Metric Labels: {__name__="node_network_transmit_multicast", device="en2"}
Metric : node_network_transmit_multicast{device="gif0"}, Metric Labels: {__name__="node_network_transmit_multicast", device="gif0"}
Metric : node_network_transmit_multicast{device="llw0"}, Metric Labels: {__name__="node_network_transmit_multicast", device="llw0"}
Metric : node_network_transmit_multicast{device="lo0"}, Metric Labels: {__name__="node_network_transmit_multicast", device="lo0"}
Metric : node_network_transmit_multicast{device="p2p0"}, Metric Labels: {__name__="node_network_transmit_multicast", device="p2p0"}
Metric : node_network_transmit_multicast{device="stf0"}, Metric Labels: {__name__="node_network_transmit_multicast", device="stf0"}
Metric : node_network_transmit_multicast{device="utun0"}, Metric Labels: {__name__="node_network_transmit_multicast", device="utun0"}
Metric : node_network_transmit_multicast{device="utun1"}, Metric Labels: {__name__="node_network_transmit_multicast", device="utun1"}
Metric Name: node_network_transmit_packets, Metric Type: gauge
Metric : node_network_transmit_packets{device="awdl0"}, Metric Labels: {__name__="node_network_transmit_packets", device="awdl0"}
Metric : node_network_transmit_packets{device="bridge0"}, Metric Labels: {__name__="node_network_transmit_packets", device="bridge0"}
Metric : node_network_transmit_packets{device="en0"}, Metric Labels: {__name__="node_network_transmit_packets", device="en0"}
Metric : node_network_transmit_packets{device="en1"}, Metric Labels: {__name__="node_network_transmit_packets", device="en1"}
Metric : node_network_transmit_packets{device="en2"}, Metric Labels: {__name__="node_network_transmit_packets", device="en2"}
Metric : node_network_transmit_packets{device="gif0"}, Metric Labels: {__name__="node_network_transmit_packets", device="gif0"}
Metric : node_network_transmit_packets{device="llw0"}, Metric Labels: {__name__="node_network_transmit_packets", device="llw0"}
Metric : node_network_transmit_packets{device="lo0"}, Metric Labels: {__name__="node_network_transmit_packets", device="lo0"}
Metric : node_network_transmit_packets{device="p2p0"}, Metric Labels: {__name__="node_network_transmit_packets", device="p2p0"}
Metric : node_network_transmit_packets{device="stf0"}, Metric Labels: {__name__="node_network_transmit_packets", device="stf0"}
Metric : node_network_transmit_packets{device="utun0"}, Metric Labels: {__name__="node_network_transmit_packets", device="utun0"}
Metric : node_network_transmit_packets{device="utun1"}, Metric Labels: {__name__="node_network_transmit_packets", device="utun1"}
Metric Name: node_scrape_collector_duration_seconds, Metric Type: gauge
Metric : node_scrape_collector_duration_seconds{collector="cpu"}, Metric Labels: {__name__="node_scrape_collector_duration_seconds", collector="cpu"}
Metric : node_scrape_collector_duration_seconds{collector="diskstats"}, Metric Labels: {__name__="node_scrape_collector_duration_seconds", collector="diskstats"}
Metric : node_scrape_collector_duration_seconds{collector="filesystem"}, Metric Labels: {__name__="node_scrape_collector_duration_seconds", collector="filesystem"}
Metric : node_scrape_collector_duration_seconds{collector="loadavg"}, Metric Labels: {__name__="node_scrape_collector_duration_seconds", collector="loadavg"}
Metric : node_scrape_collector_duration_seconds{collector="meminfo"}, Metric Labels: {__name__="node_scrape_collector_duration_seconds", collector="meminfo"}
Metric : node_scrape_collector_duration_seconds{collector="netdev"}, Metric Labels: {__name__="node_scrape_collector_duration_seconds", collector="netdev"}
Metric : node_scrape_collector_duration_seconds{collector="textfile"}, Metric Labels: {__name__="node_scrape_collector_duration_seconds", collector="textfile"}
Metric : node_scrape_collector_duration_seconds{collector="time"}, Metric Labels: {__name__="node_scrape_collector_duration_seconds", collector="time"}
Metric Name: node_scrape_collector_success, Metric Type: gauge
Metric : node_scrape_collector_success{collector="cpu"}, Metric Labels: {__name__="node_scrape_collector_success", collector="cpu"}
Metric : node_scrape_collector_success{collector="diskstats"}, Metric Labels: {__name__="node_scrape_collector_success", collector="diskstats"}
Metric : node_scrape_collector_success{collector="filesystem"}, Metric Labels: {__name__="node_scrape_collector_success", collector="filesystem"}
Metric : node_scrape_collector_success{collector="loadavg"}, Metric Labels: {__name__="node_scrape_collector_success", collector="loadavg"}
Metric : node_scrape_collector_success{collector="meminfo"}, Metric Labels: {__name__="node_scrape_collector_success", collector="meminfo"}
Metric : node_scrape_collector_success{collector="netdev"}, Metric Labels: {__name__="node_scrape_collector_success", collector="netdev"}
Metric : node_scrape_collector_success{collector="textfile"}, Metric Labels: {__name__="node_scrape_collector_success", collector="textfile"}
Metric : node_scrape_collector_success{collector="time"}, Metric Labels: {__name__="node_scrape_collector_success", collector="time"}
Metric Name: node_time, Metric Type: gauge
Metric : node_time, Metric Labels: {__name__="node_time"}



new metrics : 
Metric Name: go_gc_duration_seconds, Metric Type: summary
Metric : go_gc_duration_seconds{quantile="0"}, Metric Labels: {__name__="go_gc_duration_seconds", quantile="0"}
Metric : go_gc_duration_seconds{quantile="0.25"}, Metric Labels: {__name__="go_gc_duration_seconds", quantile="0.25"}
Metric : go_gc_duration_seconds{quantile="0.5"}, Metric Labels: {__name__="go_gc_duration_seconds", quantile="0.5"}
Metric : go_gc_duration_seconds{quantile="0.75"}, Metric Labels: {__name__="go_gc_duration_seconds", quantile="0.75"}
Metric : go_gc_duration_seconds{quantile="1"}, Metric Labels: {__name__="go_gc_duration_seconds", quantile="1"}
Metric : go_gc_duration_seconds_sum, Metric Labels: {__name__="go_gc_duration_seconds_sum"}
Metric : go_gc_duration_seconds_count, Metric Labels: {__name__="go_gc_duration_seconds_count"}
Metric Name: go_goroutines, Metric Type: gauge
Metric : go_goroutines, Metric Labels: {__name__="go_goroutines"}
Metric Name: go_info, Metric Type: gauge
Metric : go_info{version="go1.11.10"}, Metric Labels: {__name__="go_info", version="go1.11.10"}
Metric Name: go_memstats_alloc_bytes, Metric Type: gauge
Metric : go_memstats_alloc_bytes, Metric Labels: {__name__="go_memstats_alloc_bytes"}
Metric Name: go_memstats_alloc_bytes_total, Metric Type: counter
Metric : go_memstats_alloc_bytes_total, Metric Labels: {__name__="go_memstats_alloc_bytes_total"}
Metric Name: go_memstats_buck_hash_sys_bytes, Metric Type: gauge
Metric : go_memstats_buck_hash_sys_bytes, Metric Labels: {__name__="go_memstats_buck_hash_sys_bytes"}
Metric Name: go_memstats_frees_total, Metric Type: counter
Metric : go_memstats_frees_total, Metric Labels: {__name__="go_memstats_frees_total"}
Metric Name: go_memstats_gc_cpu_fraction, Metric Type: gauge
Metric : go_memstats_gc_cpu_fraction, Metric Labels: {__name__="go_memstats_gc_cpu_fraction"}
Metric Name: go_memstats_gc_sys_bytes, Metric Type: gauge
Metric : go_memstats_gc_sys_bytes, Metric Labels: {__name__="go_memstats_gc_sys_bytes"}
Metric Name: go_memstats_heap_alloc_bytes, Metric Type: gauge
Metric : go_memstats_heap_alloc_bytes, Metric Labels: {__name__="go_memstats_heap_alloc_bytes"}
Metric Name: go_memstats_heap_idle_bytes, Metric Type: gauge
Metric : go_memstats_heap_idle_bytes, Metric Labels: {__name__="go_memstats_heap_idle_bytes"}
Metric Name: go_memstats_heap_inuse_bytes, Metric Type: gauge
Metric : go_memstats_heap_inuse_bytes, Metric Labels: {__name__="go_memstats_heap_inuse_bytes"}
Metric Name: go_memstats_heap_objects, Metric Type: gauge
Metric : go_memstats_heap_objects, Metric Labels: {__name__="go_memstats_heap_objects"}
Metric Name: go_memstats_heap_released_bytes, Metric Type: gauge
Metric : go_memstats_heap_released_bytes, Metric Labels: {__name__="go_memstats_heap_released_bytes"}
Metric Name: go_memstats_heap_sys_bytes, Metric Type: gauge
Metric : go_memstats_heap_sys_bytes, Metric Labels: {__name__="go_memstats_heap_sys_bytes"}
Metric Name: go_memstats_last_gc_time_seconds, Metric Type: gauge
Metric : go_memstats_last_gc_time_seconds, Metric Labels: {__name__="go_memstats_last_gc_time_seconds"}
Metric Name: go_memstats_lookups_total, Metric Type: counter
Metric : go_memstats_lookups_total, Metric Labels: {__name__="go_memstats_lookups_total"}
Metric Name: go_memstats_mallocs_total, Metric Type: counter
Metric : go_memstats_mallocs_total, Metric Labels: {__name__="go_memstats_mallocs_total"}
Metric Name: go_memstats_mcache_inuse_bytes, Metric Type: gauge
Metric : go_memstats_mcache_inuse_bytes, Metric Labels: {__name__="go_memstats_mcache_inuse_bytes"}
Metric Name: go_memstats_mcache_sys_bytes, Metric Type: gauge
Metric : go_memstats_mcache_sys_bytes, Metric Labels: {__name__="go_memstats_mcache_sys_bytes"}
Metric Name: go_memstats_mspan_inuse_bytes, Metric Type: gauge
Metric : go_memstats_mspan_inuse_bytes, Metric Labels: {__name__="go_memstats_mspan_inuse_bytes"}
Metric Name: go_memstats_mspan_sys_bytes, Metric Type: gauge
Metric : go_memstats_mspan_sys_bytes, Metric Labels: {__name__="go_memstats_mspan_sys_bytes"}
Metric Name: go_memstats_next_gc_bytes, Metric Type: gauge
Metric : go_memstats_next_gc_bytes, Metric Labels: {__name__="go_memstats_next_gc_bytes"}
Metric Name: go_memstats_other_sys_bytes, Metric Type: gauge
Metric : go_memstats_other_sys_bytes, Metric Labels: {__name__="go_memstats_other_sys_bytes"}
Metric Name: go_memstats_stack_inuse_bytes, Metric Type: gauge
Metric : go_memstats_stack_inuse_bytes, Metric Labels: {__name__="go_memstats_stack_inuse_bytes"}
Metric Name: go_memstats_stack_sys_bytes, Metric Type: gauge
Metric : go_memstats_stack_sys_bytes, Metric Labels: {__name__="go_memstats_stack_sys_bytes"}
Metric Name: go_memstats_sys_bytes, Metric Type: gauge
Metric : go_memstats_sys_bytes, Metric Labels: {__name__="go_memstats_sys_bytes"}
Metric Name: go_threads, Metric Type: gauge
Metric : go_threads, Metric Labels: {__name__="go_threads"}
Metric Name: node_boot_time_seconds, Metric Type: gauge
Metric : node_boot_time_seconds, Metric Labels: {__name__="node_boot_time_seconds"}
Metric Name: node_cpu_seconds_total, Metric Type: counter
Metric : node_cpu_seconds_total{cpu="0",mode="idle"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="0", mode="idle"}
Metric : node_cpu_seconds_total{cpu="0",mode="nice"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="0", mode="nice"}
Metric : node_cpu_seconds_total{cpu="0",mode="system"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="0", mode="system"}
Metric : node_cpu_seconds_total{cpu="0",mode="user"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="0", mode="user"}
Metric : node_cpu_seconds_total{cpu="1",mode="idle"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="1", mode="idle"}
Metric : node_cpu_seconds_total{cpu="1",mode="nice"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="1", mode="nice"}
Metric : node_cpu_seconds_total{cpu="1",mode="system"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="1", mode="system"}
Metric : node_cpu_seconds_total{cpu="1",mode="user"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="1", mode="user"}
Metric : node_cpu_seconds_total{cpu="2",mode="idle"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="2", mode="idle"}
Metric : node_cpu_seconds_total{cpu="2",mode="nice"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="2", mode="nice"}
Metric : node_cpu_seconds_total{cpu="2",mode="system"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="2", mode="system"}
Metric : node_cpu_seconds_total{cpu="2",mode="user"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="2", mode="user"}
Metric : node_cpu_seconds_total{cpu="3",mode="idle"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="3", mode="idle"}
Metric : node_cpu_seconds_total{cpu="3",mode="nice"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="3", mode="nice"}
Metric : node_cpu_seconds_total{cpu="3",mode="system"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="3", mode="system"}
Metric : node_cpu_seconds_total{cpu="3",mode="user"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="3", mode="user"}
Metric : node_cpu_seconds_total{cpu="4",mode="idle"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="4", mode="idle"}
Metric : node_cpu_seconds_total{cpu="4",mode="nice"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="4", mode="nice"}
Metric : node_cpu_seconds_total{cpu="4",mode="system"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="4", mode="system"}
Metric : node_cpu_seconds_total{cpu="4",mode="user"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="4", mode="user"}
Metric : node_cpu_seconds_total{cpu="5",mode="idle"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="5", mode="idle"}
Metric : node_cpu_seconds_total{cpu="5",mode="nice"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="5", mode="nice"}
Metric : node_cpu_seconds_total{cpu="5",mode="system"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="5", mode="system"}
Metric : node_cpu_seconds_total{cpu="5",mode="user"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="5", mode="user"}
Metric : node_cpu_seconds_total{cpu="6",mode="idle"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="6", mode="idle"}
Metric : node_cpu_seconds_total{cpu="6",mode="nice"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="6", mode="nice"}
Metric : node_cpu_seconds_total{cpu="6",mode="system"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="6", mode="system"}
Metric : node_cpu_seconds_total{cpu="6",mode="user"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="6", mode="user"}
Metric : node_cpu_seconds_total{cpu="7",mode="idle"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="7", mode="idle"}
Metric : node_cpu_seconds_total{cpu="7",mode="nice"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="7", mode="nice"}
Metric : node_cpu_seconds_total{cpu="7",mode="system"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="7", mode="system"}
Metric : node_cpu_seconds_total{cpu="7",mode="user"}, Metric Labels: {__name__="node_cpu_seconds_total", cpu="7", mode="user"}
Metric Name: node_disk_read_bytes_total, Metric Type: counter
Metric : node_disk_read_bytes_total{device="disk0"}, Metric Labels: {__name__="node_disk_read_bytes_total", device="disk0"}
Metric Name: node_disk_read_sectors_total, Metric Type: counter
Metric : node_disk_read_sectors_total{device="disk0"}, Metric Labels: {__name__="node_disk_read_sectors_total", device="disk0"}
Metric Name: node_disk_read_time_seconds_total, Metric Type: counter
Metric : node_disk_read_time_seconds_total{device="disk0"}, Metric Labels: {__name__="node_disk_read_time_seconds_total", device="disk0"}
Metric Name: node_disk_reads_completed_total, Metric Type: counter
Metric : node_disk_reads_completed_total{device="disk0"}, Metric Labels: {__name__="node_disk_reads_completed_total", device="disk0"}
Metric Name: node_disk_write_time_seconds_total, Metric Type: counter
Metric : node_disk_write_time_seconds_total{device="disk0"}, Metric Labels: {__name__="node_disk_write_time_seconds_total", device="disk0"}
Metric Name: node_disk_writes_completed_total, Metric Type: counter
Metric : node_disk_writes_completed_total{device="disk0"}, Metric Labels: {__name__="node_disk_writes_completed_total", device="disk0"}
Metric Name: node_disk_written_bytes_total, Metric Type: counter
Metric : node_disk_written_bytes_total{device="disk0"}, Metric Labels: {__name__="node_disk_written_bytes_total", device="disk0"}
Metric Name: node_disk_written_sectors_total, Metric Type: counter
Metric : node_disk_written_sectors_total{device="disk0"}, Metric Labels: {__name__="node_disk_written_sectors_total", device="disk0"}
Metric Name: node_exporter_build_info, Metric Type: gauge
Metric : node_exporter_build_info{branch="HEAD",goversion="go1.11.10",revision="f97f01c46cfde2ff97b5539b7964f3044c04947b",version="0.18.0"}, Metric Labels: {__name__="node_exporter_build_info", branch="HEAD", goversion="go1.11.10", revision="f97f01c46cfde2ff97b5539b7964f3044c04947b", version="0.18.0"}
Metric Name: node_filesystem_avail_bytes, Metric Type: gauge
Metric : node_filesystem_avail_bytes{device="/dev/disk1s1",fstype="apfs",mountpoint="/System/Volumes/Data"}, Metric Labels: {__name__="node_filesystem_avail_bytes", device="/dev/disk1s1", fstype="apfs", mountpoint="/System/Volumes/Data"}
Metric : node_filesystem_avail_bytes{device="/dev/disk1s4",fstype="apfs",mountpoint="/private/var/vm"}, Metric Labels: {__name__="node_filesystem_avail_bytes", device="/dev/disk1s4", fstype="apfs", mountpoint="/private/var/vm"}
Metric : node_filesystem_avail_bytes{device="/dev/disk1s5",fstype="apfs",mountpoint="/"}, Metric Labels: {__name__="node_filesystem_avail_bytes", device="/dev/disk1s5", fstype="apfs", mountpoint="/"}
Metric : node_filesystem_avail_bytes{device="map auto_home",fstype="autofs",mountpoint="/System/Volumes/Data/home"}, Metric Labels: {__name__="node_filesystem_avail_bytes", device="map auto_home", fstype="autofs", mountpoint="/System/Volumes/Data/home"}
Metric Name: node_filesystem_device_error, Metric Type: gauge
Metric : node_filesystem_device_error{device="/dev/disk1s1",fstype="apfs",mountpoint="/System/Volumes/Data"}, Metric Labels: {__name__="node_filesystem_device_error", device="/dev/disk1s1", fstype="apfs", mountpoint="/System/Volumes/Data"}
Metric : node_filesystem_device_error{device="/dev/disk1s4",fstype="apfs",mountpoint="/private/var/vm"}, Metric Labels: {__name__="node_filesystem_device_error", device="/dev/disk1s4", fstype="apfs", mountpoint="/private/var/vm"}
Metric : node_filesystem_device_error{device="/dev/disk1s5",fstype="apfs",mountpoint="/"}, Metric Labels: {__name__="node_filesystem_device_error", device="/dev/disk1s5", fstype="apfs", mountpoint="/"}
Metric : node_filesystem_device_error{device="map auto_home",fstype="autofs",mountpoint="/System/Volumes/Data/home"}, Metric Labels: {__name__="node_filesystem_device_error", device="map auto_home", fstype="autofs", mountpoint="/System/Volumes/Data/home"}
Metric Name: node_filesystem_files, Metric Type: gauge
Metric : node_filesystem_files{device="/dev/disk1s1",fstype="apfs",mountpoint="/System/Volumes/Data"}, Metric Labels: {__name__="node_filesystem_files", device="/dev/disk1s1", fstype="apfs", mountpoint="/System/Volumes/Data"}
Metric : node_filesystem_files{device="/dev/disk1s4",fstype="apfs",mountpoint="/private/var/vm"}, Metric Labels: {__name__="node_filesystem_files", device="/dev/disk1s4", fstype="apfs", mountpoint="/private/var/vm"}
Metric : node_filesystem_files{device="/dev/disk1s5",fstype="apfs",mountpoint="/"}, Metric Labels: {__name__="node_filesystem_files", device="/dev/disk1s5", fstype="apfs", mountpoint="/"}
Metric : node_filesystem_files{device="map auto_home",fstype="autofs",mountpoint="/System/Volumes/Data/home"}, Metric Labels: {__name__="node_filesystem_files", device="map auto_home", fstype="autofs", mountpoint="/System/Volumes/Data/home"}
Metric Name: node_filesystem_files_free, Metric Type: gauge
Metric : node_filesystem_files_free{device="/dev/disk1s1",fstype="apfs",mountpoint="/System/Volumes/Data"}, Metric Labels: {__name__="node_filesystem_files_free", device="/dev/disk1s1", fstype="apfs", mountpoint="/System/Volumes/Data"}
Metric : node_filesystem_files_free{device="/dev/disk1s4",fstype="apfs",mountpoint="/private/var/vm"}, Metric Labels: {__name__="node_filesystem_files_free", device="/dev/disk1s4", fstype="apfs", mountpoint="/private/var/vm"}
Metric : node_filesystem_files_free{device="/dev/disk1s5",fstype="apfs",mountpoint="/"}, Metric Labels: {__name__="node_filesystem_files_free", device="/dev/disk1s5", fstype="apfs", mountpoint="/"}
Metric : node_filesystem_files_free{device="map auto_home",fstype="autofs",mountpoint="/System/Volumes/Data/home"}, Metric Labels: {__name__="node_filesystem_files_free", device="map auto_home", fstype="autofs", mountpoint="/System/Volumes/Data/home"}
Metric Name: node_filesystem_free_bytes, Metric Type: gauge
Metric : node_filesystem_free_bytes{device="/dev/disk1s1",fstype="apfs",mountpoint="/System/Volumes/Data"}, Metric Labels: {__name__="node_filesystem_free_bytes", device="/dev/disk1s1", fstype="apfs", mountpoint="/System/Volumes/Data"}
Metric : node_filesystem_free_bytes{device="/dev/disk1s4",fstype="apfs",mountpoint="/private/var/vm"}, Metric Labels: {__name__="node_filesystem_free_bytes", device="/dev/disk1s4", fstype="apfs", mountpoint="/private/var/vm"}
Metric : node_filesystem_free_bytes{device="/dev/disk1s5",fstype="apfs",mountpoint="/"}, Metric Labels: {__name__="node_filesystem_free_bytes", device="/dev/disk1s5", fstype="apfs", mountpoint="/"}
Metric : node_filesystem_free_bytes{device="map auto_home",fstype="autofs",mountpoint="/System/Volumes/Data/home"}, Metric Labels: {__name__="node_filesystem_free_bytes", device="map auto_home", fstype="autofs", mountpoint="/System/Volumes/Data/home"}
Metric Name: node_filesystem_readonly, Metric Type: gauge
Metric : node_filesystem_readonly{device="/dev/disk1s1",fstype="apfs",mountpoint="/System/Volumes/Data"}, Metric Labels: {__name__="node_filesystem_readonly", device="/dev/disk1s1", fstype="apfs", mountpoint="/System/Volumes/Data"}
Metric : node_filesystem_readonly{device="/dev/disk1s4",fstype="apfs",mountpoint="/private/var/vm"}, Metric Labels: {__name__="node_filesystem_readonly", device="/dev/disk1s4", fstype="apfs", mountpoint="/private/var/vm"}
Metric : node_filesystem_readonly{device="/dev/disk1s5",fstype="apfs",mountpoint="/"}, Metric Labels: {__name__="node_filesystem_readonly", device="/dev/disk1s5", fstype="apfs", mountpoint="/"}
Metric : node_filesystem_readonly{device="map auto_home",fstype="autofs",mountpoint="/System/Volumes/Data/home"}, Metric Labels: {__name__="node_filesystem_readonly", device="map auto_home", fstype="autofs", mountpoint="/System/Volumes/Data/home"}
Metric Name: node_filesystem_size_bytes, Metric Type: gauge
Metric : node_filesystem_size_bytes{device="/dev/disk1s1",fstype="apfs",mountpoint="/System/Volumes/Data"}, Metric Labels: {__name__="node_filesystem_size_bytes", device="/dev/disk1s1", fstype="apfs", mountpoint="/System/Volumes/Data"}
Metric : node_filesystem_size_bytes{device="/dev/disk1s4",fstype="apfs",mountpoint="/private/var/vm"}, Metric Labels: {__name__="node_filesystem_size_bytes", device="/dev/disk1s4", fstype="apfs", mountpoint="/private/var/vm"}
Metric : node_filesystem_size_bytes{device="/dev/disk1s5",fstype="apfs",mountpoint="/"}, Metric Labels: {__name__="node_filesystem_size_bytes", device="/dev/disk1s5", fstype="apfs", mountpoint="/"}
Metric : node_filesystem_size_bytes{device="map auto_home",fstype="autofs",mountpoint="/System/Volumes/Data/home"}, Metric Labels: {__name__="node_filesystem_size_bytes", device="map auto_home", fstype="autofs", mountpoint="/System/Volumes/Data/home"}
Metric Name: node_load1, Metric Type: gauge
Metric : node_load1, Metric Labels: {__name__="node_load1"}
Metric Name: node_load15, Metric Type: gauge
Metric : node_load15, Metric Labels: {__name__="node_load15"}
Metric Name: node_load5, Metric Type: gauge
Metric : node_load5, Metric Labels: {__name__="node_load5"}
Metric Name: node_memory_active_bytes, Metric Type: gauge
Metric : node_memory_active_bytes, Metric Labels: {__name__="node_memory_active_bytes"}
Metric Name: node_memory_compressed_bytes, Metric Type: gauge
Metric : node_memory_compressed_bytes, Metric Labels: {__name__="node_memory_compressed_bytes"}
Metric Name: node_memory_free_bytes, Metric Type: gauge
Metric : node_memory_free_bytes, Metric Labels: {__name__="node_memory_free_bytes"}
Metric Name: node_memory_inactive_bytes, Metric Type: gauge
Metric : node_memory_inactive_bytes, Metric Labels: {__name__="node_memory_inactive_bytes"}
Metric Name: node_memory_swapped_in_bytes_total, Metric Type: counter
Metric : node_memory_swapped_in_bytes_total, Metric Labels: {__name__="node_memory_swapped_in_bytes_total"}
Metric Name: node_memory_swapped_out_bytes_total, Metric Type: counter
Metric : node_memory_swapped_out_bytes_total, Metric Labels: {__name__="node_memory_swapped_out_bytes_total"}
Metric Name: node_memory_total_bytes, Metric Type: gauge
Metric : node_memory_total_bytes, Metric Labels: {__name__="node_memory_total_bytes"}
Metric Name: node_memory_wired_bytes, Metric Type: gauge
Metric : node_memory_wired_bytes, Metric Labels: {__name__="node_memory_wired_bytes"}
Metric Name: node_network_receive_bytes_total, Metric Type: counter
Metric : node_network_receive_bytes_total{device="awdl0"}, Metric Labels: {__name__="node_network_receive_bytes_total", device="awdl0"}
Metric : node_network_receive_bytes_total{device="bridge0"}, Metric Labels: {__name__="node_network_receive_bytes_total", device="bridge0"}
Metric : node_network_receive_bytes_total{device="en0"}, Metric Labels: {__name__="node_network_receive_bytes_total", device="en0"}
Metric : node_network_receive_bytes_total{device="en1"}, Metric Labels: {__name__="node_network_receive_bytes_total", device="en1"}
Metric : node_network_receive_bytes_total{device="en2"}, Metric Labels: {__name__="node_network_receive_bytes_total", device="en2"}
Metric : node_network_receive_bytes_total{device="gif0"}, Metric Labels: {__name__="node_network_receive_bytes_total", device="gif0"}
Metric : node_network_receive_bytes_total{device="llw0"}, Metric Labels: {__name__="node_network_receive_bytes_total", device="llw0"}
Metric : node_network_receive_bytes_total{device="lo0"}, Metric Labels: {__name__="node_network_receive_bytes_total", device="lo0"}
Metric : node_network_receive_bytes_total{device="p2p0"}, Metric Labels: {__name__="node_network_receive_bytes_total", device="p2p0"}
Metric : node_network_receive_bytes_total{device="stf0"}, Metric Labels: {__name__="node_network_receive_bytes_total", device="stf0"}
Metric : node_network_receive_bytes_total{device="utun0"}, Metric Labels: {__name__="node_network_receive_bytes_total", device="utun0"}
Metric : node_network_receive_bytes_total{device="utun1"}, Metric Labels: {__name__="node_network_receive_bytes_total", device="utun1"}
Metric Name: node_network_receive_errs_total, Metric Type: counter
Metric : node_network_receive_errs_total{device="awdl0"}, Metric Labels: {__name__="node_network_receive_errs_total", device="awdl0"}
Metric : node_network_receive_errs_total{device="bridge0"}, Metric Labels: {__name__="node_network_receive_errs_total", device="bridge0"}
Metric : node_network_receive_errs_total{device="en0"}, Metric Labels: {__name__="node_network_receive_errs_total", device="en0"}
Metric : node_network_receive_errs_total{device="en1"}, Metric Labels: {__name__="node_network_receive_errs_total", device="en1"}
Metric : node_network_receive_errs_total{device="en2"}, Metric Labels: {__name__="node_network_receive_errs_total", device="en2"}
Metric : node_network_receive_errs_total{device="gif0"}, Metric Labels: {__name__="node_network_receive_errs_total", device="gif0"}
Metric : node_network_receive_errs_total{device="llw0"}, Metric Labels: {__name__="node_network_receive_errs_total", device="llw0"}
Metric : node_network_receive_errs_total{device="lo0"}, Metric Labels: {__name__="node_network_receive_errs_total", device="lo0"}
Metric : node_network_receive_errs_total{device="p2p0"}, Metric Labels: {__name__="node_network_receive_errs_total", device="p2p0"}
Metric : node_network_receive_errs_total{device="stf0"}, Metric Labels: {__name__="node_network_receive_errs_total", device="stf0"}
Metric : node_network_receive_errs_total{device="utun0"}, Metric Labels: {__name__="node_network_receive_errs_total", device="utun0"}
Metric : node_network_receive_errs_total{device="utun1"}, Metric Labels: {__name__="node_network_receive_errs_total", device="utun1"}
Metric Name: node_network_receive_multicast_total, Metric Type: counter
Metric : node_network_receive_multicast_total{device="awdl0"}, Metric Labels: {__name__="node_network_receive_multicast_total", device="awdl0"}
Metric : node_network_receive_multicast_total{device="bridge0"}, Metric Labels: {__name__="node_network_receive_multicast_total", device="bridge0"}
Metric : node_network_receive_multicast_total{device="en0"}, Metric Labels: {__name__="node_network_receive_multicast_total", device="en0"}
Metric : node_network_receive_multicast_total{device="en1"}, Metric Labels: {__name__="node_network_receive_multicast_total", device="en1"}
Metric : node_network_receive_multicast_total{device="en2"}, Metric Labels: {__name__="node_network_receive_multicast_total", device="en2"}
Metric : node_network_receive_multicast_total{device="gif0"}, Metric Labels: {__name__="node_network_receive_multicast_total", device="gif0"}
Metric : node_network_receive_multicast_total{device="llw0"}, Metric Labels: {__name__="node_network_receive_multicast_total", device="llw0"}
Metric : node_network_receive_multicast_total{device="lo0"}, Metric Labels: {__name__="node_network_receive_multicast_total", device="lo0"}
Metric : node_network_receive_multicast_total{device="p2p0"}, Metric Labels: {__name__="node_network_receive_multicast_total", device="p2p0"}
Metric : node_network_receive_multicast_total{device="stf0"}, Metric Labels: {__name__="node_network_receive_multicast_total", device="stf0"}
Metric : node_network_receive_multicast_total{device="utun0"}, Metric Labels: {__name__="node_network_receive_multicast_total", device="utun0"}
Metric : node_network_receive_multicast_total{device="utun1"}, Metric Labels: {__name__="node_network_receive_multicast_total", device="utun1"}
Metric Name: node_network_receive_packets_total, Metric Type: counter
Metric : node_network_receive_packets_total{device="awdl0"}, Metric Labels: {__name__="node_network_receive_packets_total", device="awdl0"}
Metric : node_network_receive_packets_total{device="bridge0"}, Metric Labels: {__name__="node_network_receive_packets_total", device="bridge0"}
Metric : node_network_receive_packets_total{device="en0"}, Metric Labels: {__name__="node_network_receive_packets_total", device="en0"}
Metric : node_network_receive_packets_total{device="en1"}, Metric Labels: {__name__="node_network_receive_packets_total", device="en1"}
Metric : node_network_receive_packets_total{device="en2"}, Metric Labels: {__name__="node_network_receive_packets_total", device="en2"}
Metric : node_network_receive_packets_total{device="gif0"}, Metric Labels: {__name__="node_network_receive_packets_total", device="gif0"}
Metric : node_network_receive_packets_total{device="llw0"}, Metric Labels: {__name__="node_network_receive_packets_total", device="llw0"}
Metric : node_network_receive_packets_total{device="lo0"}, Metric Labels: {__name__="node_network_receive_packets_total", device="lo0"}
Metric : node_network_receive_packets_total{device="p2p0"}, Metric Labels: {__name__="node_network_receive_packets_total", device="p2p0"}
Metric : node_network_receive_packets_total{device="stf0"}, Metric Labels: {__name__="node_network_receive_packets_total", device="stf0"}
Metric : node_network_receive_packets_total{device="utun0"}, Metric Labels: {__name__="node_network_receive_packets_total", device="utun0"}
Metric : node_network_receive_packets_total{device="utun1"}, Metric Labels: {__name__="node_network_receive_packets_total", device="utun1"}
Metric Name: node_network_transmit_bytes_total, Metric Type: counter
Metric : node_network_transmit_bytes_total{device="awdl0"}, Metric Labels: {__name__="node_network_transmit_bytes_total", device="awdl0"}
Metric : node_network_transmit_bytes_total{device="bridge0"}, Metric Labels: {__name__="node_network_transmit_bytes_total", device="bridge0"}
Metric : node_network_transmit_bytes_total{device="en0"}, Metric Labels: {__name__="node_network_transmit_bytes_total", device="en0"}
Metric : node_network_transmit_bytes_total{device="en1"}, Metric Labels: {__name__="node_network_transmit_bytes_total", device="en1"}
Metric : node_network_transmit_bytes_total{device="en2"}, Metric Labels: {__name__="node_network_transmit_bytes_total", device="en2"}
Metric : node_network_transmit_bytes_total{device="gif0"}, Metric Labels: {__name__="node_network_transmit_bytes_total", device="gif0"}
Metric : node_network_transmit_bytes_total{device="llw0"}, Metric Labels: {__name__="node_network_transmit_bytes_total", device="llw0"}
Metric : node_network_transmit_bytes_total{device="lo0"}, Metric Labels: {__name__="node_network_transmit_bytes_total", device="lo0"}
Metric : node_network_transmit_bytes_total{device="p2p0"}, Metric Labels: {__name__="node_network_transmit_bytes_total", device="p2p0"}
Metric : node_network_transmit_bytes_total{device="stf0"}, Metric Labels: {__name__="node_network_transmit_bytes_total", device="stf0"}
Metric : node_network_transmit_bytes_total{device="utun0"}, Metric Labels: {__name__="node_network_transmit_bytes_total", device="utun0"}
Metric : node_network_transmit_bytes_total{device="utun1"}, Metric Labels: {__name__="node_network_transmit_bytes_total", device="utun1"}
Metric Name: node_network_transmit_errs_total, Metric Type: counter
Metric : node_network_transmit_errs_total{device="awdl0"}, Metric Labels: {__name__="node_network_transmit_errs_total", device="awdl0"}
Metric : node_network_transmit_errs_total{device="bridge0"}, Metric Labels: {__name__="node_network_transmit_errs_total", device="bridge0"}
Metric : node_network_transmit_errs_total{device="en0"}, Metric Labels: {__name__="node_network_transmit_errs_total", device="en0"}
Metric : node_network_transmit_errs_total{device="en1"}, Metric Labels: {__name__="node_network_transmit_errs_total", device="en1"}
Metric : node_network_transmit_errs_total{device="en2"}, Metric Labels: {__name__="node_network_transmit_errs_total", device="en2"}
Metric : node_network_transmit_errs_total{device="gif0"}, Metric Labels: {__name__="node_network_transmit_errs_total", device="gif0"}
Metric : node_network_transmit_errs_total{device="llw0"}, Metric Labels: {__name__="node_network_transmit_errs_total", device="llw0"}
Metric : node_network_transmit_errs_total{device="lo0"}, Metric Labels: {__name__="node_network_transmit_errs_total", device="lo0"}
Metric : node_network_transmit_errs_total{device="p2p0"}, Metric Labels: {__name__="node_network_transmit_errs_total", device="p2p0"}
Metric : node_network_transmit_errs_total{device="stf0"}, Metric Labels: {__name__="node_network_transmit_errs_total", device="stf0"}
Metric : node_network_transmit_errs_total{device="utun0"}, Metric Labels: {__name__="node_network_transmit_errs_total", device="utun0"}
Metric : node_network_transmit_errs_total{device="utun1"}, Metric Labels: {__name__="node_network_transmit_errs_total", device="utun1"}
Metric Name: node_network_transmit_multicast_total, Metric Type: counter
Metric : node_network_transmit_multicast_total{device="awdl0"}, Metric Labels: {__name__="node_network_transmit_multicast_total", device="awdl0"}
Metric : node_network_transmit_multicast_total{device="bridge0"}, Metric Labels: {__name__="node_network_transmit_multicast_total", device="bridge0"}
Metric : node_network_transmit_multicast_total{device="en0"}, Metric Labels: {__name__="node_network_transmit_multicast_total", device="en0"}
Metric : node_network_transmit_multicast_total{device="en1"}, Metric Labels: {__name__="node_network_transmit_multicast_total", device="en1"}
Metric : node_network_transmit_multicast_total{device="en2"}, Metric Labels: {__name__="node_network_transmit_multicast_total", device="en2"}
Metric : node_network_transmit_multicast_total{device="gif0"}, Metric Labels: {__name__="node_network_transmit_multicast_total", device="gif0"}
Metric : node_network_transmit_multicast_total{device="llw0"}, Metric Labels: {__name__="node_network_transmit_multicast_total", device="llw0"}
Metric : node_network_transmit_multicast_total{device="lo0"}, Metric Labels: {__name__="node_network_transmit_multicast_total", device="lo0"}
Metric : node_network_transmit_multicast_total{device="p2p0"}, Metric Labels: {__name__="node_network_transmit_multicast_total", device="p2p0"}
Metric : node_network_transmit_multicast_total{device="stf0"}, Metric Labels: {__name__="node_network_transmit_multicast_total", device="stf0"}
Metric : node_network_transmit_multicast_total{device="utun0"}, Metric Labels: {__name__="node_network_transmit_multicast_total", device="utun0"}
Metric : node_network_transmit_multicast_total{device="utun1"}, Metric Labels: {__name__="node_network_transmit_multicast_total", device="utun1"}
Metric Name: node_network_transmit_packets_total, Metric Type: counter
Metric : node_network_transmit_packets_total{device="awdl0"}, Metric Labels: {__name__="node_network_transmit_packets_total", device="awdl0"}
Metric : node_network_transmit_packets_total{device="bridge0"}, Metric Labels: {__name__="node_network_transmit_packets_total", device="bridge0"}
Metric : node_network_transmit_packets_total{device="en0"}, Metric Labels: {__name__="node_network_transmit_packets_total", device="en0"}
Metric : node_network_transmit_packets_total{device="en1"}, Metric Labels: {__name__="node_network_transmit_packets_total", device="en1"}
Metric : node_network_transmit_packets_total{device="en2"}, Metric Labels: {__name__="node_network_transmit_packets_total", device="en2"}
Metric : node_network_transmit_packets_total{device="gif0"}, Metric Labels: {__name__="node_network_transmit_packets_total", device="gif0"}
Metric : node_network_transmit_packets_total{device="llw0"}, Metric Labels: {__name__="node_network_transmit_packets_total", device="llw0"}
Metric : node_network_transmit_packets_total{device="lo0"}, Metric Labels: {__name__="node_network_transmit_packets_total", device="lo0"}
Metric : node_network_transmit_packets_total{device="p2p0"}, Metric Labels: {__name__="node_network_transmit_packets_total", device="p2p0"}
Metric : node_network_transmit_packets_total{device="stf0"}, Metric Labels: {__name__="node_network_transmit_packets_total", device="stf0"}
Metric : node_network_transmit_packets_total{device="utun0"}, Metric Labels: {__name__="node_network_transmit_packets_total", device="utun0"}
Metric : node_network_transmit_packets_total{device="utun1"}, Metric Labels: {__name__="node_network_transmit_packets_total", device="utun1"}
Metric Name: node_scrape_collector_duration_seconds, Metric Type: gauge
Metric : node_scrape_collector_duration_seconds{collector="boottime"}, Metric Labels: {__name__="node_scrape_collector_duration_seconds", collector="boottime"}
Metric : node_scrape_collector_duration_seconds{collector="cpu"}, Metric Labels: {__name__="node_scrape_collector_duration_seconds", collector="cpu"}
Metric : node_scrape_collector_duration_seconds{collector="diskstats"}, Metric Labels: {__name__="node_scrape_collector_duration_seconds", collector="diskstats"}
Metric : node_scrape_collector_duration_seconds{collector="filesystem"}, Metric Labels: {__name__="node_scrape_collector_duration_seconds", collector="filesystem"}
Metric : node_scrape_collector_duration_seconds{collector="loadavg"}, Metric Labels: {__name__="node_scrape_collector_duration_seconds", collector="loadavg"}
Metric : node_scrape_collector_duration_seconds{collector="meminfo"}, Metric Labels: {__name__="node_scrape_collector_duration_seconds", collector="meminfo"}
Metric : node_scrape_collector_duration_seconds{collector="netdev"}, Metric Labels: {__name__="node_scrape_collector_duration_seconds", collector="netdev"}
Metric : node_scrape_collector_duration_seconds{collector="textfile"}, Metric Labels: {__name__="node_scrape_collector_duration_seconds", collector="textfile"}
Metric : node_scrape_collector_duration_seconds{collector="time"}, Metric Labels: {__name__="node_scrape_collector_duration_seconds", collector="time"}
Metric Name: node_scrape_collector_success, Metric Type: gauge
Metric : node_scrape_collector_success{collector="boottime"}, Metric Labels: {__name__="node_scrape_collector_success", collector="boottime"}
Metric : node_scrape_collector_success{collector="cpu"}, Metric Labels: {__name__="node_scrape_collector_success", collector="cpu"}
Metric : node_scrape_collector_success{collector="diskstats"}, Metric Labels: {__name__="node_scrape_collector_success", collector="diskstats"}
Metric : node_scrape_collector_success{collector="filesystem"}, Metric Labels: {__name__="node_scrape_collector_success", collector="filesystem"}
Metric : node_scrape_collector_success{collector="loadavg"}, Metric Labels: {__name__="node_scrape_collector_success", collector="loadavg"}
Metric : node_scrape_collector_success{collector="meminfo"}, Metric Labels: {__name__="node_scrape_collector_success", collector="meminfo"}
Metric : node_scrape_collector_success{collector="netdev"}, Metric Labels: {__name__="node_scrape_collector_success", collector="netdev"}
Metric : node_scrape_collector_success{collector="textfile"}, Metric Labels: {__name__="node_scrape_collector_success", collector="textfile"}
Metric : node_scrape_collector_success{collector="time"}, Metric Labels: {__name__="node_scrape_collector_success", collector="time"}
Metric Name: node_textfile_scrape_error, Metric Type: gauge
Metric : node_textfile_scrape_error, Metric Labels: {__name__="node_textfile_scrape_error"}
Metric Name: node_time_seconds, Metric Type: gauge
Metric : node_time_seconds, Metric Labels: {__name__="node_time_seconds"}
Metric Name: promhttp_metric_handler_requests_in_flight, Metric Type: gauge
Metric : promhttp_metric_handler_requests_in_flight, Metric Labels: {__name__="promhttp_metric_handler_requests_in_flight"}
Metric Name: promhttp_metric_handler_requests_total, Metric Type: counter
Metric : promhttp_metric_handler_requests_total{code="200"}, Metric Labels: {__name__="promhttp_metric_handler_requests_total", code="200"}
Metric : promhttp_metric_handler_requests_total{code="500"}, Metric Labels: {__name__="promhttp_metric_handler_requests_total", code="500"}
Metric : promhttp_metric_handler_requests_total{code="503"}, Metric Labels: {__name__="promhttp_metric_handler_requests_total", code="503"}
```

I also wrote code to get the unit of the metric, but in the input that I had,
I guess there's no metric, and hence no output regarding units. Anyways, that
doesn't matter I guess. I'm just going to see how to store only the metric
name, metric type and then all the metric label names, excluding `__name__`,
which common for all and stores the metric name in the label value, and I don't
know why 🙈


