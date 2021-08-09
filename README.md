# Golayout
Golayout is a boilerplate project that containing the usage of the best practice and popular components. The code organization follows the [standard go project layout](https://github.com/golang-standards/project-layout).

## Why this project?
Golang is a simple programming language. The standard library has many tiny and useful wheels to help your development efficiently. But there are no frameworks such as [Ruby on Rails](https://rubyonrails.org/) or [Java Spring Boot](https://spring.io/projects/spring-boot) to help you create a runnable scaffold project fast. You must combine every tiny 'wheel' manually for your `project`.

I had experienced several Golang projects and saw many of them made things wrong in the early stage. When the project
was going to the middle of development, it has so many troubles that can't refactor the code or design painlessly. For
standing **the broken window theory** on
the book < [The Pragmatic Programmer](https://www.amazon.com/dp/B0833FBNHV?plink=KCNIUfDkqIUcf4x1&ref=adblp13nvvxx_0_2_im) >, we
should avoid the broken window in the early stages. Having a good and well-design scaffold is very important.

For learning the coding skills, we should have an environment to practice or implement our ideas to check the code whether may or may not be done. Every programmer never wants to run the pale-running code in the production without well testing. For purpose of this, we must have an easy setup environment for testing the immature code. The environment should be the same as our production, but the frustrating configurations always bothered us to running the immature code.

## Project Achieved
Suppose you want to use Golang for starting a new backend project. There are many works  might that have been achieved repeatedly.   
* Building a makefile for construction works.
* Choosing and importing an RPC framework.   
* Integrating logger.   
* Writing the service register/discovering.
* Designing the RESTFul HTTP server/doc/testing.
* CI/CD integration.
  
  [comment]: <todo> (* Tracing the calling route.   )
  
In this project, all these works have been finished. I don't write any special business code for this project. You can clone this repository and rename it for your own. It can help you set up a Golang project immediately.

The principles applying to this project:    
* DRY(Don't repeat yourself)    
* Being simple, and follow the k.i.s.s(keep it simple, stupid).    
* Package-oriented design and all packages have a public `Init` function if needed. *See [The problems of Golang init function
](https://liyafu.com/2021-07-07-the-problems-of-go-init-func/)* 
* Separated the library and business code.  
* Applying the popular components and best practices.   

## Usage
### Install
```shell
git clone git@github.com:leyafo/golayout.git your_project
cd your_project 
sh scripts/rename.sh your_project  #rename the project name for your own
```

### Build
```shell
make proto #build protobuf
make build   
```

### Build the docker images
```shell
make docker_build
```

### Run all apps
```shell
make up
```

### Testing
```shell
curl -i http://localhost:9527/v1/version
```

## Description of the project organization
```shell
├── bin             //store executable files     
├── buf.yaml        //protobuf building configurations
├── buf.gen.yaml    //protobuf building configurations 
├── cmd             //the name of binaries you want to build.    
├── conf            //application configurations, it should be ignored by git.  
├── deployment      //store the Dockerfile, docker-compose or other deployment configurations  
├── Makefile        //the only one make file that can build, test, pack, run the project  
├── pkg             //the library used in the project  
├── scripts         //some install or set up scripts  
├── tools           //the tools used in this project
└── internal        //business code that cannot be importable
```

## Contribution
If you have any advice, thoughts, and best practice, feel free to commit the pull requests. This project doesn't have any historical burden or business trad-off. You can implement your thoughts in this project without trouble.

## License
MIT License

Copyright (c) 2021 李亚夫