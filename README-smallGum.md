# Fork location

https://github.com/smallGum/microservice-agenda

# My commits

+ a8e7e72 - add cli test result
+ 4dd96c4 - create shared file for docker and my PC
+ 2b2a342 - supplement necessary script
+ ee27a77 - remove unnecessary comment
+ 7d105d6 - change cli-agenda's data storage from text files to sqlite3 database
+ 9fd6d13 - fix some bugs
+ 424a6af - modified README file
+ cd4074d - update README file
+ 5624801 - add user and meeting test for command-line agenda
+ 4ced821 - clear test contents
+ c698a8b - clear all test contents of files
+ d431ad0 - clear all contents of files
+ 4475d6e - add user test for command-line agenda
+ c110bd0 - add necessary files for test
+ 4b0f46f - update README file
+ 240b782 - fix some bugs
+ 8d253ab - fix some bugs
+ d3e3ed7 - fix some bugs
+ 61fe4e6 - fix some bugs
+ 0b06e8e - update .travis.yml file
+ 0299ace - update Dockerfile
+ 0846f8d - add user test file
+ 5cc4163 - fix some bugs of packets path
+ 2f84ef8 - fix some bugs
+ 8c92a2a - change dockerfile to Dockerfile
+ 99b6e7d - change dockerfile name
+ 3d3a0cb - finish meeting test
+ fa4953e - Merge pull request #6 from chengr25/master
+ 6807a92 - test user's handler
+ 4d08b99 - Merge pull request #5 from chengr25/master
+ e7f779e - register
+ 7301ec9 - Merge pull request #4 from chengr25/master
+ 48002c7 - --
+ 4218c63 - modify the handler directory
+ 9c4e866 - --
+ 8166f30 - Merge pull request #3 from chengr25/master
+ 54aa7df - Merge branch 'master' into master
+ 3e263cd - Merge branch 'master' into master
+ 668fbab - modify the user's part of ApiBlueprint
+ a99eee1 - Merge branch 'master' of github.com:chengr25/microservice-agenda
+ a3e53ab - some handlers about user
+ 625dd98 - add extra responses for meetings API
+ 7eb7d43 - fix some bugs
+ 4bbefbf - complete all meeting actions
+ 8294b35 - Merge pull request #2 from chengr25/master
+ 66f9795 - Merge branch 'master' into master
+ 2702cba - Merge branch 'master' into master
+ 234ee55 - complete meeting creation handler
+ 72d5732 - implement some functions of user struct
+ ca3b97d - Merge pull request #1 from chengr25/master
+ b7e8676 - define user struct,meeting struct
+ bced409 - update service for database initialization
+ 5bd5cf9 - add more descriptions for meetings' API
+ d3de005 - update API for meetings in apiary
+ a0b20c9 - add API Design v1.0
+ 8a7d117 - add apiary synchronization
+ e31cd6a - initialize the microservice agenda program repository

# Summary

Through this project, I've learned how to lauch parallel and independent development of client-side and server-side of a program, design API using API Blueprint and push my program to docker hub for others to download it as an image, so that they can run my program on any operating system platform.

API Blueprint is an amazing tool to design API for program. The powerful RESTful API allows us to develop client-side and server-side parallelly. Client-side can use mock server on [apiary](https://apiary.io/) to get response from url without waiting for server-side to finish. This greatly improves the efficiency and the independence between client and server development.

Besides, [docker hub](https://hub.docker.com/) provides graceful containers for our program to run on any platform. However, deploying our program on docker hub requires skills to code Dockerfile. Also, we should remember necessary docker commands in order to download our program's image and run it successfully. Some basic comprehension like volume of docker is of great importance.