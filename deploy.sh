 #!/bin/bash

key="$HOME/.ssh/id_ed25519"

scp -i "$key" docker-compose.yml root@5.63.152.17:/root/docker-compose.yml
ssh -i "$key" root@5.63.152.17 docker-compose up -d
