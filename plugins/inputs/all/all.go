package all

import (
	_ "github.com/influxdata/telegraf/plugins/inputs/aerospike"
	_ "github.com/influxdata/telegraf/plugins/inputs/amqp_consumer"
	_ "github.com/influxdata/telegraf/plugins/inputs/apache"
	_ "github.com/influxdata/telegraf/plugins/inputs/bcache"
	_ "github.com/influxdata/telegraf/plugins/inputs/bond"
	_ "github.com/influxdata/telegraf/plugins/inputs/cassandra"
	_ "github.com/influxdata/telegraf/plugins/inputs/ceph"
	_ "github.com/influxdata/telegraf/plugins/inputs/cgroup"
	_ "github.com/influxdata/telegraf/plugins/inputs/chrony"
	_ "github.com/influxdata/telegraf/plugins/inputs/cloudwatch"
	_ "github.com/influxdata/telegraf/plugins/inputs/conntrack"
	_ "github.com/influxdata/telegraf/plugins/inputs/consul"
	_ "github.com/influxdata/telegraf/plugins/inputs/couchbase"
	_ "github.com/influxdata/telegraf/plugins/inputs/couchdb"
	_ "github.com/influxdata/telegraf/plugins/inputs/dcos"
	_ "github.com/influxdata/telegraf/plugins/inputs/disque"
	_ "github.com/influxdata/telegraf/plugins/inputs/dmcache"
	_ "github.com/influxdata/telegraf/plugins/inputs/dns_query"
	_ "github.com/influxdata/telegraf/plugins/inputs/docker"
	_ "github.com/influxdata/telegraf/plugins/inputs/dovecot"
	_ "github.com/influxdata/telegraf/plugins/inputs/elasticsearch"
	_ "github.com/influxdata/telegraf/plugins/inputs/exec"
	_ "github.com/influxdata/telegraf/plugins/inputs/fail2ban"
	_ "github.com/influxdata/telegraf/plugins/inputs/filestat"
	_ "github.com/influxdata/telegraf/plugins/inputs/fluentd"
	_ "github.com/influxdata/telegraf/plugins/inputs/graylog"
	_ "github.com/influxdata/telegraf/plugins/inputs/haproxy"
	_ "github.com/influxdata/telegraf/plugins/inputs/hddtemp"
	_ "github.com/influxdata/telegraf/plugins/inputs/http"
	_ "github.com/influxdata/telegraf/plugins/inputs/http_listener"
	_ "github.com/influxdata/telegraf/plugins/inputs/http_response"
	_ "github.com/influxdata/telegraf/plugins/inputs/httpjson"
	_ "github.com/influxdata/telegraf/plugins/inputs/influxdb"
	_ "github.com/influxdata/telegraf/plugins/inputs/internal"
	_ "github.com/influxdata/telegraf/plugins/inputs/interrupts"
	_ "github.com/influxdata/telegraf/plugins/inputs/ipmi_sensor"
	_ "github.com/influxdata/telegraf/plugins/inputs/ipset"
	_ "github.com/influxdata/telegraf/plugins/inputs/iptables"
	_ "github.com/influxdata/telegraf/plugins/inputs/jolokia"
	_ "github.com/influxdata/telegraf/plugins/inputs/jolokia2"
	_ "github.com/influxdata/telegraf/plugins/inputs/kafka_consumer"
	_ "github.com/influxdata/telegraf/plugins/inputs/kafka_consumer_legacy"
	_ "github.com/influxdata/telegraf/plugins/inputs/kapacitor"
	_ "github.com/influxdata/telegraf/plugins/inputs/kubernetes"
	_ "github.com/influxdata/telegraf/plugins/inputs/leofs"
	_ "github.com/influxdata/telegraf/plugins/inputs/logparser"
	_ "github.com/influxdata/telegraf/plugins/inputs/lustre2"
	_ "github.com/influxdata/telegraf/plugins/inputs/mailchimp"
	_ "github.com/influxdata/telegraf/plugins/inputs/memcached"
	_ "github.com/influxdata/telegraf/plugins/inputs/mesos"
	_ "github.com/influxdata/telegraf/plugins/inputs/minecraft"
	_ "github.com/influxdata/telegraf/plugins/inputs/mongodb"
	_ "github.com/influxdata/telegraf/plugins/inputs/mqtt_consumer"
	_ "github.com/influxdata/telegraf/plugins/inputs/mysql"
	_ "github.com/influxdata/telegraf/plugins/inputs/nats"
	_ "github.com/influxdata/telegraf/plugins/inputs/nats_consumer"
	_ "github.com/influxdata/telegraf/plugins/inputs/net_response"
	_ "github.com/influxdata/telegraf/plugins/inputs/nginx"
	_ "github.com/influxdata/telegraf/plugins/inputs/nginx_plus"
	_ "github.com/influxdata/telegraf/plugins/inputs/nsq"
	_ "github.com/influxdata/telegraf/plugins/inputs/nsq_consumer"
	_ "github.com/influxdata/telegraf/plugins/inputs/nstat"
	_ "github.com/influxdata/telegraf/plugins/inputs/ntpq"
	_ "github.com/influxdata/telegraf/plugins/inputs/openldap"
	_ "github.com/influxdata/telegraf/plugins/inputs/opensmtpd"
	_ "github.com/influxdata/telegraf/plugins/inputs/passenger"
	_ "github.com/influxdata/telegraf/plugins/inputs/pf"
	_ "github.com/influxdata/telegraf/plugins/inputs/phpfpm"
	_ "github.com/influxdata/telegraf/plugins/inputs/ping"
	_ "github.com/influxdata/telegraf/plugins/inputs/postfix"
	_ "github.com/influxdata/telegraf/plugins/inputs/postgresql"
	_ "github.com/influxdata/telegraf/plugins/inputs/postgresql_extensible"
	_ "github.com/influxdata/telegraf/plugins/inputs/powerdns"
	_ "github.com/influxdata/telegraf/plugins/inputs/procstat"
	_ "github.com/influxdata/telegraf/plugins/inputs/prometheus"
	_ "github.com/influxdata/telegraf/plugins/inputs/puppetagent"
	_ "github.com/influxdata/telegraf/plugins/inputs/rabbitmq"
	_ "github.com/influxdata/telegraf/plugins/inputs/raindrops"
	_ "github.com/influxdata/telegraf/plugins/inputs/redis"
	_ "github.com/influxdata/telegraf/plugins/inputs/rethinkdb"
	_ "github.com/influxdata/telegraf/plugins/inputs/riak"
	_ "github.com/influxdata/telegraf/plugins/inputs/salesforce"
	_ "github.com/influxdata/telegraf/plugins/inputs/sensors"
	_ "github.com/influxdata/telegraf/plugins/inputs/signalfxmetadata"
	_ "github.com/influxdata/telegraf/plugins/inputs/smart"
	_ "github.com/influxdata/telegraf/plugins/inputs/snmp"
	_ "github.com/influxdata/telegraf/plugins/inputs/snmp_legacy"
	_ "github.com/influxdata/telegraf/plugins/inputs/socket_listener"
	_ "github.com/influxdata/telegraf/plugins/inputs/solr"
	_ "github.com/influxdata/telegraf/plugins/inputs/sqlserver"
	_ "github.com/influxdata/telegraf/plugins/inputs/statsd"
	_ "github.com/influxdata/telegraf/plugins/inputs/sysstat"
	_ "github.com/influxdata/telegraf/plugins/inputs/system"
	_ "github.com/influxdata/telegraf/plugins/inputs/tail"
	_ "github.com/influxdata/telegraf/plugins/inputs/tcp_listener"
	_ "github.com/influxdata/telegraf/plugins/inputs/teamspeak"
	_ "github.com/influxdata/telegraf/plugins/inputs/tomcat"
	_ "github.com/influxdata/telegraf/plugins/inputs/trig"
	_ "github.com/influxdata/telegraf/plugins/inputs/twemproxy"
	_ "github.com/influxdata/telegraf/plugins/inputs/udp_listener"
	_ "github.com/influxdata/telegraf/plugins/inputs/unbound"
	_ "github.com/influxdata/telegraf/plugins/inputs/varnish"
	_ "github.com/influxdata/telegraf/plugins/inputs/webhooks"
	_ "github.com/influxdata/telegraf/plugins/inputs/win_perf_counters"
	_ "github.com/influxdata/telegraf/plugins/inputs/win_services"
	_ "github.com/influxdata/telegraf/plugins/inputs/zfs"
	_ "github.com/influxdata/telegraf/plugins/inputs/zipkin"
	_ "github.com/influxdata/telegraf/plugins/inputs/zookeeper"
)
