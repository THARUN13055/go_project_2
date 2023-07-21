# Golang_Project_2 Using Github_action

***

## Laungage

- Golang
- Database (MYSQL)

## CICD tool

- Github Action
 
## Used Tools 

- SonarQube Cloud
- Docker && Docker-compose
- Azure container Registory
- trivy
- Azure Storage Account
- Helm
- Azure Kubernetes
- Slack-Notification

## About Golang_code
  
  This is a simple project which has Golang with connection of database as Mysql.
  And here we are use a **GET** and **POST** method.

## Steps of this project

1. First we are check wether the project is working in **localhost:8080/database** so we us **docker-compose** to verify.

2. Push to the github.

3. We use Github action as CICD pipline for this project.

4. Build Golang code

5. And make sure setup your Mysql-server. And create username as root and **password** as **password**

6. Now Scan the Code for Quality checking and bugs So we use SonarQube cloud and write a propertitits file and connect to the github action.

7. Create a Dockerfile and mension the steps and we are used Alpine image.

8. Create a Azure Container Registory and connect to the Github action.

9. Build the image & Push to the ACK.

10. We are Deploying our porject image in  Azure Kubernetes.

11. And we are using artifactory as Azure Storage Account for store our build file **library** .

12. If the Pipline is success of failed we need to intimate so we use Stack-Notification.

## For installing tools like Docker,Sonarqube,nexus,Argocd,etc.

### For all this Installation Use This 'wget'

```bash
wget https://raw.githubusercontent.com/THARUN13055/quick_install_linux/main/pacman.py
```
### How its work

```bash
wget https://raw.githubusercontent.com/THARUN13055/quick_install_linux/main/pacman.py

python3 pacman.py list

python3 pacman.py install <listed>
```
### Sample installing of Docker

```bash
wget https://raw.githubusercontent.com/THARUN13055/quick_install_linux/main/pacman.py

python3 pacman.py list

python3 pacman.py install docker
```

# Issues and Login


## To login in ACK

- create azure container registory.

- In access key there is login server.

- docker login <login server>
  username: <Username>
  password: <Password1>

- After login build our image with login server name.

```bash
docker build -t <login server>/imagename:latest .
```

## Login in Kubernetes using azure app registory

1. Create azure App Registory in Azure Directory.

2. In Azure App Registory create Secrets.

3. Copy all the tokens like **Subscription ID, Tenant ID, client_secret ID,client ID** .

5. Create secret in github action and past.

6. if your azure login is not work go to azure app registory This are the steps use and create azure credientials.

7. app registory > certificats&secrets > add credientials (organization: tharun13055) (reponame: go_project_2) > select branch and name main.

## Azure Kubernetes Service

1. Create Azure kuberentes service Resource name as **kuberentes**.

2. Cluster name as **golang**.


## Change of version in helmchart docker image
- this is the sample command for changing the version in helm chart because if we create new container we need to push it so we need to change it.

```bash
sed -i "s|image: tharun13055/library:.*|image: tharun13055/library:1|" ./golang_chart/values.yaml
```

- This is for helm chart 

```bash
sed -i "s|image: tharun13055/library:.*|image: tharun13055/library:{{ github.run_number }}|" ./golang_chart/values.yaml

```