# [](#use-the-develop-space)Use the Develop Space

- [Open the Develop Space](#open-the-develop-space)
- [A quick tour](#a-quick-tour)
- [Using the document editor](#using-the-document-editor)
  
  - [Using auto-complete](#using-auto-complete)
  - [Using script templates](#using-script-templates)
  - [Using the CREATE EXTERNAL TABLE template to import data](#using-the-create-external-table-template-to-import-data)
- [Managing scripts](#managing-scripts)
- [Running scripts and working with results](#running-scripts-and-working-with-results)
  
  - [Viewing results](#viewing-results)
  - [Viewing multi-statement script results](#viewing-multi-statement-script-results)
  - [Exporting results to a local hard drive](#exporting-results-to-a-local-hard-drive)
- [Switching between light and dark mode](#switching-between-light-and-dark-mode)
- [Keyboard shortcuts for the Develop Space](#keyboard-shortcuts-for-the-develop-space)
  
  - [Query operations](#query-operations)
  - [Script management](#script-management)
  - [Search functionality](#search-functionality)
  - [Editing text](#editing-text)

The **Firebolt Workspace** has a **Develop Space** that you use to edit and run SQL scripts and view query results.

## [](#open-the-develop-space)Open the Develop Space

You can launch the space for a database by clicking the **Develop** icon from the left navigation pane or clicking the “+” icon next to “Script 1”.

![drawing](../../assets/images/develop_workspace_ex0.png)

![drawing](../../assets/images/develop_workspace_ex4.png)

**Starting the Develop Space for the last database you worked with**

1. Choose the **&lt;/&gt;** icon from the left navigation pane.
   
   ![drawing](../../assets/images/develop_workspace_ex0.png)
   
   The space for the database that you last worked with will open, and the database will be selected from the list.
2. To switch to different database’s space, choose from the dropdown menu in the Databases panel.

## [](#a-quick-tour)A quick tour

The **Develop Space** is organized into two panels.

- The left panel is the explore panel. You can use it to navigate to different databases and to work with different scripts in your database.
- The center panel is the document editor. You can use it to edit scripts, save them, and run scripts. When you run a script, the results will be shown in the bottom part of the pane.
  
  ![drawing](../../assets/images/develop_workspace_ex5.png)

## [](#using-the-document-editor)Using the document editor

The document editor uses tabs to help you organize your SQL scripts. You can switch tabs to work with different scripts and run them. You can have multiple query statements on the same tab. Each statement must be terminated by a semi-colon (`;`).

### [](#using-auto-complete)Using auto-complete

As you enter your code in a script tab, Firebolt suggests keywords and object names from the chosen database. Press the tab key to add the first suggestion in the list to your script, or use arrow keys to select a different item from the list and then press the tab key.

### [](#using-script-templates)Using script templates

Script templates are available for common tasks, such as creating fact or dimension tables. Place the cursor in the editor where you want to insert code, choose the **&lt;/+** icon, and then select a query template from the list.

### [](#using-the-create-external-table-template-to-import-data)Using the CREATE EXTERNAL TABLE template to import data

To create an external table, which is the first step for ingesting data into Firebolt, choose the **Import Data** button from the object pane or choose the download icon and then choose **Import data** as shown in the example below.

Firebolt creates a new tab with a `CREATE EXTERNAL TABLE` statement.

## [](#managing-scripts)Managing scripts

- [To rename a script](#scriptrename)
- [To copy a script](#scriptcopy)
- [To export a script and download it as a .sql file](#scriptexport)

**Renaming a script**[]()

- Choose the vertical ellipses next to the script name in the left pane, choose **Rename script**, type a new name, and then press ENTER.

**Copying a script**[]()

- Choose the vertical ellipses next to the script name in the left pane, choose **Duplicate script**, and then press ENTER. Firebolt saves a new script with the pattern \`\_copy.

**Exporting a script and downloading it as a .sql file**[]()

- Choose the vertical ellipses next to the script name in the left pane, and then choose **Export script**.
  
  Firebolt downloads the file to your browser’s default download directory using the file pattern `<your_script_name>.sql`.

## [](#running-scripts-and-working-with-results)Running scripts and working with results

At the bottom of each script tab, you can choose **Run** to execute SQL statements. SQL statements can only run on running engines. If an engine isn’t running, you can select it from the list and then choose the **Start** button for that engine. For more information about engines, see [Operate engines](/Guides/operate-engines/operate-engines.html)

You can run all statements in a script or select snippets of SQL to run.

**Running all SQL statements in a script**

- Position the cursor anywhere in the script editor and then choose **Run**. All SQL statements must be terminated by a semi-colon (`;`) or an error occurs.

**Running a snippet of SQL as a statement**

- Select the SQL code you want to run as a statement and then choose **Run**. Behind the scenes, Firebolt automatically appends a semi-colon to the selected SQL code so it can run as a statement.

### [](#viewing-results)Viewing results

After you run a script or query statement, more results appear below the script editor, along with statistics about query execution. The statistics section will provide further information on your statement such as its status, duration, and more.

![drawing](../../assets/images/develop_workspace_ex1.png)

### [](#viewing-multi-statement-script-results)Viewing multi-statement script results

When you run a script that has multiple SQL statements with result sets (`SELECT` statements), each result is shown on a separate line with statistics about statement execution. The first statement that ran is numbered 1 and at the bottom of the list.

To view the results table for a result set, choose the table icon as shown in the example below.

![drawing](../../assets/images/develop_workspace_ex6.png)

### [](#exporting-results-to-a-local-hard-drive)Exporting results to a local hard drive

You can export up to 10,000 rows of query results to your local hard drive after you run a query.

1. Choose the download icon (see image below).
2. Choose **Export table as CSV** or **Export table as JSON**.  
   Firebolt downloads the file type that you chose to the default download location for your browser.

It is possible to export the results of a single query alongside the results summary of all queries run in your script (with the statistics).

## [](#switching-between-light-and-dark-mode)Switching between light and dark mode

Click on the toggle at the bottom of the left navigation pane to switch between light and dark mode.

![drawing](../../assets/images/develop_workspace_ex7.png)

## [](#keyboard-shortcuts-for-the-develop-space)Keyboard shortcuts for the Develop Space

- [Query operations](#query-operations)
- [Script management](#script-management)
- [Search functionality](#search-functionality)
- [Editing text](#editing-text)

**Tip:** Use the **Keyboard shortcuts panel** (`Ctrl + Shift + ?`) to quickly view available shortcuts directly within the Develop Space.

### [](#query-operations)Query operations

Function Windows &amp; Linux Shortcut Mac Shortcut **Run** the **currently selected query**. Ctrl + Enter ⌘ + Enter **Run all** queries in the current script. Ctrl + Shift + Enter ⌘ + Shift + Enter **Toggle** expanding or collapsing **query results**. Ctrl + Alt + E ⌘ + Option + E

### [](#script-management)Script management

Function Windows &amp; Linux Shortcut Mac Shortcut **Create** a new script. Ctrl + Alt + N ⌘ + Option + N **Jump** to a **previous** script. Ctrl + Alt + \[ ⌘ + Option + \[ **Jump** to the **next** script. Ctrl + Alt + ] ⌘ + Option + ] **Close** the **current** script. Ctrl + Alt + X ⌘ + Option + X **Close all** scripts. Ctrl + Alt + G ⌘ + Option + G **Close all but** the **current** script. Ctrl + Alt + O ⌘ + Option + O

### [](#search-functionality)Search functionality

Function Windows &amp; Linux Shortcut Mac Shortcut **Open** a **search** panel. Ctrl + F ⌘ + F **Find** the **next search result**. F3 F3 **Find** the **previous search result**. Shift + F3 Shift + F3

### [](#editing-text)Editing text

Function Windows &amp; Linux Shortcut Mac Shortcut **Toggle** adding or removing a **comment marker** for the current line. Ctrl + / Cmd + / **Toggle** adding or removing a **block comment marker** around a block of code or text. Shift + Alt + A Shift + Option + A **Automatically organize and indent** code for readability. Ctrl + Alt + F ⌘ + Option + F **Copy** the selected lines and paste them directly **above** the original. Alt + Shift + Up arrow Shift + Option + Up arrow **Move** the selected lines and paste them directly **above** the original without creating a duplicate. Alt + Up arrow Option + Up arrow **Copy** the selected lines and paste them directly **below** the original. Alt + Shift + Down arrow Shift + Option + Down arrow **Move** the selected lines and paste them directly **below** the original without creating a duplicate. Alt + Down arrow Option + Down arrow **Select text** to the **left** of the cursor. Alt + Shift + Left arrow Ctrl + Shift + Left arrow **Select text** to the **right** of the cursor. Alt + Shift + Right arrow Ctrl + Shift + Right arrow **Select** the **entire line**. Alt + L Ctrl + L **Decrease** the **indentation level** of the current or selected lines. Ctrl + \[ Cmd + \[ **Increase** the **indentation level** of the current or selected lines. Ctrl + ] Cmd + ] **Delete** the current or selected **lines**. Shift + Ctrl + K Shift + Cmd + K