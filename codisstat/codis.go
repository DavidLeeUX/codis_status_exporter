package codisstat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	OK    = 1
	ERROR = -1
)

type CodisData struct {
	Version string     `json:version`
	Compile string     `json:compile`
	Model   CodisModel `json:model`
	Stats   CodisStats `json:stats`
}
type CodisModel struct {
	Token        string `json:token`
	Start_Time   string `json:start_time`
	Product_Name string `json:product_name`
}
type CodisStats struct {
	Group CodisStatsGroup `json:group`
	Proxy CodisStatsProxy `json:proxy`
}

type CodisStatsGroup struct {
	Models []CodisStatsGroupModels         `json:models`
	Stats  map[string]CodisStatsGroupStats `json:stats`
}

type CodisStatsGroupModels struct {
	Id          int64                   `json:id`
	Servers     []CodisStatsGroupServer `json:servers`
	Out_of_sync bool                    `json:out_of_sync`
}
type CodisStatsGroupServer struct {
	Server        string            `json:server`
	Datacenter    string            `json:datacenter`
	Action        map[string]string `json:action`
	Replica_group bool              `json:replica_group`
}
type CodisStatsGroupStats struct {
	Stats    CodisStatsGroupStatsServerInfo `json:stats`
	Unixtime int64                          `json:unixtime`
}
type CodisStatsGroupStatsServerInfo struct {
	Aof_current_rewrite_time_sec   string `json:aof_current_rewrite_time_sec`
	Aof_enabled                    string `json:aof_enabled`
	Aof_last_bgrewrite_status      string `json:aof_last_bgrewrite_status`
	Aof_last_rewrite_time_sec      string `json:aof_last_rewrite_time_sec`
	Aof_last_write_status          string `json:aof_last_write_status`
	Aof_rewrite_in_progress        string `json:aof_rewrite_in_progress`
	Aof_rewrite_scheduled          string `json:aof_rewrite_scheduled`
	Blocked_clients                string `json:blocked_clients`
	Client_biggest_input_buf       string `json:client_biggest_input_buf`
	Client_longest_output_list     string `json:client_longest_output_list`
	Connected_clients              string `json:connected_clients`
	Connected_slaves               string `json:connected_slaves`
	Evicted_keys                   string `json:evicted_keys`
	Expired_keys                   string `json:expired_keys`
	Instantaneous_input_kbps       string `json:instantaneous_input_kbps`
	Instantaneous_ops_per_sec      string `json:instantaneous_ops_per_sec`
	Instantaneous_output_kbps      string `json:instantaneous_output_kbps`
	Keyspace_hits                  string `json:keyspace_hits`
	Keyspace_misses                string `json:keyspace_misses`
	Latest_fork_usec               string `json:latest_fork_usec`
	Loading                        string `json:loading`
	Lru_clock                      string `json:lru_clock`
	Master_repl_offset             string `json:master_repl_offset`
	Maxmemory                      string `json:maxmemory`
	Mem_fragmentation_ratio        string `json:mem_fragmentation_ratio`
	Migrate_cached_sockets         string `json:migrate_cached_sockets`
	Process_id                     string `json:process_id`
	Pubsub_channels                string `json:pubsub_channels`
	Pubsub_patterns                string `json:pubsub_patterns`
	Rdb_bgsave_in_progress         string `json:rdb_bgsave_in_progress`
	Rdb_changes_since_last_save    string `json:rdb_changes_since_last_save`
	Rdb_current_bgsave_time_sec    string `json:rdb_current_bgsave_time_sec`
	Rdb_last_bgsave_status         string `json:rdb_last_bgsave_status`
	Rdb_last_bgsave_time_sec       string `json:rdb_last_bgsave_time_sec`
	Rdb_last_save_time             string `json:rdb_last_save_time`
	Redis_build_id                 string `json:redis_build_id`
	Redis_version                  string `json:redis_version`
	Rejected_connections           string `json:rejected_connections`
	Repl_backlog_active            string `json:repl_backlog_active`
	Repl_backlog_first_byte_offset string `json:repl_backlog_first_byte_offset`
	Repl_backlog_histlen           string `json:repl_backlog_histlen`
	Repl_backlog_size              string `json:repl_backlog_size`
	Role                           string `json:role`
	Sync_full                      string `json:sync_full`
	Sync_partial_err               string `json:sync_partial_err`
	Sync_partial_ok                string `json:sync_partial_ok`
	Total_commands_processed       string `json:total_commands_processed`
	Total_connections_received     string `json:total_connections_received`
	Total_net_input_bytes          string `json:total_net_input_bytes`
	Total_net_output_bytes         string `json:total_net_output_bytes`
	Total_system_memory            string `json:total_system_memory`
	Uptime_in_seconds              string `json:uptime_in_seconds`
	Used_cpu_sys                   string `json:used_cpu_sys`
	Used_cpu_sys_children          string `json:used_cpu_sys_children`
	Used_cpu_user                  string `json:used_cpu_user`
	Used_cpu_user_children         string `json:used_cpu_user_children`
	Used_memory                    string `json:used_memory`
	Used_memory_lua                string `json:used_memory_lua`
	Used_memory_peak               string `json:used_memory_peak`
	Used_memory_rss                string `json:used_memory_rss`
	Db0                            string `json:db0`
}

type CodisStatsGroupDB struct {
	Keys    int64 `json:keys`
	Expires int64 `json:expires`
	Avg_ttl int64 `json:avg_ttl`
}
type CodisStatsProxy struct {
	Models []CodisStatsProxyModels         `json:models`
	Stats  map[string]CodisStatsProxyStats `json:stats`
}

type CodisStatsProxyModels struct {
	Id           int64  `json:id`
	Token        string `json:token`
	Start_time   string `json:start_time`
	Hostname     string `json:hostname`
	Admin_addr   string `json:admin_addr`
	Proto_type   string `json:proto_type`
	Proxy_addr   string `json:proxy_addr`
	Product_name string `json:product_name`
}
type CodisStatsProxyStats struct {
	Stats    CodisStatsProxyInfo `json:stats`
	Unixtime int64               `json:unixtime`
}
type CodisStatsProxyInfo struct {
	Online   bool          `json:online`
	Closed   bool          `json:closed`
	Ops      statsOps      `json:ops`
	Sessions statsSessions `json:sessions`
	Rusage   statsRusage   `json:rusage`
}

type statsOps struct {
	Total int64            `json:total`
	Fails int64            `json:fails`
	Redis map[string]int64 `json:redis`
	Qps   int64            `json:qps`
}
type statsSessions struct {
	Total int64 `json:total`
	Alive int64 `json:alive`
}
type statsRusage struct {
	Now string         `json:now`
	Cpu float64        `json:cpu`
	Mem int64          `json:mem`
	Raw statsRusageRaw `json:raw`
}
type statsRusageRaw struct {
	Utime       int64 `json:utime`
	Stime       int64 `json:stime`
	Cutime      int64 `json:cutime`
	Cstime      int64 `json:cstime`
	Num_Threads int64 `json:num_threads`
	VmSize      int64 `json:vm_size`
	VmRss       int64 `json:vm_rss`
}

type ProxyDataRaw struct {
	Id           int64
	Start_time   int64
	Admin_addr   string
	Proto_type   string
	Proxy_addr   string
	Product_name string
	Online       bool
	Closed       bool
	Ops_total    int64

	Ops_fails       int64
	Curr_qps        int64
	Ops_redis_error int64
	Sessions_total  int64
	Sessions_alive  int64
	Rusage_cpu      float64
	Rusage_mem      int64
	Utime           int64
	Stime           int64
	Cutime          int64
	Cstime          int64
	Num_threads     int64
	Vm_size         int64
	Vm_rss          int64
}

func NewCodisData(url string) (*CodisData, error) {

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var data CodisData
	json.Unmarshal(body, &data)
	//fmt.Printf("Results: %v\n", data)
	//fmt.Println(data.Model.Product_Name)
	//	t1 := ToNowTimeSec(data.Model.Start_Time)
	//fmt.Println(t1)

	//fmt.Printf(string(body))
	return &data, nil
}

func (codis *CodisData) GetStartTimeSec() int64 {
	return ToNowTimeSec(codis.Model.Start_Time)
}

func (codis *CodisData) ProcessProxyData() map[string]*ProxyDataRaw {

	var result = make(map[string]*ProxyDataRaw)
	for _, k := range codis.Stats.Proxy.Models {

		result[k.Proxy_addr] = &ProxyDataRaw{
			Admin_addr:      k.Admin_addr,
			Closed:          codis.Stats.Proxy.Stats[k.Token].Stats.Closed,
			Cstime:          codis.Stats.Proxy.Stats[k.Token].Stats.Rusage.Raw.Cstime,
			Curr_qps:        codis.Stats.Proxy.Stats[k.Token].Stats.Ops.Qps,
			Id:              k.Id,
			Start_time:      ToNowTimeSec(k.Start_time),
			Proto_type:      k.Proto_type,
			Proxy_addr:      k.Proxy_addr,
			Product_name:    k.Product_name,
			Online:          codis.Stats.Proxy.Stats[k.Token].Stats.Online,
			Ops_total:       codis.Stats.Proxy.Stats[k.Token].Stats.Ops.Total,
			Ops_fails:       codis.Stats.Proxy.Stats[k.Token].Stats.Ops.Fails,
			Ops_redis_error: codis.Stats.Proxy.Stats[k.Token].Stats.Ops.Redis["errors"],
			Sessions_total:  codis.Stats.Proxy.Stats[k.Token].Stats.Sessions.Total,
			Sessions_alive:  codis.Stats.Proxy.Stats[k.Token].Stats.Sessions.Alive,
			Rusage_cpu:      codis.Stats.Proxy.Stats[k.Token].Stats.Rusage.Cpu,
			Rusage_mem:      codis.Stats.Proxy.Stats[k.Token].Stats.Rusage.Mem,
			Utime:           codis.Stats.Proxy.Stats[k.Token].Stats.Rusage.Raw.Utime,
			Stime:           codis.Stats.Proxy.Stats[k.Token].Stats.Rusage.Raw.Stime,
			Cutime:          codis.Stats.Proxy.Stats[k.Token].Stats.Rusage.Raw.Cutime,
			Num_threads:     codis.Stats.Proxy.Stats[k.Token].Stats.Rusage.Raw.Num_Threads,
			Vm_size:         codis.Stats.Proxy.Stats[k.Token].Stats.Rusage.Raw.VmSize,
			Vm_rss:          codis.Stats.Proxy.Stats[k.Token].Stats.Rusage.Raw.VmRss,
		}

	}
	return result

}

func (codis *CodisData) GetTotalQps() int64 {
	var total int64
	for _, v := range codis.Stats.Proxy.Stats {
		total += v.Stats.Ops.Qps
	}
	return total
}

func GetDb0Data(dbStr string) (CodisStatsGroupDB, error) {
	var result = CodisStatsGroupDB{}

	for _, v := range strings.Split(dbStr, ",") {
		dbSlice := strings.Split(v, "=")
		if len(dbSlice) >= 2 {
			dbSliceInt, err := strconv.ParseInt(dbSlice[1], 10, 64)
			if err != nil {
				return result, err
			}
			switch dbSlice[0] {
			case "keys":
				result.Keys = dbSliceInt
			case "expires":
				result.Expires = dbSliceInt
			case "avg_ttl":
				result.Avg_ttl = dbSliceInt
			}
		} else {
			return result, errors.New("db0 data error")
		}

	}
	return result, nil
}
func GetTimeCST(t1 string) (ts int64) {
	fmt.Println(t1)
	layout := "2006-01-02 15:04:05 -0700 MST"
	t, err := time.Parse(layout, t1)
	if err != nil {
		return
	}
	ts = t.Unix()
	return
}

func ToNowTimeSec(t1 string) (ts int64) {
	ts = GetTimeCST(t1)
	return time.Now().Unix() - ts
}
