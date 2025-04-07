# [](#creating-an-access-key-and-secret-id-in-aws)Creating an access key and secret ID in AWS

This section will walk you through the steps to create security credentials in AWS. These credentials will be used to load data from AWS S3 into Firebolt.

In order to enable Firebolt to load data from your S3 buckets, you must:

1. Create a user.
2. Create appropriate permissions for this user.
3. Create access credentials to authenticate this user.

## [](#create-a-user)Create a user

1. Log into your AWS Management console and go to the IAM section. You can do this by typing “IAM” in the search bar.
2. Once you are in the IAM section, select the **Create User** button.
   
   ![Create IAM User](/assets/images/Create_User_Dialog.png)
3. Enter a name for the user and select **Next**.
   
   ![Specify User Name](/assets/images/Specify_User_Name.png)
4. You can have the default permission option set to **Add user to group** and select **Next**.
   
   ![Set Permissions](/assets/images/Set_Permissions.png)
5. Select **Create User**.
   
   ![Review and Create User](/assets/images/Review_Create_User.png)
6. You will see a message **User created successfully**.
   
   ![User created successfully](/assets/images/User_Created_Successfully.png)

## [](#create-s3-access-permissions)Create S3 access permissions

Now that you have created the user, you will now assign this user appropriate permissions for S3.

1. Select on the user name as shown below.
   
   ![Click User](/assets/images/Click_User.png)
2. In the Permissions tab, select the **Add Permissions** drop-down and choose **Create inline policy**.
   
   ![Choose Inline Policy](/assets/images/Choose_Iniline_Permissions.png)
3. In **Specify Permissions** choose S3 as the service.
   
   ![Choose S3](/assets/images/Choose_S3.png)
4. Choose **JSON**, paste the following JSON code in the policy editor, and select **Next**.
   
   ![Set Permissions](/assets/images/Specify_Permissions.png)
   
   ```
   {
      "Version": "2012-10-17",
      "Statement": [
          {
              "Effect": "Allow",
              "Action": [
                  "s3:GetObject",
                  "s3:GetObjectVersion"
              ],
              "Resource": "arn:aws:s3:::<bucket>/<prefix>/*"
          },
          {
              "Effect": "Allow",
              "Action": "s3:GetBucketLocation",
              "Resource": "arn:aws:s3:::<bucket>"
          },
          {
              "Effect": "Allow",
              "Action": "s3:PutObject",
              "Resource": "arn:aws:s3:::<bucket>/*"
          },
          {
              "Effect": "Allow",
              "Action": "s3:ListBucket",
              "Resource": "arn:aws:s3:::<bucket>",
              "Condition": {
                  "StringLike": {
                      "s3:prefix": [
                          "<prefix>/*"
                      ]
                  }
              }
          }
      ]
   }
   ```
   
   **IMPORTANT:** Replace “&lt;bucket&gt;” with the S3 bucket that you want to provide access to.
5. Enter a description for the policy and select **Create Policy**.
   
   ![Create policy](/assets/images/Create_Policy.png)
6. You will see a message that the policy has been successfully created.

## [](#create-access-key-and-secret-id)Create access key and secret ID

Now that you have created a user, authorized the user with the appropriate S3 permissions, you will create access credentials for this user. These credentials will be used to authenticate the user.

1. Select the **Security Credentials** tab, as shown in the following image:
   
   ![Security Credentials](/assets/images/Choose_Security_Credentials.png)
2. In the **Access Keys** section, select the **Create Access Key** button.
   
   ![Create Access Key](/assets/images/Create_Access_Keys.png)
3. For the use case, choose the **Application running on AWS compute service**. You will see an alternative recommendation. You can check the box that says “I understand the above recommendation and want to proceed to create an access key” and select **Next**.
   
   ![Use case warning](/assets/images/Access_Key_Use_Case.png)
4. Set a description tag for the access key and select \*\*C, ![Set Desc tag Access Key](/assets/images/Description_Tag_Access_Key.png)
5. You will see a message indicating that the access key was created. Make sure to download the access key. You will need these credentials when you load S3 data into Firebolt.
   
   ![Access Key created successfully](/assets/images/Download_CSV_Access_Key.png)