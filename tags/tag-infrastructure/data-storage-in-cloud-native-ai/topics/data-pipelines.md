### Data pipelines - bridging from Transactions to Analytics in modern AI world

For decades the architecture for transaction processing and analytics processing has been evolving to meet the demands of that era. Lately, industry drivers for AI and near-real-time insights have made the batch model a bottleneck. Simultaneously, cloud-native paradigms push systems towards asynchronous, loosely-coupled interactions and an architecture where data is emitted as a stream of events.

In addition to traditional real-time transactions accessing databases and data stores, there are now asynchronous events, AI agents, and applications data (e.g. logs) accessing and updating these data stores that need to be factored. Many data stores (e.g. Postgres) have evolved to multi-model data stores that include structure, vector, time series, geo-spatial, and graph.

These drivers have given rise to many options for many streaming Extract-transform and load. Many patterns are emerging to fit diverse use case needs. This section must expand open-source and CNCF infrastructures that support data analytics in 2026 and beyond.

Examples of infrastructure that need to investigated are change data capture CDC (e.g. Debezium), Apache Kafka, workflow orchestration (Directed acyclic graphs DAGs) tools such as Apache Airflow, data quality tools (e.g. Great Expectations), Query engines (e.g. Apache Spark) and analytics store Apache Iceberg open table format on Parquet file format.
