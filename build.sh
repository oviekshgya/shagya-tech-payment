sudo git pull origin v.1
sudo docker container stop shagya-tech-payment
sudo docker container rm shagya-tech-payment
sudo docker image rm shagya-tech-payment
sudo docker-compose up --build -d
sudo service nginx reload
sudo service nginx restart