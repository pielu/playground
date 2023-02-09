# Challenge? Sure

## Assumptions

- The `ModifiedAt` of the S3 object key is used to determine the date of deployment
- Disregard apparent issue with trusting `ModifiedAt` as always accurate (per <https://amacal.medium.com/aws-s3-last-modified-date-caff5a485ce0>)

## Caveats

- The testing is not super robust and could hit some edge case when it comes to "latest" objects
- The amount of abstraction is not as nice as it could, should that be more than quick script
- The AWS SDK configuration / authn / authz chain is not tested beyond localstack

## How to test run

1. From within the sure directory

    ```text
    export AWS_REGION=us-east-1
    export AWS_ENDPOINT=http://0.0.0.0:4566
    export AWS_ACCESS_KEY_ID=test
    export AWS_SECRET_ACCESS_KEY=test
    export S3_BUCKET=sure
    export S3_BUCKET_POPULATE_PATH=./s3bucket
    ```

    ```text
    go test -v
    ```

2. From within vscode

    Just open `sure` folder and `main_test.go` file, assuming `golang` plugin configured, click `run test` over `TestSure`. To make vscode include vars, use the following configuration

    ```text
    "go.testEnvFile": "${workspaceFolder}/.env",
    ```

## How to run as a program accepting the X parameter
