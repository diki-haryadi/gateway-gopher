build:
	go build -o db-tunnel-dialer

run-cobra-key-pg:
	go run main.go gw \
      --ssh-host="example.com" \
      --ssh-port=22 \
      --ssh-user="username" \
      --ssh-password="password" \
      --ssh-key="~/.ssh/id_rsa" \
      --ssh-auth-type="key" \
      --db-host="dbserver.internal" \
      --db-port=5432 \
      --local-host="127.0.0.1" \
      --local-port=5432 \
      --toggle=false

run-cobra-basic-pg:
	go run main.go gw \
      --ssh-host="example.com" \
      --ssh-port=22 \
      --ssh-user="username" \
      --ssh-password="password" \
      --ssh-auth-type="basic" \
      --db-host="dbserver.internal" \
      --db-port=5432 \
      --local-host="127.0.0.1" \
      --local-port=5432 \
      --toggle=false

run-cobra-key-mysql:
	go run main.go gw \
      --ssh-host="example.com" \
      --ssh-port=22 \
      --ssh-user="username" \
      --ssh-password="password" \
      --ssh-key="~/.ssh/id_rsa" \
      --ssh-auth-type="key" \
      --db-host="dbserver.internal" \
      --db-port=3306 \
      --local-host="127.0.0.1" \
      --local-port=3306 \
      --toggle=false

run-cobra-basic-mysql:
	go run main.go gw \
      --ssh-host="example.com" \
      --ssh-port=22 \
      --ssh-user="username" \
      --ssh-password="password" \
      --ssh-auth-type="basic" \
      --db-host="dbserver.internal" \
      --db-port=3306 \
      --local-host="127.0.0.1" \
      --local-port=3306 \
      --toggle=false