##
# bootcamp 2020 , Redis 
# Fri Jun  5 11:42:16 JST 2020

## Document
# https://hub.docker.com/_/redis/
# https://github.com/dockerfile/redis

# Pull base image.
FROM redis:5.0.9

# Install tools
RUN apt-get update && \ 
    apt-get install -y git curl net-tools unzip procps && \
    rm -rf /var/lib/apt/lists/*
    # mkdir -p /etc/redis
    # cp -f *.conf /etc/redis && \
    # sed -i 's/^\(bind .*\)$/# \1/' /etc/redis/redis.conf && \
    # sed -i 's/^\(daemonize .*\)$/# \1/' /etc/redis/redis.conf && \
    # sed -i 's/^\(dir .*\)$/# \1\ndir \/data/' /etc/redis/redis.conf && \
    # sed -i 's/^\(logfile .*\)$/# \1/' /etc/redis/redis.conf
         
# Define mountable directories.
# VOLUME ["/data"]

# Define working directory.
WORKDIR /data

# Sample DATA and go binary
COPY redis.conf /etc/redis/
COPY list_* /data/
COPY main* /data/

# Define default command.
# CMD ["redis-server", "/etc/redis/redis.conf"]
CMD ["redis-server"]

# Expose ports.
# EXPOSE 6379
