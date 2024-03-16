FROM  golang:1.22 as build

WORKDIR /app

# Copy dependencies list
COPY . .

# Build with optional lambda.norpc tag as newer images do not require rpc
RUN go build -tags lambda.nrpc -o main ./app1/cmd

# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2023

COPY --from=build /app/main /var/task/bootstrap
