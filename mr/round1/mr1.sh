#!/bin/sh
set -e
INPUT=""
OUTPUT=""
echo "INPUT = ${INPUT},  OUTPUT= ${OUTPUT}"
hadoop fs -rm -r ${OUTPUT} || true
hadoop jar /opt/hadoop-streaming/hadoop-streaming.jar \
  -files "mapper,reducer" \
  -input ${INPUT} \
  -output ${OUTPUT} \
  -jobconf num.key.fields.for.partition=1 \
  -jobconf stream.num.map.output.key.fields=2 \
  -mapper mapper \
  -reducer reducer \
  -libjars /ldap_home/regdeep_general/hadoop/hadoop-4mc-2.0.0.jar \
  -inputformat com.hadoop.mapred.FourMcTextInputFormat \
  -jobconf mapreduce.map.memory.mb=4800 \
  -jobconf mapreduce.reduce.memory.mb=9600 \
  -jobconf mapreduce.job.name="mr:round1" \
  -jobconf mapred.reduce.tasks=200
