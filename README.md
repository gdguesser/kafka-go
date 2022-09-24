# Useful commands

- Create topic command:
    
    ```docker
    docker exec broker \
    kafka-topics --bootstrap-server broker:9092 \
                 --create \
                 --topic quickstart
    ```
    
- Write message to the topic:
    
    ```
    docker exec --interactive --tty broker \
    kafka-console-producer --bootstrap-server broker:9092 \
                           --topic quickstart
    ```
    
    - Type in some lines of text. Each line is a new message.
- Read messages from the topic:
    
    ```
    docker exec --interactive --tty broker \
    kafka-console-consumer --bootstrap-server broker:9092 \
                           --topic quickstart \
                           --from-beginning
    ```