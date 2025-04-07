# [](#connecting-to-airflow)Connecting to Airflow

[Apache Airflow](https://airflow.apache.org/) is a data orchestration tool that allows you to programmatically create, schedule, and monitor workflows. You can connect a Firebolt database into your data pipeline using the Airflow provider package for Firebolt. For example, you can schedule automatic incremental data ingestion into Firebolt.

This guide explains how to install the [Airflow provider package](https://pypi.org/project/airflow-provider-firebolt/) for Firebolt, set up a connection to Firebolt resources using the Airflow user interface (UI), and create an example Directed Acyclic Graph (DAG) for common Firebolt tasks. The source code for the Airflow provider package for Firebolt is available in the [airflow-provider-firebolt](https://github.com/firebolt-db/airflow-provider-firebolt) repository on GitHub.

## [](#prerequisites)Prerequisites

Make sure that you have:

- A Firebolt account. [Create a new account](/Guides/managing-your-organization/managing-accounts.html#create-a-new-account).
- A Firebolt database and engine.
- [Python](https://www.python.org/downloads/) version 3.8 or later.
- An installation of Airflow. See the [Airflow installation guide](https://airflow.apache.org/docs/apache-airflow/stable/installation/index.html).

## [](#install-the-airflow-provider-package-for-firebolt)Install the Airflow provider package for Firebolt

You need to install the Airflow provider package for Firebolt. This package enables Firebolt as a **Connection type** in the Airflow UI.

1. Install the package.
   
   Run the following command to install the package:
   
   ```
   pip install airflow-provider-firebolt
   ```
2. Upgrade to the latest version.
   
   Run the latest version of the provider package. [Release history](https://pypi.org/project/airflow-provider-firebolt/#history) is available on PyPI.
   
   Use the following command to upgrade:
   
   ```
   pip install airflow-provider-firebolt --upgrade
   ```
   
   Restart Airflow after upgrading to apply the new changes.
3. Install a specific version.
   
   If a specific version is required, replace `1.0.0` with the desired version:
   
   ```
   pip install airflow-provider-firebolt==1.0.0
   ```
4. Install the provider for AWS Managed Airflow (MWAA).
   
   Ensure you are using version 2 of AWS Managed Airflow (MWAA) when working with the Firebolt Airflow provider. Add `airflow-provider-firebolt` to the `requirements.txt` file following the instructions in the [MWAA Documentation.](https://docs.aws.amazon.com/mwaa/latest/userguide/working-dags-dependencies.html)

## [](#connect-airflow-to-firebolt)Connect Airflow to Firebolt

Create a connection object in the Airflow UI to integrate Firebolt with Airflow.

### [](#steps-to-configure-a-connection)Steps to configure a connection

1. Open the Airflow UI and log in.
2. Select the **Admin** menu.
3. Choose **Connections**.
4. Select the **+** button to add a new connection.
5. Choose Firebolt from the **Connection Type** list
6. Provide the details in the following table. These connection parameters correspond to built-in Airflow variables.
   
   Parameter Description Example value Connection id The name of the connection for the UI. `My_Firebolt_Connection` Description Information about the connection. `Connection to Firebolt database MyDatabase using engine MyFireboltDatabase_general_purpose.` Database The name of the Firebolt database to connect to. `MyFireboltDatabase` Engine The name of the engine to run queries `MyFireboltEngine` Client ID The ID of your service account. `XyZ83JSuhsua82hs` Client Secret The [secret](/Guides/loading-data/creating-access-keys-aws.html) for your service account authentication. `yy7h&993))29&%j` Account The name of your account. `developer` Extra The additional properties that you may need to set (optional). `{"property1": "value1", "property2": "value2"}`
   
   Client ID and secret credentials can be obtained by registering a [service account](/Guides/managing-your-organization/service-accounts.html).
7. Choose **Test** to verify the connection.
8. Once the test succeeds, select **Save**.

## [](#create-a-dag-for-data-processing-with-firebolt)Create a DAG for data processing with Firebolt

A DAG file in Airflow is a Python script that defines tasks and their execution order for a data workflow. The following example is an example DAG for performing the following tasks:

- Start a Firebolt [engine](/Overview/engine-fundamentals.html).
- Create an [external table](/Guides/loading-data/working-with-external-tables.html) linked to an Amazon S3 data source.
- Create a fact table for ingested data. For more information, see [Firebolt-managed tables](/Overview/indexes/using-indexes.html#firebolt-managed-tables).
- Insert data into the fact table.
- Stop the Firebolt engine. This task is not required if your engine has `AUTO_STOP` configured

### [](#dag-script-example)DAG script example

The following DAG script creates a DAG named `firebolt_provider_trip_data`. It uses an Airflow connection to Firebolt named `my_firebolt_connection`. For the contents of the SQL scripts that the DAG runs, see the following [SQL script examples](#sql-script-examples). You can run this example with your own database and engine by updating the connector values in Airfow, setting the `FIREBOLT_CONN_ID` to match your connector, and creating the necessary custom variables in Airflow.

```
import time
import airflow
from airflow.models import DAG, Variable
from firebolt_provider.operators.firebolt \
    import FireboltOperator, FireboltStartEngineOperator, FireboltStopEngineOperator

# Define function to get Firebolt connection parameters
def connection_params(conn_opp, field):
    connector = FireboltOperator(
        firebolt_conn_id=conn_opp, sql="", task_id="CONNECT")
    conn_parameters = connector.get_db_hook()._get_conn_params()
    return getattr(conn_parameters, field)

# Set up the Firebolt connection ID
firebolt_conn_id = 'firebolt'
firebolt_engine_name = connection_params(firebolt_conn_id, 'engine_name')
tmpl_search_path = Variable.get("firebolt_sql_path")
default_args = {
    'owner': 'airflow',
    'start_date': airflow.utils.dates.days_ago(1)
}

# Function to open query files saved locally.
def get_query(query_file):
    return open(query_file, "r").read()

# Define a variable based on an Airflow DAG class.
# For class parameters, see https://airflow.apache.org/docs/apache-airflow/stable/_api/airflow/models/dag/index.html#airflow.models.dag.DAG.
with DAG('firebolt_provider_startstop_trip_data',
          default_args=default_args,
          template_searchpath=tmpl_search_path,
          schedule_interval=None,
          catchup=False,
          tags=["firebolt"]) as dag:

    # Define DAG tasks and task sequence.
    # Where necessary, read local sql files using the Airflow variable.
    task_start_engine = FireboltStartEngineOperator(
        dag=dag,
        task_id="START_ENGINE",
        firebolt_conn_id=firebolt_conn_id,
        engine_name=firebolt_engine_name)

    task_trip_data__external_table = FireboltOperator(
        dag=dag,
        task_id="task_trip_data__external_table",
        sql=get_query(f'{tmpl_search_path}/trip_data__create_external_table.sql'),
        firebolt_conn_id=firebolt_conn_id
    )

    task_trip_data__create_table = FireboltOperator(
        dag=dag,
        task_id="task_trip_data__create_table",
        sql=get_query(f'{tmpl_search_path}/trip_data__create_table.sql'),
        firebolt_conn_id=firebolt_conn_id
    )
    task_trip_data__create_table.post_execute = lambda **x: time.sleep(10)

    task_trip_data__process_data = FireboltOperator(
        dag=dag,
        task_id="task_trip_data__process_data",
        sql=get_query(f'{tmpl_search_path}/trip_data__process.sql'),
        firebolt_conn_id=firebolt_conn_id
    )

    task_stop_engine = FireboltStopEngineOperator(
        dag=dag,
        task_id="STOP_ENGINE",
        firebolt_conn_id=firebolt_conn_id,
        engine_name=firebolt_engine_name)

    (task_start_engine >> task_trip_data__external_table >>
     task_trip_data__create_table >> task_trip_data__process_data >> task_stop_engine)
```

This DAG showcases various Firebolt tasks as an example and is not intended to represent a typical real-world workflow or pipeline.

### [](#define-airflow-variables)Define Airflow variables

Airflow variables store-key value pairs that DAGs can use during execution. You can create and manage variables through the Airflow user interface (UI) or JSON documents. For detailed instructions, check out Airflow’s [Variables](https://airflow.apache.org/docs/apache-airflow/stable/concepts/variables.html) and [Managing Variables](https://airflow.apache.org/docs/apache-airflow/stable/howto/variable.html) documentation pages.

**Example variable for SQL files**  
The DAG example uses the custom variable `firebolt_sql_path` to define the directory within your Airflow home directory where SQL files are stored. The DAG reads these files to execute tasks in Firebolt.

- **Key**: `firebolt_sql_path`
- **Value**: Path to the directory containing SQL scripts. For example, `~/airflow/sql_store`.

**Using the variable in the DAG**  
A python function in the DAG reads the SQL scripts stored in the directory defined by `firebolt_sql_path`. This allows the DAG to dynamically execute the SQL files as tasks in Firebolt.

The following example demonstrates how the variable is accessed in the DAG script:

```
tmpl_search_path = Variable.get("firebolt_sql_path")

def get_query(query_file):
    with open(query_file, "r") as file:
        return file.read()
```

### [](#sql-script-examples)SQL script examples

Save the following SQL scripts to your `tmpl_search_path`.

#### [](#trip_data__create_external_tablesql)trip\_data\_\_create\_external\_table.sql

This example creates the `ex_trip_data` fact table to connect to a public Amazon S3 data store.

```
CREATE EXTERNAL TABLE IF NOT EXISTS ex_trip_data(
   vendorid INTEGER,
   lpep_pickup_datetime TIMESTAMP,
   lpep_dropoff_datetime TIMESTAMP,
   passenger_count INTEGER,
   trip_distance REAL,
   ratecodeid INTEGER,
   store_and_fwd_flag TEXT,
   pu_location_id INTEGER,
   do_location_id INTEGER,
   payment_type INTEGER,
   fare_amount REAL,
   extra REAL,
   mta_tax REAL,
   tip_amount REAL,
   tolls_amount REAL,
   improvement_surcharge REAL,
   total_amount REAL,
   congestion_surcharge REAL
)
url = 's3://firebolt-publishing-public/samples/taxi/'
object_pattern = '*yellow*2020*.csv'
type = (CSV SKIP_HEADER_ROWS = true);
```

#### [](#trip_data__create_tablesql)trip\_data\_\_create\_table.sql

This example creates the `my_taxi_trip_data` fact table, to receive ingested data.

```
DROP TABLE IF EXISTS my_taxi_trip_data;
CREATE FACT TABLE IF NOT EXISTS my_taxi_trip_data(
   vendorid INTEGER,
   lpep_pickup_datetime TIMESTAMP,
   lpep_dropoff_datetime TIMESTAMP,
   passenger_count INTEGER,
   trip_distance REAL,
   ratecodeid INTEGER,
   store_and_fwd_flag TEXT,
   pu_location_id INTEGER,
   do_location_id INTEGER,
   payment_type INTEGER,
   fare_amount REAL,
   extra REAL,
   mta_tax REAL,
   tip_amount REAL,
   tolls_amount REAL,
   improvement_surcharge REAL,
   total_amount REAL,
   congestion_surcharge REAL,
   SOURCE_FILE_NAME TEXT,
   SOURCE_FILE_TIMESTAMP TIMESTAMP
) PRIMARY INDEX vendorid;
```

#### [](#trip_data__processsql)trip\_data\_\_process.sql

An `INSERT INTO` operation ingests data into the `my_taxi_trip_data` fact table using the `ex_trip_data` external table. This example uses the external table metadata column, `$source_file_timestamp`, to retrieve records exclusively from the latest file.

```
INSERT INTO my_taxi_trip_data
SELECT
   vendorid,
   lpep_pickup_datetime,
   lpep_dropoff_datetime,
   passenger_count,
   trip_distance,
   ratecodeid,
   store_and_fwd_flag,
   pu_location_id,
   do_location_id,
   payment_type,
   fare_amount,
   extra,
   mta_tax,
   tip_amount,
   tolls_amount,
   improvement_surcharge,
   total_amount,
   congestion_surcharge,
   $source_file_name,
   $source_file_timestamp
FROM ex_trip_data
WHERE coalesce($source_file_timestamp > (SELECT MAX(source_file_timestamp) FROM my_taxi_trip_data), true);
```

## [](#example-working-with-query-results)Example: Working with query results

The `FireboltOperator` is designed to execute SQL queries but does not return query results. To retrieve query results, use the `FireboltHook` class. The following example demonstrates how to use `FireboltHook` to execute a query and log the row count in the `my_taxi_trip_data` table.

### [](#python-code-example-retrieiving-query-results)Python code example: Retrieiving query results

```
import logging

import airflow
from airflow import DAG
from airflow.operators.python import PythonOperator
from firebolt_provider.hooks.firebolt import FireboltHook
from airflow.providers.common.sql.hooks.sql import fetch_one_handler

# Set up the Firebolt connection ID
firebolt_conn_id = 'firebolt'
default_args = {
    'owner': 'airflow',
    'start_date': airflow.utils.dates.days_ago(1)
}


# Function to notify the team about the data
def notify(message: str):
    logging.info(message)


# Function to fetch data from Firebolt and notify the team
def fetch_firebolt_data():
    hook = FireboltHook(firebolt_conn_id=firebolt_conn_id)
    results = hook.run(
        "SELECT count(*) FROM my_taxi_trip_data",
        handler=fetch_one_handler
    )
    count = results[0]
    notify("Amount of data in Firebolt: " + str(count))


with DAG(
    'return_result_dag',
    default_args=default_args,
    schedule_interval=None,  # Run manually
    catchup=False
) as dag:
    # Define a Python operator to fetch data from Firebolt and notify the team
    monitor_firebolt_data = PythonOperator(
        task_id='monitor_firebolt_data',
        python_callable=fetch_firebolt_data
    )

    monitor_firebolt_data
```

## [](#example-controlling-query-execution-timeout)Example: Controlling query execution timeout

The Firebolt provider includes parameters to control query execution time and behavior when a timeout occurs:

- `query_timeout`: Sets the maximum duration (in seconds) that a query can run
- `fail_on_query_timeout` - If `True`, a timeout raises a `QueryTimeoutError`. If `False`, the task terminates quietly, and the task proceeds without raising an error.

### [](#python-code-example-using-timeout-settings)Python code example: Using timeout settings

In this example, the `FireboltOperator` task stops execution after one second and proceeds without error. The `PythonOperator` task fetches data from Firebolt with a timeout of 0.5 seconds and raises an error if the query times out.

```
import airflow
from airflow.models import DAG, Variable
from airflow.operators.python import PythonOperator
from firebolt_provider.hooks.firebolt import FireboltHook
from airflow.providers.common.sql.hooks.sql import fetch_one_handler
from firebolt_provider.operators.firebolt import FireboltOperator

# Set up the Firebolt connection ID
firebolt_conn_id = 'firebolt'
default_args = {
    'owner': 'airflow',
    'start_date': airflow.utils.dates.days_ago(1)
}
tmpl_search_path = Variable.get("firebolt_sql_path")


def get_query(query_file):
    return open(query_file, "r").read()

# Function to fetch data with a timeout
def fetch_with_timeout():
    hook = FireboltHook(
        firebolt_conn_id=firebolt_conn_id,
        query_timeout=0.5, 
        fail_on_query_timeout=True, 
    )
    results = hook.run(
        "SELECT count(*) FROM my_taxi_trip_data",
        handler=fetch_one_handler,
    )
    print(f"Results: {results}")

# Define the DAG
with DAG(
    'timeout_dag',
    default_args=default_args,
    schedule_interval=None,  # Run manually
    catchup=False
) as dag:

    # Firebolt operator with a timeout

    firebolt_operator_with_timeout = FireboltOperator(
        dag=dag,
        task_id="insert_with_timeout",
        sql=get_query(f'{tmpl_search_path}/trip_data__process.sql'),
        firebolt_conn_id=firebolt_conn_id,
        query_timeout=1,
        # Task will not fail if query times out, and will proceed to the next task
        fail_on_query_timeout=False,
    )

    # Python operator to fetch data with a timeout
    operator_with_hook_timeout = PythonOperator(
        dag=dag,
        task_id='select_with_hook_timeout',
        python_callable=fetch_with_timeout,
    )

    firebolt_operator_with_timeout >> operator_with_hook_timeout
```

## [](#additional-resources)Additional resources

For more information about connecting to Airflow, refer to the following resources:

- [Managing Connections in Airflow](https://airflow.apache.org/docs/apache-airflow/stable/howto/connection.html)
- [Firebolt Airflow provider on Pypi](https://pypi.org/project/airflow-provider-firebolt/)
- [DAGs](https://airflow.apache.org/docs/apache-airflow/stable/core-concepts/dags.html)
- [airflow.models.dag](https://airflow.apache.org/docs/apache-airflow/stable/_api/airflow/models/dag/index.html#module-airflow.models.dag)