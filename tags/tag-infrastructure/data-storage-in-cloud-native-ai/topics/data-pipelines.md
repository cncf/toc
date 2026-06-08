# Data Pipelines - from Transactions to Analytics in the AI era

For decades the data storage architecture for transaction processing and analytics/AI processing has been evolving to meet the demands of that era. Lately, industry drivers for AI and near-real-time insights have made the batch model a bottleneck. Simultaneously, cloud-native paradigms push systems towards asynchronous, loosely-coupled interactions and an architecture where data is emitted as a stream of events⁵.

In addition to traditional real-time transactions accessing databases and data stores, there are now asynchronous events, AI agents, and applications data (e.g. logs) accessing and updating these data stores that need to be factored. Many data stores (e.g. PostgreSQL) have evolved to multi-modal data stores required for AI use cases, that include, relational, vector, time series, geo-spatial, and graph⁶. Many of these infrastructures can be set up as cloud native deployments with native Kubernetes support or via operators providing all the benefits of activation, scaling, and monitoring.

These drivers have given rise to many options for batch or streaming Extract, transform and load for analytics and agentic AI. There are many patterns to fit diverse use case needs. This section expands on open-source and CNCF infrastructures that support data analytics in 2026 and beyond.

There are 3 major patterns that an organization can adopt as a roadmap:
- Batch processing data from operational databases to warehouse style databases. Lower complexity and resources.
- Micro-batching from operational databases to Lakehouse engines for richer and higher performance analytics as well as MLops for training AI models - balance capabilities, performance, and resources.
- For near real-time analytics: Real-time streaming from operational data stores to analytics databases or lakehouses typically using change data capture and Kafka.

Infrastructures that support one or more of the above patterns are:
- **Data flow orchestration**: Platforms designed to create, schedule and monitor workflows for many enterprise data centric use cases - business ops, ETL/ELT, MLOps and Infrastructure management. They include directed acyclic graphs (DAGs), built-in retries, and idempotency (examples Apache Airflow, Apache NiFi, Prefect, Dagster, and Knative).
- **Change data capture**: Captures database change events (insert/update/delete) and publishes them to a stream, allowing near real-time replication of operational data into analytical systems (examples Debezium, GoldenGate, StreamSets, and native CDC in Postgres).
- **Event streaming platforms**: High-throughput, fault-tolerant event bus that decouples data producers from consumers. Enables real-time distribution of changes from databases, sensors, logs, and other data sources to multiple subscribers with minimal latency. Notable ones are Apache Kafka, Apache Flink, and Strimzi.
- **Data Quality**: Critical criteria for analytics and training data (examples Great Expectations, Deequ, and Soda)
- **Apache Iceberg Data Lakehouse**: Legacy data lake with ACID compliance. It's a vendor neutral standard with open table format and REST-based API catalog interface. The table format handles schema evolution, partition evolution, and time travel, and partitioning without sacrificing openness.
- **Query Engines**: Required for data processing in both batch and stream modes. Typical capabilities are distributed, multi-language, handle large datasets, high throughput, low latency (examples Apache Spark, Trino, DuckDB, and PySpark)

The top three architecture patterns for tackling data movement and transformation with pipelines are documented with proof of concept (POC) in this GitHub site⁷.

---

## References

5. "Building a Modern Data Platform based on Data Lakehouse architecture and Cloud-Native Ecosystem", Ahmed AbouZaid, Sep 2024, https://doi.org/10.21203/rs.3.rs-4824797/v1
6. "Just use Postgres for Everything", Stephan Schmidt, updated Dec 2025, https://www.amazingcto.com/postgres-for-everything/
7. Modern ETL Infrastructure, https://github.com/kalluripradeep/modern-etl-stack
