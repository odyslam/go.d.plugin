package mysql

import "github.com/netdata/go.d.plugin/modules"

type (
	// Charts is an alias for modules.Charts
	Charts = modules.Charts
	// Dims is an alias for modules.Dims
	Dims = modules.Dims
)

var charts = Charts{
	{
		ID:    "net",
		Title: "Bandwidth",
		Units: "kilobits/s",
		Fam:   "bandwidth",
		Ctx:   "mysql.net",
		Type:  modules.Area,
		Dims: Dims{
			{ID: "bytes_received", Name: "in", Algo: modules.Incremental, Mul: 8, Div: 1000},
			{ID: "bytes_sent", Name: "out", Algo: modules.Incremental, Mul: -8, Div: 1000},
		},
	},
	{
		ID:    "queries",
		Title: "Queries",
		Units: "queries/s",
		Fam:   "queries",
		Ctx:   "mysql.queries",
		Dims: Dims{
			{ID: "queries", Name: "queries", Algo: modules.Incremental},
			{ID: "questions", Name: "questions", Algo: modules.Incremental},
			{ID: "slow_queries", Name: "slow queries", Algo: modules.Incremental},
		},
	},
	{
		ID:    "queries_type",
		Title: "Queries By Type",
		Units: "queries/s",
		Fam:   "queries",
		Ctx:   "mysql.queries_type",
		Type:  modules.Stacked,
		Dims: Dims{
			{ID: "com_select", Name: "select", Algo: modules.Incremental},
			{ID: "com_delete", Name: "delete", Algo: modules.Incremental},
			{ID: "com_update", Name: "update", Algo: modules.Incremental},
			{ID: "com_insert", Name: "insert", Algo: modules.Incremental},
			{ID: "qcache_hits", Name: "cache hits", Algo: modules.Incremental},
			{ID: "com_replace", Name: "replace", Algo: modules.Incremental},
		},
	},
	{
		ID:    "handlers",
		Title: "Handlers",
		Units: "handlers/s",
		Fam:   "handlers",
		Ctx:   "mysql.handlers",
		Dims: Dims{
			{ID: "handler_commit", Name: "commit", Algo: modules.Incremental},
			{ID: "handler_delete", Name: "delete", Algo: modules.Incremental},
			{ID: "handler_prepare", Name: "prepare", Algo: modules.Incremental},
			{ID: "handler_read_first", Name: "read first", Algo: modules.Incremental},
			{ID: "handler_read_key", Name: "read key", Algo: modules.Incremental},
			{ID: "handler_read_next", Name: "read next", Algo: modules.Incremental},
			{ID: "handler_read_prev", Name: "read prev", Algo: modules.Incremental},
			{ID: "handler_read_rnd", Name: "read rnd", Algo: modules.Incremental},
			{ID: "handler_read_rnd_next", Name: "read rnd next", Algo: modules.Incremental},
			{ID: "handler_rollback", Name: "rollback", Algo: modules.Incremental},
			{ID: "handler_savepoint", Name: "savepoint", Algo: modules.Incremental},
			{ID: "handler_savepoint_rollback", Name: "savepointrollback", Algo: modules.Incremental},
			{ID: "handler_update", Name: "update", Algo: modules.Incremental},
			{ID: "handler_write", Name: "write", Algo: modules.Incremental},
		},
	},
	{
		ID:    "table_locks",
		Title: "Table Locks",
		Units: "locks/s",
		Fam:   "locks",
		Ctx:   "mysql.table_locks",
		Dims: Dims{
			{ID: "table_locks_immediate", Name: "immediate", Algo: modules.Incremental},
			{ID: "table_locks_waited", Name: "waited", Algo: modules.Incremental, Mul: -1},
		},
	},
	{
		ID:    "join_issues",
		Title: "Table Select Join Issues",
		Units: "joins/s",
		Fam:   "issues",
		Ctx:   "mysql.join_issues",
		Dims: Dims{
			{ID: "select_full_join", Name: "full join", Algo: modules.Incremental},
			{ID: "select_full_range_join", Name: "full range join", Algo: modules.Incremental},
			{ID: "select_range", Name: "range", Algo: modules.Incremental},
			{ID: "select_range_check", Name: "range check", Algo: modules.Incremental},
			{ID: "select_scan", Name: "scan", Algo: modules.Incremental},
		},
	},
	{
		ID:    "sort_issues",
		Title: "Table Sort Issues",
		Units: "issues/s",
		Fam:   "issues",
		Ctx:   "mysql.sort_issues",
		Dims: Dims{
			{ID: "sort_merge_passes", Name: "merge passes", Algo: modules.Incremental},
			{ID: "sort_range", Name: "range", Algo: modules.Incremental},
			{ID: "sort_scan", Name: "scan", Algo: modules.Incremental},
		},
	},
	{
		ID:    "tmp",
		Title: "Tmp Operations",
		Units: "events/s",
		Fam:   "temporaries",
		Ctx:   "mysql.tmp",
		Dims: Dims{
			{ID: "created_tmp_disk_tables", Name: "disk tables", Algo: modules.Incremental},
			{ID: "created_tmp_files", Name: "files", Algo: modules.Incremental},
			{ID: "created_tmp_tables", Name: "tables", Algo: modules.Incremental},
		},
	},
	{
		ID:    "connections",
		Title: "Connections",
		Units: "connections/s",
		Fam:   "connections",
		Ctx:   "mysql.connections",
		Dims: Dims{
			{ID: "connections", Name: "all", Algo: modules.Incremental},
			{ID: "aborted_connects", Name: "aborted", Algo: modules.Incremental},
		},
	},
	{
		ID:    "connections_active",
		Title: "Connections Active",
		Units: "connections",
		Fam:   "connections",
		Ctx:   "mysql.connections_active",
		Dims: Dims{
			{ID: "threads_connected", Name: "active"},
			{ID: "max_connections", Name: "limit"},
			{ID: "max_used_connections", Name: "max active"},
		},
	},
	{
		ID:    "binlog_cache",
		Title: "Binlog Cache",
		Units: "transactions/s",
		Fam:   "binlog",
		Ctx:   "mysql.binlog_cache",
		Dims: Dims{
			{ID: "binlog_cache_disk_use", Name: "disk", Algo: modules.Incremental},
			{ID: "binlog_cache_use", Name: "all", Algo: modules.Incremental},
		},
	},
	{
		ID:    "threads",
		Title: "Threads",
		Units: "threads",
		Fam:   "threads",
		Ctx:   "mysql.threads",
		Dims: Dims{
			{ID: "threads_connected", Name: "connected"},
			{ID: "threads_cached", Name: "cached", Mul: -1},
			{ID: "threads_running", Name: "running"},
		},
	},
	{
		ID:    "threads_creation_rate",
		Title: "Threads Creation Rate",
		Units: "threads/s",
		Fam:   "threads",
		Ctx:   "mysql.threads",
		Dims: Dims{
			{ID: "threads_created", Name: "created", Algo: modules.Incremental},
		},
	},
	{
		ID:    "thread_cache_misses",
		Title: "Threads Cache Misses",
		Units: "misses",
		Fam:   "threads",
		Ctx:   "mysql.thread_cache_misses",
		Type:  modules.Area,
		Dims: Dims{
			{ID: "thread_cache_misses", Name: "misses", Div: 100},
		},
	},
	{
		ID:    "innodb_io",
		Title: "InnoDB I/O Bandwidth",
		Units: "KiB/s",
		Fam:   "innodb",
		Ctx:   "mysql.innodb_io",
		Type:  modules.Area,
		Dims: Dims{
			{ID: "innodb_data_read", Name: "read", Algo: modules.Incremental, Div: 1024},
			{ID: "innodb_data_written", Name: "write", Algo: modules.Incremental, Div: 1024},
		},
	},
	{
		ID:    "innodb_io_ops",
		Title: "InnoDB I/O Operations",
		Units: "operations/s",
		Fam:   "innodb",
		Ctx:   "mysql.innodb_io_ops",
		Dims: Dims{
			{ID: "innodb_data_reads", Name: "reads", Algo: modules.Incremental},
			{ID: "innodb_data_writes", Name: "writes", Algo: modules.Incremental, Mul: -1},
			{ID: "innodb_data_fsyncs", Name: "fsyncs", Algo: modules.Incremental},
		},
	},
	{
		ID:    "innodb_io_pending_ops",
		Title: "InnoDB Pending I/O Operations",
		Units: "operations",
		Fam:   "innodb",
		Ctx:   "mysql.innodb_io_pending_ops",
		Dims: Dims{
			{ID: "innodb_data_pending_reads", Name: "reads"},
			{ID: "innodb_data_pending_writes", Name: "writes", Mul: -1},
			{ID: "innodb_data_pending_fsyncs", Name: "fsyncs"},
		},
	},
	{
		ID:    "innodb_log",
		Title: "InnoDB Log Operations",
		Units: "operations/s",
		Fam:   "innodb",
		Ctx:   "mysql.innodb_log",
		Dims: Dims{
			{ID: "innodb_log_waits", Name: "waits", Algo: modules.Incremental},
			{ID: "innodb_log_write_requests", Name: "write requests", Algo: modules.Incremental, Mul: -1},
			{ID: "innodb_log_writes", Name: "writes", Algo: modules.Incremental, Mul: -1},
		},
	},
	{
		ID:    "innodb_os_log",
		Title: "InnoDB OS Log Pending Operations",
		Units: "operations",
		Fam:   "innodb",
		Ctx:   "mysql.innodb_os_log",
		Dims: Dims{
			{ID: "innodb_os_log_pending_fsyncs", Name: "fsyncs"},
			{ID: "innodb_os_log_pending_writes", Name: "writes", Mul: -1},
		},
	},
	{
		ID:    "innodb_os_log_fsync_writes",
		Title: "InnoDB OS Log Operations",
		Units: "operations/s",
		Fam:   "innodb",
		Ctx:   "mysql.innodb_os_log",
		Dims: Dims{
			{ID: "innodb_os_log_fsyncs", Name: "fsyncs", Algo: modules.Incremental},
		},
	},
	{
		ID:    "innodb_os_log_io",
		Title: "InnoDB OS Log Bandwidth",
		Units: "KiB/s",
		Fam:   "innodb",
		Ctx:   "mysql.innodb_os_log_io",
		Type:  modules.Area,
		Dims: Dims{
			{ID: "innodb_os_log_written", Name: "write", Algo: modules.Incremental, Mul: -1, Div: 1024},
		},
	},
	{
		ID:    "innodb_cur_row_lock",
		Title: "InnoDB Current Row Locks",
		Units: "operations",
		Fam:   "innodb",
		Ctx:   "mysql.innodb_cur_row_lock",
		Type:  modules.Area,
		Dims: Dims{
			{ID: "innodb_row_lock_current_waits", Name: "current waits"},
		},
	},
	{
		ID:    "innodb_rows",
		Title: "InnoDB Row Operations",
		Units: "operations/s",
		Fam:   "innodb",
		Ctx:   "mysql.innodb_rows",
		Type:  modules.Area,
		Dims: Dims{
			{ID: "innodb_rows_inserted", Name: "inserted", Algo: modules.Incremental},
			{ID: "innodb_rows_read", Name: "read", Algo: modules.Incremental},
			{ID: "innodb_rows_updated", Name: "updated", Algo: modules.Incremental},
			{ID: "innodb_rows_deleted", Name: "deleted", Algo: modules.Incremental, Mul: -1},
		},
	},
	{
		ID:    "innodb_buffer_pool_pages",
		Title: "InnoDB Buffer Pool Pages",
		Units: "pages",
		Fam:   "innodb",
		Ctx:   "mysql.innodb_buffer_pool_pages",
		Dims: Dims{
			{ID: "innodb_buffer_pool_pages_data", Name: "data"},
			{ID: "innodb_buffer_pool_pages_dirty", Name: "dirty", Mul: -1},
			{ID: "innodb_buffer_pool_pages_free", Name: "free"},
			{ID: "innodb_buffer_pool_pages_misc", Name: "misc", Mul: -1},
			{ID: "innodb_buffer_pool_pages_total", Name: "total"},
		},
	},
	{
		ID:    "innodb_buffer_pool_flush_pages_requests",
		Title: "InnoDB Buffer Pool Flush Pages Requests",
		Units: "requests/s",
		Fam:   "innodb",
		Ctx:   "mysql.innodb_buffer_pool_pages",
		Dims: Dims{
			{ID: "innodb_buffer_pool_pages_flushed", Name: "flush pages", Algo: modules.Incremental},
		},
	},
	{
		ID:    "innodb_buffer_pool_bytes",
		Title: "InnoDB Buffer Pool Bytes",
		Units: "MiB",
		Fam:   "innodb",
		Ctx:   "mysql.innodb_buffer_pool_bytes",
		Type:  modules.Area,
		Dims: Dims{
			{ID: "innodb_buffer_pool_bytes_data", Name: "data", Div: 1024 * 1024},
			{ID: "innodb_buffer_pool_bytes_dirty", Name: "dirty", Mul: -1, Div: 1024 * 1024},
		},
	},
	{
		ID:    "innodb_buffer_pool_read_ahead",
		Title: "InnoDB Buffer Pool Bytes",
		Units: "operations/s",
		Fam:   "innodb",
		Ctx:   "mysql.innodb_buffer_pool_read_ahead",
		Type:  modules.Area,
		Dims: Dims{
			{ID: "innodb_buffer_pool_read_ahead", Name: "all", Algo: modules.Incremental},
			{ID: "innodb_buffer_pool_read_ahead_evicted", Name: "evicted", Algo: modules.Incremental, Mul: -1},
			{ID: "innodb_buffer_pool_read_ahead_rnd", Name: "random", Algo: modules.Incremental},
		},
	},
	{
		ID:    "innodb_buffer_pool_ops",
		Title: "InnoDB Buffer Pool Operations",
		Units: "operations/s",
		Fam:   "innodb",
		Ctx:   "mysql.innodb_buffer_pool_ops",
		Type:  modules.Area,
		Dims: Dims{
			{ID: "innodb_buffer_pool_reads", Name: "disk reads", Algo: modules.Incremental},
			{ID: "innodb_buffer_pool_wait_free", Name: "wait free", Algo: modules.Incremental, Mul: -1, Div: 1},
		},
	},
	{
		ID:    "qcache_ops",
		Title: "QCache Operations",
		Units: "queries/s",
		Fam:   "qcache",
		Ctx:   "mysql.qcache_ops",
		Dims: Dims{
			{ID: "qcache_hits", Name: "hits", Algo: modules.Incremental},
			{ID: "qcache_lowmem_prunes", Name: "lowmem prunes", Algo: modules.Incremental, Mul: -1},
			{ID: "qcache_inserts", Name: "inserts", Algo: modules.Incremental},
			{ID: "qcache_not_cached", Name: "not cached", Algo: modules.Incremental, Mul: -1},
		},
	},
	{
		ID:    "qcache",
		Title: "QCache Queries in Cache",
		Units: "queries",
		Fam:   "qcache",
		Ctx:   "mysql.qcache",
		Dims: Dims{
			{ID: "qcache_queries_in_cache", Name: "queries", Algo: modules.Absolute},
		},
	},
	{
		ID:    "qcache_freemem",
		Title: "QCache Free Memory",
		Units: "MiB",
		Fam:   "qcache",
		Ctx:   "mysql.qcache_freemem",
		Type:  modules.Area,
		Dims: Dims{
			{ID: "qcache_free_memory", Name: "free", Div: 1024 * 1024},
		},
	},
	{
		ID:    "qcache_memblocks",
		Title: "QCache Memory Blocks",
		Units: "blocks",
		Fam:   "qcache",
		Ctx:   "mysql.qcache_memblocks",
		Dims: Dims{
			{ID: "qcache_free_blocks", Name: "free"},
			{ID: "qcache_total_blocks", Name: "total"},
		},
	},
	{
		ID:    "key_blocks",
		Title: "MyISAM Key Cache Blocks",
		Units: "blocks",
		Fam:   "myisam",
		Ctx:   "mysql.key_blocks",
		Dims: Dims{
			{ID: "key_blocks_unused", Name: "unused"},
			{ID: "key_blocks_used", Name: "used", Mul: -1},
			{ID: "key_blocks_not_flushed", Name: "not flushed"},
		},
	},
	{
		ID:    "key_requests",
		Title: "MyISAM Key Cache Requests",
		Units: "requests/s",
		Fam:   "myisam",
		Ctx:   "mysql.key_requests",
		Type:  modules.Area,
		Dims: Dims{
			{ID: "key_read_requests", Name: "reads", Algo: modules.Incremental},
			{ID: "key_write_requests", Name: "writes", Algo: modules.Incremental, Mul: -1},
		},
	},
	{
		ID:    "key_disk_ops",
		Title: "MyISAM Key Cache Disk Operations",
		Units: "operations/s",
		Fam:   "myisam",
		Ctx:   "mysql.key_disk_ops",
		Type:  modules.Area,
		Dims: Dims{
			{ID: "key_reads", Name: "reads", Algo: modules.Incremental},
			{ID: "key_writes", Name: "writes", Algo: modules.Incremental, Mul: -1},
		},
	},
	{
		ID:    "files",
		Title: "Open Files",
		Units: "files",
		Fam:   "files",
		Ctx:   "mysql.files",
		Dims: Dims{
			{ID: "open_files", Name: "files"},
		},
	},
	{
		ID:    "files_rate",
		Title: "Opened Files Rate",
		Units: "files/s",
		Fam:   "files",
		Ctx:   "mysql.files_rate",
		Dims: Dims{
			{ID: "opened_files", Name: "files", Algo: modules.Incremental},
		},
	},
	{
		ID:    "binlog_stmt_cache",
		Title: "Binlog Statement Cache",
		Units: "statements/s",
		Fam:   "binlog",
		Ctx:   "mysql.binlog_stmt_cache",
		Dims: Dims{
			{ID: "binlog_stmt_cache_disk_use", Name: "disk", Algo: modules.Incremental},
			{ID: "binlog_stmt_cache_use", Name: "all", Algo: modules.Incremental},
		},
	},
	{
		ID:    "connection_errors",
		Title: "Connection Errors",
		Units: "errors/s",
		Fam:   "connections",
		Ctx:   "mysql.connection_errors",
		Dims: Dims{
			{ID: "connection_errors_accept", Name: "accept", Algo: modules.Incremental},
			{ID: "connection_errors_internal", Name: "internal", Algo: modules.Incremental},
			{ID: "connection_errors_max_connections", Name: "max", Algo: modules.Incremental},
			{ID: "connection_errors_peer_address", Name: "peer addr", Algo: modules.Incremental},
			{ID: "connection_errors_select", Name: "select", Algo: modules.Incremental},
			{ID: "connection_errors_tcpwrap", Name: "tcpwrap", Algo: modules.Incremental},
		},
	},
}

var slaveCharts = Charts{
	{
		ID:    "slave_behind",
		Title: "Slave Behind Seconds",
		Units: "seconds",
		Fam:   "slave",
		Ctx:   "mysql.slave_behind",
		Type:  modules.Line,
		Dims: Dims{
			{ID: "seconds_behind_master", Name: "time"},
		},
	},
	{
		ID:    "slave_thread_running",
		Title: "I/O / SQL Thread Running State",
		Units: "bool",
		Fam:   "slave",
		Ctx:   "mysql.slave_thread_running",
		Dims: Dims{
			{ID: "slave_sql_running", Name: "sql"},
			{ID: "slave_io_running", Name: "io"},
		},
	},
}

var galeraCharts = Charts{
	{
		ID:    "galera_writesets",
		Title: "Replicated Writesets",
		Units: "writesets/s",
		Fam:   "galera",
		Ctx:   "mysql.galera_writesets",
		Dims: Dims{
			{ID: "wsrep_received", Name: "rx", Algo: modules.Incremental},
			{ID: "wsrep_replicated", Name: "tx", Algo: modules.Incremental, Mul: -1},
		},
	},
	{
		ID:    "galera_bytes",
		Title: "Replicated Bytes",
		Units: "KiB/s",
		Fam:   "galera",
		Ctx:   "mysql.galera_bytes",
		Type:  modules.Area,
		Dims: Dims{
			{ID: "wsrep_received_bytes", Name: "rx", Algo: modules.Incremental, Div: 1024},
			{ID: "wsrep_replicated_bytes", Name: "tx", Algo: modules.Incremental, Mul: -1, Div: 1024},
		},
	},
	{
		ID:    "galera_queue",
		Title: "Galera Queue",
		Units: "writesets",
		Fam:   "galera",
		Ctx:   "mysql.galera_queue",
		Dims: Dims{
			{ID: "wsrep_local_recv_queue", Name: "rx"},
			{ID: "wsrep_local_send_queue", Name: "tx", Mul: -1},
		},
	},
	{
		ID:    "galera_conflicts",
		Title: "Replication Conflicts",
		Units: "transactions",
		Fam:   "galera",
		Ctx:   "mysql.galera_conflicts",
		Type:  modules.Area,
		Dims: Dims{
			{ID: "wsrep_local_bf_aborts", Name: "bf aborts", Algo: modules.Incremental},
			{ID: "wsrep_local_cert_failures", Name: "cert fails", Algo: modules.Incremental, Mul: -1},
		},
	},
	{
		ID:    "galera_flow_control",
		Title: "Flow Control",
		Units: "ms",
		Fam:   "galera",
		Ctx:   "mysql.galera_flow_control",
		Type:  modules.Area,
		Dims: Dims{
			{ID: "wsrep_flow_control_paused_ns", Name: "paused", Algo: modules.Incremental, Div: 1000000},
		},
	},
}
