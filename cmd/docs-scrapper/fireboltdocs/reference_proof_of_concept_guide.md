# [](#firebolt-proof-of-concept-technical-guidelines)Firebolt proof of concept technical guidelines

> This guide should be used after you have met with your Firebolt designated sales team. If you are looking to try out Firebolt and have yet to contact us, [get started here.](https://www.firebolt.io/getting-started-now)

- [Step 1 - Provide access to S3](#step-1---provide-access-to-s3)
- [Step 2 - Provide your Firebolt team with relevant assets](#step-2---provide-your-firebolt-team-with-relevant-assets)

## [](#step-1---provide-access-to-s3)Step 1 - Provide access to S3

The data used in the proof of concept should be made available in an S3 region ([any region](/Reference/available-regions.html) is fine). To ensure the proof of concept accurately reflects your production workload, it’s recommended to have all or a significant portion of the data available.

Firebolt pays for data transfer costs if such occurs. You can find additional information in [this AWS guide.](https://docs.aws.amazon.com/AmazonS3/latest/userguide/RequesterPaysBuckets.html)

Please go through the following steps to grant Firebolt access to the relevant S3 bucket. You can find additional information in [this AWS guide.](https://docs.aws.amazon.com/AmazonS3/latest/userguide/add-bucket-policy.html)

1. Connect to the AWS console and click on the relevant S3 bucket.
2. Go to *Permissions* and then scroll down and edit *Bucket Policy*.
3. Copy the following policy. Make sure you replace `<bucket>` with the actual bucket name.
   
   ```
    {
     "Version": "2012-10-17",
     "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::231290928314:root"
      },
      "Action": [
        "s3:GetObject",
        "s3:GetObjectTagging",
        "s3:PutObject",
        "s3:PutObjectTagging",
        "s3:ListBucket",
        "s3:GetBucketLocation"
      ],
      "Resource": [
        "arn:aws:s3:::<bucket>",
        "arn:aws:s3:::<bucket>/*"
      ]
    }
     ]
   }
   
   ```
4. Save changes

## [](#step-2---provide-your-firebolt-team-with-relevant-assets)Step 2 - Provide your Firebolt team with relevant assets

Please share the following assets with your Firebolt team through your designated Slack channel or email.

Please make sure all files are saved as a `sql` file type (i.e, `file_name.sql`)

1. **Schema** - this file should include the DDL commands that we should use to create the desired database schema.
2. **Sample queries &amp; average durations** - this file should include a set of queries to be executed over the data set. These queries might be:
   
   - A good representation of queries you usually run
   - Queries that are slow and you would like to see improved
   
   For each query, add a comment in the file describing the average duration, or any additional metric that might be relevant.