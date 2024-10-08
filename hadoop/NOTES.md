1. ensure to have passwordless ssh key, create one if you don't have it
2. override ssh options in hadoop with `export HADOOP_SSH_OPTS="~/.ssh/hadoop_id_rsa"`
3. in macos, ensure your sharing setting already allows you to ssh localhost.
3. follow the command here. https://hadoop.apache.org/docs/r2.8.5/hadoop-project-dist/hadoop-common/SingleCluster.html


port ui: http://localhost:9870

use `jps` to confirm if datanode and namenodes is working

	•	NameNode Web UI: http://localhost:9870
	•	DataNode Web UI: http://localhost:9864
	•	SecondaryNameNode Web UI: http://localhost:9868