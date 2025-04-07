# [](#help-menu)Help menu

The **Firebolt Help menu** is found in the bottom left corner of the screen, appearing as a question mark (`?`).

![Help Menu](/assets/images/getting_to_help.png)

## [](#menu-options)Menu options

The help menu has the following options:

- Status Page
- Release Notes
- Contact Support
- Firebolt Support Access
- Documentation

### [](#status-page)Status page

The status page allows you to view the operational status of systems in real time, as well as descriptions of historical incidents.

You can subscribe to notifications whenever Firebolt **creates**, **updates**, or **resolves** an incident through the following options:

1. **Email notification** - by providing your email address.
2. **Text notifications** - by providing your phone number. *You will only receive text notifications when Firebolt either creates or resolves an incident.*
3. **Slack notifications** - by connecting via Slack. *You can receive maintenance status updates through Slack*
4. **RSS updates** - by connecting via your RSS/Atom application.

### [](#release-notes)Release Notes

The Release Notes link brings you to Firebolt’s [latest version release notes](/Reference/release-notes/release-notes.html).

### [](#contact-support---reaching-out-to-firebolt-support)Contact Support - Reaching out to Firebolt Support

**Contact Support** allows you to create a support case for Firebolt’s support team.

To create a case:

1. Select **Contact Support**.
2. A support form will appear, with the following information populated **automatically**:
   
   - **First Name**
   - **Last Name**
   - **Organization**
   - **Email**
3. Fill in the following:
   
   - **Account** - Select the name of your account within your organization.
   - **Severity** - Select the severity of your case according to the following categories:
     
     - *Critical*: You are currently experiencing a loss of critical functionality. For example, your engine is unable to start.
     - *Urgent*: Your critical functionality is being intermittently impacted. For example, your engine intermittently becomes unresponsive.
     - *Tolerable*: Your services remain operational, though non-critical functionality may be impacted. For example, a specific SQL query generates an error.
     - *Question*: Your operations are running smoothly with no disruptions. You may have a question or want to report a minor issue with the user interface.
   - **Engine name** - Enter the name of the engine with the issue.
   - **Subject** - Provide a clear and descriptive subject.
   - **Description** - Provide any relevant details including the following:
     
     - What were you trying to do when the issue occurred?
     - What was your expected outcome?
     - What actually happened?
     - What errors were returned?
4. Select **Submit**.
5. Your case will be sent to our Support team, and you will receive a confirmation email.

### [](#firebolt-support-access)Firebolt Support Access

Firebolt Support Access allows you to manage and control the level of access that Firebolt’s support team has to your account.

For new accounts, support access is enabled by default; however you can revoke it at any time.

All support team activity is logged and can be viewed in the [query history](../sql_reference/information-schema/engine-query-history).

To grant access to the support team:

1. Select “Firebolt Support Access”
2. A support access form will appear, with the following information populated **automatically**:
   
   - **Account Name**
   - **Duration**
   - **Assign Roles**
3. Fill in the following:
   
   - **Duration** - Specify the duration for which you’d like to grant access to the support team. Once this period ends, access will be automatically revoked.
   - **Assign Roles** - Select the role(s) that you want to assign to the support team. Their access will be limited to these roles. You can create a dedicated role by following the [instructions here](/Overview/Security/Role-Based%20Access%20Control/role-management/custom-roles.html). Note that the support team will have full access data at the organization level.
4. Select **Grant Access**.

Once access is granted, an additional icon will appear in the bottom left corner of the screen, indicating that Firebolt’s support team has access to your account, as shown below:

![Support Access](/assets/images/support_access_menu.png)

To edit support team access:

1. Select the icon that appeared when you granted access to the support team. You can find this icon in the bottom left corner of your **Firebolt Workspace**.
2. Select **Manage access**. You can also open the same screen by selecting **Firebolt Support Access** in the **Help** menu.
3. Select **Edit Access**.
4. After you have made your changes, select **Update Access**.

To revoke support team access:

1. Select the icon that appeared when you granted access to the support team. You can find this icon in the bottom left corner of your **Firebolt Workspace**.
2. Select **Manage access**. You can also open the same screen by selecting **Firebolt Support Access** in the **Help** menu.
3. Select **Revoke Access**.

Once revoked, the support access icon will disappear from the bottom left corner of the screen.

### [](#documentation)Documentation

The Documentation link brings you to Firebolt’s docs (where you are now!)