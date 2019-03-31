[![CircleCI](https://circleci.com/gh/iamabhishek-dubey/cloud-auditor.svg?style=svg)](https://circleci.com/gh/iamabhishek-dubey/cloud-auditor)

# OpsTree Cloud Auditor
The main goal of creating this tool is to provide easier and faster security scan of the cloud environment. As a result, the consumer will get to know exactly where his/her security is lacking.

The program is written in Golang which enhances the performance of tool in terms of speed.

## Supported Cloud Environment
The cloud environment which is supported right now is:-

- [X] **Amazon Web Services(AWS)**
- [ ] **Google Cloud Platform(GCP)**
- [ ] **Microsoft Azure**

### Amazon Web Services(AWS) Checks
In AWS, we are checking for these services:-

- [X] **EC2**
- [ ] **S3**
- [ ] **IAM**

## Requirments
The requirements for using this tools are:-

- **Golang**

You can skip golang installation if you are using the *Dockerized Setup*.

## Parameters
Here is the list of parameters which is accepted by this tool.

|**Parameter**|**Supported Values**|**Description**|
|-------------|--------------------|---------------|
|--service|<ul><li>ec2</li><li>s3</li><li>iam</li></ul>| Name of the service on which you want to perform the audit. Right now it supports **ec2** only|
|--region| Any AWS region|Name of the region in which you want to perform the audit|

## How to Use
This tool is pretty much straight forward for use. We have categorized it in two parts i.e. **Manual Setup** and **Dockerized Setup**

### Manual Setup
Steps for Manual Setup

```shell
git clone git@github.com:iamabhishek-dubey/cloud-auditor.git
cd cloud-auditor
make get-depends
make build-code
```

### Dockerized Setup
Steps for Dockerized Setup

```shell
make build-image
docker run -itd --name cloud-auditor -v ${HOME}/.aws/:/root/.aws/ cloud-auditor:latest
```
