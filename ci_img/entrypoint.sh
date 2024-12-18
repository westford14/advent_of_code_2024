#! /bin/bash

# Since we want to use the CircleCI docker images instead of having to use a machine
# executor from scratch we need to build our own bazel-remote image. Then we need to
# start the service so we can start interacting with our backend cache in S3.  This
# then requires us to pump the credentials into the docker container without using
# volume mounts. This fun little script is a way around this by opening up a separate
# port and posting in the needed credentials and then starting the service once the
# credentials have been consumed.
# 
# More information on this can be found in the README
function handleRequest {
    while IFS= read -r INPUT || [ "$INPUT" ]; do
        if [[ $INPUT == *"AWS_ACCESS_KEY_ID"* ]]; then
            echo "Received start request"
            FULL_DATA=$INPUT
            AWS_ACCESS_KEY_ID=$(jq -r  '.AWS_ACCESS_KEY_ID' <<< "$FULL_DATA")
            AWS_SECRET_ACCESS_KEY=$(jq -r  '.AWS_SECRET_ACCESS_KEY' <<< "$FULL_DATA")
            AWS_SESSION_TOKEN=$(jq -r  '.AWS_SESSION_TOKEN' <<< "$FULL_DATA")
            echo "writing credential file"
            mkdir -p etc
            touch etc/credentials
            printf "[default]\naws_access_key_id = $AWS_ACCESS_KEY_ID\naws_secret_access_key = $AWS_SECRET_ACCESS_KEY\naws_session_token = $AWS_SESSION_TOKEN\n" >> etc/credentials
            echo "creating cache directory"
            mkdir -p data
            echo "starting the bazel-remote server"
            /bazel-remote --dir=data \
                --http_address=:8080 \
                --s3.bucket=bazel-artifacts \
                --s3.auth_method=aws_credentials_file \
                --s3.region=eu-north-1 \
                --s3.endpoint=s3.eu-north-1.amazonaws.com \
                --s3.disable_ssl=true \
                --max_size=100 \
                --s3.aws_shared_credentials_file=etc/credentials
        fi
    done;
}

echo 'Listening on 3000...'

while true; do
    echo -e 'HTTP/1.1 200 OK\r\nContent-Length: 0\r\n\r\n\r\n</h1>PONG</h1>' | nc -l 3000 | handleRequest
done