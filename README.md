# Export env file

export $(<.env)

# Create file .github/workflows
    main.yaml

# Check docker log
    /var/lib/docker/containers/

# publisher/server attribute
    deploy:
      mode: global
# consumer attribute
    deploy:
      mode: replicated
      replicas: 1

# Pull image from docker hub after cicd github action 
    -- when 32bit
        docker pull --platform=linux/amd64 wiliehidayat87/testmq-publisher:latest
    -- else
        docker pull wiliehidayat87/testmq-listener:latest
        docker pull wiliehidayat87/testmq-consumer:latest
        docker pull wiliehidayat87/testmq-publisher:latest
# Init service 
    docker swarm init
# Add task to service 
    docker stack deploy -c docker-compose.core.yaml core
    docker stack deploy -c docker-compose.listener.yaml listener
    docker stack deploy -c docker-compose.consumer.yaml consumer
    docker stack deploy -c docker-compose.publisher.yaml publisher
# Check tasks from service  
    docker service ls
# Remove task from service 
    docker service rm publisher_listener-service
# Add worker task
    docker service scale consumer_listener-service=3

