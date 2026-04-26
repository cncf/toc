### Data Warehouses, Data Lakes, and Data Lake Houses

Data warehouses typically store structured data from transaction processing systems and other business applications. In most cases, the data is cleansed and curated before going into a data warehouse. For data warehouses, data is traditionally stored in relational databases using conventional disk storage, often using a columnar format that stores data as compressed arrays. Modern data warehouses generally use shared object storage to handle large datasets.

Data lakes can handle a combination of structured, semistructured and unstructured data, which commonly is often stored in its native format to make the full sets of raw data available for analysis. Data lakes are primarily used for data science applications that involve machine learning, predictive modeling and other advanced analytics techniques. Data is typically stored in NoSQL databases such as Cassandra, object storage, or distributed file systems such as Hadoop, etc⁵.

Data lakes and data warehouses are merging for many users. A data lake house is a hybrid model of data warehouses and data lakes⁶. It can support both structured and unstructured data and has the characteristics of both data warehouses and data lakes. For example, it is common to store information in open file formats like Parquet that offer columnar format, compression, and indexes to facilitate fast query. Data lake houses also utilize open table formats like Apache Iceberg that organize files into tabular representation complete with features like schema management and transactions. Data lakes enable safe, shared access to data from many types of analytic applications–principally data warehouses, AI/ML, and data science.

---

## References

5. Data lake vs. data warehouse: Key differences explained
6. Lakehouse: A New Generation of Open Platforms that Unify Data Warehousing and Advanced Analytics, CIDR 2021
