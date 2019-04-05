# PD server
docker run -d --name pd1 \
  --net=host \
  -v /etc/localtime:/etc/localtime:ro \
  -v /data:/data \
  pingcap/pd:latest \
  --name="pd1" \
  --data-dir="/data/pd1" \
  --client-urls="http://0.0.0.0:2379" \
  --advertise-client-urls="http://192.168.140.94:2379" \
  --peer-urls="http://0.0.0.0:2380" \
  --advertise-peer-urls="http://192.168.140.94:2380" \
  --initial-cluster="pd1=http://192.168.140.94:2380,pd2=http://192.168.140.94:2390,pd3=http://192.168.140.94:2400"

  docker run -d --name pd2 \
  --net=host \
  -v /etc/localtime:/etc/localtime:ro \
  -v /data:/data \
  pingcap/pd:latest \
  --name="pd2" \
  --data-dir="/data/pd2" \
  --client-urls="http://0.0.0.0:2389" \
  --advertise-client-urls="http://192.168.140.94:2389" \
  --peer-urls="http://0.0.0.0:2390" \
  --advertise-peer-urls="http://192.168.140.94:2390" \
  --initial-cluster="pd1=http://192.168.140.94:2380,pd2=http://192.168.140.94:2390,pd3=http://192.168.140.94:2400"

  docker run -d --name pd3 \
  --net=host \
  -v /etc/localtime:/etc/localtime:ro \
  -v /data:/data \
  pingcap/pd:latest \
  --name="pd3" \
  --data-dir="/data/pd3" \
  --client-urls="http://0.0.0.0:2399" \
  --advertise-client-urls="http://192.168.140.94:2399" \
  --peer-urls="http://0.0.0.0:2400" \
  --advertise-peer-urls="http://192.168.140.94:2400" \
  --initial-cluster="pd1=http://192.168.140.94:2380,pd2=http://192.168.140.94:2390,pd3=http://192.168.140.94:2400"

# TIKV server
docker run -d --name tikv1 \
  --net=host \
  --ulimit nofile=1000000:1000000 \
  -v /etc/localtime:/etc/localtime:ro \
  -v /data/kv1:/data \
  pingcap/tikv:latest \
  --addr="0.0.0.0:20160" \
  --advertise-addr="192.168.140.94:20160" \
  --data-dir="/data/tikv1" \
  --pd="192.168.140.94:2379,192.168.140.94:2389,192.168.140.94:2399"

docker run -d --name tikv2 \
  --net=host \
  --ulimit nofile=1000000:1000000 \
  -v /etc/localtime:/etc/localtime:ro \
  -v /data/kv2:/data \
  pingcap/tikv:latest \
  --addr="0.0.0.0:20161" \
  --advertise-addr="192.168.140.94:20161" \
  --data-dir="/data/tikv2" \
  --pd="192.168.140.94:2379,192.168.140.94:2389,192.168.140.94:2399"

docker run -d --name tikv3 \
  --net=host \
  --ulimit nofile=1000000:1000000 \
  -v /etc/localtime:/etc/localtime:ro \
  -v /data/kv3:/data \
  pingcap/tikv:latest \
  --addr="0.0.0.0:20162" \
  --advertise-addr="192.168.140.94:20162" \
  --data-dir="/data/tikv3" \
  --pd="192.168.140.94:2379,192.168.140.94:2389,192.168.140.94:2399"

# TIDB server
docker run -d --name tidb \
  -p 4000:4000 \
  -p 10080:10080 \
  pingcap/tidb:latest \
  --store=tikv \
  --path="192.168.140.94:2379,192.168.140.94:2389,192.168.140.94:2399"
