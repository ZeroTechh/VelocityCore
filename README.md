# VelocityCore
The central codebase of velocity containing common utilities imported by each service

## config
This folder contains configurations for whole of velocity. The config is written in go instead of yaml due to file reading issues. In future the config has to change to yaml

## proto
This folder contains proto definations for each service of velocity

## logger
This is a utility which returns a zap logger. Each service uses this to get a zap logger. This is created to reduce zap logger initialization code in each service and also an easy way to change logging configs such as logs location

## utils
This folder contains lot of random utilities used by each service

## services
A simple record of ip address of each service. Kind of acts like a service discovery. However in future this will be replaced by a more robust service discover