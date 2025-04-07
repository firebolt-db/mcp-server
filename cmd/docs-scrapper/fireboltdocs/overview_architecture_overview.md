# [](#firebolt-architecture)Firebolt architecture

The following diagram gives a high-level overview of Firebolt’s structural architecture. Firebolt’s components, which include management, compute, and storage layers, interact with a variety of workloads to enhance performance, scalability, and resource efficiency.

![Firebolt's architecture has management, compute, and storage layers that work with various workloads.](../../assets/images/fireboltarchitecture.png)

## [](#management-layer)Management layer

Firebolt’s **management layer** handles key administrative functions including managing metadata, security settings, and observability, all in one place. Through this layer, administrators can oversee and control user access, permissions, and roles, ensuring robust security. It also provides insights into system performance and resource usage, providing operational visibility. The management layer also streamlines workspace management, enabling seamless organization and monitoring of data environments, further simplifying the overall administration process.

Firebolt’s management layer consists of the following:

- **Administration** - Manage access control, user roles, permissions, and system configurations.
- **Metadata** - Track and manage data definitions, schema information, and relationships within the database.
- **Security** - Configure and enforce access controls, authentication, encryption, and user permissions to protect data and ensure compliance.
- **Observability** - Monitor system performance, track query performance, and analyze resource utilization so that you can optimize operations and troubleshoot issues.
- **Workspace** - Organize and manage data environments, resources, and project-specific configurations, enabling collaboration and streamlining data workflows.

## [](#compute-layer)Compute layer

Firebolt’s **compute layer** is responsible for running queries and processing data through its scalable engines. Engines use parallel processing to deliver high performance and efficiency. You can create and configure multiple engines tailored to different workflows, such as data integration or analytical queries supporting customer facing analytics. Engines can be configured for different needs like query latency, throughput, and concurrency, adapting to specific data processing tasks.

Firebolt’s compute layer features the following:

- **Scalable engines** - Firebolt’s compute layer provides scalable engines that process queries in parallel, ensuring high-speed performance and efficiency for diverse workloads.
- **Customizable configurations** - You can tailor engine configurations to match your specific workflow needs.
- **Multiple engines for workload isolation** - You can create multiple engines to handle mixed workloads. Each workload runs in an isolated compute environment, so that you can avoid performance issues caused by noisy neighbors.
- **On-demand resource allocation** - Engines can be started, stopped, or resized based to match current workload demands.

## [](#storage-layer)Storage layer

Firebolt’s **storage layer** efficiently manages large amounts of data by keeping storage separate from the compute process, which means that you can store as much data as needed without impacting compute resources. It uses cloud-based storage for high availability and durability, while also reducing costs. Data is stored in a compressed, column-based format to save space and improve query performance. Firebolt’s indexing features further speed up data retrieval by reducing the need to scan large datasets. This separation between storage and compute allows users to scale their storage needs independently, making resource management more flexible and cost-efficient.

Firebolt’s storage layer features the following:

- **Separation of storage and compute** - You can scale data storage independently of the computing resources used for processing queries. This separation lets you adjust storage without affecting compute resources, allowing for flexible and efficient resource management.
- **Cloud-based and cost-effective** - Firebolt leverages cloud storage for high availability, durability, and reduced costs.
- **Columnar format** - Data is stored in a compressed, column-based format, optimizing both storage efficiency and query performance.
- **Fast data retrieval with advanced indexing** - Indexing features including primary and aggregating indexes speed up queries by minimizing the need to scan large datasets.