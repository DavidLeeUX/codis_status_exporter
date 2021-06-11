package exporter

import (
	"strconv"

	"all_kinds_monitoring/codis_status_exporter/codisstat"

	"github.com/prometheus/client_golang/prometheus"
)

// 初始化group监控信息
func (e *Exporter) InitGroupMetric() {
	var (
		GroupMetricLabelNames = []string{
			"addr",
			"product",
			"redis_version",
			"redis_build_id",
			"serveraddr",
			"role",
		}
	)
	e.GroupMetric = map[string]*prometheus.GaugeVec{}
	e.GroupMetric["group_aof_current_rewrite_time_sec"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_aof_current_rewrite_time_sec",
		Help:      "group_aof_current_rewrite_time_sec",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_aof_enabled"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_aof_enabled",
		Help:      "group_aof_enabled",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_aof_last_bgrewrite_status"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_aof_last_bgrewrite_status",
		Help:      "group_aof_last_bgrewrite_status 1=OK , -1=ERROR",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_aof_last_rewrite_time_sec"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_aof_last_rewrite_time_sec",
		Help:      "group_aof_last_rewrite_time_sec ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_aof_last_write_status"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_aof_last_write_status",
		Help:      "group_aof_last_write_status ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_aof_rewrite_in_progress"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_aof_rewrite_in_progress",
		Help:      "group_aof_rewrite_in_progress ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_aof_rewrite_scheduled"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_aof_rewrite_scheduled",
		Help:      "group_aof_rewrite_scheduled ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_blocked_clients"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_blocked_clients",
		Help:      "group_blocked_clients ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_client_biggest_input_buf"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_client_biggest_input_buf",
		Help:      "group_client_biggest_input_buf ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_client_longest_output_list"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_client_longest_output_list",
		Help:      "group_client_longest_output_list ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_connected_clients"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_connected_clients",
		Help:      "group_connected_clients ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_connected_slaves"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_connected_slaves",
		Help:      "group_connected_slaves ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_evicted_keys"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_evicted_keys",
		Help:      "group_evicted_keys ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_expired_keys"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_expired_keys",
		Help:      "group_expired_keys ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_instantaneous_input_kbps"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_instantaneous_input_kbps",
		Help:      "group_instantaneous_input_kbps ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_instantaneous_ops_per_sec"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_instantaneous_ops_per_sec",
		Help:      "group_instantaneous_ops_per_sec ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_instantaneous_output_kbps"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_instantaneous_output_kbps",
		Help:      "group_instantaneous_output_kbps ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_keyspace_hits"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_keyspace_hits",
		Help:      "group_keyspace_hits ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_keyspace_misses"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_keyspace_misses",
		Help:      "group_keyspace_misses ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_latest_fork_usec"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_latest_fork_usec",
		Help:      "group_latest_fork_usec ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_loading"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_loading",
		Help:      "group_loading ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_lru_clock"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_lru_clock",
		Help:      "group_lru_clock ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_master_repl_offset"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_master_repl_offset",
		Help:      "group_master_repl_offset ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_maxmemory"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_maxmemory",
		Help:      "group_maxmemory ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_mem_fragmentation_ratio"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_mem_fragmentation_ratio",
		Help:      "group_mem_fragmentation_ratio ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_migrate_cached_sockets"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_migrate_cached_sockets",
		Help:      "group_migrate_cached_sockets ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_pubsub_channels"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_pubsub_channels",
		Help:      "group_pubsub_channels ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_pubsub_patterns"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_pubsub_patterns",
		Help:      "group_pubsub_patterns ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_rdb_bgsave_in_progress"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_rdb_bgsave_in_progress",
		Help:      "group_rdb_bgsave_in_progress ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_rdb_changes_since_last_save"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_rdb_changes_since_last_save",
		Help:      "group_rdb_changes_since_last_save ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_rdb_current_bgsave_time_sec"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_rdb_current_bgsave_time_sec",
		Help:      "group_rdb_current_bgsave_time_sec ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_rdb_last_bgsave_status"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_rdb_last_bgsave_status",
		Help:      "group_rdb_last_bgsave_status ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_rdb_last_bgsave_time_sec"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_rdb_last_bgsave_time_sec",
		Help:      "group_rdb_last_bgsave_time_sec ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_rdb_last_save_time"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_rdb_last_save_time",
		Help:      "group_rdb_last_save_time ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_rejected_connections"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_rejected_connections",
		Help:      "group_rejected_connections ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_repl_backlog_active"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_repl_backlog_active",
		Help:      "group_repl_backlog_active ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_repl_backlog_first_byte_offset"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_repl_backlog_first_byte_offset",
		Help:      "group_repl_backlog_first_byte_offset ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_repl_backlog_histlen"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_repl_backlog_histlen",
		Help:      "group_repl_backlog_histlen ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_repl_backlog_size"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_repl_backlog_size",
		Help:      "group_repl_backlog_size ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_sync_full"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_sync_full",
		Help:      "group_sync_full ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_sync_partial_err"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_sync_partial_err",
		Help:      "group_sync_partial_err ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_sync_partial_ok"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_sync_partial_ok",
		Help:      "group_sync_partial_ok ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_total_commands_processed"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_total_commands_processed",
		Help:      "group_total_commands_processed ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_total_connections_received"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_total_connections_received",
		Help:      "group_total_connections_received ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_total_net_input_bytes"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_total_net_input_bytes",
		Help:      "group_total_net_input_bytes ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_total_net_output_bytes"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_total_net_output_bytes",
		Help:      "group_total_net_output_bytes ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_total_system_memory"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_total_system_memory",
		Help:      "group_total_system_memory ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_uptime_in_seconds"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_uptime_in_seconds",
		Help:      "group_uptime_in_seconds ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_used_cpu_sys"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_used_cpu_sys",
		Help:      "group_used_cpu_sys ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_used_cpu_sys_children"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_used_cpu_sys_children",
		Help:      "group_used_cpu_sys_children ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_used_cpu_user"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_used_cpu_user",
		Help:      "group_used_cpu_user ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_used_cpu_user_children"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_used_cpu_user_children",
		Help:      "group_used_cpu_user_children ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_used_memory"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_used_memory",
		Help:      "group_used_memory ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_used_memory_lua"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_used_memory_lua",
		Help:      "group_used_memory_lua ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_used_memory_peak"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_used_memory_peak",
		Help:      "group_used_memory_peak ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_used_memory_rss"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_used_memory_rss",
		Help:      "group_used_memory_rss ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_db0_keys"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_db0_keys",
		Help:      "group_db0_keys",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_db0_expires"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_db0_expires",
		Help:      "group_db0_expires ",
	}, GroupMetricLabelNames)

	e.GroupMetric["group_db0_avg_ttl"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "group_db0_avg_ttl",
		Help:      "group_db0_avg_ttl ",
	}, GroupMetricLabelNames)
}

func (e *Exporter) GetGroupMetric(ch chan<- prometheus.Metric) {
	GroupMetricLabel := make([]string, 6)
	GroupMetricLabel[0] = e.Addr
	GroupMetricLabel[1] = e.ProductName

	for k, v := range e.CodisData.Stats.Group.Stats {
		GroupMetricLabel[2] = v.Stats.Redis_version
		GroupMetricLabel[3] = v.Stats.Redis_build_id
		// redis group address+port
		GroupMetricLabel[4] = k
		GroupMetricLabel[5] = v.Stats.Role
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Aof_current_rewrite_time_sec, "group_aof_current_rewrite_time_sec")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Aof_enabled, "group_aof_enabled")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Aof_last_bgrewrite_status, "group_aof_last_bgrewrite_status")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Aof_last_rewrite_time_sec, "group_aof_last_rewrite_time_sec")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Aof_last_write_status, "group_aof_last_write_status")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Aof_rewrite_in_progress, "group_aof_rewrite_in_progress")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Aof_rewrite_scheduled, "group_aof_rewrite_scheduled")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Blocked_clients, "group_blocked_clients")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Client_biggest_input_buf, "group_client_biggest_input_buf")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Client_longest_output_list, "group_client_longest_output_list")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Connected_clients, "group_connected_clients")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Connected_slaves, "group_connected_slaves")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Evicted_keys, "group_evicted_keys")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Expired_keys, "group_expired_keys")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Instantaneous_input_kbps, "group_instantaneous_input_kbps")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Instantaneous_ops_per_sec, "group_instantaneous_ops_per_sec")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Instantaneous_output_kbps, "group_instantaneous_output_kbps")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Keyspace_hits, "group_keyspace_hits")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Keyspace_misses, "group_keyspace_misses")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Latest_fork_usec, "group_latest_fork_usec")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Loading, "group_loading")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Lru_clock, "group_lru_clock")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Master_repl_offset, "group_master_repl_offset")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Maxmemory, "group_maxmemory")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Mem_fragmentation_ratio, "group_mem_fragmentation_ratio")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Migrate_cached_sockets, "group_migrate_cached_sockets")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Pubsub_channels, "group_pubsub_channels")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Pubsub_patterns, "group_pubsub_patterns")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Rdb_bgsave_in_progress, "group_rdb_bgsave_in_progress")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Rdb_changes_since_last_save, "group_rdb_changes_since_last_save")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Rdb_current_bgsave_time_sec, "group_rdb_current_bgsave_time_sec")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Rdb_last_bgsave_status, "group_rdb_last_bgsave_status")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Rdb_last_bgsave_time_sec, "group_rdb_last_bgsave_time_sec")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Rdb_last_save_time, "group_rdb_last_save_time")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Rejected_connections, "group_rejected_connections")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Repl_backlog_active, "group_repl_backlog_active")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Repl_backlog_first_byte_offset, "group_repl_backlog_first_byte_offset")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Repl_backlog_histlen, "group_repl_backlog_histlen")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Repl_backlog_size, "group_repl_backlog_size")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Sync_full, "group_sync_full")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Sync_partial_err, "group_sync_partial_err")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Sync_partial_ok, "group_sync_partial_ok")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Total_commands_processed, "group_total_commands_processed")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Total_connections_received, "group_total_connections_received")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Total_net_input_bytes, "group_total_net_input_bytes")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Total_net_output_bytes, "group_total_net_output_bytes")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Total_system_memory, "group_total_system_memory")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Uptime_in_seconds, "group_uptime_in_seconds")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Used_cpu_sys, "group_used_cpu_sys")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Used_cpu_sys_children, "group_used_cpu_sys_children")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Used_cpu_user, "group_used_cpu_user")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Used_cpu_user_children, "group_used_cpu_user_children")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Used_memory, "group_used_memory")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Used_memory_lua, "group_used_memory_lua")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Used_memory_peak, "group_used_memory_peak")
		e.GroupProcessValues(GroupMetricLabel, v.Stats.Used_memory_rss, "group_used_memory_rss")
		DbData, err := codisstat.GetDb0Data(v.Stats.Db0)
		if err == nil {
			e.GroupMetricsMtx.Lock()
			e.GroupMetric["group_db0_keys"].WithLabelValues(GroupMetricLabel...).Set(float64(DbData.Keys))
			e.GroupMetric["group_db0_expires"].WithLabelValues(GroupMetricLabel...).Set(float64(DbData.Expires))
			e.GroupMetric["group_db0_avg_ttl"].WithLabelValues(GroupMetricLabel...).Set(float64(DbData.Avg_ttl))
			e.GroupMetricsMtx.Unlock()
		}

	}

	for _, vm := range e.GroupMetric {
		vm.Collect(ch)
	}
}
func (e *Exporter) GroupProcessValues(GroupMetricLabel []string, value string, mapKey string) {
	e.GroupMetricsMtx.Lock()
	if value == "ok" {
		e.GroupMetric[mapKey].WithLabelValues(GroupMetricLabel...).Set(codisstat.OK)
	} else {
		if val, err := strconv.ParseFloat(value, 64); err == nil {
			e.GroupMetric[mapKey].WithLabelValues(GroupMetricLabel...).Set(val)
		} else {
			e.GroupMetric[mapKey].WithLabelValues(GroupMetricLabel...).Set(codisstat.ERROR)
		}
	}
	e.GroupMetricsMtx.Unlock()

}
