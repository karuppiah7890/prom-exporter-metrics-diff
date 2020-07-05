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

