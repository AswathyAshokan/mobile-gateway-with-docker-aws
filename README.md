# mobile-gateway-with-docker-aws
#it is a simple golang api using gorilla mux

#first install  docker on your machine and  move to the current directory where your project reside,run the below code
#sudo docker login 
#sudo docker build -t  username/imagename .
#docker run --rm -p 8080:8080 username/imagename:latest

#for running the mysql image
#docker run --name=test-mysql mysql
#docker ps -a
#docker run --name=test-mysql --env="MYSQL_ROOT_PASSWORD=password" mysql


#create the instance in the aws add custom tcp with acess for port 8080 and create rds of mysql,update the endpoint in the codebase (where marked as hostname)
 #connect with the aws instance using(public ip)
 #ssh -i "myApp.pem" ec2-user@public ip
 #run docker image in the aws instance using docker run --rm -p 8080:8080 username/imagename:latest
 
