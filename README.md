# web-bridge

ngrok http --domain=sharp-happy-grouse.ngrok-free.app 80

firebase emulators:start --only firestore --project dessage


<script src="https://cdn.ethers.io/lib/ethers-5.2.umd.min.js"></script>
const message = "你的消息内容";
const messageBytes = ethers.utils.toUtf8Bytes(message);
const hash = ethers.utils.keccak256(messageBytes);

console.log("哈希值: ", hash);

#nginx 
sudo setsebool -P httpd_can_network_connect 1
sudo setsebool -P httpd_read_user_content 1


#firewall
sudo firewall-cmd --permanent --add-service=http
sudo firewall-cmd --permanent --add-service=https
sudo firewall-cmd --zone=public --add-port=8880/tcp --permanent
sudo firewall-cmd --zone=public --add-port=8881/tcp --permanent
sudo firewall-cmd --reload
