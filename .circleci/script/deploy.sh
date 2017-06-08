#!/usr/bin/env bash
# variable
AWS_DEFAULT_REGION=ap-northeast-1
AWS_ECS_TASKDEF_NAME=familog-production
AWS_ECS_CLUSTER_NAME=familog-production
AWS_ECS_SERVICE_NAME=familog-production
AWS_ECR_REP_NAME=familog

push_ecr_image(){
	eval $(aws ecr get-login --region ${AWS_DEFAULT_REGION})
	docker push $AWS_ACCOUNT_ID.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com/${AWS_ECR_REP_NAME}:$CIRCLE_SHA1
	docker tag $AWS_ACCOUNT_ID.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com/${AWS_ECR_REP_NAME}:$CIRCLE_SHA1 $AWS_ACCOUNT_ID.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com/${AWS_ECR_REP_NAME}:latest
	docker push $AWS_ACCOUNT_ID.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com/${AWS_ECR_REP_NAME}:latest
}

docker build -t $AWS_ACCOUNT_ID.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com/${AWS_ECR_REP_NAME}:$CIRCLE_SHA1 .
push_ecr_image
wget https://github.com/crowdworks/ecs-goploy/releases/download/v0.2.0/ecs-goploy_v0.2.0_linux_amd64.zip
unzip ecs-goploy_v0.2.0_linux_amd64.zip
./ecs-goploy deploy -c $AWS_ECS_CLUSTER_NAME -n $AWS_ECS_SERVICE_NAME -i $AWS_ACCOUNT_ID.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com/${AWS_ECR_REP_NAME}:$CIRCLE_SHA1 -t 300
