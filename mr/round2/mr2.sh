#!/bin/sh
set -e
INPUT=""
OUTPUT=""
echo "INPUT = ${INPUT},  OUTPUT= ${OUTPUT}"
hadoop fs -rm -r ${OUTPUT} || true
hadoop jar /opt/hadoop-streaming/hadoop-streaming.jar \
  -files "reducer" \
  -input ${INPUT} \
  -output ${OUTPUT} \
  -mapper "cat" \
  -reducer reducer \
  -jobconf mapreduce.map.memory.mb=4800 \
  -jobconf mapreduce.reduce.memory.mb=9600 \
  -jobconf mapreduce.job.name="mr:round2" \
  -jobconf mapred.reduce.tasks=10
