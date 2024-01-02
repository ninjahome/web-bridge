# web-bridge

ngrok http --domain=sharp-happy-grouse.ngrok-free.app 80

firebase emulators:start --only firestore --project dessage


#nginx 
sudo setsebool -P httpd_can_network_connect 1
sudo setsebool -P httpd_read_user_content 1


#firewall
sudo firewall-cmd --permanent --add-service=http
sudo firewall-cmd --permanent --add-service=https
sudo firewall-cmd --zone=public --add-port=8880/tcp --permanent
sudo firewall-cmd --zone=public --add-port=8881/tcp --permanent
sudo firewall-cmd --reload
