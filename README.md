# sonar cloud

here we use sonar cloud which make ease to config



# ACR

create azure container registory

In access key there is login server

docker login <login server>
  username: <Username>
  password: <Password1>


docker build -t <login server>/imagename:latest

# change of version in helmchart docker image
this is the sample command for changing the version in helm chart because if we create new container we need to push it so we need to change it 

sed -i "s|image: tharun13055/library:.*|image: tharun13055/library:1|" ./golang_chart/values.yaml

this is for helm chart 

sed -i "s|image: tharun13055/library:.*|image: tharun13055/library:{{ github.run_number }}|" ./golang_chart/values.yaml