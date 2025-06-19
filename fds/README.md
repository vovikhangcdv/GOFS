This Go project monitors blockchain transactions and detects abnormal activity, specifically focusing on ERC20 token transfers and automatic blacklisting of suspicious addresses.

## FDS Dashboard
![alt text](./public/image.png)

### 1. Detect suspicious transactions in network 
![alt text](./public/image2.png)
Details relate supspicious transaction selected

![alt text](./public/image3.png)

### 2. View details transaction related addresses in network

![alt text](./public/image4.png)


![alt text](./public/Animation.gif)

### 3. Define rules to detect and take action relate suspicious transactions
![alt text](./public/image6.png)

### 4. Manage blacklisted address block by auto/manual 

![alt text](./public/image8.png)

###
To start fds:
```
docker-compose up --build -d
```